// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package controlplane

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/gardener/gardener/pkg/apis/core/v1beta1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/ironcore-dev/controller-utils/buildutils"
	"github.com/ironcore-dev/controller-utils/modutils"
	storagev1alpha1 "github.com/ironcore-dev/ironcore/api/storage/v1alpha1"
	utilsenvtest "github.com/ironcore-dev/ironcore/utils/envtest"
	"github.com/ironcore-dev/ironcore/utils/envtest/apiserver"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/config"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/komega"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	ironcoreextensionv1alpha1 "github.com/ironcore-dev/gardener-extension-provider-ironcore/pkg/apis/ironcore/v1alpha1"
)

const (
	pollingInterval      = 50 * time.Millisecond
	eventuallyTimeout    = 10 * time.Second
	consistentlyDuration = 1 * time.Second
	apiServiceTimeout    = 5 * time.Minute
)

var (
	testEnv    *envtest.Environment
	testEnvExt *utilsenvtest.EnvironmentExtensions
	cfg        *rest.Config
	k8sClient  client.Client
)

func TestAPIs(t *testing.T) {
	SetDefaultConsistentlyPollingInterval(pollingInterval)
	SetDefaultEventuallyPollingInterval(pollingInterval)
	SetDefaultEventuallyTimeout(eventuallyTimeout)
	SetDefaultConsistentlyDuration(consistentlyDuration)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Infrastructure Controller Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	var err error

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "..", "example", "20-crd-resources.gardener.cloud_managedresources.yaml"),
			filepath.Join("..", "..", "..", "example", "20-crd-extensions.gardener.cloud_clusters.yaml"),
			filepath.Join("..", "..", "..", "example", "20-crd-extensions.gardener.cloud_infrastructures.yaml"),
			filepath.Join("..", "..", "..", "example", "20-crd-extensions.gardener.cloud_controlplanes.yaml"),
		},
		ErrorIfCRDPathMissing: true,

		// The BinaryAssetsDirectory is only required if you want to run the tests directly
		// without call the makefile target test. If not informed it will look for the
		// default path defined in controller-apiruntime which is /usr/local/kubebuilder/.
		// Note that you must have the required binaries setup under the bin directory to perform
		// the tests directly. When we run make test it will be setup and used automatically.
		BinaryAssetsDirectory: filepath.Join("..", "..", "bin", "k8s",
			fmt.Sprintf("1.32.0-%s-%s", runtime.GOOS, runtime.GOARCH)),
	}

	testEnvExt = &utilsenvtest.EnvironmentExtensions{
		APIServiceDirectoryPaths: []string{
			modutils.Dir("github.com/ironcore-dev/ironcore", "config", "apiserver", "apiservice", "bases"),
		},
		ErrorIfAPIServicePathIsMissing: true,
	}

	cfg, err = utilsenvtest.StartWithExtensions(testEnv, testEnvExt)
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())
	DeferCleanup(utilsenvtest.StopWithExtensions, testEnv, testEnvExt)

	Expect(extensionsv1alpha1.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(ironcoreextensionv1alpha1.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(corev1.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(storagev1alpha1.AddToScheme(scheme.Scheme)).To(Succeed())

	// Init package-level k8sClient
	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())

	komega.SetClient(k8sClient)

	apiSrv, err := apiserver.New(cfg, apiserver.Options{
		MainPath:     "github.com/ironcore-dev/ironcore/cmd/ironcore-apiserver",
		BuildOptions: []buildutils.BuildOption{buildutils.ModModeMod},
		ETCDServers:  []string{testEnv.ControlPlane.Etcd.URL.String()},
		Host:         testEnvExt.APIServiceInstallOptions.LocalServingHost,
		Port:         testEnvExt.APIServiceInstallOptions.LocalServingPort,
		CertDir:      testEnvExt.APIServiceInstallOptions.LocalServingCertDir,
	})
	Expect(err).NotTo(HaveOccurred())

	Expect(apiSrv.Start()).To(Succeed())
	DeferCleanup(apiSrv.Stop)

	err = utilsenvtest.WaitUntilAPIServicesReadyWithTimeout(apiServiceTimeout, testEnvExt, cfg, k8sClient, scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())
})

