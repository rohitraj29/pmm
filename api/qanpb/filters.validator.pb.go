// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qanpb/filters.proto

package qanv1beta1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *FiltersRequest) Validate() error {
	if this.PeriodStartFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeriodStartFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeriodStartFrom", err)
		}
	}
	if this.PeriodStartTo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeriodStartTo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeriodStartTo", err)
		}
	}
	for _, item := range this.Labels {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Labels", err)
			}
		}
	}
	return nil
}
func (this *FiltersReply) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ListLabels) Validate() error {
	for _, item := range this.Name {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
			}
		}
	}
	return nil
}
func (this *Values) Validate() error {
	return nil
}
