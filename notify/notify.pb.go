// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: notify/notify.proto

package notify

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type OrderOperationType int32

const (
	OrderOperationType_INIT_ORDER   OrderOperationType = 0
	OrderOperationType_CANCEL_ORDER OrderOperationType = 1
)

// Enum value maps for OrderOperationType.
var (
	OrderOperationType_name = map[int32]string{
		0: "INIT_ORDER",
		1: "CANCEL_ORDER",
	}
	OrderOperationType_value = map[string]int32{
		"INIT_ORDER":   0,
		"CANCEL_ORDER": 1,
	}
)

func (x OrderOperationType) Enum() *OrderOperationType {
	p := new(OrderOperationType)
	*p = x
	return p
}

func (x OrderOperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderOperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_notify_notify_proto_enumTypes[0].Descriptor()
}

func (OrderOperationType) Type() protoreflect.EnumType {
	return &file_notify_notify_proto_enumTypes[0]
}

func (x OrderOperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderOperationType.Descriptor instead.
func (OrderOperationType) EnumDescriptor() ([]byte, []int) {
	return file_notify_notify_proto_rawDescGZIP(), []int{0}
}

type OrderNotifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OrderNotifyResponse) Reset() {
	*x = OrderNotifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notify_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderNotifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderNotifyResponse) ProtoMessage() {}

func (x *OrderNotifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notify_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderNotifyResponse.ProtoReflect.Descriptor instead.
func (*OrderNotifyResponse) Descriptor() ([]byte, []int) {
	return file_notify_notify_proto_rawDescGZIP(), []int{0}
}

func (x *OrderNotifyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product   string               `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	UserID    uint32               `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Operation OrderOperationType   `protobuf:"varint,3,opt,name=operation,proto3,enum=notify.OrderOperationType" json:"operation,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notify_notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_notify_notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_notify_notify_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *Order) GetUserID() uint32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Order) GetOperation() OrderOperationType {
	if x != nil {
		return x.Operation
	}
	return OrderOperationType_INIT_ORDER
}

func (x *Order) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type OrderNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *OrderNotification) Reset() {
	*x = OrderNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notify_notify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderNotification) ProtoMessage() {}

func (x *OrderNotification) ProtoReflect() protoreflect.Message {
	mi := &file_notify_notify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderNotification.ProtoReflect.Descriptor instead.
func (*OrderNotification) Descriptor() ([]byte, []int) {
	return file_notify_notify_proto_rawDescGZIP(), []int{2}
}

func (x *OrderNotification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OrderNotification) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Subscriber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberID uint32 `protobuf:"varint,1,opt,name=subscriberID,proto3" json:"subscriberID,omitempty"`
}

func (x *Subscriber) Reset() {
	*x = Subscriber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notify_notify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscriber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscriber) ProtoMessage() {}

func (x *Subscriber) ProtoReflect() protoreflect.Message {
	mi := &file_notify_notify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscriber.ProtoReflect.Descriptor instead.
func (*Subscriber) Descriptor() ([]byte, []int) {
	return file_notify_notify_proto_rawDescGZIP(), []int{3}
}

func (x *Subscriber) GetSubscriberID() uint32 {
	if x != nil {
		return x.SubscriberID
	}
	return 0
}

var File_notify_notify_proto protoreflect.FileDescriptor

var file_notify_notify_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f,
	0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xad, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x67, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x30, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x44, 0x2a, 0x36, 0x0a, 0x12, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x49, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x10, 0x01, 0x32, 0x80, 0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x36, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x0d, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x3c, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x12, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x1a, 0x19, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x6f, 0x70, 0x65, 0x72, 0x32, 0x30, 0x37, 0x34, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_notify_notify_proto_rawDescOnce sync.Once
	file_notify_notify_proto_rawDescData = file_notify_notify_proto_rawDesc
)

func file_notify_notify_proto_rawDescGZIP() []byte {
	file_notify_notify_proto_rawDescOnce.Do(func() {
		file_notify_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_notify_notify_proto_rawDescData)
	})
	return file_notify_notify_proto_rawDescData
}

var file_notify_notify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notify_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_notify_notify_proto_goTypes = []interface{}{
	(OrderOperationType)(0),     // 0: notify.OrderOperationType
	(*OrderNotifyResponse)(nil), // 1: notify.OrderNotifyResponse
	(*Order)(nil),               // 2: notify.Order
	(*OrderNotification)(nil),   // 3: notify.OrderNotification
	(*Subscriber)(nil),          // 4: notify.Subscriber
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_notify_notify_proto_depIdxs = []int32{
	0, // 0: notify.Order.operation:type_name -> notify.OrderOperationType
	5, // 1: notify.Order.timestamp:type_name -> google.protobuf.Timestamp
	5, // 2: notify.OrderNotification.timestamp:type_name -> google.protobuf.Timestamp
	2, // 3: notify.Notifier.Notify:input_type -> notify.Order
	4, // 4: notify.Notifier.Subscribe:input_type -> notify.Subscriber
	1, // 5: notify.Notifier.Notify:output_type -> notify.OrderNotifyResponse
	3, // 6: notify.Notifier.Subscribe:output_type -> notify.OrderNotification
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_notify_notify_proto_init() }
func file_notify_notify_proto_init() {
	if File_notify_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notify_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderNotifyResponse); i {
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
		file_notify_notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_notify_notify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderNotification); i {
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
		file_notify_notify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscriber); i {
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
			RawDescriptor: file_notify_notify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notify_notify_proto_goTypes,
		DependencyIndexes: file_notify_notify_proto_depIdxs,
		EnumInfos:         file_notify_notify_proto_enumTypes,
		MessageInfos:      file_notify_notify_proto_msgTypes,
	}.Build()
	File_notify_notify_proto = out.File
	file_notify_notify_proto_rawDesc = nil
	file_notify_notify_proto_goTypes = nil
	file_notify_notify_proto_depIdxs = nil
}