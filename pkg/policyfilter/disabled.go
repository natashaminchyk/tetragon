// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

// nolint:revive // prevent unused-parameter alert, disabled method obviously don't use args
package policyfilter

import (
	"fmt"

	slimv1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
	"github.com/cilium/tetragon/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

func DisabledState() State {
	return &disabled{}
}

type disabled struct {
}

func (s *disabled) AddPolicy(polID PolicyID, namespace string, podSelector *slimv1.LabelSelector) error {
	return fmt.Errorf("policyfilter is disabled")
}

func (s *disabled) DelPolicy(polID PolicyID) error {
	return fmt.Errorf("policyfilter is disabled")
}

func (s *disabled) AddPodContainer(podID PodID, namespace string, podLabels labels.Labels, containerID string, cgIDp CgroupID) error {
	return nil
}

func (s *disabled) UpdatePod(podID PodID, namespace string, podLabels labels.Labels, containerIDs []string) error {
	return nil
}

func (s *disabled) DelPodContainer(podID PodID, containerID string) error {
	return nil
}

func (s *disabled) DelPod(podID PodID) error {
	return nil
}

func (s *disabled) RegisterPodHandlers(podInformer cache.SharedIndexInformer) {
}

func (s *disabled) Close() error {
	return nil
}