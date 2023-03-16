// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: managementpb/backup/restores.proto

package backupv1

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RestoreStatus shows the current status of execution of restore.
type RestoreStatus int32

const (
	RestoreStatus_RESTORE_STATUS_INVALID     RestoreStatus = 0
	RestoreStatus_RESTORE_STATUS_IN_PROGRESS RestoreStatus = 1
	RestoreStatus_RESTORE_STATUS_SUCCESS     RestoreStatus = 2
	RestoreStatus_RESTORE_STATUS_ERROR       RestoreStatus = 3
)

// Enum value maps for RestoreStatus.
var (
	RestoreStatus_name = map[int32]string{
		0: "RESTORE_STATUS_INVALID",
		1: "RESTORE_STATUS_IN_PROGRESS",
		2: "RESTORE_STATUS_SUCCESS",
		3: "RESTORE_STATUS_ERROR",
	}
	RestoreStatus_value = map[string]int32{
		"RESTORE_STATUS_INVALID":     0,
		"RESTORE_STATUS_IN_PROGRESS": 1,
		"RESTORE_STATUS_SUCCESS":     2,
		"RESTORE_STATUS_ERROR":       3,
	}
)

func (x RestoreStatus) Enum() *RestoreStatus {
	p := new(RestoreStatus)
	*p = x
	return p
}

func (x RestoreStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RestoreStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_backup_restores_proto_enumTypes[0].Descriptor()
}

func (RestoreStatus) Type() protoreflect.EnumType {
	return &file_managementpb_backup_restores_proto_enumTypes[0]
}

func (x RestoreStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RestoreStatus.Descriptor instead.
func (RestoreStatus) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_backup_restores_proto_rawDescGZIP(), []int{0}
}

// RestoreHistoryItem represents single backup restore item.
type RestoreHistoryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Machine-readable restore id.
	RestoreId string `protobuf:"bytes,1,opt,name=restore_id,json=restoreId,proto3" json:"restore_id,omitempty"`
	// ID of the artifact used for restore.
	ArtifactId string `protobuf:"bytes,2,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
	// Artifact name used for restore.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Database vendor e.g. PostgreSQL, MongoDB, MySQL.
	Vendor string `protobuf:"bytes,4,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Machine-readable location ID.
	LocationId string `protobuf:"bytes,5,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	// Location name.
	LocationName string `protobuf:"bytes,6,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	// Machine-readable service ID.
	ServiceId string `protobuf:"bytes,7,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// Service name.
	ServiceName string `protobuf:"bytes,8,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Backup data model.
	DataModel DataModel `protobuf:"varint,9,opt,name=data_model,json=dataModel,proto3,enum=backup.v1.DataModel" json:"data_model,omitempty"`
	// Restore status.
	Status RestoreStatus `protobuf:"varint,10,opt,name=status,proto3,enum=backup.v1.RestoreStatus" json:"status,omitempty"`
	// Restore start time.
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// Restore finish time.
	FinishedAt *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	// PITR timestamp is filled for PITR restores, empty otherwise.
	PitrTimestamp *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=pitr_timestamp,json=pitrTimestamp,proto3" json:"pitr_timestamp,omitempty"`
}

func (x *RestoreHistoryItem) Reset() {
	*x = RestoreHistoryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_restores_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreHistoryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreHistoryItem) ProtoMessage() {}

func (x *RestoreHistoryItem) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_restores_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreHistoryItem.ProtoReflect.Descriptor instead.
func (*RestoreHistoryItem) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_restores_proto_rawDescGZIP(), []int{0}
}

func (x *RestoreHistoryItem) GetRestoreId() string {
	if x != nil {
		return x.RestoreId
	}
	return ""
}

func (x *RestoreHistoryItem) GetArtifactId() string {
	if x != nil {
		return x.ArtifactId
	}
	return ""
}

func (x *RestoreHistoryItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RestoreHistoryItem) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *RestoreHistoryItem) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *RestoreHistoryItem) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

func (x *RestoreHistoryItem) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *RestoreHistoryItem) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *RestoreHistoryItem) GetDataModel() DataModel {
	if x != nil {
		return x.DataModel
	}
	return DataModel_DATA_MODEL_INVALID
}

func (x *RestoreHistoryItem) GetStatus() RestoreStatus {
	if x != nil {
		return x.Status
	}
	return RestoreStatus_RESTORE_STATUS_INVALID
}

func (x *RestoreHistoryItem) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *RestoreHistoryItem) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *RestoreHistoryItem) GetPitrTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.PitrTimestamp
	}
	return nil
}

type ListRestoreHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRestoreHistoryRequest) Reset() {
	*x = ListRestoreHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_restores_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRestoreHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRestoreHistoryRequest) ProtoMessage() {}

