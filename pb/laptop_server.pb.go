// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: laptop_server.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateLaptopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptop *Laptop `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *CreateLaptopRequest) Reset() {
	*x = CreateLaptopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLaptopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLaptopRequest) ProtoMessage() {}

func (x *CreateLaptopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLaptopRequest.ProtoReflect.Descriptor instead.
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return file_laptop_server_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLaptopRequest) GetLaptop() *Laptop {
	if x != nil {
		return x.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateLaptopResponse) Reset() {
	*x = CreateLaptopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLaptopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLaptopResponse) ProtoMessage() {}

func (x *CreateLaptopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLaptopResponse.ProtoReflect.Descriptor instead.
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return file_laptop_server_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLaptopResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SearchLaptopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *SearchLaptopRequest) Reset() {
	*x = SearchLaptopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLaptopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLaptopRequest) ProtoMessage() {}

func (x *SearchLaptopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLaptopRequest.ProtoReflect.Descriptor instead.
func (*SearchLaptopRequest) Descriptor() ([]byte, []int) {
	return file_laptop_server_proto_rawDescGZIP(), []int{2}
}

func (x *SearchLaptopRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type SearchLaptopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptop *Laptop `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *SearchLaptopResponse) Reset() {
	*x = SearchLaptopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLaptopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLaptopResponse) ProtoMessage() {}

func (x *SearchLaptopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLaptopResponse.ProtoReflect.Descriptor instead.
func (*SearchLaptopResponse) Descriptor() ([]byte, []int) {
	return file_laptop_server_proto_rawDescGZIP(), []int{3}
}

func (x *SearchLaptopResponse) GetLaptop() *Laptop {
	if x != nil {
		return x.Laptop
	}
	return nil
}

type UploadImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*UploadImageRequest_Info
	//	*UploadImageRequest_ChunkData
	Data isUploadImageRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_laptop_server_proto_rawDescGZIP(), []int{4}
}

func (m *UploadImageRequest) GetData() isUploadImageRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadImageRequest) GetInfo() *ImageInfo {
	if x, ok := x.GetData().(*UploadImageRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadImageRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*UploadImageRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isUploadImageRequest_Data interface {
	isUploadImageRequest_Data()
}

type UploadImageRequest_Info struct {
	Info *ImageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"` //第一个请求中将只包含元数据。
}

type UploadImageRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"` //后面的请求包含图像数据块。
}

func (*UploadImageRequest_Info) isUploadImageRequest_Data() {}

func (*UploadImageRequest_ChunkData) isUploadImageRequest_Data() {}

type ImageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaptopId  string `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`    //电脑id
	ImageType string `protobuf:"bytes,2,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"` //图像类型如.jpg或.png
}

func (x *ImageInfo) Reset() {
	*x = ImageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfo) ProtoMessage() {}

func (x *ImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfo.ProtoReflect.Descriptor instead.
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return file_laptop_server_proto_rawDescGZIP(), []int{5}
}

func (x *ImageInfo) GetLaptopId() string {
	if x != nil {
		return x.LaptopId
	}
	return ""
}

func (x *ImageInfo) GetImageType() string {
	if x != nil {
		return x.ImageType
	}
	return ""
}

type UploadImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`      //生成的ID。
	Size uint32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"` //图像字节大小
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_laptop_server_proto_rawDescGZIP(), []int{6}
}

func (x *UploadImageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UploadImageResponse) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type RateLaptopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaptopId string  `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	Score    float64 `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"` //我们将为客户端编写一个API,以从1~10的分数对电脑流进行评分。服务器将响应每台笔记本电脑的平均分数流
}

func (x *RateLaptopRequest) Reset() {
	*x = RateLaptopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLaptopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLaptopRequest) ProtoMessage() {}

