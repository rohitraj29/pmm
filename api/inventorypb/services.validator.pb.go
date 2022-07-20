// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: inventorypb/services.proto

package inventorypb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *MySQLService) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *MongoDBService) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *PostgreSQLService) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ProxySQLService) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *HAProxyService) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ExternalService) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ListServicesRequest) Validate() error {
	return nil
}
func (this *ListServicesResponse) Validate() error {
	for _, item := range this.Mysql {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Mysql", err)
			}
		}
	}
	for _, item := range this.Mongodb {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Mongodb", err)
			}
		}
	}
	for _, item := range this.Postgresql {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Postgresql", err)
			}
		}
	}
	for _, item := range this.Proxysql {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Proxysql", err)
			}
		}
	}
	for _, item := range this.Haproxy {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Haproxy", err)
			}
		}
	}
	for _, item := range this.External {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("External", err)
			}
		}
	}
	return nil
}
func (this *GetServiceRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	return nil
}
func (this *GetServiceResponse) Validate() error {
	if oneOfNester, ok := this.GetService().(*GetServiceResponse_Mysql); ok {
		if oneOfNester.Mysql != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Mysql); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Mysql", err)
			}
		}
	}
	if oneOfNester, ok := this.GetService().(*GetServiceResponse_Mongodb); ok {
		if oneOfNester.Mongodb != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Mongodb); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Mongodb", err)
			}
		}
	}
	if oneOfNester, ok := this.GetService().(*GetServiceResponse_Postgresql); ok {
		if oneOfNester.Postgresql != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Postgresql); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Postgresql", err)
			}
		}
	}
	if oneOfNester, ok := this.GetService().(*GetServiceResponse_Proxysql); ok {
		if oneOfNester.Proxysql != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Proxysql); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Proxysql", err)
			}
		}
	}
	if oneOfNester, ok := this.GetService().(*GetServiceResponse_Haproxy); ok {
		if oneOfNester.Haproxy != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Haproxy); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Haproxy", err)
			}
		}
	}
	if oneOfNester, ok := this.GetService().(*GetServiceResponse_External); ok {
		if oneOfNester.External != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.External); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("External", err)
			}
		}
	}
	return nil
}
func (this *AddMySQLServiceRequest) Validate() error {
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddMySQLServiceResponse) Validate() error {
	if this.Mysql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Mysql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Mysql", err)
		}
	}
	return nil
}
func (this *AddMongoDBServiceRequest) Validate() error {
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddMongoDBServiceResponse) Validate() error {
	if this.Mongodb != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Mongodb); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Mongodb", err)
		}
	}
	return nil
}
func (this *AddPostgreSQLServiceRequest) Validate() error {
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddPostgreSQLServiceResponse) Validate() error {
	if this.Postgresql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Postgresql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Postgresql", err)
		}
	}
	return nil
}
func (this *AddProxySQLServiceRequest) Validate() error {
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddProxySQLServiceResponse) Validate() error {
	if this.Proxysql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proxysql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proxysql", err)
		}
	}
	return nil
}
func (this *AddHAProxyServiceRequest) Validate() error {
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddHAProxyServiceResponse) Validate() error {
	if this.Haproxy != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Haproxy); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Haproxy", err)
		}
	}
	return nil
}
func (this *AddExternalServiceRequest) Validate() error {
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddExternalServiceResponse) Validate() error {
	if this.External != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.External); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("External", err)
		}
	}
	return nil
}
func (this *RemoveServiceRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	return nil
}
func (this *RemoveServiceResponse) Validate() error {
	return nil
}
