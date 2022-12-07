// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: managementpb/dbaas/dbaas.proto

package dbaasv1beta1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DBClusterType represents database cluster type.
type DBClusterType int32

const (
	// DB_CLUSTER_TYPE_INVALID represents unknown cluster type.
	DBClusterType_DB_CLUSTER_TYPE_INVALID DBClusterType = 0
	// DB_CLUSTER_TYPE_PXC represents pxc cluster type.
	DBClusterType_DB_CLUSTER_TYPE_PXC DBClusterType = 1
	// DB_CLUSTER_TYPE_PSMDB represents psmdb cluster type.
	DBClusterType_DB_CLUSTER_TYPE_PSMDB DBClusterType = 2
)

// Enum value maps for DBClusterType.
var (
	DBClusterType_name = map[int32]string{
		0: "DB_CLUSTER_TYPE_INVALID",
		1: "DB_CLUSTER_TYPE_PXC",
		2: "DB_CLUSTER_TYPE_PSMDB",
	}
	DBClusterType_value = map[string]int32{
		"DB_CLUSTER_TYPE_INVALID": 0,
		"DB_CLUSTER_TYPE_PXC":     1,
		"DB_CLUSTER_TYPE_PSMDB":   2,
	}
)

func (x DBClusterType) Enum() *DBClusterType {
	p := new(DBClusterType)
	*p = x
	return p
}

func (x DBClusterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DBClusterType) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_dbaas_dbaas_proto_enumTypes[0].Descriptor()
}

func (DBClusterType) Type() protoreflect.EnumType {
	return &file_managementpb_dbaas_dbaas_proto_enumTypes[0]
}

func (x DBClusterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DBClusterType.Descriptor instead.
func (DBClusterType) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_dbaas_dbaas_proto_rawDescGZIP(), []int{0}
}

// OperatorsStatus defines status of operators installed in Kubernetes cluster.
type OperatorsStatus int32

const (
	// OPERATORS_STATUS_INVALID represents unknown state.
	OperatorsStatus_OPERATORS_STATUS_INVALID OperatorsStatus = 0
	// OPERATORS_STATUS_OK represents that operators are installed and have supported API version.
	OperatorsStatus_OPERATORS_STATUS_OK OperatorsStatus = 1
	// OPERATORS_STATUS_UNSUPPORTED represents that operators are installed, but doesn't have supported API version.
	OperatorsStatus_OPERATORS_STATUS_UNSUPPORTED OperatorsStatus = 2
	// OPERATORS_STATUS_NOT_INSTALLED represents that operators are not installed.
	OperatorsStatus_OPERATORS_STATUS_NOT_INSTALLED OperatorsStatus = 3
)

// Enum value maps for OperatorsStatus.
var (
	OperatorsStatus_name = map[int32]string{
		0: "OPERATORS_STATUS_INVALID",
		1: "OPERATORS_STATUS_OK",
		2: "OPERATORS_STATUS_UNSUPPORTED",
		3: "OPERATORS_STATUS_NOT_INSTALLED",
	}
	OperatorsStatus_value = map[string]int32{
		"OPERATORS_STATUS_INVALID":       0,
		"OPERATORS_STATUS_OK":            1,
		"OPERATORS_STATUS_UNSUPPORTED":   2,
		"OPERATORS_STATUS_NOT_INSTALLED": 3,
	}
)

func (x OperatorsStatus) Enum() *OperatorsStatus {
	p := new(OperatorsStatus)
	*p = x
	return p
}

func (x OperatorsStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperatorsStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_dbaas_dbaas_proto_enumTypes[1].Descriptor()
}

func (OperatorsStatus) Type() protoreflect.EnumType {
	return &file_managementpb_dbaas_dbaas_proto_enumTypes[1]
}

func (x OperatorsStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperatorsStatus.Descriptor instead.
func (OperatorsStatus) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_dbaas_dbaas_proto_rawDescGZIP(), []int{1}
}

// RunningOperation respresents a long-running operation.
type RunningOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Finished steps of the operaion; can decrease or increase compared to the previous value.
	FinishedSteps int32 `protobuf:"varint,1,opt,name=finished_steps,json=finishedSteps,proto3" json:"finished_steps,omitempty"`
	// Text describing the current operation progress step.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Total steps needed to finish the operation; can decrease or increase compared to the previous value.
	TotalSteps int32 `protobuf:"varint,3,opt,name=total_steps,json=totalSteps,proto3" json:"total_steps,omitempty"`
}

func (x *RunningOperation) Reset() {
	*x = RunningOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_dbaas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunningOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunningOperation) ProtoMessage() {}

func (x *RunningOperation) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_dbaas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunningOperation.ProtoReflect.Descriptor instead.
func (*RunningOperation) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_dbaas_proto_rawDescGZIP(), []int{0}
}

func (x *RunningOperation) GetFinishedSteps() int32 {
	if x != nil {
		return x.FinishedSteps
	}
	return 0
}

func (x *RunningOperation) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RunningOperation) GetTotalSteps() int32 {
	if x != nil {
		return x.TotalSteps
	}
	return 0
}

// ComputeResources represents container computer resources requests or limits.
type ComputeResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CpuM int32 `protobuf:"varint,1,opt,name=cpu_m,json=cpuM,proto3" json:"cpu_m,omitempty"`
	// Memory in bytes.
	MemoryBytes int64 `protobuf:"varint,2,opt,name=memory_bytes,json=memoryBytes,proto3" json:"memory_bytes,omitempty"`
}

func (x *ComputeResources) Reset() {
	*x = ComputeResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_dbaas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeResources) ProtoMessage() {}

