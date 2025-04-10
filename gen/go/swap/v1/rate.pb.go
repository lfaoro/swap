// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: swap/v1/rate.proto

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

// SwapRateRequest defines the parameters for a swap or payment request.
type SwapRateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ticker of the coin you want to sell (e.g., "btc").
	TickerFrom string `protobuf:"bytes,1,opt,name=ticker_from,json=tickerFrom,proto3" json:"ticker_from,omitempty"`
	// The network of the coin you want to sell (e.g., "Mainnet").
	NetworkFrom string `protobuf:"bytes,2,opt,name=network_from,json=networkFrom,proto3" json:"network_from,omitempty"`
	// The ticker of the coin you want to buy (e.g., "xmr").
	TickerTo string `protobuf:"bytes,3,opt,name=ticker_to,json=tickerTo,proto3" json:"ticker_to,omitempty"`
	// The network of the coin you want to buy (e.g., "Mainnet").
	NetworkTo string `protobuf:"bytes,4,opt,name=network_to,json=networkTo,proto3" json:"network_to,omitempty"`
	// The amount of the coin you want to sell (required for standard swaps).
	AmountFrom float64 `protobuf:"fixed64,5,opt,name=amount_from,json=amountFrom,proto3" json:"amount_from,omitempty"`
	// The amount of the coin you want to receive (required for payments).
	AmountTo float64 `protobuf:"fixed64,6,opt,name=amount_to,json=amountTo,proto3" json:"amount_to,omitempty"`
	// Whether to create a fixed-rate payment or a standard swap.
	Payment bool `protobuf:"varint,7,opt,name=payment,proto3" json:"payment,omitempty"`
	// Minimum KYC rating required for the exchange (e.g., "A", "B", "C", "D").
	// Whether to return only the best rate for the provided parameters.
	BestOnly      bool `protobuf:"varint,11,opt,name=best_only,json=bestOnly,proto3" json:"best_only,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SwapRateRequest) Reset() {
	*x = SwapRateRequest{}
	mi := &file_swap_v1_rate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SwapRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapRateRequest) ProtoMessage() {}

func (x *SwapRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_swap_v1_rate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapRateRequest.ProtoReflect.Descriptor instead.
func (*SwapRateRequest) Descriptor() ([]byte, []int) {
	return file_swap_v1_rate_proto_rawDescGZIP(), []int{0}
}

func (x *SwapRateRequest) GetTickerFrom() string {
	if x != nil {
		return x.TickerFrom
	}
	return ""
}

func (x *SwapRateRequest) GetNetworkFrom() string {
	if x != nil {
		return x.NetworkFrom
	}
	return ""
}

func (x *SwapRateRequest) GetTickerTo() string {
	if x != nil {
		return x.TickerTo
	}
	return ""
}

func (x *SwapRateRequest) GetNetworkTo() string {
	if x != nil {
		return x.NetworkTo
	}
	return ""
}

func (x *SwapRateRequest) GetAmountFrom() float64 {
	if x != nil {
		return x.AmountFrom
	}
	return 0
}

func (x *SwapRateRequest) GetAmountTo() float64 {
	if x != nil {
		return x.AmountTo
	}
	return 0
}

func (x *SwapRateRequest) GetPayment() bool {
	if x != nil {
		return x.Payment
	}
	return false
}

func (x *SwapRateRequest) GetBestOnly() bool {
	if x != nil {
		return x.BestOnly
	}
	return false
}

// SwapRateResponse defines the response for a swap or payment request.
type SwapRateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TradeId       string                 `protobuf:"bytes,1,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	TickerFrom    string                 `protobuf:"bytes,3,opt,name=ticker_from,json=tickerFrom,proto3" json:"ticker_from,omitempty"`
	TickerTo      string                 `protobuf:"bytes,4,opt,name=ticker_to,json=tickerTo,proto3" json:"ticker_to,omitempty"`
	CoinFrom      string                 `protobuf:"bytes,5,opt,name=coin_from,json=coinFrom,proto3" json:"coin_from,omitempty"`
	CoinTo        string                 `protobuf:"bytes,6,opt,name=coin_to,json=coinTo,proto3" json:"coin_to,omitempty"`
	NetworkFrom   string                 `protobuf:"bytes,7,opt,name=network_from,json=networkFrom,proto3" json:"network_from,omitempty"`
	NetworkTo     string                 `protobuf:"bytes,8,opt,name=network_to,json=networkTo,proto3" json:"network_to,omitempty"`
	AmountFrom    float64                `protobuf:"fixed64,9,opt,name=amount_from,json=amountFrom,proto3" json:"amount_from,omitempty"`
	AmountTo      float64                `protobuf:"fixed64,10,opt,name=amount_to,json=amountTo,proto3" json:"amount_to,omitempty"`
	Provider      string                 `protobuf:"bytes,11,opt,name=provider,proto3" json:"provider,omitempty"`
	Fixed         bool                   `protobuf:"varint,12,opt,name=fixed,proto3" json:"fixed,omitempty"`
	Payment       bool                   `protobuf:"varint,13,opt,name=payment,proto3" json:"payment,omitempty"`
	Status        string                 `protobuf:"bytes,14,opt,name=status,proto3" json:"status,omitempty"`
	Quotes        *QuoteList             `protobuf:"bytes,15,opt,name=quotes,proto3" json:"quotes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SwapRateResponse) Reset() {
	*x = SwapRateResponse{}
	mi := &file_swap_v1_rate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SwapRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapRateResponse) ProtoMessage() {}

func (x *SwapRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_swap_v1_rate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapRateResponse.ProtoReflect.Descriptor instead.
func (*SwapRateResponse) Descriptor() ([]byte, []int) {
	return file_swap_v1_rate_proto_rawDescGZIP(), []int{1}
}

func (x *SwapRateResponse) GetTradeId() string {
	if x != nil {
		return x.TradeId
	}
	return ""
}

func (x *SwapRateResponse) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *SwapRateResponse) GetTickerFrom() string {
	if x != nil {
		return x.TickerFrom
	}
	return ""
}

func (x *SwapRateResponse) GetTickerTo() string {
	if x != nil {
		return x.TickerTo
	}
	return ""
}

func (x *SwapRateResponse) GetCoinFrom() string {
	if x != nil {
		return x.CoinFrom
	}
	return ""
}

func (x *SwapRateResponse) GetCoinTo() string {
	if x != nil {
		return x.CoinTo
	}
	return ""
}

func (x *SwapRateResponse) GetNetworkFrom() string {
	if x != nil {
		return x.NetworkFrom
	}
	return ""
}

func (x *SwapRateResponse) GetNetworkTo() string {
	if x != nil {
		return x.NetworkTo
	}
	return ""
}

func (x *SwapRateResponse) GetAmountFrom() float64 {
	if x != nil {
		return x.AmountFrom
	}
	return 0
}

func (x *SwapRateResponse) GetAmountTo() float64 {
	if x != nil {
		return x.AmountTo
	}
	return 0
}

func (x *SwapRateResponse) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *SwapRateResponse) GetFixed() bool {
	if x != nil {
		return x.Fixed
	}
	return false
}

func (x *SwapRateResponse) GetPayment() bool {
	if x != nil {
		return x.Payment
	}
	return false
}

func (x *SwapRateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SwapRateResponse) GetQuotes() *QuoteList {
	if x != nil {
		return x.Quotes
	}
	return nil
}

var File_swap_v1_rate_proto protoreflect.FileDescriptor

var file_swap_v1_rate_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x73, 0x77, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x73, 0x77, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x86, 0x02, 0x0a, 0x0f, 0x53, 0x77, 0x61, 0x70, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x54, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x62, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xe1, 0x03, 0x0a,
	0x10, 0x53, 0x77, 0x61, 0x70, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x54, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f,
	0x69, 0x6e, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6f, 0x69, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x69, 0x6e, 0x5f,
	0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e, 0x54, 0x6f,
	0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x54, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73,
	0x42, 0x5e, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x42,
	0x09, 0x52, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x07, 0x73, 0x77,
	0x61, 0x70, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x53, 0x77,
	0x61, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x53, 0x77, 0x61, 0x70, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x13, 0x53, 0x77, 0x61, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x53, 0x77, 0x61, 0x70, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_swap_v1_rate_proto_rawDescOnce sync.Once
	file_swap_v1_rate_proto_rawDescData []byte
)

func file_swap_v1_rate_proto_rawDescGZIP() []byte {
	file_swap_v1_rate_proto_rawDescOnce.Do(func() {
		file_swap_v1_rate_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_swap_v1_rate_proto_rawDesc), len(file_swap_v1_rate_proto_rawDesc)))
	})
	return file_swap_v1_rate_proto_rawDescData
}

var file_swap_v1_rate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_swap_v1_rate_proto_goTypes = []any{
	(*SwapRateRequest)(nil),       // 0: swap.v1.SwapRateRequest
	(*SwapRateResponse)(nil),      // 1: swap.v1.SwapRateResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*QuoteList)(nil),             // 3: swap.v1.QuoteList
}
var file_swap_v1_rate_proto_depIdxs = []int32{
	2, // 0: swap.v1.SwapRateResponse.date:type_name -> google.protobuf.Timestamp
	3, // 1: swap.v1.SwapRateResponse.quotes:type_name -> swap.v1.QuoteList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_swap_v1_rate_proto_init() }
func file_swap_v1_rate_proto_init() {
	if File_swap_v1_rate_proto != nil {
		return
	}
	file_swap_v1_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_swap_v1_rate_proto_rawDesc), len(file_swap_v1_rate_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_swap_v1_rate_proto_goTypes,
		DependencyIndexes: file_swap_v1_rate_proto_depIdxs,
		MessageInfos:      file_swap_v1_rate_proto_msgTypes,
	}.Build()
	File_swap_v1_rate_proto = out.File
	file_swap_v1_rate_proto_goTypes = nil
	file_swap_v1_rate_proto_depIdxs = nil
}
