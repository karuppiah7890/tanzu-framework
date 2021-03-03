// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"strings"

	"golang.org/x/mod/semver"
)

// VersionLatest is the latest version.
var VersionLatest = "latest"

// VersionSelector returns a version from a set of versions.
type VersionSelector func(versions []string) string

// DefaultVersionSelector is the default version selector.
var DefaultVersionSelector = SelectVersionStable

// FilterVersions returns the list of valid versions depending on whether
// unstable versions are requested or not.
func FilterVersions(versions []string, includeUnstable bool) (results []string) {
	for _, version := range versions {
		if !semver.IsValid(version) {
			continue
		}
		split := strings.Split(version, "-")
		if len(split) > 1 && !includeUnstable {
			continue
		}
		results = append(results, version)
	}
	return
}

// SelectVersionStable returns the latest stable version from a list of
// versions. If there are no stable versions it will return an empty string.
func SelectVersionStable(versions []string) (v string) {
	for _, version := range FilterVersions(versions, false) {
		c := semver.Compare(v, version)
		if c == -1 {
			v = version
		}
	}
	return
}

// SelectVersionAny returns the latest version from a list of versions including prereleases.
func SelectVersionAny(versions []string) (v string) {
	for _, version := range FilterVersions(versions, true) {
		c := semver.Compare(v, version)
		if c == -1 {
			v = version
		}
	}
	return
}

// TODO: SelectVersionSaaS
