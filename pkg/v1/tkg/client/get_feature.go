// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package client

// ListTKGFeatures

import (
	"github.com/pkg/errors"

	configv1alpha1 "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
)

// ListTKGClustersOptions contains options supported by ListClusters
type ListTKGFeaturesOptions struct {
	activated       bool
	available       bool
	immutable       bool
	nondiscoverable bool
}

// FeatureInfo defines features
type FeatureInfo struct {
	Name         string `json:"name" yaml:"name"`
	Description  string `json:"description" yaml:"description"`
	Activated    bool   `json:"activated" yaml:"activated"`
	Discoverable bool   `json:"discoverable" yaml:"discoverable"`
	Immutable    bool   `json:"immutable" yaml:"immutable"`
	Gate         string `json:"gate" yaml:"gate"`
}

// GetTKGFeatureState reconciles the current state of Features and Gates into an easy to parse list of Features
func (c *TkgClient) GetFeatures(options ListTKGClustersOptions) ([]*configv1alpha1.Feature, error) {
	features := &configv1alpha1.FeatureList{}
	err := c.clusterClient.ListResources(&features, listOptions)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get list of Feature objects")
	}

	return features, nil
}

func (c *TkgClient) GetFeatureGates(options ListTKGClustersOptions) ([]*configv1alpha1.FeatureGate, error) {
	gates := &configv1alpha1.FeatureGateList{}
	err := c.clusterClient.ListResources(&features, listOptions)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get list of FeatureGate objects")
	}

	return gates, nil
}
