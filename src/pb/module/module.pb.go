// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: src/pb/module/module.proto

package module

import (
	ebpf "github.com/tricorder/src/pb/module/ebpf"
	wasm "github.com/tricorder/src/pb/module/wasm"
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

type Module_TransmissionParadigm int32

const (
	Module_PER_EVENT Module_TransmissionParadigm = 0
)

// Enum value maps for Module_TransmissionParadigm.
var (
	Module_TransmissionParadigm_name = map[int32]string{
		0: "PER_EVENT",
	}
	Module_TransmissionParadigm_value = map[string]int32{
		"PER_EVENT": 0,
	}
)

func (x Module_TransmissionParadigm) Enum() *Module_TransmissionParadigm {
	p := new(Module_TransmissionParadigm)
	*p = x
	return p
}

func (x Module_TransmissionParadigm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Module_TransmissionParadigm) Descriptor() protoreflect.EnumDescriptor {
	return file_src_pb_module_module_proto_enumTypes[0].Descriptor()
}

func (Module_TransmissionParadigm) Type() protoreflect.EnumType {
	return &file_src_pb_module_module_proto_enumTypes[0]
}

func (x Module_TransmissionParadigm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Module_TransmissionParadigm.Descriptor instead.
func (Module_TransmissionParadigm) EnumDescriptor() ([]byte, []int) {
	return file_src_pb_module_module_proto_rawDescGZIP(), []int{0, 0}
}

type Module_EncodingParadigm int32

const (
	Module_NONE Module_EncodingParadigm = 0
	Module_TLV  Module_EncodingParadigm = 1
	Module_JSON Module_EncodingParadigm = 2
)

// Enum value maps for Module_EncodingParadigm.
var (
	Module_EncodingParadigm_name = map[int32]string{
		0: "NONE",
		1: "TLV",
		2: "JSON",
	}
	Module_EncodingParadigm_value = map[string]int32{
		"NONE": 0,
		"TLV":  1,
		"JSON": 2,
	}
)

func (x Module_EncodingParadigm) Enum() *Module_EncodingParadigm {
	p := new(Module_EncodingParadigm)
	*p = x
	return p
}

func (x Module_EncodingParadigm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Module_EncodingParadigm) Descriptor() protoreflect.EnumDescriptor {
	return file_src_pb_module_module_proto_enumTypes[1].Descriptor()
}

func (Module_EncodingParadigm) Type() protoreflect.EnumType {
	return &file_src_pb_module_module_proto_enumTypes[1]
}

func (x Module_EncodingParadigm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Module_EncodingParadigm.Descriptor instead.
func (Module_EncodingParadigm) EnumDescriptor() ([]byte, []int) {
	return file_src_pb_module_module_proto_rawDescGZIP(), []int{0, 1}
}

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ebpf               *ebpf.Program           `protobuf:"bytes,2,opt,name=ebpf,proto3" json:"ebpf,omitempty"`
	Wasm               *wasm.Program           `protobuf:"bytes,3,opt,name=wasm,proto3" json:"wasm,omitempty"`
	WasmOutputEncoding Module_EncodingParadigm `protobuf:"varint,4,opt,name=wasm_output_encoding,json=wasmOutputEncoding,proto3,enum=tricorder.pb.module.Module_EncodingParadigm" json:"wasm_output_encoding,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pb_module_module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_src_pb_module_module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_src_pb_module_module_proto_rawDescGZIP(), []int{0}
}

func (x *Module) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Module) GetEbpf() *ebpf.Program {
	if x != nil {
		return x.Ebpf
	}
	return nil
}

func (x *Module) GetWasm() *wasm.Program {
	if x != nil {
		return x.Wasm
	}
	return nil
}

func (x *Module) GetWasmOutputEncoding() Module_EncodingParadigm {
	if x != nil {
		return x.WasmOutputEncoding
	}
	return Module_NONE
}

var File_src_pb_module_module_proto protoreflect.FileDescriptor

var file_src_pb_module_module_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x72,
	0x69, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x1a, 0x1d, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x65, 0x62, 0x70, 0x66, 0x2f, 0x65, 0x62, 0x70, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f,
	0x77, 0x61, 0x73, 0x6d, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc2, 0x02, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x65, 0x62, 0x70, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74,
	0x72, 0x69, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x65, 0x62, 0x70, 0x66, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52,
	0x04, 0x65, 0x62, 0x70, 0x66, 0x12, 0x35, 0x0a, 0x04, 0x77, 0x61, 0x73, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x04, 0x77, 0x61, 0x73, 0x6d, 0x12, 0x5e, 0x0a, 0x14,
	0x77, 0x61, 0x73, 0x6d, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x74, 0x72, 0x69,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x72, 0x61, 0x64, 0x69, 0x67, 0x6d, 0x52, 0x12, 0x77, 0x61, 0x73, 0x6d, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x25, 0x0a, 0x14,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x64, 0x69, 0x67, 0x6d, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x45, 0x52, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x10, 0x00, 0x22, 0x2f, 0x0a, 0x10, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x50,
	0x61, 0x72, 0x61, 0x64, 0x69, 0x67, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x4c, 0x56, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53,
	0x4f, 0x4e, 0x10, 0x02, 0x42, 0x08, 0x5a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_pb_module_module_proto_rawDescOnce sync.Once
	file_src_pb_module_module_proto_rawDescData = file_src_pb_module_module_proto_rawDesc
)

func file_src_pb_module_module_proto_rawDescGZIP() []byte {
	file_src_pb_module_module_proto_rawDescOnce.Do(func() {
		file_src_pb_module_module_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_pb_module_module_proto_rawDescData)
	})
	return file_src_pb_module_module_proto_rawDescData
}

var file_src_pb_module_module_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_src_pb_module_module_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_src_pb_module_module_proto_goTypes = []interface{}{
	(Module_TransmissionParadigm)(0), // 0: tricorder.pb.module.Module.TransmissionParadigm
	(Module_EncodingParadigm)(0),     // 1: tricorder.pb.module.Module.EncodingParadigm
	(*Module)(nil),                   // 2: tricorder.pb.module.Module
	(*ebpf.Program)(nil),             // 3: tricorder.pb.module.ebpf.Program
	(*wasm.Program)(nil),             // 4: tricorder.pb.module.wasm.Program
}
var file_src_pb_module_module_proto_depIdxs = []int32{
	3, // 0: tricorder.pb.module.Module.ebpf:type_name -> tricorder.pb.module.ebpf.Program
	4, // 1: tricorder.pb.module.Module.wasm:type_name -> tricorder.pb.module.wasm.Program
	1, // 2: tricorder.pb.module.Module.wasm_output_encoding:type_name -> tricorder.pb.module.Module.EncodingParadigm
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_src_pb_module_module_proto_init() }
func file_src_pb_module_module_proto_init() {
	if File_src_pb_module_module_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_pb_module_module_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
			RawDescriptor: file_src_pb_module_module_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_pb_module_module_proto_goTypes,
		DependencyIndexes: file_src_pb_module_module_proto_depIdxs,
		EnumInfos:         file_src_pb_module_module_proto_enumTypes,
		MessageInfos:      file_src_pb_module_module_proto_msgTypes,
	}.Build()
	File_src_pb_module_module_proto = out.File
	file_src_pb_module_module_proto_rawDesc = nil
	file_src_pb_module_module_proto_goTypes = nil
	file_src_pb_module_module_proto_depIdxs = nil
}
