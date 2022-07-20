// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/dbaas/logs.proto

package dbaasv1beta1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Logs) Validate() error {
	return nil
}
func (this *GetLogsRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.ClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.ClusterName))
	}
	return nil
}
func (this *GetLogsResponse) Validate() error {
	for _, item := range this.Logs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Logs", err)
			}
		}
	}
	return nil
}
