// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0--rc3
// source: proto/products_manager.proto

package products_manager

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

type SortingOptions_Order int32

const (
	SortingOptions_ASC  SortingOptions_Order = 0
	SortingOptions_DESC SortingOptions_Order = 1
)

// Enum value maps for SortingOptions_Order.
var (
	SortingOptions_Order_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	SortingOptions_Order_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x SortingOptions_Order) Enum() *SortingOptions_Order {
	p := new(SortingOptions_Order)
	*p = x
	return p
}

func (x SortingOptions_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortingOptions_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_products_manager_proto_enumTypes[0].Descriptor()
}

func (SortingOptions_Order) Type() protoreflect.EnumType {
	return &file_proto_products_manager_proto_enumTypes[0]
}

func (x SortingOptions_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortingOptions_Order.Descriptor instead.
func (SortingOptions_Order) EnumDescriptor() ([]byte, []int) {
	return file_proto_products_manager_proto_rawDescGZIP(), []int{2, 0}
}

type ProductItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Price float32 `protobuf:"fixed32,2,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *ProductItem) Reset() {
	*x = ProductItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductItem) ProtoMessage() {}

func (x *ProductItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductItem.ProtoReflect.Descriptor instead.
func (*ProductItem) Descriptor() ([]byte, []int) {
	return file_proto_products_manager_proto_rawDescGZIP(), []int{0}
}

func (x *ProductItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type PaginationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PaginationOptions) Reset() {
	*x = PaginationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationOptions) ProtoMessage() {}

func (x *PaginationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationOptions.ProtoReflect.Descriptor instead.
func (*PaginationOptions) Descriptor() ([]byte, []int) {
	return file_proto_products_manager_proto_rawDescGZIP(), []int{1}
}

func (x *PaginationOptions) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PaginationOptions) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type SortingOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order SortingOptions_Order `protobuf:"varint,1,opt,name=order,proto3,enum=products_manager.SortingOptions_Order" json:"order,omitempty"`
	Field string               `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *SortingOptions) Reset() {
	*x = SortingOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortingOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortingOptions) ProtoMessage() {}

func (x *SortingOptions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortingOptions.ProtoReflect.Descriptor instead.
func (*SortingOptions) Descriptor() ([]byte, []int) {
	return file_proto_products_manager_proto_rawDescGZIP(), []int{2}
}

func (x *SortingOptions) GetOrder() SortingOptions_Order {
	if x != nil {
		return x.Order
	}
	return SortingOptions_ASC
}

func (x *SortingOptions) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaginOpts *PaginationOptions `protobuf:"bytes,1,opt,name=paginOpts,proto3" json:"paginOpts,omitempty"`
	SortOpts  *SortingOptions    `protobuf:"bytes,2,opt,name=sortOpts,proto3" json:"sortOpts,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_products_manager_proto_rawDescGZIP(), []int{3}
}

func (x *ListRequest) GetPaginOpts() *PaginationOptions {
	if x != nil {
		return x.PaginOpts
	}
	return nil
}

func (x *ListRequest) GetSortOpts() *SortingOptions {
	if x != nil {
		return x.SortOpts
	}
	return nil
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ProductItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_products_manager_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetData() []*ProductItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_proto_products_manager_proto_rawDescGZIP(), []int{5}
}

func (x *FetchRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_products_manager_proto_rawDescGZIP(), []int{6}
}

var File_proto_products_manager_proto protoreflect.FileDescriptor

var file_proto_products_manager_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x22, 0x37, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x41, 0x0a, 0x11, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x80, 0x01, 0x0a,
	0x0e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x3c, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x22, 0x1a, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x22,
	0x8e, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x41, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x4f, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x4f, 0x70,
	0x74, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x73,
	0x22, 0x41, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x20, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xa4,
	0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_products_manager_proto_rawDescOnce sync.Once
	file_proto_products_manager_proto_rawDescData = file_proto_products_manager_proto_rawDesc
)

func file_proto_products_manager_proto_rawDescGZIP() []byte {
	file_proto_products_manager_proto_rawDescOnce.Do(func() {
		file_proto_products_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_products_manager_proto_rawDescData)
	})
	return file_proto_products_manager_proto_rawDescData
}

var file_proto_products_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_products_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_products_manager_proto_goTypes = []interface{}{
	(SortingOptions_Order)(0), // 0: products_manager.SortingOptions.Order
	(*ProductItem)(nil),       // 1: products_manager.ProductItem
	(*PaginationOptions)(nil), // 2: products_manager.PaginationOptions
	(*SortingOptions)(nil),    // 3: products_manager.SortingOptions
	(*ListRequest)(nil),       // 4: products_manager.ListRequest
	(*ListResponse)(nil),      // 5: products_manager.ListResponse
	(*FetchRequest)(nil),      // 6: products_manager.FetchRequest
	(*Empty)(nil),             // 7: products_manager.Empty
}
var file_proto_products_manager_proto_depIdxs = []int32{
	0, // 0: products_manager.SortingOptions.order:type_name -> products_manager.SortingOptions.Order
	2, // 1: products_manager.ListRequest.paginOpts:type_name -> products_manager.PaginationOptions
	3, // 2: products_manager.ListRequest.sortOpts:type_name -> products_manager.SortingOptions
	1, // 3: products_manager.ListResponse.data:type_name -> products_manager.ProductItem
	6, // 4: products_manager.ProductManagerService.Fetch:input_type -> products_manager.FetchRequest
	4, // 5: products_manager.ProductManagerService.List:input_type -> products_manager.ListRequest
	7, // 6: products_manager.ProductManagerService.Fetch:output_type -> products_manager.Empty
	5, // 7: products_manager.ProductManagerService.List:output_type -> products_manager.ListResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_products_manager_proto_init() }
func file_proto_products_manager_proto_init() {
	if File_proto_products_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_products_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductItem); i {
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
		file_proto_products_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationOptions); i {
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
		file_proto_products_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortingOptions); i {
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
		file_proto_products_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_proto_products_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_proto_products_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_proto_products_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_proto_products_manager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_products_manager_proto_goTypes,
		DependencyIndexes: file_proto_products_manager_proto_depIdxs,
		EnumInfos:         file_proto_products_manager_proto_enumTypes,
		MessageInfos:      file_proto_products_manager_proto_msgTypes,
	}.Build()
	File_proto_products_manager_proto = out.File
	file_proto_products_manager_proto_rawDesc = nil
	file_proto_products_manager_proto_goTypes = nil
	file_proto_products_manager_proto_depIdxs = nil
}
