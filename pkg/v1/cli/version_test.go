// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelectVersionStable(t *testing.T) {
	for _, test := range []struct {
		name     string
		versions []string
		max      string
	}{
		{
			name:     "basic patch",
			versions: []string{"v0.0.1", "v0.0.2"},
			max:      "v0.0.2",
		},
		{
			name:     "release candidates",
			versions: []string{"v0.0.1", "v1.3.0-rc.1", "v1.3.0-pre-alpha-1"},
			max:      "v0.0.1",
		},
		{
			name:     "release candidates same",
			versions: []string{"v0.0.1", "v1.3.0", "v1.3.0-rc.1", "v1.3.0-pre-alpha-1"},
			max:      "v1.3.0",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			p := Plugin{
				Name:     "foo",
				Versions: test.versions,
			}
			v := p.FindVersion(SelectVersionStable)
			require.Equal(t, test.max, v)
		})
	}
}

func TestSelectVersionAny(t *testing.T) {
	for _, test := range []struct {
		name     string
		versions []string
		max      string
	}{
		{
			name:     "basic patch",
			versions: []string{"v0.0.1", "v0.0.2"},
			max:      "v0.0.2",
		},
		{
			name:     "release candidates",
			versions: []string{"v0.0.1", "v1.3.0-rc.1", "v1.3.0-pre-alpha-1"},
			max:      "v1.3.0-rc.1",
		},
		{
			name:     "release candidates same",
			versions: []string{"v0.0.1", "v1.3.0", "v1.3.0-rc.1", "v1.3.0-pre-alpha-1"},
			max:      "v1.3.0",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			p := Plugin{
				Name:     "foo",
				Versions: test.versions,
			}
			v := p.FindVersion(SelectVersionAny)
			require.Equal(t, test.max, v)
		})
	}
}

func TestFilterVersions(t *testing.T) {
	for _, test := range []struct {
		name            string
		includeUnstable bool
		versions        []string
		results         []string
	}{
		{
			name:            "ignore bad versions",
			includeUnstable: false,
			versions:        []string{"v0.0.1", "dev", "bad.version.not.semver", "v0.0.2"},
			results:         []string{"v0.0.1", "v0.0.2"},
		},
		{
			name:            "all versions including unstable",
			includeUnstable: true,
			versions:        []string{"v0.0.1", "v1.3.0-rc.1", "v1.3.0-pre-alpha-1"},
			results:         []string{"v0.0.1", "v1.3.0-rc.1", "v1.3.0-pre-alpha-1"},
		},
		{
			name:            "no unstable versions",
			includeUnstable: false,
			versions:        []string{"v0.0.1", "v1.3.0-rc.1", "v1.3.0"},
			results:         []string{"v0.0.1", "v1.3.0"},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			v := FilterVersions(test.versions, test.includeUnstable)
			require.Equal(t, test.results, v)
		})
	}
}
