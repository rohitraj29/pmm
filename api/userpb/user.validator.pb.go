// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: userpb/user.proto

package userpb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *UserDetailsRequest) Validate() error {
	return nil
}
func (this *UserDetailsResponse) Validate() error {
	return nil
}
func (this *UserUpdateRequest) Validate() error {
	return nil
}
func (this *ListUsersRequest) Validate() error {
	return nil
}
func (this *ListUsersResponse) Validate() error {
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	return nil
}
func (this *ListUsersResponse_UserDetail) Validate() error {
	return nil
}
