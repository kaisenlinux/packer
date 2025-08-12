// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package version

import (
	"github.com/hashicorp/packer-plugin-sdk/version"
	packerVersion "github.com/hashicorp/packer/version"
)

var FilePluginVersion *version.PluginVersion

func init() {
	FilePluginVersion = version.NewPluginVersion(
		packerVersion.Version, packerVersion.VersionPrerelease, packerVersion.VersionMetadata)
}