func SetupTest() (*corev1.Namespace, *valuesProvider, *extensionsv1alpha1.Cluster) {
	var (
		namespace = &corev1.Namespace{}
		cluster   = &extensionsv1alpha1.Cluster{}
		vp        = &valuesProvider{}
	)

	BeforeEach(func(ctx SpecContext) {
		var mgrCtx context.Context
		mgrCtx, cancel := context.WithCancel(context.Background())
		DeferCleanup(cancel)

		*namespace = corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "testns-",
			},
		}
		Expect(k8sClient.Create(ctx, namespace)).To(Succeed(), "failed to create test namespace")
		DeferCleanup(k8sClient.Delete, namespace)

		shoot := v1beta1.Shoot{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace.Name,
				Name:      "foo",
			},
			Spec: v1beta1.ShootSpec{
				Kubernetes: v1beta1.Kubernetes{
					Version: "1.26.0",
				},
				Region: "foo",
				Provider: v1beta1.Provider{
					Workers: []v1beta1.Worker{
						{Name: "foo"},
						{Name: "bar"},
					},
				},
				Networking: &v1beta1.Networking{
					Nodes: ptr.To[string]("10.0.0.0/24"),
				},
			},
		}
		shootJson, err := json.Marshal(shoot)
		Expect(err).NotTo(HaveOccurred())

		*cluster = extensionsv1alpha1.Cluster{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace.Name,
			},
			Spec: extensionsv1alpha1.ClusterSpec{
				CloudProfile: apiruntime.RawExtension{Raw: []byte("{}")},
				Seed:         apiruntime.RawExtension{Raw: []byte("{}")},
				Shoot:        apiruntime.RawExtension{Raw: shootJson},
			},
		}
		Expect(k8sClient.Create(ctx, cluster)).Should(Succeed())
		DeferCleanup(k8sClient.Delete, cluster)

		mgr, err := manager.New(cfg, manager.Options{
			Scheme:  scheme.Scheme,
			Metrics: metricsserver.Options{BindAddress: "0"},
			WebhookServer: &webhook.DefaultServer{Options: webhook.Options{
				CertDir: testEnv.WebhookInstallOptions.LocalServingCertDir,
			}},
			Controller: config.Controller{
				// need to skip unique controller name validation
				// since all tests need a dedicated controller
				SkipNameValidation: ptr.To(true),
			},
		})
		Expect(err).NotTo(HaveOccurred())

		secret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace.Name,
				Name:      "my-infra-creds",
			},
			Data: map[string][]byte{
				"namespace": []byte(namespace.Name),
				"token":     []byte("foo"),
			},
		}
		Expect(k8sClient.Create(ctx, secret)).To(Succeed())

		user, err := testEnv.AddUser(envtest.User{
			Name:   "dummy",
			Groups: []string{"system:authenticated", "system:masters"},
		}, cfg)
		Expect(err).NotTo(HaveOccurred())

		kubeconfig, err := user.KubeConfig()
		Expect(err).NotTo(HaveOccurred())

		By("creating a test cloudprovider secret")
		cloudproviderSecret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace.Name,
				Name:      "cloudprovider",
			},
			Data: map[string][]byte{
				"namespace":  []byte(namespace.Name),
				"token":      []byte("foo"),
				"kubeconfig": kubeconfig,
			},
		}
		Expect(k8sClient.Create(ctx, cloudproviderSecret)).To(Succeed())
		DeferCleanup(k8sClient.Delete, cloudproviderSecret)

		Expect(AddToManagerWithOptions(mgrCtx, mgr, AddOptions{
			IgnoreOperationAnnotation: true,
		})).NotTo(HaveOccurred())

		*vp = valuesProvider{
			client:  mgr.GetClient(),
			decoder: serializer.NewCodecFactory(mgr.GetScheme(), serializer.EnableStrict).UniversalDecoder(),
		}

		go func() {
			defer GinkgoRecover()
			Expect(mgr.Start(mgrCtx)).To(Succeed(), "failed to start manager")
		}()
	})

	return namespace, vp, cluster
}
