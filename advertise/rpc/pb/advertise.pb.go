// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: advertise.proto

package pb

import (
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

type Advertise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" db:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" db:"id"`
	// @inject_tag: json:"name" db:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" db:"name"`
	// @inject_tag: json:"img" db:"img"
	Img string `protobuf:"bytes,3,opt,name=img,proto3" json:"img" db:"img"`
	// @inject_tag: json:"pos" db:"pos"
	Pos int32 `protobuf:"varint,4,opt,name=pos,proto3" json:"pos" db:"pos"`
	// @inject_tag: json:"category_id" db:"category_id"
	CategoryId int64 `protobuf:"varint,5,opt,name=category_id,json=categoryId,proto3" json:"category_id" db:"category_id"`
	// @inject_tag: json:"type" db:"type"
	Type int32 `protobuf:"varint,6,opt,name=type,proto3" json:"type" db:"type"`
	// @inject_tag: json:"url" db:"url"
	Url string `protobuf:"bytes,7,opt,name=url,proto3" json:"url" db:"url"`
	// @inject_tag: json:"go_id" db:"go_id"
	GoId int32 `protobuf:"varint,8,opt,name=go_id,json=goId,proto3" json:"go_id" db:"go_id"`
	// @inject_tag: json:"status" db:"status"
	Status int32 `protobuf:"varint,9,opt,name=status,proto3" json:"status" db:"status"`
	// @inject_tag: json:"start_time" db:"start_time"
	StartTime string `protobuf:"bytes,10,opt,name=start_time,json=startTime,proto3" json:"start_time" db:"start_time"`
	// @inject_tag: json:"end_time" db:"end_time"
	EndTime string `protobuf:"bytes,11,opt,name=end_time,json=endTime,proto3" json:"end_time" db:"end_time"`
	// @inject_tag: json:"add_time" db:"add_time"
	AddTime string `protobuf:"bytes,12,opt,name=add_time,json=addTime,proto3" json:"add_time" db:"add_time"`
}

func (x *Advertise) Reset() {
	*x = Advertise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advertise_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Advertise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Advertise) ProtoMessage() {}

func (x *Advertise) ProtoReflect() protoreflect.Message {
	mi := &file_advertise_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Advertise.ProtoReflect.Descriptor instead.
func (*Advertise) Descriptor() ([]byte, []int) {
	return file_advertise_proto_rawDescGZIP(), []int{0}
}

func (x *Advertise) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Advertise) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Advertise) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *Advertise) GetPos() int32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

func (x *Advertise) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *Advertise) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Advertise) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Advertise) GetGoId() int32 {
	if x != nil {
		return x.GoId
	}
	return 0
}

func (x *Advertise) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Advertise) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *Advertise) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *Advertise) GetAddTime() string {
	if x != nil {
		return x.AddTime
	}
	return ""
}

type AdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: form:"pos"
	Pos int32 `protobuf:"varint,1,opt,name=pos,proto3" json:"pos,omitempty" form:"pos"` //
	// @inject_tag: form:"category_id"
	CategoryId int64 `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty" form:"category_id"`
	// @inject_tag: form:"status"
	Status int32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty" form:"status"`
	// @inject_tag: form:"page_size"
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" form:"page_size"`
	// @inject_tag: form:"page_num"
	PageNum int32 `protobuf:"varint,5,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty" form:"page_num"`
	// @inject_tag: form:"start_time"
	StartTime string `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" form:"start_time"`
	// @inject_tag: form:"end_time"
	EndTime string `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" form:"end_time"`
	// @inject_tag: form:"name"
	Name string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty" form:"name"`
}

func (x *AdRequest) Reset() {
	*x = AdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advertise_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdRequest) ProtoMessage() {}

func (x *AdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advertise_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdRequest.ProtoReflect.Descriptor instead.
func (*AdRequest) Descriptor() ([]byte, []int) {
	return file_advertise_proto_rawDescGZIP(), []int{1}
}

func (x *AdRequest) GetPos() int32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

func (x *AdRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *AdRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AdRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AdRequest) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *AdRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *AdRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *AdRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AdOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: form:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id"`
}

func (x *AdOneRequest) Reset() {
	*x = AdOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advertise_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdOneRequest) ProtoMessage() {}

