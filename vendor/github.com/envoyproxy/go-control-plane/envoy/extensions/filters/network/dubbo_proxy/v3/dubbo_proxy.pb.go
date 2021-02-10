// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/extensions/filters/network/dubbo_proxy/v3/dubbo_proxy.proto

package envoy_extensions_filters_network_dubbo_proxy_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// Dubbo Protocol types supported by Envoy.
type ProtocolType int32

const (
	// the default protocol.
	ProtocolType_Dubbo ProtocolType = 0
)

// Enum value maps for ProtocolType.
var (
	ProtocolType_name = map[int32]string{
		0: "Dubbo",
	}
	ProtocolType_value = map[string]int32{
		"Dubbo": 0,
	}
)

func (x ProtocolType) Enum() *ProtocolType {
	p := new(ProtocolType)
	*p = x
	return p
}

func (x ProtocolType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtocolType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_enumTypes[0].Descriptor()
}

func (ProtocolType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_enumTypes[0]
}

func (x ProtocolType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtocolType.Descriptor instead.
func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDescGZIP(), []int{0}
}

// Dubbo Serialization types supported by Envoy.
type SerializationType int32

const (
	// the default serialization protocol.
	SerializationType_Hessian2 SerializationType = 0
)

// Enum value maps for SerializationType.
var (
	SerializationType_name = map[int32]string{
		0: "Hessian2",
	}
	SerializationType_value = map[string]int32{
		"Hessian2": 0,
	}
)

func (x SerializationType) Enum() *SerializationType {
	p := new(SerializationType)
	*p = x
	return p
}

func (x SerializationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SerializationType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_enumTypes[1].Descriptor()
}

func (SerializationType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_enumTypes[1]
}

func (x SerializationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SerializationType.Descriptor instead.
func (SerializationType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDescGZIP(), []int{1}
}

// [#next-free-field: 6]
type DubboProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human readable prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Configure the protocol used.
	ProtocolType ProtocolType `protobuf:"varint,2,opt,name=protocol_type,json=protocolType,proto3,enum=envoy.extensions.filters.network.dubbo_proxy.v3.ProtocolType" json:"protocol_type,omitempty"`
	// Configure the serialization protocol used.
	SerializationType SerializationType `protobuf:"varint,3,opt,name=serialization_type,json=serializationType,proto3,enum=envoy.extensions.filters.network.dubbo_proxy.v3.SerializationType" json:"serialization_type,omitempty"`
	// The route table for the connection manager is static and is specified in this property.
	RouteConfig []*RouteConfiguration `protobuf:"bytes,4,rep,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	// A list of individual Dubbo filters that make up the filter chain for requests made to the
	// Dubbo proxy. Order matters as the filters are processed sequentially. For backwards
	// compatibility, if no dubbo_filters are specified, a default Dubbo router filter
	// (`envoy.filters.dubbo.router`) is used.
	DubboFilters []*DubboFilter `protobuf:"bytes,5,rep,name=dubbo_filters,json=dubboFilters,proto3" json:"dubbo_filters,omitempty"`
}

func (x *DubboProxy) Reset() {
	*x = DubboProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DubboProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DubboProxy) ProtoMessage() {}

func (x *DubboProxy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DubboProxy.ProtoReflect.Descriptor instead.
func (*DubboProxy) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *DubboProxy) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *DubboProxy) GetProtocolType() ProtocolType {
	if x != nil {
		return x.ProtocolType
	}
	return ProtocolType_Dubbo
}

func (x *DubboProxy) GetSerializationType() SerializationType {
	if x != nil {
		return x.SerializationType
	}
	return SerializationType_Hessian2
}

func (x *DubboProxy) GetRouteConfig() []*RouteConfiguration {
	if x != nil {
		return x.RouteConfig
	}
	return nil
}

func (x *DubboProxy) GetDubboFilters() []*DubboFilter {
	if x != nil {
		return x.DubboFilters
	}
	return nil
}

// DubboFilter configures a Dubbo filter.
type DubboFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the filter to instantiate. The name must match a supported
	// filter.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being
	// instantiated. See the supported filters for further documentation.
	Config *any.Any `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *DubboFilter) Reset() {
	*x = DubboFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DubboFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DubboFilter) ProtoMessage() {}

