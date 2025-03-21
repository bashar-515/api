// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: app/mlinference/v1/ml_inference.proto

package v1

import (
	v1 "go.viam.com/api/app/data/v1"
	v11 "go.viam.com/api/service/mlmodel/v1"
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

type GetInferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The model framework and model type are inferred from the ML model registry item;
	// For valid model types (classification, detections) we will return the formatted
	// labels or annotations from the associated inference outputs.
	RegistryItemId      string       `protobuf:"bytes,1,opt,name=registry_item_id,json=registryItemId,proto3" json:"registry_item_id,omitempty"`
	RegistryItemVersion string       `protobuf:"bytes,2,opt,name=registry_item_version,json=registryItemVersion,proto3" json:"registry_item_version,omitempty"`
	BinaryId            *v1.BinaryID `protobuf:"bytes,3,opt,name=binary_id,json=binaryId,proto3" json:"binary_id,omitempty"`
	BinaryDataId        string       `protobuf:"bytes,5,opt,name=binary_data_id,json=binaryDataId,proto3" json:"binary_data_id,omitempty"`
	OrganizationId      string       `protobuf:"bytes,4,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *GetInferenceRequest) Reset() {
	*x = GetInferenceRequest{}
	mi := &file_app_mlinference_v1_ml_inference_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInferenceRequest) ProtoMessage() {}

func (x *GetInferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_mlinference_v1_ml_inference_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInferenceRequest.ProtoReflect.Descriptor instead.
func (*GetInferenceRequest) Descriptor() ([]byte, []int) {
	return file_app_mlinference_v1_ml_inference_proto_rawDescGZIP(), []int{0}
}

func (x *GetInferenceRequest) GetRegistryItemId() string {
	if x != nil {
		return x.RegistryItemId
	}
	return ""
}

func (x *GetInferenceRequest) GetRegistryItemVersion() string {
	if x != nil {
		return x.RegistryItemVersion
	}
	return ""
}

func (x *GetInferenceRequest) GetBinaryId() *v1.BinaryID {
	if x != nil {
		return x.BinaryId
	}
	return nil
}

func (x *GetInferenceRequest) GetBinaryDataId() string {
	if x != nil {
		return x.BinaryDataId
	}
	return ""
}

func (x *GetInferenceRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

type GetInferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutputTensors *v11.FlatTensors `protobuf:"bytes,1,opt,name=output_tensors,json=outputTensors,proto3" json:"output_tensors,omitempty"`
	Annotations   *v1.Annotations  `protobuf:"bytes,2,opt,name=annotations,proto3" json:"annotations,omitempty"`
}

func (x *GetInferenceResponse) Reset() {
	*x = GetInferenceResponse{}
	mi := &file_app_mlinference_v1_ml_inference_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInferenceResponse) ProtoMessage() {}

func (x *GetInferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_mlinference_v1_ml_inference_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInferenceResponse.ProtoReflect.Descriptor instead.
func (*GetInferenceResponse) Descriptor() ([]byte, []int) {
	return file_app_mlinference_v1_ml_inference_proto_rawDescGZIP(), []int{1}
}

func (x *GetInferenceResponse) GetOutputTensors() *v11.FlatTensors {
	if x != nil {
		return x.OutputTensors
	}
	return nil
}

func (x *GetInferenceResponse) GetAnnotations() *v1.Annotations {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_app_mlinference_v1_ml_inference_proto protoreflect.FileDescriptor

var file_app_mlinference_v1_ml_inference_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x6c, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x6d, 0x6c, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x16, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x6d, 0x6c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6c, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x01, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x37, 0x0a, 0x09, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x44, 0x52,
	0x08, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x49, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x69, 0x61, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6c, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x61, 0x74, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x52,
	0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x3f,
	0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32,
	0x81, 0x01, 0x0a, 0x12, 0x4d, 0x4c, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2c, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x6d, 0x6c, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x6d, 0x6c, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x6c, 0x69, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_app_mlinference_v1_ml_inference_proto_rawDescOnce sync.Once
	file_app_mlinference_v1_ml_inference_proto_rawDescData = file_app_mlinference_v1_ml_inference_proto_rawDesc
)

func file_app_mlinference_v1_ml_inference_proto_rawDescGZIP() []byte {
	file_app_mlinference_v1_ml_inference_proto_rawDescOnce.Do(func() {
		file_app_mlinference_v1_ml_inference_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_mlinference_v1_ml_inference_proto_rawDescData)
	})
	return file_app_mlinference_v1_ml_inference_proto_rawDescData
}

var file_app_mlinference_v1_ml_inference_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_mlinference_v1_ml_inference_proto_goTypes = []any{
	(*GetInferenceRequest)(nil),  // 0: viam.app.mlinference.v1.GetInferenceRequest
	(*GetInferenceResponse)(nil), // 1: viam.app.mlinference.v1.GetInferenceResponse
	(*v1.BinaryID)(nil),          // 2: viam.app.data.v1.BinaryID
	(*v11.FlatTensors)(nil),      // 3: viam.service.mlmodel.v1.FlatTensors
	(*v1.Annotations)(nil),       // 4: viam.app.data.v1.Annotations
}
var file_app_mlinference_v1_ml_inference_proto_depIdxs = []int32{
	2, // 0: viam.app.mlinference.v1.GetInferenceRequest.binary_id:type_name -> viam.app.data.v1.BinaryID
	3, // 1: viam.app.mlinference.v1.GetInferenceResponse.output_tensors:type_name -> viam.service.mlmodel.v1.FlatTensors
	4, // 2: viam.app.mlinference.v1.GetInferenceResponse.annotations:type_name -> viam.app.data.v1.Annotations
	0, // 3: viam.app.mlinference.v1.MLInferenceService.GetInference:input_type -> viam.app.mlinference.v1.GetInferenceRequest
	1, // 4: viam.app.mlinference.v1.MLInferenceService.GetInference:output_type -> viam.app.mlinference.v1.GetInferenceResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_app_mlinference_v1_ml_inference_proto_init() }
func file_app_mlinference_v1_ml_inference_proto_init() {
	if File_app_mlinference_v1_ml_inference_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_mlinference_v1_ml_inference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_mlinference_v1_ml_inference_proto_goTypes,
		DependencyIndexes: file_app_mlinference_v1_ml_inference_proto_depIdxs,
		MessageInfos:      file_app_mlinference_v1_ml_inference_proto_msgTypes,
	}.Build()
	File_app_mlinference_v1_ml_inference_proto = out.File
	file_app_mlinference_v1_ml_inference_proto_rawDesc = nil
	file_app_mlinference_v1_ml_inference_proto_goTypes = nil
	file_app_mlinference_v1_ml_inference_proto_depIdxs = nil
}
