// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: swap/v1/main.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_swap_v1_main_proto protoreflect.FileDescriptor

var file_swap_v1_main_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x73, 0x77, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x77, 0x61, 0x70,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x73, 0x77, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x73, 0x77, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x77, 0x61, 0x70, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x93, 0x02,
	0x0a, 0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x53, 0x77, 0x61, 0x70, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x70,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x77,
	0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x53, 0x77, 0x61, 0x70, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x77,
	0x61, 0x70, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x53, 0x77,
	0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x77, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x42, 0x5e, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e,
	0x76, 0x31, 0x42, 0x09, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x07, 0x73, 0x77, 0x61, 0x70, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02,
	0x07, 0x53, 0x77, 0x61, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x53, 0x77, 0x61, 0x70, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x13, 0x53, 0x77, 0x61, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x53, 0x77, 0x61, 0x70, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_swap_v1_main_proto_goTypes = []any{
	(*emptypb.Empty)(nil),      // 0: google.protobuf.Empty
	(*SwapRateRequest)(nil),    // 1: swap.v1.SwapRateRequest
	(*SwapTradeRequest)(nil),   // 2: swap.v1.SwapTradeRequest
	(*SwapStatusRequest)(nil),  // 3: swap.v1.SwapStatusRequest
	(*CoinList)(nil),           // 4: swap.v1.CoinList
	(*SwapRateResponse)(nil),   // 5: swap.v1.SwapRateResponse
	(*SwapTradeResponse)(nil),  // 6: swap.v1.SwapTradeResponse
	(*SwapStatusResponse)(nil), // 7: swap.v1.SwapStatusResponse
}
var file_swap_v1_main_proto_depIdxs = []int32{
	0, // 0: swap.v1.CoinService.ListCoins:input_type -> google.protobuf.Empty
	1, // 1: swap.v1.CoinService.SwapRate:input_type -> swap.v1.SwapRateRequest
	2, // 2: swap.v1.CoinService.SwapTrade:input_type -> swap.v1.SwapTradeRequest
	3, // 3: swap.v1.CoinService.SwapStatus:input_type -> swap.v1.SwapStatusRequest
	4, // 4: swap.v1.CoinService.ListCoins:output_type -> swap.v1.CoinList
	5, // 5: swap.v1.CoinService.SwapRate:output_type -> swap.v1.SwapRateResponse
	6, // 6: swap.v1.CoinService.SwapTrade:output_type -> swap.v1.SwapTradeResponse
	7, // 7: swap.v1.CoinService.SwapStatus:output_type -> swap.v1.SwapStatusResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_swap_v1_main_proto_init() }
func file_swap_v1_main_proto_init() {
	if File_swap_v1_main_proto != nil {
		return
	}
	file_swap_v1_coin_proto_init()
	file_swap_v1_rate_proto_init()
	file_swap_v1_trade_proto_init()
	file_swap_v1_status_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_swap_v1_main_proto_rawDesc), len(file_swap_v1_main_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_swap_v1_main_proto_goTypes,
		DependencyIndexes: file_swap_v1_main_proto_depIdxs,
	}.Build()
	File_swap_v1_main_proto = out.File
	file_swap_v1_main_proto_goTypes = nil
	file_swap_v1_main_proto_depIdxs = nil
}
