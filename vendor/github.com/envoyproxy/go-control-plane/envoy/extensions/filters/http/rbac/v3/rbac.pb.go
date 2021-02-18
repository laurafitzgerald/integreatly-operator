// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/extensions/filters/http/rbac/v3/rbac.proto

package envoy_extensions_filters_http_rbac_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v3"
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

// RBAC filter config.
type RBAC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the RBAC rules to be applied globally.
	// If absent, no enforcing RBAC policy will be applied.
	Rules *v3.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	// Shadow rules are not enforced by the filter (i.e., returning a 403)
	// but will emit stats and logs and can be used for rule testing.
	// If absent, no shadow RBAC policy will be applied.
	ShadowRules *v3.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
}

func (x *RBAC) Reset() {
	*x = RBAC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_rbac_v3_rbac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RBAC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RBAC) ProtoMessage() {}

func (x *RBAC) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_rbac_v3_rbac_proto_msgTypes[0]
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
	return file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDescGZIP(), []int{0}
}

func (x *RBAC) GetRules() *v3.RBAC {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *RBAC) GetShadowRules() *v3.RBAC {
	if x != nil {
		return x.ShadowRules
	}
	return nil
}

type RBACPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Override the global configuration of the filter with this new config.
	// If absent, the global RBAC policy will be disabled for this route.
	Rbac *RBAC `protobuf:"bytes,2,opt,name=rbac,proto3" json:"rbac,omitempty"`
}

func (x *RBACPerRoute) Reset() {
	*x = RBACPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_rbac_v3_rbac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RBACPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RBACPerRoute) ProtoMessage() {}

func (x *RBACPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_rbac_v3_rbac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RBACPerRoute.ProtoReflect.Descriptor instead.
func (*RBACPerRoute) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDescGZIP(), []int{1}
}

func (x *RBACPerRoute) GetRbac() *RBAC {
	if x != nil {
		return x.Rbac
	}
	return nil
}

var File_envoy_extensions_filters_http_rbac_v3_rbac_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x72, 0x62, 0x61, 0x63, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x76, 0x33, 0x2f,
	0x72, 0x62, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x04, 0x52, 0x42, 0x41, 0x43, 0x12, 0x30,
	0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x62, 0x61,
	0x63, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x42, 0x41, 0x43, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x42,
	0x41, 0x43, 0x52, 0x0b, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x3a,
	0x2c, 0x9a, 0xc5, 0x88, 0x1e, 0x27, 0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x42, 0x41, 0x43, 0x22, 0x8b, 0x01,
	0x0a, 0x0c, 0x52, 0x42, 0x41, 0x43, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x3f,
	0x0a, 0x04, 0x72, 0x62, 0x61, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x62, 0x61,
	0x63, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x42, 0x41, 0x43, 0x52, 0x04, 0x72, 0x62, 0x61, 0x63, 0x3a,
	0x34, 0x9a, 0xc5, 0x88, 0x1e, 0x2f, 0x0a, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x42, 0x41, 0x43, 0x50, 0x65, 0x72,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x42, 0x4a, 0x0a, 0x33, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e,
	0x76, 0x33, 0x42, 0x09, 0x52, 0x62, 0x61, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDescData = file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDesc
)

func file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDescData
}

var file_envoy_extensions_filters_http_rbac_v3_rbac_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_rbac_v3_rbac_proto_goTypes = []interface{}{
	(*RBAC)(nil),         // 0: envoy.extensions.filters.http.rbac.v3.RBAC
	(*RBACPerRoute)(nil), // 1: envoy.extensions.filters.http.rbac.v3.RBACPerRoute
	(*v3.RBAC)(nil),      // 2: envoy.config.rbac.v3.RBAC
}
var file_envoy_extensions_filters_http_rbac_v3_rbac_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.http.rbac.v3.RBAC.rules:type_name -> envoy.config.rbac.v3.RBAC
	2, // 1: envoy.extensions.filters.http.rbac.v3.RBAC.shadow_rules:type_name -> envoy.config.rbac.v3.RBAC
	0, // 2: envoy.extensions.filters.http.rbac.v3.RBACPerRoute.rbac:type_name -> envoy.extensions.filters.http.rbac.v3.RBAC
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_rbac_v3_rbac_proto_init() }
func file_envoy_extensions_filters_http_rbac_v3_rbac_proto_init() {
	if File_envoy_extensions_filters_http_rbac_v3_rbac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_rbac_v3_rbac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_envoy_extensions_filters_http_rbac_v3_rbac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RBACPerRoute); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_rbac_v3_rbac_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_rbac_v3_rbac_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_rbac_v3_rbac_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_rbac_v3_rbac_proto = out.File
	file_envoy_extensions_filters_http_rbac_v3_rbac_proto_rawDesc = nil
	file_envoy_extensions_filters_http_rbac_v3_rbac_proto_goTypes = nil
	file_envoy_extensions_filters_http_rbac_v3_rbac_proto_depIdxs = nil
}