func (x *ListRestoreHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_restores_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRestoreHistoryRequest.ProtoReflect.Descriptor instead.
func (*ListRestoreHistoryRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_restores_proto_rawDescGZIP(), []int{1}
}

type ListRestoreHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*RestoreHistoryItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListRestoreHistoryResponse) Reset() {
	*x = ListRestoreHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_restores_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRestoreHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRestoreHistoryResponse) ProtoMessage() {}

func (x *ListRestoreHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_restores_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRestoreHistoryResponse.ProtoReflect.Descriptor instead.
func (*ListRestoreHistoryResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_restores_proto_rawDescGZIP(), []int{2}
}

func (x *ListRestoreHistoryResponse) GetItems() []*RestoreHistoryItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_managementpb_backup_restores_proto protoreflect.FileDescriptor

var file_managementpb_backup_restores_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xaa, 0x04, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a,
	0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x69,
	0x74, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d,
	0x70, 0x69, 0x74, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x1b, 0x0a,
	0x19, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x1a, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x81, 0x01,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x16, 0x52, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x52,
	0x45, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e,
	0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x52,
	0x45, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x54, 0x4f,
	0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x03, 0x32, 0xaa, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x97, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e,
	0x3a, 0x01, 0x2a, 0x22, 0x29, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x52, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x9c,
	0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31,
	0x42, 0x0d, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65,
	0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x3b, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58,
	0xaa, 0x02, 0x09, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_backup_restores_proto_rawDescOnce sync.Once
	file_managementpb_backup_restores_proto_rawDescData = file_managementpb_backup_restores_proto_rawDesc
)

func file_managementpb_backup_restores_proto_rawDescGZIP() []byte {
	file_managementpb_backup_restores_proto_rawDescOnce.Do(func() {
		file_managementpb_backup_restores_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_backup_restores_proto_rawDescData)
	})
	return file_managementpb_backup_restores_proto_rawDescData
}

var (
	file_managementpb_backup_restores_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_managementpb_backup_restores_proto_msgTypes  = make([]protoimpl.MessageInfo, 3)
	file_managementpb_backup_restores_proto_goTypes   = []interface{}{
		(RestoreStatus)(0),                 // 0: backup.v1.RestoreStatus
		(*RestoreHistoryItem)(nil),         // 1: backup.v1.RestoreHistoryItem
		(*ListRestoreHistoryRequest)(nil),  // 2: backup.v1.ListRestoreHistoryRequest
		(*ListRestoreHistoryResponse)(nil), // 3: backup.v1.ListRestoreHistoryResponse
		(DataModel)(0),                     // 4: backup.v1.DataModel
		(*timestamppb.Timestamp)(nil),      // 5: google.protobuf.Timestamp
	}
)

var file_managementpb_backup_restores_proto_depIdxs = []int32{
	4, // 0: backup.v1.RestoreHistoryItem.data_model:type_name -> backup.v1.DataModel
	0, // 1: backup.v1.RestoreHistoryItem.status:type_name -> backup.v1.RestoreStatus
	5, // 2: backup.v1.RestoreHistoryItem.started_at:type_name -> google.protobuf.Timestamp
	5, // 3: backup.v1.RestoreHistoryItem.finished_at:type_name -> google.protobuf.Timestamp
	5, // 4: backup.v1.RestoreHistoryItem.pitr_timestamp:type_name -> google.protobuf.Timestamp
	1, // 5: backup.v1.ListRestoreHistoryResponse.items:type_name -> backup.v1.RestoreHistoryItem
	2, // 6: backup.v1.RestoreHistory.ListRestoreHistory:input_type -> backup.v1.ListRestoreHistoryRequest
	3, // 7: backup.v1.RestoreHistory.ListRestoreHistory:output_type -> backup.v1.ListRestoreHistoryResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_managementpb_backup_restores_proto_init() }
func file_managementpb_backup_restores_proto_init() {
	if File_managementpb_backup_restores_proto != nil {
		return
	}
	file_managementpb_backup_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_managementpb_backup_restores_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreHistoryItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_backup_restores_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRestoreHistoryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_backup_restores_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRestoreHistoryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_backup_restores_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_backup_restores_proto_goTypes,
		DependencyIndexes: file_managementpb_backup_restores_proto_depIdxs,
		EnumInfos:         file_managementpb_backup_restores_proto_enumTypes,
		MessageInfos:      file_managementpb_backup_restores_proto_msgTypes,
	}.Build()
	File_managementpb_backup_restores_proto = out.File
	file_managementpb_backup_restores_proto_rawDesc = nil
	file_managementpb_backup_restores_proto_goTypes = nil
	file_managementpb_backup_restores_proto_depIdxs = nil
}
