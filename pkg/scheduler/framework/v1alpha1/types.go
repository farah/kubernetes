/*
Copyright 2015 The Kubernetes Authors.

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

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type NodeInfo = framework.NodeInfo

func NewNodeInfo(pods ...*v1.Pod) *NodeInfo {
	return framework.NewNodeInfo(pods...)
}

func NewPodInfo(pod *v1.Pod) *PodInfo {
	return framework.NewPodInfo(pod)
}

type Resource = framework.Resource

type QueuedPodInfo = framework.QueuedPodInfo

type PodInfo = framework.PodInfo

type AffinityTerm = framework.AffinityTerm

type WeightedAffinityTerm = framework.WeightedAffinityTerm

type ImageStateSummary = framework.ImageStateSummary

type TransientSchedulerInfo = framework.TransientSchedulerInfo

type HostPortInfo = framework.HostPortInfo

type ProtocolPort = framework.ProtocolPort

func NewResource(rl v1.ResourceList) *Resource {
	return framework.NewResource(rl)
}