func (x *RateLaptopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLaptopRequest.ProtoReflect.Descriptor instead.
func (*RateLaptopRequest) Descriptor() ([]byte, []int) {
	return file_laptop_server_proto_rawDescGZIP(), []int{7}
}

func (x *RateLaptopRequest) GetLaptopId() string {
	if x != nil {
		return x.LaptopId
	}
	return ""
}

func (x *RateLaptopRequest) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type RateLaptopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaptopId     string  `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	RatedCount   uint32  `protobuf:"varint,2,opt,name=rated_count,json=ratedCount,proto3" json:"rated_count,omitempty"`        //这台电脑被服务器评分的次数
	AverageScore float64 `protobuf:"fixed64,3,opt,name=average_score,json=averageScore,proto3" json:"average_score,omitempty"` //平均评分
}

func (x *RateLaptopResponse) Reset() {
	*x = RateLaptopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLaptopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLaptopResponse) ProtoMessage() {}

func (x *RateLaptopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLaptopResponse.ProtoReflect.Descriptor instead.
func (*RateLaptopResponse) Descriptor() ([]byte, []int) {
	return file_laptop_server_proto_rawDescGZIP(), []int{8}
}

func (x *RateLaptopResponse) GetLaptopId() string {
	if x != nil {
		return x.LaptopId
	}
	return ""
}

func (x *RateLaptopResponse) GetRatedCount() uint32 {
	if x != nil {
		return x.RatedCount
	}
	return 0
}

func (x *RateLaptopResponse) GetAverageScore() float64 {
	if x != nil {
		return x.AverageScore
	}
	return 0
}

var File_laptop_server_proto protoreflect.FileDescriptor

var file_laptop_server_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x14, 0x6c, 0x61, 0x70, 0x74, 0x6f,
	0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x06,
	0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70,
	0x22, 0x26, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x6c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x22,
	0x62, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x39, 0x0a, 0x13,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x46, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22,
	0x77, 0x0a, 0x12, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x61, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x32, 0xa2, 0x02, 0x0a, 0x0d, 0x4c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12,
	0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x42, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x41, 0x0a, 0x0a, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x05, 0x5a,
	0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_laptop_server_proto_rawDescOnce sync.Once
	file_laptop_server_proto_rawDescData = file_laptop_server_proto_rawDesc
)

func file_laptop_server_proto_rawDescGZIP() []byte {
	file_laptop_server_proto_rawDescOnce.Do(func() {
		file_laptop_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_laptop_server_proto_rawDescData)
	})
	return file_laptop_server_proto_rawDescData
}

var file_laptop_server_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_laptop_server_proto_goTypes = []interface{}{
	(*CreateLaptopRequest)(nil),  // 0: pb.CreateLaptopRequest
	(*CreateLaptopResponse)(nil), // 1: pb.CreateLaptopResponse
	(*SearchLaptopRequest)(nil),  // 2: pb.SearchLaptopRequest
	(*SearchLaptopResponse)(nil), // 3: pb.SearchLaptopResponse
	(*UploadImageRequest)(nil),   // 4: pb.UploadImageRequest
	(*ImageInfo)(nil),            // 5: pb.ImageInfo
	(*UploadImageResponse)(nil),  // 6: pb.UploadImageResponse
	(*RateLaptopRequest)(nil),    // 7: pb.RateLaptopRequest
	(*RateLaptopResponse)(nil),   // 8: pb.RateLaptopResponse
	(*Laptop)(nil),               // 9: pb.Laptop
	(*Filter)(nil),               // 10: pb.Filter
}
var file_laptop_server_proto_depIdxs = []int32{
	9,  // 0: pb.CreateLaptopRequest.laptop:type_name -> pb.Laptop
	10, // 1: pb.SearchLaptopRequest.filter:type_name -> pb.Filter
	9,  // 2: pb.SearchLaptopResponse.laptop:type_name -> pb.Laptop
	5,  // 3: pb.UploadImageRequest.info:type_name -> pb.ImageInfo
	0,  // 4: pb.LaptopService.CreateLaptop:input_type -> pb.CreateLaptopRequest
	2,  // 5: pb.LaptopService.SearchLaptop:input_type -> pb.SearchLaptopRequest
	4,  // 6: pb.LaptopService.UploadImage:input_type -> pb.UploadImageRequest
	7,  // 7: pb.LaptopService.RateLaptop:input_type -> pb.RateLaptopRequest
	1,  // 8: pb.LaptopService.CreateLaptop:output_type -> pb.CreateLaptopResponse
	3,  // 9: pb.LaptopService.SearchLaptop:output_type -> pb.SearchLaptopResponse
	6,  // 10: pb.LaptopService.UploadImage:output_type -> pb.UploadImageResponse
	8,  // 11: pb.LaptopService.RateLaptop:output_type -> pb.RateLaptopResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_laptop_server_proto_init() }
func file_laptop_server_proto_init() {
	if File_laptop_server_proto != nil {
		return
	}
	file_laptop_message_proto_init()
	file_filter_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_laptop_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLaptopRequest); i {
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
		file_laptop_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLaptopResponse); i {
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
		file_laptop_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchLaptopRequest); i {
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
		file_laptop_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchLaptopResponse); i {
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
		file_laptop_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageRequest); i {
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
		file_laptop_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageInfo); i {
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
		file_laptop_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageResponse); i {
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
		file_laptop_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLaptopRequest); i {
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
		file_laptop_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLaptopResponse); i {
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
	file_laptop_server_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*UploadImageRequest_Info)(nil),
		(*UploadImageRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_laptop_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_laptop_server_proto_goTypes,
		DependencyIndexes: file_laptop_server_proto_depIdxs,
		MessageInfos:      file_laptop_server_proto_msgTypes,
	}.Build()
	File_laptop_server_proto = out.File
	file_laptop_server_proto_rawDesc = nil
	file_laptop_server_proto_goTypes = nil
	file_laptop_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LaptopServiceClient is the client API for LaptopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LaptopServiceClient interface {
	CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error)
	SearchLaptop(ctx context.Context, in *SearchLaptopRequest, opts ...grpc.CallOption) (LaptopService_SearchLaptopClient, error)
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (LaptopService_UploadImageClient, error)
	RateLaptop(ctx context.Context, opts ...grpc.CallOption) (LaptopService_RateLaptopClient, error)
}

type laptopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLaptopServiceClient(cc grpc.ClientConnInterface) LaptopServiceClient {
	return &laptopServiceClient{cc}
}

func (c *laptopServiceClient) CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error) {
	out := new(CreateLaptopResponse)
	err := c.cc.Invoke(ctx, "/pb.LaptopService/CreateLaptop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laptopServiceClient) SearchLaptop(ctx context.Context, in *SearchLaptopRequest, opts ...grpc.CallOption) (LaptopService_SearchLaptopClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LaptopService_serviceDesc.Streams[0], "/pb.LaptopService/SearchLaptop", opts...)
	if err != nil {
		return nil, err
	}
	x := &laptopServiceSearchLaptopClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LaptopService_SearchLaptopClient interface {
	Recv() (*SearchLaptopResponse, error)
	grpc.ClientStream
}

type laptopServiceSearchLaptopClient struct {
	grpc.ClientStream
}

func (x *laptopServiceSearchLaptopClient) Recv() (*SearchLaptopResponse, error) {
	m := new(SearchLaptopResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *laptopServiceClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (LaptopService_UploadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LaptopService_serviceDesc.Streams[1], "/pb.LaptopService/UploadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &laptopServiceUploadImageClient{stream}
	return x, nil
}

type LaptopService_UploadImageClient interface {
	Send(*UploadImageRequest) error
	CloseAndRecv() (*UploadImageResponse, error)
	grpc.ClientStream
}

type laptopServiceUploadImageClient struct {
	grpc.ClientStream
}

func (x *laptopServiceUploadImageClient) Send(m *UploadImageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *laptopServiceUploadImageClient) CloseAndRecv() (*UploadImageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *laptopServiceClient) RateLaptop(ctx context.Context, opts ...grpc.CallOption) (LaptopService_RateLaptopClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LaptopService_serviceDesc.Streams[2], "/pb.LaptopService/RateLaptop", opts...)
	if err != nil {
		return nil, err
	}
	x := &laptopServiceRateLaptopClient{stream}
	return x, nil
}

type LaptopService_RateLaptopClient interface {
	Send(*RateLaptopRequest) error
	Recv() (*RateLaptopResponse, error)
	grpc.ClientStream
}

type laptopServiceRateLaptopClient struct {
	grpc.ClientStream
}

func (x *laptopServiceRateLaptopClient) Send(m *RateLaptopRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *laptopServiceRateLaptopClient) Recv() (*RateLaptopResponse, error) {
	m := new(RateLaptopResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LaptopServiceServer is the server API for LaptopService service.
type LaptopServiceServer interface {
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
	SearchLaptop(*SearchLaptopRequest, LaptopService_SearchLaptopServer) error
	UploadImage(LaptopService_UploadImageServer) error
	RateLaptop(LaptopService_RateLaptopServer) error
}

// UnimplementedLaptopServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLaptopServiceServer struct {
}

func (*UnimplementedLaptopServiceServer) CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaptop not implemented")
}
func (*UnimplementedLaptopServiceServer) SearchLaptop(*SearchLaptopRequest, LaptopService_SearchLaptopServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchLaptop not implemented")
}
func (*UnimplementedLaptopServiceServer) UploadImage(LaptopService_UploadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (*UnimplementedLaptopServiceServer) RateLaptop(LaptopService_RateLaptopServer) error {
	return status.Errorf(codes.Unimplemented, "method RateLaptop not implemented")
}

func RegisterLaptopServiceServer(s *grpc.Server, srv LaptopServiceServer) {
	s.RegisterService(&_LaptopService_serviceDesc, srv)
}

func _LaptopService_CreateLaptop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLaptopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LaptopService/CreateLaptop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, req.(*CreateLaptopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaptopService_SearchLaptop_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchLaptopRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LaptopServiceServer).SearchLaptop(m, &laptopServiceSearchLaptopServer{stream})
}

type LaptopService_SearchLaptopServer interface {
	Send(*SearchLaptopResponse) error
	grpc.ServerStream
}

type laptopServiceSearchLaptopServer struct {
	grpc.ServerStream
}

func (x *laptopServiceSearchLaptopServer) Send(m *SearchLaptopResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LaptopService_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LaptopServiceServer).UploadImage(&laptopServiceUploadImageServer{stream})
}

type LaptopService_UploadImageServer interface {
	SendAndClose(*UploadImageResponse) error
	Recv() (*UploadImageRequest, error)
	grpc.ServerStream
}

type laptopServiceUploadImageServer struct {
	grpc.ServerStream
}

func (x *laptopServiceUploadImageServer) SendAndClose(m *UploadImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *laptopServiceUploadImageServer) Recv() (*UploadImageRequest, error) {
	m := new(UploadImageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LaptopService_RateLaptop_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LaptopServiceServer).RateLaptop(&laptopServiceRateLaptopServer{stream})
}

type LaptopService_RateLaptopServer interface {
	Send(*RateLaptopResponse) error
	Recv() (*RateLaptopRequest, error)
	grpc.ServerStream
}

type laptopServiceRateLaptopServer struct {
	grpc.ServerStream
}

func (x *laptopServiceRateLaptopServer) Send(m *RateLaptopResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *laptopServiceRateLaptopServer) Recv() (*RateLaptopRequest, error) {
	m := new(RateLaptopRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LaptopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LaptopService",
	HandlerType: (*LaptopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaptop",
			Handler:    _LaptopService_CreateLaptop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchLaptop",
			Handler:       _LaptopService_SearchLaptop_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadImage",
			Handler:       _LaptopService_UploadImage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RateLaptop",
			Handler:       _LaptopService_RateLaptop_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "laptop_server.proto",
}
