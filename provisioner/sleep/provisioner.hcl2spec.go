// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package sleep

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatProvisioner is an auto-generated flat version of Provisioner.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatProvisioner struct {
	Duration *string `cty:"duration" hcl:"duration"`
}

// FlatMapstructure returns a new FlatProvisioner.
// FlatProvisioner is an auto-generated flat version of Provisioner.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Provisioner) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatProvisioner)
}

// HCL2Spec returns the hcl spec of a Provisioner.
// This spec is used by HCL to read the fields of Provisioner.
// The decoded values from this spec will then be applied to a FlatProvisioner.
func (*FlatProvisioner) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"duration": &hcldec.AttrSpec{Name: "duration", Type: cty.String, Required: false},
	}
	return s
}
