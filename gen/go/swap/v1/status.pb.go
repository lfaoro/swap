// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: swap/v1/status.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type SwapStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TradeId       string                 `protobuf:"bytes,1,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SwapStatusRequest) Reset() {
	*x = SwapStatusRequest{}
	mi := &file_swap_v1_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SwapStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapStatusRequest) ProtoMessage() {}

func (x *SwapStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_swap_v1_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapStatusRequest.ProtoReflect.Descriptor instead.
func (*SwapStatusRequest) Descriptor() ([]byte, []int) {
	return file_swap_v1_status_proto_rawDescGZIP(), []int{0}
}

func (x *SwapStatusRequest) GetTradeId() string {
	if x != nil {
		return x.TradeId
	}
	return ""
}

type SwapStatusResponse struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	TradeId             string                 `protobuf:"bytes,1,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
	Date                *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Type                string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	TickerFrom          string                 `protobuf:"bytes,4,opt,name=ticker_from,json=tickerFrom,proto3" json:"ticker_from,omitempty"`
	TickerTo            string                 `protobuf:"bytes,5,opt,name=ticker_to,json=tickerTo,proto3" json:"ticker_to,omitempty"`
	CoinFrom            string                 `protobuf:"bytes,6,opt,name=coin_from,json=coinFrom,proto3" json:"coin_from,omitempty"`
	CoinTo              string                 `protobuf:"bytes,7,opt,name=coin_to,json=coinTo,proto3" json:"coin_to,omitempty"`
	NetworkFrom         string                 `protobuf:"bytes,8,opt,name=network_from,json=networkFrom,proto3" json:"network_from,omitempty"`
	NetworkTo           string                 `protobuf:"bytes,9,opt,name=network_to,json=networkTo,proto3" json:"network_to,omitempty"`
	AmountFrom          float64                `protobuf:"fixed64,10,opt,name=amount_from,json=amountFrom,proto3" json:"amount_from,omitempty"`
	AmountTo            float64                `protobuf:"fixed64,11,opt,name=amount_to,json=amountTo,proto3" json:"amount_to,omitempty"`
	Provider            string                 `protobuf:"bytes,12,opt,name=provider,proto3" json:"provider,omitempty"`
	Payment             bool                   `protobuf:"varint,14,opt,name=payment,proto3" json:"payment,omitempty"`
	Fixed               bool                   `protobuf:"varint,13,opt,name=fixed,proto3" json:"fixed,omitempty"`
	Status              string                 `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	AddressProvider     string                 `protobuf:"bytes,16,opt,name=address_provider,json=addressProvider,proto3" json:"address_provider,omitempty"`
	AddressProviderMemo string                 `protobuf:"bytes,17,opt,name=address_provider_memo,json=addressProviderMemo,proto3" json:"address_provider_memo,omitempty"`
	AddressUser         string                 `protobuf:"bytes,18,opt,name=address_user,json=addressUser,proto3" json:"address_user,omitempty"`
	AddressUserMemo     string                 `protobuf:"bytes,19,opt,name=address_user_memo,json=addressUserMemo,proto3" json:"address_user_memo,omitempty"`
	RefundAddress       string                 `protobuf:"bytes,20,opt,name=refund_address,json=refundAddress,proto3" json:"refund_address,omitempty"`
	RefundAddressMemo   string                 `protobuf:"bytes,21,opt,name=refund_address_memo,json=refundAddressMemo,proto3" json:"refund_address_memo,omitempty"`
	Password            string                 `protobuf:"bytes,22,opt,name=password,proto3" json:"password,omitempty"`
	IdProvider          string                 `protobuf:"bytes,23,opt,name=id_provider,json=idProvider,proto3" json:"id_provider,omitempty"`
	Details             *Details               `protobuf:"bytes,26,opt,name=details,proto3" json:"details,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *SwapStatusResponse) Reset() {
	*x = SwapStatusResponse{}
	mi := &file_swap_v1_status_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SwapStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapStatusResponse) ProtoMessage() {}

func (x *SwapStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_swap_v1_status_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapStatusResponse.ProtoReflect.Descriptor instead.
func (*SwapStatusResponse) Descriptor() ([]byte, []int) {
	return file_swap_v1_status_proto_rawDescGZIP(), []int{1}
}

func (x *SwapStatusResponse) GetTradeId() string {
	if x != nil {
		return x.TradeId
	}
	return ""
}

func (x *SwapStatusResponse) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *SwapStatusResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SwapStatusResponse) GetTickerFrom() string {
	if x != nil {
		return x.TickerFrom
	}
	return ""
}

func (x *SwapStatusResponse) GetTickerTo() string {
	if x != nil {
		return x.TickerTo
	}
	return ""
}

func (x *SwapStatusResponse) GetCoinFrom() string {
	if x != nil {
		return x.CoinFrom
	}
	return ""
}

func (x *SwapStatusResponse) GetCoinTo() string {
	if x != nil {
		return x.CoinTo
	}
	return ""
}

func (x *SwapStatusResponse) GetNetworkFrom() string {
	if x != nil {
		return x.NetworkFrom
	}
	return ""
}

func (x *SwapStatusResponse) GetNetworkTo() string {
	if x != nil {
		return x.NetworkTo
	}
	return ""
}

func (x *SwapStatusResponse) GetAmountFrom() float64 {
	if x != nil {
		return x.AmountFrom
	}
	return 0
}

func (x *SwapStatusResponse) GetAmountTo() float64 {
	if x != nil {
		return x.AmountTo
	}
	return 0
}

func (x *SwapStatusResponse) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *SwapStatusResponse) GetPayment() bool {
	if x != nil {
		return x.Payment
	}
	return false
}

func (x *SwapStatusResponse) GetFixed() bool {
	if x != nil {
		return x.Fixed
	}
	return false
}

func (x *SwapStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SwapStatusResponse) GetAddressProvider() string {
	if x != nil {
		return x.AddressProvider
	}
	return ""
}

func (x *SwapStatusResponse) GetAddressProviderMemo() string {
	if x != nil {
		return x.AddressProviderMemo
	}
	return ""
}

func (x *SwapStatusResponse) GetAddressUser() string {
	if x != nil {
		return x.AddressUser
	}
	return ""
}

func (x *SwapStatusResponse) GetAddressUserMemo() string {
	if x != nil {
		return x.AddressUserMemo
	}
	return ""
}

func (x *SwapStatusResponse) GetRefundAddress() string {
	if x != nil {
		return x.RefundAddress
	}
	return ""
}

func (x *SwapStatusResponse) GetRefundAddressMemo() string {
	if x != nil {
		return x.RefundAddressMemo
	}
	return ""
}

func (x *SwapStatusResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SwapStatusResponse) GetIdProvider() string {
	if x != nil {
		return x.IdProvider
	}
	return ""
}

func (x *SwapStatusResponse) GetDetails() *Details {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_swap_v1_status_proto protoreflect.FileDescriptor

var file_swap_v1_status_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x73, 0x77, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x73, 0x77, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x11, 0x53, 0x77, 0x61, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x49, 0x64, 0x22, 0xb9, 0x06, 0x0a, 0x12, 0x53, 0x77, 0x61, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x54, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69,
	0x6e, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x69, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x74,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x12,
	0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x6f,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54,
	0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x32,
	0x0a, 0x15, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4d, 0x65,
	0x6d, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x6d,
	0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x42, 0x60, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31,
	0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x07, 0x73, 0x77, 0x61, 0x70, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02,
	0x07, 0x53, 0x77, 0x61, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x53, 0x77, 0x61, 0x70, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x13, 0x53, 0x77, 0x61, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x53, 0x77, 0x61, 0x70, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_swap_v1_status_proto_rawDescOnce sync.Once
	file_swap_v1_status_proto_rawDescData []byte
)

func file_swap_v1_status_proto_rawDescGZIP() []byte {
	file_swap_v1_status_proto_rawDescOnce.Do(func() {
		file_swap_v1_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_swap_v1_status_proto_rawDesc), len(file_swap_v1_status_proto_rawDesc)))
	})
	return file_swap_v1_status_proto_rawDescData
}

var file_swap_v1_status_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_swap_v1_status_proto_goTypes = []any{
	(*SwapStatusRequest)(nil),     // 0: swap.v1.SwapStatusRequest
	(*SwapStatusResponse)(nil),    // 1: swap.v1.SwapStatusResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*Details)(nil),               // 3: swap.v1.Details
}
var file_swap_v1_status_proto_depIdxs = []int32{
	2, // 0: swap.v1.SwapStatusResponse.date:type_name -> google.protobuf.Timestamp
	3, // 1: swap.v1.SwapStatusResponse.details:type_name -> swap.v1.Details
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_swap_v1_status_proto_init() }
func file_swap_v1_status_proto_init() {
	if File_swap_v1_status_proto != nil {
		return
	}
	file_swap_v1_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_swap_v1_status_proto_rawDesc), len(file_swap_v1_status_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_swap_v1_status_proto_goTypes,
		DependencyIndexes: file_swap_v1_status_proto_depIdxs,
		MessageInfos:      file_swap_v1_status_proto_msgTypes,
	}.Build()
	File_swap_v1_status_proto = out.File
	file_swap_v1_status_proto_goTypes = nil
	file_swap_v1_status_proto_depIdxs = nil
}