func (x *ComputeResources) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_dbaas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeResources.ProtoReflect.Descriptor instead.
func (*ComputeResources) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_dbaas_proto_rawDescGZIP(), []int{1}
}

func (x *ComputeResources) GetCpuM() int32 {
	if x != nil {
		return x.CpuM
	}
	return 0
}

func (x *ComputeResources) GetMemoryBytes() int64 {
	if x != nil {
		return x.MemoryBytes
	}
	return 0
}

// Resources contains Kubernetes cluster resources.
type Resources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Memory in bytes.
	MemoryBytes uint64 `protobuf:"varint,1,opt,name=memory_bytes,json=memoryBytes,proto3" json:"memory_bytes,omitempty"`
	// CPU in millicpus. For example 0.1 of CPU is equivalent to 100 millicpus.
	// See https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.
	CpuM uint64 `protobuf:"varint,2,opt,name=cpu_m,json=cpuM,proto3" json:"cpu_m,omitempty"`
	// Disk size in bytes.
	DiskSize uint64 `protobuf:"varint,3,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
}

func (x *Resources) Reset() {
	*x = Resources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_dbaas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resources) ProtoMessage() {}

func (x *Resources) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_dbaas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resources.ProtoReflect.Descriptor instead.
func (*Resources) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_dbaas_proto_rawDescGZIP(), []int{2}
}

func (x *Resources) GetMemoryBytes() uint64 {
	if x != nil {
		return x.MemoryBytes
	}
	return 0
}

func (x *Resources) GetCpuM() uint64 {
	if x != nil {
		return x.CpuM
	}
	return 0
}

func (x *Resources) GetDiskSize() uint64 {
	if x != nil {
		return x.DiskSize
	}
	return 0
}

var File_managementpb_dbaas_dbaas_proto protoreflect.FileDescriptor

var file_managementpb_dbaas_dbaas_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x64,
	0x62, 0x61, 0x61, 0x73, 0x2f, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22,
	0x74, 0x0a, 0x10, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f,
	0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x53, 0x74, 0x65, 0x70, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74,
	0x65, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x53, 0x74, 0x65, 0x70, 0x73, 0x22, 0x4a, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x70, 0x75,
	0x5f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x70, 0x75, 0x4d, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x22, 0x60, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x70, 0x75, 0x5f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x63, 0x70, 0x75, 0x4d, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x53,
	0x69, 0x7a, 0x65, 0x2a, 0x60, 0x0a, 0x0d, 0x44, 0x42, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x42, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x42, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x58, 0x43, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x42,
	0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x53,
	0x4d, 0x44, 0x42, 0x10, 0x02, 0x2a, 0x8e, 0x01, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x4f, 0x52, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x4f, 0x52, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x01,
	0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x53, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x53, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41,
	0x4c, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0xb0, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x44, 0x62,
	0x61, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70,
	0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2f, 0x64, 0x62, 0x61, 0x61, 0x73, 0x3b, 0x64, 0x62, 0x61, 0x61, 0x73, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x44,
	0x62, 0x61, 0x61, 0x73, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x0d, 0x44,
	0x62, 0x61, 0x61, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x19, 0x44,
	0x62, 0x61, 0x61, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x44, 0x62, 0x61, 0x61, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_managementpb_dbaas_dbaas_proto_rawDescOnce sync.Once
	file_managementpb_dbaas_dbaas_proto_rawDescData = file_managementpb_dbaas_dbaas_proto_rawDesc
)

func file_managementpb_dbaas_dbaas_proto_rawDescGZIP() []byte {
	file_managementpb_dbaas_dbaas_proto_rawDescOnce.Do(func() {
		file_managementpb_dbaas_dbaas_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_dbaas_dbaas_proto_rawDescData)
	})
	return file_managementpb_dbaas_dbaas_proto_rawDescData
}

var (
	file_managementpb_dbaas_dbaas_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
	file_managementpb_dbaas_dbaas_proto_msgTypes  = make([]protoimpl.MessageInfo, 3)
	file_managementpb_dbaas_dbaas_proto_goTypes   = []interface{}{
		(DBClusterType)(0),       // 0: dbaas.v1beta1.DBClusterType
		(OperatorsStatus)(0),     // 1: dbaas.v1beta1.OperatorsStatus
		(*RunningOperation)(nil), // 2: dbaas.v1beta1.RunningOperation
		(*ComputeResources)(nil), // 3: dbaas.v1beta1.ComputeResources
		(*Resources)(nil),        // 4: dbaas.v1beta1.Resources
	}
)

var file_managementpb_dbaas_dbaas_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_managementpb_dbaas_dbaas_proto_init() }
func file_managementpb_dbaas_dbaas_proto_init() {
	if File_managementpb_dbaas_dbaas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managementpb_dbaas_dbaas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunningOperation); i {
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
		file_managementpb_dbaas_dbaas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeResources); i {
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
		file_managementpb_dbaas_dbaas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resources); i {
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
			RawDescriptor: file_managementpb_dbaas_dbaas_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_managementpb_dbaas_dbaas_proto_goTypes,
		DependencyIndexes: file_managementpb_dbaas_dbaas_proto_depIdxs,
		EnumInfos:         file_managementpb_dbaas_dbaas_proto_enumTypes,
		MessageInfos:      file_managementpb_dbaas_dbaas_proto_msgTypes,
	}.Build()
	File_managementpb_dbaas_dbaas_proto = out.File
	file_managementpb_dbaas_dbaas_proto_rawDesc = nil
	file_managementpb_dbaas_dbaas_proto_goTypes = nil
	file_managementpb_dbaas_dbaas_proto_depIdxs = nil
}
