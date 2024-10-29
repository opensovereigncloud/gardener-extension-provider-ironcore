// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package controlplane

import (
	"context"
	"strconv"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/coreos/go-systemd/v22/unit"
	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	gcontext "github.com/gardener/gardener/extensions/pkg/webhook/context"
	"github.com/gardener/gardener/extensions/pkg/webhook/controlplane/genericmutator"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	gardenerutils "github.com/gardener/gardener/pkg/utils/gardener"
	imagevectorutils "github.com/gardener/gardener/pkg/utils/imagevector"
	testutils "github.com/gardener/gardener/pkg/utils/test"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	vpaautoscalingv1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	"k8s.io/utils/ptr"
)

const (
	namespace               = "foo"
	portProviderMetrics     = 10259
	portNameProviderMetrics = "providermetrics"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ControlPlane Webhook Suite")
}

var _ = Describe("Ensurer", func() {
	var (
		ctx = context.TODO()

		ctrl *gomock.Controller

		ensurer genericmutator.Ensurer

		dummyContext = gcontext.NewGardenContext(nil, nil)

		eContextK8s = gcontext.NewInternalGardenContext(
			&extensionscontroller.Cluster{
				Shoot: &gardencorev1beta1.Shoot{
					Spec: gardencorev1beta1.ShootSpec{
						Kubernetes: gardencorev1beta1.Kubernetes{
							Version: "1.26.0",
						},
					},
				},
			},
		)
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ensurer = NewEnsurer(logger, false)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("#EnsureKubeAPIServerDeployment", func() {
		var dep *appsv1.Deployment

		BeforeEach(func() {
			dep = &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: v1beta1constants.DeploymentNameKubeAPIServer},
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "kube-apiserver",
								},
							},
						},
					},
				},
			}
		})

		It("should add missing elements to kube-apiserver deployment", func() {
			err := ensurer.EnsureKubeAPIServerDeployment(ctx, eContextK8s, dep, nil)
			Expect(err).To(Not(HaveOccurred()))

			checkKubeAPIServerDeployment(dep)
		})

		It("should modify existing elements of kube-apiserver deployment", func() {
			dep = &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: v1beta1constants.DeploymentNameKubeAPIServer},
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "kube-apiserver",
									Command: []string{
										"--cloud-provider=?",
										"--cloud-config=?",
									},
								},
							},
						},
					},
				},
			}

			err := ensurer.EnsureKubeAPIServerDeployment(ctx, eContextK8s, dep, nil)
			Expect(err).To(Not(HaveOccurred()))

			checkKubeAPIServerDeployment(dep)
		})
	})

	Describe("#EnsureKubeControllerManagerDeployment", func() {
		var dep *appsv1.Deployment

		BeforeEach(func() {
			dep = &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: v1beta1constants.DeploymentNameKubeControllerManager},
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "kube-controller-manager",
								},
							},
						},
					},
				},
			}
		})

		It("should add missing elements to kube-controller-manager deployment", func() {
			err := ensurer.EnsureKubeControllerManagerDeployment(ctx, eContextK8s, dep, nil)
			Expect(err).To(Not(HaveOccurred()))

			checkKubeControllerManagerDeployment(dep)
		})

		It("should modify existing elements of kube-controller-manager deployment", func() {
			dep = &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: v1beta1constants.DeploymentNameKubeControllerManager},
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								v1beta1constants.LabelNetworkPolicyToBlockedCIDRs: v1beta1constants.LabelNetworkPolicyAllowed,
							},
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name: "kube-controller-manager",
									Command: []string{
										"--cloud-provider=?",
										"--cloud-config=?",
									},
								},
							},
						},
					},
				},
			}

			err := ensurer.EnsureKubeControllerManagerDeployment(ctx, eContextK8s, dep, nil)
			Expect(err).To(Not(HaveOccurred()))

			checkKubeControllerManagerDeployment(dep)
		})
	})

	Describe("#EnsureKubeletServiceUnitOptions", func() {
		var (
			oldUnitOptions        []*unit.UnitOption
			hostnamectlUnitOption *unit.UnitOption
		)

		BeforeEach(func() {
			oldUnitOptions = []*unit.UnitOption{
				{
					Section: "Service",
					Name:    "ExecStart",
					Value: `/opt/bin/hyperkube kubelet \	
	    --config=/var/lib/kubelet/config/kubelet`,
				},
			}

			hostnamectlUnitOption = &unit.UnitOption{
				Section: "Service",
				Name:    "ExecStartPre",
				Value:   `/bin/sh -c 'hostnamectl set-hostname $(hostname -f)'`,
			}
		})

		It("should modify existing elements of kubelet.service unit options",
			func() {
				newUnitOptions := []*unit.UnitOption{
					{
						Section: "Service",
						Name:    "ExecStart",
						Value:   "/opt/bin/hyperkube kubelet \\\n    --config=/var/lib/kubelet/config/kubelet \\\n    --cloud-provider=external",
					},
					hostnamectlUnitOption,
				}

				opts, err := ensurer.EnsureKubeletServiceUnitOptions(ctx, dummyContext, semver.MustParse("1.23.0"), oldUnitOptions, nil)
				Expect(err).To(Not(HaveOccurred()))
				Expect(opts).To(Equal(newUnitOptions))
			},
		)
	})

	Describe("#EnsureMachineControllerManagerDeployment", func() {
		var (
			ensurer    genericmutator.Ensurer
			deployment *appsv1.Deployment
		)

		BeforeEach(func() {
			deployment = &appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{Namespace: "foo"}}
		})

		Context("when gardenlet manages MCM", func() {
			BeforeEach(func() {
				ensurer = NewEnsurer(logger, true)
				DeferCleanup(testutils.WithVar(&ImageVector, imagevectorutils.ImageVector{{
					Name:       "machine-controller-manager-provider-ironcore",
					Repository: ptr.To("foo"),
					Tag:        ptr.To[string]("bar"),
				}}))
			})

			It("should inject the sidecar container", func() {
				Expect(deployment.Spec.Template.Spec.Containers).To(BeEmpty())
				Expect(ensurer.EnsureMachineControllerManagerDeployment(ctx, nil, deployment, nil)).To(Succeed())
				Expect(deployment.Spec.Template.Spec.Containers).To(ConsistOf(corev1.Container{
					Name:            "machine-controller-manager-provider-ironcore",
					Image:           "foo:bar",
					ImagePullPolicy: corev1.PullIfNotPresent,
					Args: []string{
						"--control-kubeconfig=inClusterConfig",
						"--machine-creation-timeout=20m",
						"--machine-drain-timeout=2h",
						"--machine-health-timeout=10m",
						"--machine-safety-apiserver-statuscheck-timeout=30s",
						"--machine-safety-apiserver-statuscheck-period=1m",
						"--machine-safety-orphan-vms-period=30m",
						"--namespace=" + namespace,
						"--port=" + strconv.Itoa(portProviderMetrics),
						"--target-kubeconfig=" + gardenerutils.PathGenericKubeconfig,
						"--v=3",
						"--ironcore-kubeconfig=/etc/ironcore/kubeconfig",
					},
					LivenessProbe: &corev1.Probe{
						ProbeHandler: corev1.ProbeHandler{
							HTTPGet: &corev1.HTTPGetAction{
								Path:   "/healthz",
								Port:   intstr.FromInt32(portProviderMetrics),
								Scheme: corev1.URISchemeHTTP,
							},
						},
						InitialDelaySeconds: 30,
						TimeoutSeconds:      5,
						PeriodSeconds:       10,
						SuccessThreshold:    1,
						FailureThreshold:    3,
					},
					Ports: []corev1.ContainerPort{{
						Name:          portNameProviderMetrics,
						ContainerPort: portProviderMetrics,
						Protocol:      corev1.ProtocolTCP,
					}},
					VolumeMounts: []corev1.VolumeMount{{
						Name:      "kubeconfig",
						MountPath: gardenerutils.VolumeMountPathGenericKubeconfig,
						ReadOnly:  true,
					},
						{
							Name:      "cloudprovider",
							MountPath: "/etc/ironcore",
							ReadOnly:  true,
						},
					},
				}))
				Expect(deployment.Spec.Template.Spec.Volumes).To(ContainElement(corev1.Volume{
					Name: "cloudprovider",
					VolumeSource: corev1.VolumeSource{
						Secret: &corev1.SecretVolumeSource{
							SecretName: "cloudprovider",
						},
					},
				}))
			})
		})
	})

	Describe("#EnsureMachineControllerManagerVPA", func() {
		var (
			ensurer genericmutator.Ensurer
			vpa     *vpaautoscalingv1.VerticalPodAutoscaler
		)

		BeforeEach(func() {
			vpa = &vpaautoscalingv1.VerticalPodAutoscaler{}
		})

		Context("when gardenlet manages MCM", func() {
			BeforeEach(func() {
				ensurer = NewEnsurer(logger, true)
			})

			It("should inject the sidecar container policy", func() {
				Expect(vpa.Spec.ResourcePolicy).To(BeNil())
				Expect(ensurer.EnsureMachineControllerManagerVPA(ctx, nil, vpa, nil)).To(Succeed())

				ccv := vpaautoscalingv1.ContainerControlledValuesRequestsOnly
				Expect(vpa.Spec.ResourcePolicy.ContainerPolicies).To(ConsistOf(vpaautoscalingv1.ContainerResourcePolicy{
					ContainerName:    "machine-controller-manager-provider-ironcore",
					ControlledValues: &ccv,
					MinAllowed: corev1.ResourceList{
						corev1.ResourceCPU:    resource.MustParse("30m"),
						corev1.ResourceMemory: resource.MustParse("64Mi"),
					},
					MaxAllowed: corev1.ResourceList{
						corev1.ResourceCPU:    resource.MustParse("2"),
						corev1.ResourceMemory: resource.MustParse("5G"),
					},
				}))
			})
		})
	})
})

func checkKubeAPIServerDeployment(dep *appsv1.Deployment) {
	// Check that the kube-apiserver container still exists and contains all needed command line args,
	c := extensionswebhook.ContainerWithName(dep.Spec.Template.Spec.Containers, "kube-apiserver")
	Expect(c).To(Not(BeNil()))

	Expect(c.Command).NotTo(ContainElement("--cloud-provider=onemtal"))
	Expect(c.Command).NotTo(ContainElement("--cloud-config=/etc/kubernetes/cloudprovider/cloudprovider.conf"))
}

func checkKubeControllerManagerDeployment(dep *appsv1.Deployment) {
	// Check that the kube-controller-manager container still exists and contains all needed command line args,
	c := extensionswebhook.ContainerWithName(dep.Spec.Template.Spec.Containers, "kube-controller-manager")
	Expect(c).To(Not(BeNil()))

	Expect(c.Command).To(ContainElement("--cloud-provider=external"))
	Expect(c.Command).NotTo(ContainElement("--cloud-config=/etc/kubernetes/cloudprovider/cloudprovider.conf"))
}
