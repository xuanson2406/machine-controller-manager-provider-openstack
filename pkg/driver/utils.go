// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"errors"
	"fmt"

	mcmv1alpha1 "github.com/xuanson2406/machine-controller-manager/pkg/apis/machine/v1alpha1"
	"github.com/xuanson2406/machine-controller-manager/pkg/util/provider/machinecodes/codes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	"github.com/gardener/machine-controller-manager-provider-openstack/pkg/apis/openstack"
	"github.com/gardener/machine-controller-manager-provider-openstack/pkg/apis/openstack/install"
	"github.com/gardener/machine-controller-manager-provider-openstack/pkg/apis/openstack/v1alpha1"
	client "github.com/gardener/machine-controller-manager-provider-openstack/pkg/client"
	"github.com/gardener/machine-controller-manager-provider-openstack/pkg/driver/executor"
)

const (
	openstackProvider = "OpenStack"
)

// Decoder is a decoder for a scheme containing the mcm-openstack APIs.
var Decoder = serializer.NewCodecFactory(install.Install(runtime.NewScheme())).UniversalDecoder()

// DecodeProviderSpec can decode raw to a MachineProviderConfig.
func DecodeProviderSpec(decoder runtime.Decoder, raw runtime.RawExtension) (*openstack.MachineProviderConfig, error) {
	json, err := raw.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to decode provider spec: %v", err)
	}

	cfg := &openstack.MachineProviderConfig{}
	_, _, err = decoder.Decode(json, nil, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode provider spec: %v", err)
	}

	return cfg, nil
}

func (p *OpenstackDriver) decodeProviderSpec(raw runtime.RawExtension) (*openstack.MachineProviderConfig, error) {
	return DecodeProviderSpec(p.decoder, raw)
}

func migrateMachineClass(os *mcmv1alpha1.OpenStackMachineClass, machineClass *mcmv1alpha1.MachineClass) error {
	migratedNetworks := []v1alpha1.OpenStackNetwork{}
	for _, nw := range os.Spec.Networks {
		migratedNetworks = append(migratedNetworks, v1alpha1.OpenStackNetwork{
			Id:         nw.Id,
			Name:       nw.Name,
			PodNetwork: nw.PodNetwork,
		})
	}

	cfg := &v1alpha1.MachineProviderConfig{
		TypeMeta: metav1.TypeMeta{
			Kind:       "MachineProviderConfig",
			APIVersion: v1alpha1.SchemeGroupVersion.String(),
		},
		Spec: v1alpha1.MachineProviderConfigSpec{
			ImageID:          os.Spec.ImageID,
			ImageName:        os.Spec.ImageName,
			Region:           os.Spec.Region,
			AvailabilityZone: os.Spec.AvailabilityZone,
			FlavorName:       os.Spec.FlavorName,
			KeyName:          os.Spec.KeyName,
			SecurityGroups:   os.Spec.SecurityGroups,
			Tags:             os.Spec.Tags,
			NetworkID:        os.Spec.NetworkID,
			SubnetID:         os.Spec.SubnetID,
			PodNetworkCidr:   os.Spec.PodNetworkCidr,
			RootDiskSize:     os.Spec.RootDiskSize,
			UseConfigDrive:   os.Spec.UseConfigDrive,
			ServerGroupID:    os.Spec.ServerGroupID,
			Networks:         migratedNetworks,
		},
	}

	machineClass.Name = os.Name
	machineClass.Labels = os.Labels
	machineClass.Annotations = os.Annotations

	machineClass.Finalizers = os.Finalizers
	machineClass.ProviderSpec = runtime.RawExtension{
		Object: cfg,
	}
	machineClass.SecretRef = os.Spec.SecretRef
	machineClass.CredentialsSecretRef = os.Spec.CredentialsSecretRef
	machineClass.Provider = openstackProvider

	return nil
}

func mapErrorToCode(err error) codes.Code {
	if errors.Is(err, executor.ErrNotFound) {
		return codes.NotFound
	}

	if errors.Is(err, executor.ErrMultipleFound) {
		return codes.OutOfRange
	}

	if client.IsUnauthenticated(err) {
		return codes.Unauthenticated
	}

	if client.IsUnauthorized(err) {
		return codes.PermissionDenied
	}

	return codes.Internal
}