func (x *AdOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advertise_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdOneRequest.ProtoReflect.Descriptor instead.
func (*AdOneRequest) Descriptor() ([]byte, []int) {
	return file_advertise_proto_rawDescGZIP(), []int{2}
}

func (x *AdOneRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"data"
	Data []*Advertise `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
	// @inject_tag: json:"total"
	Total int64 `protobuf:"varint,3,opt,name=total,proto3" json:"total"`
	// @inject_tag: json:"detail"
	Detail string `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail"`
}

func (x *AdResponse) Reset() {
	*x = AdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advertise_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdResponse) ProtoMessage() {}

func (x *AdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_advertise_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdResponse.ProtoReflect.Descriptor instead.
func (*AdResponse) Descriptor() ([]byte, []int) {
	return file_advertise_proto_rawDescGZIP(), []int{3}
}

func (x *AdResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AdResponse) GetData() []*Advertise {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AdResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *AdResponse) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type AdInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"data"
	Data *Advertise `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
	// @inject_tag: json:"detail"
	Detail string `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail"`
}

func (x *AdInfoResponse) Reset() {
	*x = AdInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advertise_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdInfoResponse) ProtoMessage() {}

func (x *AdInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_advertise_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdInfoResponse.ProtoReflect.Descriptor instead.
func (*AdInfoResponse) Descriptor() ([]byte, []int) {
	return file_advertise_proto_rawDescGZIP(), []int{4}
}

func (x *AdInfoResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AdInfoResponse) GetData() *Advertise {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AdInfoResponse) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"detail"
	Detail string `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advertise_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_advertise_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_advertise_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

var File_advertise_proto protoreflect.FileDescriptor

var file_advertise_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x9c, 0x02, 0x0a, 0x09, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x13, 0x0a, 0x05, 0x67, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x67, 0x6f, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0xdc, 0x01, 0x0a, 0x09, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x70, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x41, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x71, 0x0a, 0x0a, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x5f, 0x0a, 0x0e, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e,
	0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x32,
	0x94, 0x02, 0x0a, 0x10, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_advertise_proto_rawDescOnce sync.Once
	file_advertise_proto_rawDescData = file_advertise_proto_rawDesc
)

func file_advertise_proto_rawDescGZIP() []byte {
	file_advertise_proto_rawDescOnce.Do(func() {
		file_advertise_proto_rawDescData = protoimpl.X.CompressGZIP(file_advertise_proto_rawDescData)
	})
	return file_advertise_proto_rawDescData
}

var file_advertise_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_advertise_proto_goTypes = []interface{}{
	(*Advertise)(nil),      // 0: pb.Advertise
	(*AdRequest)(nil),      // 1: pb.AdRequest
	(*AdOneRequest)(nil),   // 2: pb.AdOneRequest
	(*AdResponse)(nil),     // 3: pb.AdResponse
	(*AdInfoResponse)(nil), // 4: pb.AdInfoResponse
	(*Response)(nil),       // 5: pb.Response
}
var file_advertise_proto_depIdxs = []int32{
	0, // 0: pb.AdResponse.data:type_name -> pb.Advertise
	0, // 1: pb.AdInfoResponse.data:type_name -> pb.Advertise
	1, // 2: pb.AdvertiseService.AdminGetList:input_type -> pb.AdRequest
	1, // 3: pb.AdvertiseService.ApiGetList:input_type -> pb.AdRequest
	2, // 4: pb.AdvertiseService.GetInfo:input_type -> pb.AdOneRequest
	0, // 5: pb.AdvertiseService.Add:input_type -> pb.Advertise
	2, // 6: pb.AdvertiseService.Delete:input_type -> pb.AdOneRequest
	0, // 7: pb.AdvertiseService.Update:input_type -> pb.Advertise
	3, // 8: pb.AdvertiseService.AdminGetList:output_type -> pb.AdResponse
	3, // 9: pb.AdvertiseService.ApiGetList:output_type -> pb.AdResponse
	4, // 10: pb.AdvertiseService.GetInfo:output_type -> pb.AdInfoResponse
	5, // 11: pb.AdvertiseService.Add:output_type -> pb.Response
	5, // 12: pb.AdvertiseService.Delete:output_type -> pb.Response
	5, // 13: pb.AdvertiseService.Update:output_type -> pb.Response
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_advertise_proto_init() }
func file_advertise_proto_init() {
	if File_advertise_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_advertise_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Advertise); i {
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
		file_advertise_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdRequest); i {
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
		file_advertise_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdOneRequest); i {
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
		file_advertise_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdResponse); i {
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
		file_advertise_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdInfoResponse); i {
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
		file_advertise_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_advertise_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_advertise_proto_goTypes,
		DependencyIndexes: file_advertise_proto_depIdxs,
		MessageInfos:      file_advertise_proto_msgTypes,
	}.Build()
	File_advertise_proto = out.File
	file_advertise_proto_rawDesc = nil
	file_advertise_proto_goTypes = nil
	file_advertise_proto_depIdxs = nil
}
