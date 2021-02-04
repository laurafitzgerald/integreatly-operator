// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/config/filter/network/rbac/v2/rbac.proto

package envoy_config_filter_network_rbac_v2

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RBAC_EnforcementType int32

const (
	// Apply RBAC policies when the first byte of data arrives on the connection.
	RBAC_ONE_TIME_ON_FIRST_BYTE RBAC_EnforcementType = 0
	// Continuously apply RBAC policies as data arrives. Use this mode when
	// using RBAC with message oriented protocols such as Mongo, MySQL, Kafka,
	// etc. when the protocol decoders emit dynamic metadata such as the
	// resources being accessed and the operations on the resources.
	RBAC_CONTINUOUS RBAC_EnforcementType = 1
)

// Enum value maps for RBAC_EnforcementType.
var (
	RBAC_EnforcementType_name = map[int32]string{
		0: "ONE_TIME_ON_FIRST_BYTE",
		1: "CONTINUOUS",
	}
	RBAC_EnforcementType_value = map[string]int32{
		"ONE_TIME_ON_FIRST_BYTE": 0,
		"CONTINUOUS":             1,
	}
)

func (x RBAC_EnforcementType) Enum() *RBAC_EnforcementType {
	p := new(RBAC_EnforcementType)
	*p = x
	return p
}

func (x RBAC_EnforcementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RBAC_EnforcementType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_config_filter_network_rbac_v2_rbac_proto_enumTypes[0].Descriptor()
}

func (RBAC_EnforcementType) Type() protoreflect.EnumType {
	return &file_envoy_config_filter_network_rbac_v2_rbac_proto_enumTypes[0]
}

func (x RBAC_EnforcementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RBAC_EnforcementType.Descriptor instead.
func (RBAC_EnforcementType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDescGZIP(), []int{0, 0}
}

// RBAC network filter config.
//
// Header should not be used in rules/shadow_rules in RBAC network filter as
// this information is only available in :ref:`RBAC http filter <config_http_filters_rbac>`.
type RBAC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the RBAC rules to be applied globally.
	// If absent, no enforcing RBAC policy will be applied.
	Rules *v2.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	// Shadow rules are not enforced by the filter but will emit stats and logs
	// and can be used for rule testing.
	// If absent, no shadow RBAC policy will be applied.
	ShadowRules *v2.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	// The prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// RBAC enforcement strategy. By default RBAC will be enforced only once
	// when the first byte of data arrives from the downstream. When used in
	// conjunction with filters that emit dynamic metadata after decoding
	// every payload (e.g., Mongo, MySQL, Kafka) set the enforcement type to
	// CONTINUOUS to enforce RBAC policies on every message boundary.
	EnforcementType RBAC_EnforcementType `protobuf:"varint,4,opt,name=enforcement_type,json=enforcementType,proto3,enum=envoy.config.filter.network.rbac.v2.RBAC_EnforcementType" json:"enforcement_type,omitempty"`
}

func (x *RBAC) Reset() {
	*x = RBAC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_network_rbac_v2_rbac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RBAC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RBAC) ProtoMessage() {}

func (x *RBAC) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_rbac_v2_rbac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RBAC.ProtoReflect.Descriptor instead.
func (*RBAC) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDescGZIP(), []int{0}
}

func (x *RBAC) GetRules() *v2.RBAC {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *RBAC) GetShadowRules() *v2.RBAC {
	if x != nil {
		return x.ShadowRules
	}
	return nil
}

func (x *RBAC) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *RBAC) GetEnforcementType() RBAC_EnforcementType {
	if x != nil {
		return x.EnforcementType
	}
	return RBAC_ONE_TIME_ON_FIRST_BYTE
}

var File_envoy_config_filter_network_rbac_v2_rbac_proto protoreflect.FileDescriptor

var file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x62,
	0x61, 0x63, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x72, 0x62,
	0x61, 0x63, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x62, 0x61, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6,
	0x02, 0x0a, 0x04, 0x52, 0x42, 0x41, 0x43, 0x12, 0x30, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x42,
	0x41, 0x43, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72,
	0x62, 0x61, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x42, 0x41, 0x43, 0x52, 0x0b, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x12, 0x64, 0x0a, 0x10, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x42, 0x41, 0x43, 0x2e, 0x45, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x45, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4f,
	0x4e, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54,
	0x5f, 0x42, 0x59, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x54, 0x49,
	0x4e, 0x55, 0x4f, 0x55, 0x53, 0x10, 0x01, 0x42, 0x78, 0x0a, 0x31, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x32, 0x42, 0x09, 0x52, 0x62,
	0x61, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x2a, 0x12,
	0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDescOnce sync.Once
	file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDescData = file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDesc
)

func file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDescData)
	})
	return file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDescData
}

var file_envoy_config_filter_network_rbac_v2_rbac_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_config_filter_network_rbac_v2_rbac_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_filter_network_rbac_v2_rbac_proto_goTypes = []interface{}{
	(RBAC_EnforcementType)(0), // 0: envoy.config.filter.network.rbac.v2.RBAC.EnforcementType
	(*RBAC)(nil),              // 1: envoy.config.filter.network.rbac.v2.RBAC
	(*v2.RBAC)(nil),           // 2: envoy.config.rbac.v2.RBAC
}
var file_envoy_config_filter_network_rbac_v2_rbac_proto_depIdxs = []int32{
	2, // 0: envoy.config.filter.network.rbac.v2.RBAC.rules:type_name -> envoy.config.rbac.v2.RBAC
	2, // 1: envoy.config.filter.network.rbac.v2.RBAC.shadow_rules:type_name -> envoy.config.rbac.v2.RBAC
	0, // 2: envoy.config.filter.network.rbac.v2.RBAC.enforcement_type:type_name -> envoy.config.filter.network.rbac.v2.RBAC.EnforcementType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_network_rbac_v2_rbac_proto_init() }
func file_envoy_config_filter_network_rbac_v2_rbac_proto_init() {
	if File_envoy_config_filter_network_rbac_v2_rbac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_network_rbac_v2_rbac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RBAC); i {
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
			RawDescriptor: file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_network_rbac_v2_rbac_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_network_rbac_v2_rbac_proto_depIdxs,
		EnumInfos:         file_envoy_config_filter_network_rbac_v2_rbac_proto_enumTypes,
		MessageInfos:      file_envoy_config_filter_network_rbac_v2_rbac_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_network_rbac_v2_rbac_proto = out.File
	file_envoy_config_filter_network_rbac_v2_rbac_proto_rawDesc = nil
	file_envoy_config_filter_network_rbac_v2_rbac_proto_goTypes = nil
	file_envoy_config_filter_network_rbac_v2_rbac_proto_depIdxs = nil
}
