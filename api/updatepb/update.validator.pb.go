// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: updatepb/update.proto

package updatepb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

func (this *StartUpdateRequest) Validate() error {
	if this.ContainerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ContainerId", fmt.Errorf(`value '%v' must not be an empty string`, this.ContainerId))
	}
	return nil
}

func (this *StartUpdateResponse) Validate() error {
	return nil
}

func (this *UpdateStatusRequest) Validate() error {
	return nil
}

func (this *UpdateStatusResponse) Validate() error {
	return nil
}
