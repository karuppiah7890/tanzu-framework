// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package registry_test

import (
	"os"
	"strconv"
	"testing"

	ctlimg "github.com/k14s/imgpkg/pkg/imgpkg/image"
	"gopkg.in/yaml.v2"

	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/registry"
	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/tkgconfigbom"
)

func TestRegistryPullBOM(t *testing.T) {
	if integTest := os.Getenv("TKR_INTEG_TEST"); integTest != "true" {
		return
	}
	bomRegistry := os.Getenv("TEST_BOM_IMAGE_PATH")
	if bomRegistry == "" {
		t.Fatal("TEST_BOM_IMAGE_PATH must be set")
	}

	numTags, err := strconv.Atoi(os.Getenv("TEST_BOM_IMAGE_NUM"))
	if err != nil {
		t.Fatalf("wrong TEST_BOM_IMAGE_NUM, %s", err.Error())
	}

	ro := &ctlimg.RegistryOpts{}

	reg, err := registry.New(ro)
	if err != nil {
		t.Fatalf("error creating registry client")
	}

	tags, err := reg.ListImageTags(bomRegistry)
	if err != nil {
		t.Fatalf("error listing tags should not occurs %s", err.Error())
	}

	if numTags != len(tags) {
		t.Fatal("number of tags does not match")
	}

	t.Log(tags[0])

	content, err := reg.GetFile(bomRegistry, tags[0], "")
	if err != nil {
		t.Fatalf("error getting image content should not occurs %s", err.Error())
	}

	bomConfiguration := &tkgconfigbom.BOMConfiguration{}
	if err := yaml.Unmarshal(content, bomConfiguration); err != nil {
		t.Fatalf("error parsing bom content should not occurs %s", err.Error())
	}

	t.Log(bomConfiguration.Default.TKRVersion)
}
