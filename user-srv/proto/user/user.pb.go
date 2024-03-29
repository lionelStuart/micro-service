// Code generate	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)d by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package mu_micro_book_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	CreatedTime          uint64   `protobuf:"varint,4,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          uint64   `protobuf:"varint,5,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *User) GetCreatedTime() uint64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *User) GetUpdatedTime() uint64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserPwd              string   `protobuf:"bytes,3,opt,name=userPwd,proto3" json:"userPwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetUserPwd() string {
	if m != nil {
		return m.UserPwd
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "mu.micro.book.srv.user.user")
	proto.RegisterType((*Error)(nil), "mu.micro.book.srv.user.Error")
	proto.RegisterType((*Request)(nil), "mu.micro.book.srv.user.Request")
	proto.RegisterType((*Response)(nil), "mu.micro.book.srv.user.Response")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x75, 0xdb, 0xa4, 0x1f, 0x53, 0x50, 0x59, 0xb0, 0x84, 0xa2, 0x18, 0x72, 0xea, 0x69, 0x95,
	0xf6, 0x1f, 0x88, 0x1e, 0xbc, 0x88, 0x2e, 0x7e, 0xdc, 0x84, 0x36, 0x3b, 0x87, 0xa0, 0xe9, 0xc6,
	0xdd, 0xac, 0xd2, 0x9b, 0x77, 0xff, 0xb4, 0xcc, 0x24, 0xad, 0x3d, 0xd8, 0xcb, 0xf0, 0xde, 0xe4,
	0xcd, 0xc7, 0x9b, 0x2c, 0x9c, 0x54, 0xce, 0xd6, 0xf6, 0x22, 0x78, 0x74, 0x1c, 0x14, 0x73, 0x39,
	0x2e, 0x83, 0x2a, 0x8b, 0xdc, 0x59, 0xb5, 0xb4, 0xf6, 0x4d, 0x79, 0xf7, 0xa9, 0xe8, 0x6b, 0xf6,
	0x2d, 0x20, 0x22, 0x20, 0x0f, 0xa1, 0x53, 0x98, 0x44, 0xa4, 0x62, 0xda, 0xd5, 0x9d, 0xc2, 0x48,
	0x09, 0xd1, 0x6a, 0x51, 0x62, 0xd2, 0x49, 0xc5, 0x74, 0xa8, 0x19, 0xcb, 0x63, 0xe8, 0x56, 0x5f,
	0x26, 0xe9, 0x72, 0x8a, 0xa0, 0x4c, 0x61, 0x94, 0x3b, 0x5c, 0xd4, 0x68, 0x1e, 0x8b, 0x12, 0x93,
	0x28, 0x15, 0xd3, 0x48, 0xef, 0xa6, 0x48, 0x11, 0x2a, 0xb3, 0x55, 0xc4, 0x8d, 0x62, 0x27, 0x95,
	0xcd, 0x21, 0xbe, 0x71, 0xce, 0x3a, 0x1a, 0x99, 0x5b, 0x83, 0xbc, 0x44, 0xac, 0x19, 0xcb, 0x31,
	0xf4, 0x0c, 0xd6, 0x8b, 0xe2, 0xbd, 0x5d, 0xa4, 0x65, 0xd9, 0x0b, 0xf4, 0x35, 0x7e, 0x04, 0xf4,
	0x35, 0x49, 0xc8, 0xc1, 0xed, 0x35, 0x17, 0x0e, 0x75, 0xcb, 0xe4, 0x04, 0x06, 0x84, 0xee, 0xfe,
	0x5c, 0x6c, 0xb9, 0x4c, 0xa0, 0x4f, 0xf8, 0x7e, 0xeb, 0x66, 0x43, 0xb3, 0x1f, 0x01, 0x03, 0x8d,
	0xbe, 0xb2, 0x2b, 0xcf, 0x32, 0x1f, 0xf2, 0x1c, 0xbd, 0xe7, 0xde, 0x03, 0xbd, 0xa1, 0x72, 0x0e,
	0x31, 0xd2, 0xd2, 0xdc, 0x79, 0x34, 0x3b, 0x53, 0xff, 0xdf, 0x57, 0xb1, 0x33, 0xdd, 0x68, 0xe5,
	0x65, 0x73, 0x6b, 0x1e, 0x39, 0x9a, 0x9d, 0xee, 0xab, 0xa1, 0xa0, 0x59, 0x39, 0x7b, 0x85, 0xe8,
	0x89, 0xfe, 0xce, 0x33, 0x1c, 0x3d, 0x04, 0x74, 0x6b, 0x22, 0x57, 0x6b, 0xb6, 0x70, 0xbe, 0xaf,
	0xbc, 0xbd, 0xcb, 0x24, 0xdd, 0x2f, 0x68, 0xec, 0x65, 0x07, 0xcb, 0x1e, 0xbf, 0x8e, 0xf9, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0x92, 0xed, 0x80, 0x36, 0x02, 0x00, 0x00,
}
