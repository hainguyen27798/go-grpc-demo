// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/mail.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendMailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	To            []string               `protobuf:"bytes,1,rep,name=to,proto3" json:"to,omitempty"`
	Subject       string                 `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Text          *string                `protobuf:"bytes,3,opt,name=text,proto3,oneof" json:"text,omitempty"`
	TemplateName  *string                `protobuf:"bytes,4,opt,name=templateName,proto3,oneof" json:"templateName,omitempty"`
	TemplateData  map[string]string      `protobuf:"bytes,5,rep,name=templateData,proto3" json:"templateData,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMailRequest) Reset() {
	*x = SendMailRequest{}
	mi := &file_proto_mail_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailRequest) ProtoMessage() {}

func (x *SendMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mail_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailRequest.ProtoReflect.Descriptor instead.
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return file_proto_mail_proto_rawDescGZIP(), []int{0}
}

func (x *SendMailRequest) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SendMailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendMailRequest) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

func (x *SendMailRequest) GetTemplateName() string {
	if x != nil && x.TemplateName != nil {
		return *x.TemplateName
	}
	return ""
}

func (x *SendMailRequest) GetTemplateData() map[string]string {
	if x != nil {
		return x.TemplateData
	}
	return nil
}

type SendMailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMailResponse) Reset() {
	*x = SendMailResponse{}
	mi := &file_proto_mail_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailResponse) ProtoMessage() {}

func (x *SendMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mail_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailResponse.ProtoReflect.Descriptor instead.
func (*SendMailResponse) Descriptor() ([]byte, []int) {
	return file_proto_mail_proto_rawDescGZIP(), []int{1}
}

func (x *SendMailResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *SendMailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_mail_proto protoreflect.FileDescriptor

const file_proto_mail_proto_rawDesc = "" +
	"\n" +
	"\x10proto/mail.proto\x12\x05proto\"\xa6\x02\n" +
	"\x0fSendMailRequest\x12\x0e\n" +
	"\x02to\x18\x01 \x03(\tR\x02to\x12\x18\n" +
	"\asubject\x18\x02 \x01(\tR\asubject\x12\x17\n" +
	"\x04text\x18\x03 \x01(\tH\x00R\x04text\x88\x01\x01\x12'\n" +
	"\ftemplateName\x18\x04 \x01(\tH\x01R\ftemplateName\x88\x01\x01\x12L\n" +
	"\ftemplateData\x18\x05 \x03(\v2(.proto.SendMailRequest.TemplateDataEntryR\ftemplateData\x1a?\n" +
	"\x11TemplateDataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\a\n" +
	"\x05_textB\x0f\n" +
	"\r_templateName\"L\n" +
	"\x10SendMailResponse\x12\x1e\n" +
	"\n" +
	"statusCode\x18\x01 \x01(\x05R\n" +
	"statusCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2F\n" +
	"\vMailService\x127\n" +
	"\x04Send\x12\x16.proto.SendMailRequest\x1a\x17.proto.SendMailResponseB\n" +
	"Z\b./;protob\x06proto3"

var (
	file_proto_mail_proto_rawDescOnce sync.Once
	file_proto_mail_proto_rawDescData []byte
)

func file_proto_mail_proto_rawDescGZIP() []byte {
	file_proto_mail_proto_rawDescOnce.Do(func() {
		file_proto_mail_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_mail_proto_rawDesc), len(file_proto_mail_proto_rawDesc)))
	})
	return file_proto_mail_proto_rawDescData
}

var file_proto_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_mail_proto_goTypes = []any{
	(*SendMailRequest)(nil),  // 0: proto.SendMailRequest
	(*SendMailResponse)(nil), // 1: proto.SendMailResponse
	nil,                      // 2: proto.SendMailRequest.TemplateDataEntry
}
var file_proto_mail_proto_depIdxs = []int32{
	2, // 0: proto.SendMailRequest.templateData:type_name -> proto.SendMailRequest.TemplateDataEntry
	0, // 1: proto.MailService.Send:input_type -> proto.SendMailRequest
	1, // 2: proto.MailService.Send:output_type -> proto.SendMailResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_mail_proto_init() }
func file_proto_mail_proto_init() {
	if File_proto_mail_proto != nil {
		return
	}
	file_proto_mail_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_mail_proto_rawDesc), len(file_proto_mail_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mail_proto_goTypes,
		DependencyIndexes: file_proto_mail_proto_depIdxs,
		MessageInfos:      file_proto_mail_proto_msgTypes,
	}.Build()
	File_proto_mail_proto = out.File
	file_proto_mail_proto_goTypes = nil
	file_proto_mail_proto_depIdxs = nil
}
