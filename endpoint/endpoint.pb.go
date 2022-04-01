// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: endpoint/endpoint.proto

package endpoint

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServiceEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	PortName string `protobuf:"bytes,3,opt,name=port_name,json=portName,proto3" json:"port_name,omitempty"`
}

func (x *ServiceEndpoint) Reset() {
	*x = ServiceEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endpoint_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceEndpoint) ProtoMessage() {}

func (x *ServiceEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_endpoint_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceEndpoint.ProtoReflect.Descriptor instead.
func (*ServiceEndpoint) Descriptor() ([]byte, []int) {
	return file_endpoint_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceEndpoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceEndpoint) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ServiceEndpoint) GetPortName() string {
	if x != nil {
		return x.PortName
	}
	return ""
}

type EndpointInNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NamespaceId int64 `protobuf:"varint,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
}

func (x *EndpointInNamespaceRequest) Reset() {
	*x = EndpointInNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endpoint_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointInNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointInNamespaceRequest) ProtoMessage() {}

func (x *EndpointInNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_endpoint_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointInNamespaceRequest.ProtoReflect.Descriptor instead.
func (*EndpointInNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_endpoint_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointInNamespaceRequest) GetNamespaceId() int64 {
	if x != nil {
		return x.NamespaceId
	}
	return 0
}

type EndpointInNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ServiceEndpoint `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *EndpointInNamespaceResponse) Reset() {
	*x = EndpointInNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endpoint_endpoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointInNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointInNamespaceResponse) ProtoMessage() {}

func (x *EndpointInNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_endpoint_endpoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointInNamespaceResponse.ProtoReflect.Descriptor instead.
func (*EndpointInNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_endpoint_endpoint_proto_rawDescGZIP(), []int{2}
}

func (x *EndpointInNamespaceResponse) GetItems() []*ServiceEndpoint {
	if x != nil {
		return x.Items
	}
	return nil
}

type EndpointInProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *EndpointInProjectRequest) Reset() {
	*x = EndpointInProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endpoint_endpoint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointInProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointInProjectRequest) ProtoMessage() {}

func (x *EndpointInProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_endpoint_endpoint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointInProjectRequest.ProtoReflect.Descriptor instead.
func (*EndpointInProjectRequest) Descriptor() ([]byte, []int) {
	return file_endpoint_endpoint_proto_rawDescGZIP(), []int{3}
}

func (x *EndpointInProjectRequest) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type EndpointInProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ServiceEndpoint `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *EndpointInProjectResponse) Reset() {
	*x = EndpointInProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endpoint_endpoint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointInProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointInProjectResponse) ProtoMessage() {}

func (x *EndpointInProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_endpoint_endpoint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointInProjectResponse.ProtoReflect.Descriptor instead.
func (*EndpointInProjectResponse) Descriptor() ([]byte, []int) {
	return file_endpoint_endpoint_proto_rawDescGZIP(), []int{4}
}

func (x *EndpointInProjectResponse) GetItems() []*ServiceEndpoint {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_endpoint_endpoint_proto protoreflect.FileDescriptor

var file_endpoint_endpoint_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x54, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f,
	0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x1a, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x49, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22,
	0x02, 0x20, 0x00, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x45, 0x0a, 0x1b, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x42, 0x0a, 0x18, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x19, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x32, 0xf8, 0x01, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x7a, 0x0a,
	0x0b, 0x49, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12,
	0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x70, 0x0a, 0x09, 0x49, 0x6e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x63, 0x2d, 0x63, 0x6e,
	0x7a, 0x6a, 0x2f, 0x6d, 0x61, 0x72, 0x73, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x34, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_endpoint_endpoint_proto_rawDescOnce sync.Once
	file_endpoint_endpoint_proto_rawDescData = file_endpoint_endpoint_proto_rawDesc
)

func file_endpoint_endpoint_proto_rawDescGZIP() []byte {
	file_endpoint_endpoint_proto_rawDescOnce.Do(func() {
		file_endpoint_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_endpoint_endpoint_proto_rawDescData)
	})
	return file_endpoint_endpoint_proto_rawDescData
}

var file_endpoint_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_endpoint_endpoint_proto_goTypes = []interface{}{
	(*ServiceEndpoint)(nil),             // 0: ServiceEndpoint
	(*EndpointInNamespaceRequest)(nil),  // 1: EndpointInNamespaceRequest
	(*EndpointInNamespaceResponse)(nil), // 2: EndpointInNamespaceResponse
	(*EndpointInProjectRequest)(nil),    // 3: EndpointInProjectRequest
	(*EndpointInProjectResponse)(nil),   // 4: EndpointInProjectResponse
}
var file_endpoint_endpoint_proto_depIdxs = []int32{
	0, // 0: EndpointInNamespaceResponse.items:type_name -> ServiceEndpoint
	0, // 1: EndpointInProjectResponse.items:type_name -> ServiceEndpoint
	1, // 2: Endpoint.InNamespace:input_type -> EndpointInNamespaceRequest
	3, // 3: Endpoint.InProject:input_type -> EndpointInProjectRequest
	2, // 4: Endpoint.InNamespace:output_type -> EndpointInNamespaceResponse
	4, // 5: Endpoint.InProject:output_type -> EndpointInProjectResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_endpoint_endpoint_proto_init() }
func file_endpoint_endpoint_proto_init() {
	if File_endpoint_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_endpoint_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceEndpoint); i {
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
		file_endpoint_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointInNamespaceRequest); i {
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
		file_endpoint_endpoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointInNamespaceResponse); i {
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
		file_endpoint_endpoint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointInProjectRequest); i {
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
		file_endpoint_endpoint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointInProjectResponse); i {
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
			RawDescriptor: file_endpoint_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_endpoint_endpoint_proto_goTypes,
		DependencyIndexes: file_endpoint_endpoint_proto_depIdxs,
		MessageInfos:      file_endpoint_endpoint_proto_msgTypes,
	}.Build()
	File_endpoint_endpoint_proto = out.File
	file_endpoint_endpoint_proto_rawDesc = nil
	file_endpoint_endpoint_proto_goTypes = nil
	file_endpoint_endpoint_proto_depIdxs = nil
}
