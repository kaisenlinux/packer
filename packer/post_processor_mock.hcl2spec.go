// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package packer

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/zclconf/go-cty/cty"
)

// FlatMockPostProcessor is an auto-generated flat version of MockPostProcessor.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatMockPostProcessor struct {
	ArtifactId          *string         `cty:"artifact_id" hcl:"artifact_id"`
	Keep                *bool           `cty:"keep" hcl:"keep"`
	ForceOverride       *bool           `cty:"force_override" hcl:"force_override"`
	Error               error           `cty:"error" hcl:"error"`
	ConfigureCalled     *bool           `cty:"configure_called" hcl:"configure_called"`
	ConfigureConfigs    []interface{}   `cty:"configure_configs" hcl:"configure_configs"`
	ConfigureError      error           `cty:"configure_error" hcl:"configure_error"`
	PostProcessCalled   *bool           `cty:"post_process_called" hcl:"post_process_called"`
	PostProcessArtifact packer.Artifact `cty:"post_process_artifact" hcl:"post_process_artifact"`
	PostProcessUi       packer.Ui       `cty:"post_process_ui" hcl:"post_process_ui"`
}

// FlatMapstructure returns a new FlatMockPostProcessor.
// FlatMockPostProcessor is an auto-generated flat version of MockPostProcessor.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*MockPostProcessor) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatMockPostProcessor)
}

// HCL2Spec returns the hcl spec of a MockPostProcessor.
// This spec is used by HCL to read the fields of MockPostProcessor.
// The decoded values from this spec will then be applied to a FlatMockPostProcessor.
func (*FlatMockPostProcessor) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"artifact_id":           &hcldec.AttrSpec{Name: "artifact_id", Type: cty.String, Required: false},
		"keep":                  &hcldec.AttrSpec{Name: "keep", Type: cty.Bool, Required: false},
		"force_override":        &hcldec.AttrSpec{Name: "force_override", Type: cty.Bool, Required: false},
		"error":                 &hcldec.AttrSpec{Name: "error", Type: cty.Bool, Required: false}, /* TODO(azr): could not find type */
		"configure_called":      &hcldec.AttrSpec{Name: "configure_called", Type: cty.Bool, Required: false},
		"configure_configs":     &hcldec.AttrSpec{Name: "configure_configs", Type: cty.Bool, Required: false}, /* TODO(azr): could not find type */
		"configure_error":       &hcldec.AttrSpec{Name: "configure_error", Type: cty.Bool, Required: false},   /* TODO(azr): could not find type */
		"post_process_called":   &hcldec.AttrSpec{Name: "post_process_called", Type: cty.Bool, Required: false},
		"post_process_artifact": &hcldec.AttrSpec{Name: "post_process_artifact", Type: cty.Bool, Required: false}, /* TODO(azr): could not find type */
		"post_process_ui":       &hcldec.AttrSpec{Name: "post_process_ui", Type: cty.Bool, Required: false},       /* TODO(azr): could not find type */
	}
	return s
}
