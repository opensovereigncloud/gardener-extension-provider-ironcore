//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	ironcore "github.com/ironcore-dev/gardener-extension-provider-ironcore/pkg/apis/ironcore"
	commonv1alpha1 "github.com/ironcore-dev/ironcore/api/common/v1alpha1"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CloudControllerManagerConfig)(nil), (*ironcore.CloudControllerManagerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudControllerManagerConfig_To_ironcore_CloudControllerManagerConfig(a.(*CloudControllerManagerConfig), b.(*ironcore.CloudControllerManagerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.CloudControllerManagerConfig)(nil), (*CloudControllerManagerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(a.(*ironcore.CloudControllerManagerConfig), b.(*CloudControllerManagerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CloudProfileConfig)(nil), (*ironcore.CloudProfileConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudProfileConfig_To_ironcore_CloudProfileConfig(a.(*CloudProfileConfig), b.(*ironcore.CloudProfileConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.CloudProfileConfig)(nil), (*CloudProfileConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(a.(*ironcore.CloudProfileConfig), b.(*CloudProfileConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlaneConfig)(nil), (*ironcore.ControlPlaneConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControlPlaneConfig_To_ironcore_ControlPlaneConfig(a.(*ControlPlaneConfig), b.(*ironcore.ControlPlaneConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.ControlPlaneConfig)(nil), (*ControlPlaneConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(a.(*ironcore.ControlPlaneConfig), b.(*ControlPlaneConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureConfig)(nil), (*ironcore.InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureConfig_To_ironcore_InfrastructureConfig(a.(*InfrastructureConfig), b.(*ironcore.InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.InfrastructureConfig)(nil), (*InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(a.(*ironcore.InfrastructureConfig), b.(*InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureStatus)(nil), (*ironcore.InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureStatus_To_ironcore_InfrastructureStatus(a.(*InfrastructureStatus), b.(*ironcore.InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.InfrastructureStatus)(nil), (*InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(a.(*ironcore.InfrastructureStatus), b.(*InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachineImage)(nil), (*ironcore.MachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MachineImage_To_ironcore_MachineImage(a.(*MachineImage), b.(*ironcore.MachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.MachineImage)(nil), (*MachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_MachineImage_To_v1alpha1_MachineImage(a.(*ironcore.MachineImage), b.(*MachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachineImageVersion)(nil), (*ironcore.MachineImageVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MachineImageVersion_To_ironcore_MachineImageVersion(a.(*MachineImageVersion), b.(*ironcore.MachineImageVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.MachineImageVersion)(nil), (*MachineImageVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_MachineImageVersion_To_v1alpha1_MachineImageVersion(a.(*ironcore.MachineImageVersion), b.(*MachineImageVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachineImages)(nil), (*ironcore.MachineImages)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MachineImages_To_ironcore_MachineImages(a.(*MachineImages), b.(*ironcore.MachineImages), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.MachineImages)(nil), (*MachineImages)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_MachineImages_To_v1alpha1_MachineImages(a.(*ironcore.MachineImages), b.(*MachineImages), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RegionConfig)(nil), (*ironcore.RegionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RegionConfig_To_ironcore_RegionConfig(a.(*RegionConfig), b.(*ironcore.RegionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.RegionConfig)(nil), (*RegionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_RegionConfig_To_v1alpha1_RegionConfig(a.(*ironcore.RegionConfig), b.(*RegionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*StorageClass)(nil), (*ironcore.StorageClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_StorageClass_To_ironcore_StorageClass(a.(*StorageClass), b.(*ironcore.StorageClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.StorageClass)(nil), (*StorageClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_StorageClass_To_v1alpha1_StorageClass(a.(*ironcore.StorageClass), b.(*StorageClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*StorageClasses)(nil), (*ironcore.StorageClasses)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_StorageClasses_To_ironcore_StorageClasses(a.(*StorageClasses), b.(*ironcore.StorageClasses), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.StorageClasses)(nil), (*StorageClasses)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_StorageClasses_To_v1alpha1_StorageClasses(a.(*ironcore.StorageClasses), b.(*StorageClasses), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*WorkerStatus)(nil), (*ironcore.WorkerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WorkerStatus_To_ironcore_WorkerStatus(a.(*WorkerStatus), b.(*ironcore.WorkerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ironcore.WorkerStatus)(nil), (*WorkerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ironcore_WorkerStatus_To_v1alpha1_WorkerStatus(a.(*ironcore.WorkerStatus), b.(*WorkerStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CloudControllerManagerConfig_To_ironcore_CloudControllerManagerConfig(in *CloudControllerManagerConfig, out *ironcore.CloudControllerManagerConfig, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_v1alpha1_CloudControllerManagerConfig_To_ironcore_CloudControllerManagerConfig is an autogenerated conversion function.
func Convert_v1alpha1_CloudControllerManagerConfig_To_ironcore_CloudControllerManagerConfig(in *CloudControllerManagerConfig, out *ironcore.CloudControllerManagerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudControllerManagerConfig_To_ironcore_CloudControllerManagerConfig(in, out, s)
}

func autoConvert_ironcore_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in *ironcore.CloudControllerManagerConfig, out *CloudControllerManagerConfig, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_ironcore_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig is an autogenerated conversion function.
func Convert_ironcore_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in *ironcore.CloudControllerManagerConfig, out *CloudControllerManagerConfig, s conversion.Scope) error {
	return autoConvert_ironcore_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in, out, s)
}

func autoConvert_v1alpha1_CloudProfileConfig_To_ironcore_CloudProfileConfig(in *CloudProfileConfig, out *ironcore.CloudProfileConfig, s conversion.Scope) error {
	out.MachineImages = *(*[]ironcore.MachineImages)(unsafe.Pointer(&in.MachineImages))
	out.RegionConfigs = *(*[]ironcore.RegionConfig)(unsafe.Pointer(&in.RegionConfigs))
	if err := Convert_v1alpha1_StorageClasses_To_ironcore_StorageClasses(&in.StorageClasses, &out.StorageClasses, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_CloudProfileConfig_To_ironcore_CloudProfileConfig is an autogenerated conversion function.
func Convert_v1alpha1_CloudProfileConfig_To_ironcore_CloudProfileConfig(in *CloudProfileConfig, out *ironcore.CloudProfileConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudProfileConfig_To_ironcore_CloudProfileConfig(in, out, s)
}

func autoConvert_ironcore_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in *ironcore.CloudProfileConfig, out *CloudProfileConfig, s conversion.Scope) error {
	out.MachineImages = *(*[]MachineImages)(unsafe.Pointer(&in.MachineImages))
	out.RegionConfigs = *(*[]RegionConfig)(unsafe.Pointer(&in.RegionConfigs))
	if err := Convert_ironcore_StorageClasses_To_v1alpha1_StorageClasses(&in.StorageClasses, &out.StorageClasses, s); err != nil {
		return err
	}
	return nil
}

// Convert_ironcore_CloudProfileConfig_To_v1alpha1_CloudProfileConfig is an autogenerated conversion function.
func Convert_ironcore_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in *ironcore.CloudProfileConfig, out *CloudProfileConfig, s conversion.Scope) error {
	return autoConvert_ironcore_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in, out, s)
}

func autoConvert_v1alpha1_ControlPlaneConfig_To_ironcore_ControlPlaneConfig(in *ControlPlaneConfig, out *ironcore.ControlPlaneConfig, s conversion.Scope) error {
	out.CloudControllerManager = (*ironcore.CloudControllerManagerConfig)(unsafe.Pointer(in.CloudControllerManager))
	return nil
}

// Convert_v1alpha1_ControlPlaneConfig_To_ironcore_ControlPlaneConfig is an autogenerated conversion function.
func Convert_v1alpha1_ControlPlaneConfig_To_ironcore_ControlPlaneConfig(in *ControlPlaneConfig, out *ironcore.ControlPlaneConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControlPlaneConfig_To_ironcore_ControlPlaneConfig(in, out, s)
}

func autoConvert_ironcore_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in *ironcore.ControlPlaneConfig, out *ControlPlaneConfig, s conversion.Scope) error {
	out.CloudControllerManager = (*CloudControllerManagerConfig)(unsafe.Pointer(in.CloudControllerManager))
	return nil
}

// Convert_ironcore_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig is an autogenerated conversion function.
func Convert_ironcore_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in *ironcore.ControlPlaneConfig, out *ControlPlaneConfig, s conversion.Scope) error {
	return autoConvert_ironcore_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureConfig_To_ironcore_InfrastructureConfig(in *InfrastructureConfig, out *ironcore.InfrastructureConfig, s conversion.Scope) error {
	out.NetworkRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.NetworkRef))
	out.NATPortsPerNetworkInterface = (*int32)(unsafe.Pointer(in.NATPortsPerNetworkInterface))
	out.NetworkPolicyRef = (*commonv1alpha1.LocalUIDReference)(unsafe.Pointer(in.NetworkPolicyRef))
	return nil
}

// Convert_v1alpha1_InfrastructureConfig_To_ironcore_InfrastructureConfig is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureConfig_To_ironcore_InfrastructureConfig(in *InfrastructureConfig, out *ironcore.InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureConfig_To_ironcore_InfrastructureConfig(in, out, s)
}

func autoConvert_ironcore_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *ironcore.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	out.NetworkRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.NetworkRef))
	out.NATPortsPerNetworkInterface = (*int32)(unsafe.Pointer(in.NATPortsPerNetworkInterface))
	out.NetworkPolicyRef = (*commonv1alpha1.LocalUIDReference)(unsafe.Pointer(in.NetworkPolicyRef))
	return nil
}

// Convert_ironcore_InfrastructureConfig_To_v1alpha1_InfrastructureConfig is an autogenerated conversion function.
func Convert_ironcore_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *ironcore.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_ironcore_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureStatus_To_ironcore_InfrastructureStatus(in *InfrastructureStatus, out *ironcore.InfrastructureStatus, s conversion.Scope) error {
	out.NetworkRef = in.NetworkRef
	out.NATGatewayRef = in.NATGatewayRef
	out.PrefixRef = in.PrefixRef
	out.NetworkPolicyRef = in.NetworkPolicyRef
	return nil
}

// Convert_v1alpha1_InfrastructureStatus_To_ironcore_InfrastructureStatus is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureStatus_To_ironcore_InfrastructureStatus(in *InfrastructureStatus, out *ironcore.InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureStatus_To_ironcore_InfrastructureStatus(in, out, s)
}

func autoConvert_ironcore_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *ironcore.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	out.NetworkRef = in.NetworkRef
	out.NATGatewayRef = in.NATGatewayRef
	out.PrefixRef = in.PrefixRef
	out.NetworkPolicyRef = in.NetworkPolicyRef
	return nil
}

// Convert_ironcore_InfrastructureStatus_To_v1alpha1_InfrastructureStatus is an autogenerated conversion function.
func Convert_ironcore_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *ironcore.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_ironcore_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in, out, s)
}

func autoConvert_v1alpha1_MachineImage_To_ironcore_MachineImage(in *MachineImage, out *ironcore.MachineImage, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.Image = in.Image
	out.Architecture = (*string)(unsafe.Pointer(in.Architecture))
	return nil
}

// Convert_v1alpha1_MachineImage_To_ironcore_MachineImage is an autogenerated conversion function.
func Convert_v1alpha1_MachineImage_To_ironcore_MachineImage(in *MachineImage, out *ironcore.MachineImage, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineImage_To_ironcore_MachineImage(in, out, s)
}

func autoConvert_ironcore_MachineImage_To_v1alpha1_MachineImage(in *ironcore.MachineImage, out *MachineImage, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.Image = in.Image
	out.Architecture = (*string)(unsafe.Pointer(in.Architecture))
	return nil
}

// Convert_ironcore_MachineImage_To_v1alpha1_MachineImage is an autogenerated conversion function.
func Convert_ironcore_MachineImage_To_v1alpha1_MachineImage(in *ironcore.MachineImage, out *MachineImage, s conversion.Scope) error {
	return autoConvert_ironcore_MachineImage_To_v1alpha1_MachineImage(in, out, s)
}

func autoConvert_v1alpha1_MachineImageVersion_To_ironcore_MachineImageVersion(in *MachineImageVersion, out *ironcore.MachineImageVersion, s conversion.Scope) error {
	out.Version = in.Version
	out.Image = in.Image
	out.Architecture = (*string)(unsafe.Pointer(in.Architecture))
	return nil
}

// Convert_v1alpha1_MachineImageVersion_To_ironcore_MachineImageVersion is an autogenerated conversion function.
func Convert_v1alpha1_MachineImageVersion_To_ironcore_MachineImageVersion(in *MachineImageVersion, out *ironcore.MachineImageVersion, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineImageVersion_To_ironcore_MachineImageVersion(in, out, s)
}

func autoConvert_ironcore_MachineImageVersion_To_v1alpha1_MachineImageVersion(in *ironcore.MachineImageVersion, out *MachineImageVersion, s conversion.Scope) error {
	out.Version = in.Version
	out.Image = in.Image
	out.Architecture = (*string)(unsafe.Pointer(in.Architecture))
	return nil
}

// Convert_ironcore_MachineImageVersion_To_v1alpha1_MachineImageVersion is an autogenerated conversion function.
func Convert_ironcore_MachineImageVersion_To_v1alpha1_MachineImageVersion(in *ironcore.MachineImageVersion, out *MachineImageVersion, s conversion.Scope) error {
	return autoConvert_ironcore_MachineImageVersion_To_v1alpha1_MachineImageVersion(in, out, s)
}

func autoConvert_v1alpha1_MachineImages_To_ironcore_MachineImages(in *MachineImages, out *ironcore.MachineImages, s conversion.Scope) error {
	out.Name = in.Name
	out.Versions = *(*[]ironcore.MachineImageVersion)(unsafe.Pointer(&in.Versions))
	return nil
}

// Convert_v1alpha1_MachineImages_To_ironcore_MachineImages is an autogenerated conversion function.
func Convert_v1alpha1_MachineImages_To_ironcore_MachineImages(in *MachineImages, out *ironcore.MachineImages, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineImages_To_ironcore_MachineImages(in, out, s)
}

func autoConvert_ironcore_MachineImages_To_v1alpha1_MachineImages(in *ironcore.MachineImages, out *MachineImages, s conversion.Scope) error {
	out.Name = in.Name
	out.Versions = *(*[]MachineImageVersion)(unsafe.Pointer(&in.Versions))
	return nil
}

// Convert_ironcore_MachineImages_To_v1alpha1_MachineImages is an autogenerated conversion function.
func Convert_ironcore_MachineImages_To_v1alpha1_MachineImages(in *ironcore.MachineImages, out *MachineImages, s conversion.Scope) error {
	return autoConvert_ironcore_MachineImages_To_v1alpha1_MachineImages(in, out, s)
}

func autoConvert_v1alpha1_RegionConfig_To_ironcore_RegionConfig(in *RegionConfig, out *ironcore.RegionConfig, s conversion.Scope) error {
	out.Name = in.Name
	out.Server = in.Server
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	return nil
}

// Convert_v1alpha1_RegionConfig_To_ironcore_RegionConfig is an autogenerated conversion function.
func Convert_v1alpha1_RegionConfig_To_ironcore_RegionConfig(in *RegionConfig, out *ironcore.RegionConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_RegionConfig_To_ironcore_RegionConfig(in, out, s)
}

func autoConvert_ironcore_RegionConfig_To_v1alpha1_RegionConfig(in *ironcore.RegionConfig, out *RegionConfig, s conversion.Scope) error {
	out.Name = in.Name
	out.Server = in.Server
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	return nil
}

// Convert_ironcore_RegionConfig_To_v1alpha1_RegionConfig is an autogenerated conversion function.
func Convert_ironcore_RegionConfig_To_v1alpha1_RegionConfig(in *ironcore.RegionConfig, out *RegionConfig, s conversion.Scope) error {
	return autoConvert_ironcore_RegionConfig_To_v1alpha1_RegionConfig(in, out, s)
}

func autoConvert_v1alpha1_StorageClass_To_ironcore_StorageClass(in *StorageClass, out *ironcore.StorageClass, s conversion.Scope) error {
	out.Name = in.Name
	out.Type = in.Type
	return nil
}

// Convert_v1alpha1_StorageClass_To_ironcore_StorageClass is an autogenerated conversion function.
func Convert_v1alpha1_StorageClass_To_ironcore_StorageClass(in *StorageClass, out *ironcore.StorageClass, s conversion.Scope) error {
	return autoConvert_v1alpha1_StorageClass_To_ironcore_StorageClass(in, out, s)
}

func autoConvert_ironcore_StorageClass_To_v1alpha1_StorageClass(in *ironcore.StorageClass, out *StorageClass, s conversion.Scope) error {
	out.Name = in.Name
	out.Type = in.Type
	return nil
}

// Convert_ironcore_StorageClass_To_v1alpha1_StorageClass is an autogenerated conversion function.
func Convert_ironcore_StorageClass_To_v1alpha1_StorageClass(in *ironcore.StorageClass, out *StorageClass, s conversion.Scope) error {
	return autoConvert_ironcore_StorageClass_To_v1alpha1_StorageClass(in, out, s)
}

func autoConvert_v1alpha1_StorageClasses_To_ironcore_StorageClasses(in *StorageClasses, out *ironcore.StorageClasses, s conversion.Scope) error {
	out.Default = (*ironcore.StorageClass)(unsafe.Pointer(in.Default))
	out.Additional = *(*[]ironcore.StorageClass)(unsafe.Pointer(&in.Additional))
	return nil
}

// Convert_v1alpha1_StorageClasses_To_ironcore_StorageClasses is an autogenerated conversion function.
func Convert_v1alpha1_StorageClasses_To_ironcore_StorageClasses(in *StorageClasses, out *ironcore.StorageClasses, s conversion.Scope) error {
	return autoConvert_v1alpha1_StorageClasses_To_ironcore_StorageClasses(in, out, s)
}

func autoConvert_ironcore_StorageClasses_To_v1alpha1_StorageClasses(in *ironcore.StorageClasses, out *StorageClasses, s conversion.Scope) error {
	out.Default = (*StorageClass)(unsafe.Pointer(in.Default))
	out.Additional = *(*[]StorageClass)(unsafe.Pointer(&in.Additional))
	return nil
}

// Convert_ironcore_StorageClasses_To_v1alpha1_StorageClasses is an autogenerated conversion function.
func Convert_ironcore_StorageClasses_To_v1alpha1_StorageClasses(in *ironcore.StorageClasses, out *StorageClasses, s conversion.Scope) error {
	return autoConvert_ironcore_StorageClasses_To_v1alpha1_StorageClasses(in, out, s)
}

func autoConvert_v1alpha1_WorkerStatus_To_ironcore_WorkerStatus(in *WorkerStatus, out *ironcore.WorkerStatus, s conversion.Scope) error {
	out.MachineImages = *(*[]ironcore.MachineImage)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_v1alpha1_WorkerStatus_To_ironcore_WorkerStatus is an autogenerated conversion function.
func Convert_v1alpha1_WorkerStatus_To_ironcore_WorkerStatus(in *WorkerStatus, out *ironcore.WorkerStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_WorkerStatus_To_ironcore_WorkerStatus(in, out, s)
}

func autoConvert_ironcore_WorkerStatus_To_v1alpha1_WorkerStatus(in *ironcore.WorkerStatus, out *WorkerStatus, s conversion.Scope) error {
	out.MachineImages = *(*[]MachineImage)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_ironcore_WorkerStatus_To_v1alpha1_WorkerStatus is an autogenerated conversion function.
func Convert_ironcore_WorkerStatus_To_v1alpha1_WorkerStatus(in *ironcore.WorkerStatus, out *WorkerStatus, s conversion.Scope) error {
	return autoConvert_ironcore_WorkerStatus_To_v1alpha1_WorkerStatus(in, out, s)
}
