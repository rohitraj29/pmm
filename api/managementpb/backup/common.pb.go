// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: managementpb/backup/common.proto

package backupv1

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

// DataModel is a model used for performing a backup.
type DataModel int32

const (
	DataModel_DATA_MODEL_INVALID DataModel = 0
	DataModel_PHYSICAL           DataModel = 1
	DataModel_LOGICAL            DataModel = 2
)

// Enum value maps for DataModel.
var (
	DataModel_name = map[int32]string{
		0: "DATA_MODEL_INVALID",
		1: "PHYSICAL",
		2: "LOGICAL",
	}
	DataModel_value = map[string]int32{
		"DATA_MODEL_INVALID": 0,
		"PHYSICAL":           1,
		"LOGICAL":            2,
	}
)

func (x DataModel) Enum() *DataModel {
	p := new(DataModel)
	*p = x
	return p
}

func (x DataModel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataModel) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_backup_common_proto_enumTypes[0].Descriptor()
}

func (DataModel) Type() protoreflect.EnumType {
	return &file_managementpb_backup_common_proto_enumTypes[0]
}

func (x DataModel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataModel.Descriptor instead.
func (DataModel) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_backup_common_proto_rawDescGZIP(), []int{0}
}

// BackupMode specifies backup mode.
type BackupMode int32

const (
	BackupMode_BACKUP_MODE_INVALID BackupMode = 0
	BackupMode_SNAPSHOT            BackupMode = 1
	BackupMode_INCREMENTAL         BackupMode = 2
	BackupMode_PITR                BackupMode = 3
)

// Enum value maps for BackupMode.
var (
	BackupMode_name = map[int32]string{
		0: "BACKUP_MODE_INVALID",
		1: "SNAPSHOT",
		2: "INCREMENTAL",
		3: "PITR",
	}
	BackupMode_value = map[string]int32{
		"BACKUP_MODE_INVALID": 0,
		"SNAPSHOT":            1,
		"INCREMENTAL":         2,
		"PITR":                3,
	}
)

func (x BackupMode) Enum() *BackupMode {
	p := new(BackupMode)
	*p = x
	return p
}

func (x BackupMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackupMode) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_backup_common_proto_enumTypes[1].Descriptor()
}

func (BackupMode) Type() protoreflect.EnumType {
	return &file_managementpb_backup_common_proto_enumTypes[1]
}

func (x BackupMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackupMode.Descriptor instead.
func (BackupMode) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_backup_common_proto_rawDescGZIP(), []int{1}
}

var File_managementpb_backup_common_proto protoreflect.FileDescriptor

var file_managementpb_backup_common_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2a, 0x3e, 0x0a,
	0x09, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x48, 0x59, 0x53, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x4f, 0x47, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x02, 0x2a, 0x4e, 0x0a,
	0x0a, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x42,
	0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x43, 0x52, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x41,
	0x4c, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x54, 0x52, 0x10, 0x03, 0x42, 0x9a, 0x01,
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x42,
	0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f,
	0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x3b, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x09,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_managementpb_backup_common_proto_rawDescOnce sync.Once
	file_managementpb_backup_common_proto_rawDescData = file_managementpb_backup_common_proto_rawDesc
)

func file_managementpb_backup_common_proto_rawDescGZIP() []byte {
	file_managementpb_backup_common_proto_rawDescOnce.Do(func() {
		file_managementpb_backup_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_backup_common_proto_rawDescData)
	})
	return file_managementpb_backup_common_proto_rawDescData
}

var (
	file_managementpb_backup_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
	file_managementpb_backup_common_proto_goTypes   = []interface{}{
		(DataModel)(0),  // 0: backup.v1.DataModel
		(BackupMode)(0), // 1: backup.v1.BackupMode
	}
)

var file_managementpb_backup_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_managementpb_backup_common_proto_init() }
func file_managementpb_backup_common_proto_init() {
	if File_managementpb_backup_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_backup_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_managementpb_backup_common_proto_goTypes,
		DependencyIndexes: file_managementpb_backup_common_proto_depIdxs,
		EnumInfos:         file_managementpb_backup_common_proto_enumTypes,
	}.Build()
	File_managementpb_backup_common_proto = out.File
	file_managementpb_backup_common_proto_rawDesc = nil
	file_managementpb_backup_common_proto_goTypes = nil
	file_managementpb_backup_common_proto_depIdxs = nil
}
