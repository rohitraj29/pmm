// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/ia/templates.proto

package iav1beta1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/percona/pmm/api/managementpb"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *BoolParamDefinition) Validate() error {
	return nil
}
func (this *FloatParamDefinition) Validate() error {
	return nil
}
func (this *StringParamDefinition) Validate() error {
	return nil
}
func (this *ParamDefinition) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Summary == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Summary", fmt.Errorf(`value '%v' must not be an empty string`, this.Summary))
	}
	if oneOfNester, ok := this.GetValue().(*ParamDefinition_Bool); ok {
		if oneOfNester.Bool != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Bool); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Bool", err)
			}
		}
	}
	if oneOfNester, ok := this.GetValue().(*ParamDefinition_Float); ok {
		if oneOfNester.Float != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Float); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Float", err)
			}
		}
	}
	if oneOfNester, ok := this.GetValue().(*ParamDefinition_String_); ok {
		if oneOfNester.String_ != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.String_); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("String_", err)
			}
		}
	}
	return nil
}
func (this *Template) Validate() error {
	for _, item := range this.Params {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
			}
		}
	}
	if this.For != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.For); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("For", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	return nil
}
func (this *ListTemplatesRequest) Validate() error {
	if this.PageParams != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PageParams); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PageParams", err)
		}
	}
	return nil
}
func (this *ListTemplatesResponse) Validate() error {
	for _, item := range this.Templates {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Templates", err)
			}
		}
	}
	if this.Totals != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Totals); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Totals", err)
		}
	}
	return nil
}
func (this *CreateTemplateRequest) Validate() error {
	if this.Yaml == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Yaml", fmt.Errorf(`value '%v' must not be an empty string`, this.Yaml))
	}
	return nil
}
func (this *CreateTemplateResponse) Validate() error {
	return nil
}
func (this *UpdateTemplateRequest) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Yaml == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Yaml", fmt.Errorf(`value '%v' must not be an empty string`, this.Yaml))
	}
	return nil
}
func (this *UpdateTemplateResponse) Validate() error {
	return nil
}
func (this *DeleteTemplateRequest) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *DeleteTemplateResponse) Validate() error {
	return nil
}