func (x *DubboFilter) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DubboFilter.ProtoReflect.Descriptor instead.
func (*DubboFilter) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *DubboFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DubboFilter) GetConfig() *any.Any {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDesc = []byte{
	0x0a, 0x41, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76,
	0x33, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x33, 0x1a, 0x3b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x04, 0x0a, 0x0a, 0x44, 0x75, 0x62, 0x62,
	0x6f, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x12, 0x6c, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x7b,
	0x0a, 0x12, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x11, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x66, 0x0a, 0x0c, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x43, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x61, 0x0a, 0x0d, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x75, 0x62,
	0x62, 0x6f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x3a, 0x42, 0x9a, 0xc5, 0x88, 0x1e, 0x3d, 0x0a, 0x3b, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x44, 0x75, 0x62, 0x62, 0x6f, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x9d, 0x01, 0x0a, 0x0b, 0x44,
	0x75, 0x62, 0x62, 0x6f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x43, 0x9a, 0xc5, 0x88, 0x1e, 0x3e, 0x0a, 0x3c, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44,
	0x75, 0x62, 0x62, 0x6f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2a, 0x19, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x75,
	0x62, 0x62, 0x6f, 0x10, 0x00, 0x2a, 0x21, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x65,
	0x73, 0x73, 0x69, 0x61, 0x6e, 0x32, 0x10, 0x00, 0x42, 0x5a, 0x0a, 0x3d, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x0f, 0x44, 0x75, 0x62, 0x62, 0x6f,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDescData = file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDesc
)

func file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDescData
}

var file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_goTypes = []interface{}{
	(ProtocolType)(0),          // 0: envoy.extensions.filters.network.dubbo_proxy.v3.ProtocolType
	(SerializationType)(0),     // 1: envoy.extensions.filters.network.dubbo_proxy.v3.SerializationType
	(*DubboProxy)(nil),         // 2: envoy.extensions.filters.network.dubbo_proxy.v3.DubboProxy
	(*DubboFilter)(nil),        // 3: envoy.extensions.filters.network.dubbo_proxy.v3.DubboFilter
	(*RouteConfiguration)(nil), // 4: envoy.extensions.filters.network.dubbo_proxy.v3.RouteConfiguration
	(*any.Any)(nil),            // 5: google.protobuf.Any
}
var file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_depIdxs = []int32{
	0, // 0: envoy.extensions.filters.network.dubbo_proxy.v3.DubboProxy.protocol_type:type_name -> envoy.extensions.filters.network.dubbo_proxy.v3.ProtocolType
	1, // 1: envoy.extensions.filters.network.dubbo_proxy.v3.DubboProxy.serialization_type:type_name -> envoy.extensions.filters.network.dubbo_proxy.v3.SerializationType
	4, // 2: envoy.extensions.filters.network.dubbo_proxy.v3.DubboProxy.route_config:type_name -> envoy.extensions.filters.network.dubbo_proxy.v3.RouteConfiguration
	3, // 3: envoy.extensions.filters.network.dubbo_proxy.v3.DubboProxy.dubbo_filters:type_name -> envoy.extensions.filters.network.dubbo_proxy.v3.DubboFilter
	5, // 4: envoy.extensions.filters.network.dubbo_proxy.v3.DubboFilter.config:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_init() }
func file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_init() {
	if File_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto != nil {
		return
	}
	file_envoy_extensions_filters_network_dubbo_proxy_v3_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DubboProxy); i {
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
		file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DubboFilter); i {
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
			RawDescriptor: file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto = out.File
	file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_rawDesc = nil
	file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_goTypes = nil
	file_envoy_extensions_filters_network_dubbo_proxy_v3_dubbo_proxy_proto_depIdxs = nil
}
