// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/annotation.proto

package managementpb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AddAnnotationRequest) Validate() error {
	if this.Text == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Text", fmt.Errorf(`value '%v' must not be an empty string`, this.Text))
	}
	return nil
}
func (this *AddAnnotationResponse) Validate() error {
	return nil
}
