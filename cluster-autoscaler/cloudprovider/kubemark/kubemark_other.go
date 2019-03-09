// +build !linux

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Dummy implementation. Real one should be built on linux.

package kubemark

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
)

const (
	// ProviderName is the cloud provider name for kubemark
	ProviderName = "kubemark"

	// GPULabel is the label added to nodes with GPU resource.
	GPULabel = "cloud.google.com/gke-accelerator"
)

var (
	availableGPUTypes = map[string]struct{}{
		"nvidia-tesla-k80":  {},
		"nvidia-tesla-p100": {},
		"nvidia-tesla-v100": {},
	}
)

type KubemarkCloudProvider struct{}

func BuildKubemarkCloudProvider(kubemarkController interface{}, specs []string, resourceLimiter *cloudprovider.ResourceLimiter) (*KubemarkCloudProvider, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (kubemark *KubemarkCloudProvider) Name() string { return "" }

<<<<<<< HEAD
=======
// GPULabel returns the label added to nodes with GPU resource.
func (kubemark *KubemarkCloudProvider) GPULabel() string {
	return GPULabel
}

// GetAvailableGPUTypes return all available GPU types cloud provider supports
func (kubemark *KubemarkCloudProvider) GetAvailableGPUTypes() map[string]struct{} {
	return availableGPUTypes
}

// NodeGroups returns all node groups configured for this cloud provider.
>>>>>>> 90666881d... Move GPULabel and GPUTypes to cloud provider
func (kubemark *KubemarkCloudProvider) NodeGroups() []cloudprovider.NodeGroup {
	return []cloudprovider.NodeGroup{}
}

func (kubemark *KubemarkCloudProvider) Pricing() (cloudprovider.PricingModel, errors.AutoscalerError) {
	return nil, cloudprovider.ErrNotImplemented
}

func (kubemark *KubemarkCloudProvider) NodeGroupForNode(node *apiv1.Node) (cloudprovider.NodeGroup, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (kubemark *KubemarkCloudProvider) GetAvailableMachineTypes() ([]string, error) {
	return []string{}, cloudprovider.ErrNotImplemented
}

func (kubemark *KubemarkCloudProvider) NewNodeGroup(machineType string, labels map[string]string, systemLabels map[string]string,
	taints []apiv1.Taint,
	extraResources map[string]resource.Quantity) (cloudprovider.NodeGroup, error) {
	return nil, cloudprovider.ErrNotImplemented
}

// GetResourceLimiter returns struct containing limits (max, min) for resources (cores, memory etc.).
func (kubemark *KubemarkCloudProvider) GetResourceLimiter() (*cloudprovider.ResourceLimiter, error) {
	return nil, cloudprovider.ErrNotImplemented
}

// Refresh is called before every main loop and can be used to dynamically update cloud provider state.
// In particular the list of node groups returned by NodeGroups can change as a result of CloudProvider.Refresh().
func (kubemark *KubemarkCloudProvider) Refresh() error {
	return cloudprovider.ErrNotImplemented
}

// GetInstanceID gets the instance ID for the specified node.
func (kubemark *KubemarkCloudProvider) GetInstanceID(node *apiv1.Node) string {
	return ""
}

// Cleanup cleans up all resources before the cloud provider is removed
func (kubemark *KubemarkCloudProvider) Cleanup() error {
	return cloudprovider.ErrNotImplemented
}
