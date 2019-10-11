// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/accesslog/v3alpha/als.proto

package envoy_config_accesslog_v3alpha

import (
	fmt "fmt"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Configuration for the built-in *envoy.http_grpc_access_log*
// :ref:`AccessLog <envoy_api_msg_config.filter.accesslog.v3alpha.AccessLog>`. This configuration
// will populate :ref:`StreamAccessLogsMessage.http_logs
// <envoy_api_field_service.accesslog.v3alpha.StreamAccessLogsMessage.http_logs>`.
type HttpGrpcAccessLogConfig struct {
	CommonConfig *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// Additional request headers to log in :ref:`HTTPRequestProperties.request_headers
	// <envoy_api_field_data.accesslog.v3alpha.HTTPRequestProperties.request_headers>`.
	AdditionalRequestHeadersToLog []string `protobuf:"bytes,2,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	// Additional response headers to log in :ref:`HTTPResponseProperties.response_headers
	// <envoy_api_field_data.accesslog.v3alpha.HTTPResponseProperties.response_headers>`.
	AdditionalResponseHeadersToLog []string `protobuf:"bytes,3,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	// Additional response trailers to log in :ref:`HTTPResponseProperties.response_trailers
	// <envoy_api_field_data.accesslog.v3alpha.HTTPResponseProperties.response_trailers>`.
	AdditionalResponseTrailersToLog []string `protobuf:"bytes,4,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *HttpGrpcAccessLogConfig) Reset()         { *m = HttpGrpcAccessLogConfig{} }
func (m *HttpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*HttpGrpcAccessLogConfig) ProtoMessage()    {}
func (*HttpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_938e3858beb7bdc4, []int{0}
}
func (m *HttpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpGrpcAccessLogConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpGrpcAccessLogConfig.Merge(m, src)
}
func (m *HttpGrpcAccessLogConfig) XXX_Size() int {
	return m.Size()
}
func (m *HttpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *HttpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

// Configuration for the built-in *envoy.tcp_grpc_access_log* type. This configuration will
// populate *StreamAccessLogsMessage.tcp_logs*.
type TcpGrpcAccessLogConfig struct {
	CommonConfig         *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TcpGrpcAccessLogConfig) Reset()         { *m = TcpGrpcAccessLogConfig{} }
func (m *TcpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*TcpGrpcAccessLogConfig) ProtoMessage()    {}
func (*TcpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_938e3858beb7bdc4, []int{1}
}
func (m *TcpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TcpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TcpGrpcAccessLogConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TcpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpGrpcAccessLogConfig.Merge(m, src)
}
func (m *TcpGrpcAccessLogConfig) XXX_Size() int {
	return m.Size()
}
func (m *TcpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *TcpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

// Common configuration for gRPC access logs.
type CommonGrpcAccessLogConfig struct {
	// The friendly name of the access log to be returned in :ref:`StreamAccessLogsMessage.Identifier
	// <envoy_api_msg_service.accesslog.v3alpha.StreamAccessLogsMessage.Identifier>`. This allows the
	// access log server to differentiate between different access logs coming from the same Envoy.
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// The gRPC service for the access log service.
	GrpcService *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	// Interval for flushing access logs to the gRPC stream. Logger will flush requests every time
	// this interval is elapsed, or when batch size limit is hit, whichever comes first. Defaults to
	// 1 second.
	BufferFlushInterval *types.Duration `protobuf:"bytes,3,opt,name=buffer_flush_interval,json=bufferFlushInterval,proto3" json:"buffer_flush_interval,omitempty"`
	// Soft size limit in bytes for access log entries buffer. Logger will buffer requests until
	// this limit it hit, or every time flush interval is elapsed, whichever comes first. Setting it
	// to zero effectively disables the batching. Defaults to 16384.
	BufferSizeBytes      *types.UInt32Value `protobuf:"bytes,4,opt,name=buffer_size_bytes,json=bufferSizeBytes,proto3" json:"buffer_size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CommonGrpcAccessLogConfig) Reset()         { *m = CommonGrpcAccessLogConfig{} }
func (m *CommonGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*CommonGrpcAccessLogConfig) ProtoMessage()    {}
func (*CommonGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_938e3858beb7bdc4, []int{2}
}
func (m *CommonGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommonGrpcAccessLogConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommonGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonGrpcAccessLogConfig.Merge(m, src)
}
func (m *CommonGrpcAccessLogConfig) XXX_Size() int {
	return m.Size()
}
func (m *CommonGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonGrpcAccessLogConfig proto.InternalMessageInfo

func (m *CommonGrpcAccessLogConfig) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *CommonGrpcAccessLogConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferFlushInterval() *types.Duration {
	if m != nil {
		return m.BufferFlushInterval
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferSizeBytes() *types.UInt32Value {
	if m != nil {
		return m.BufferSizeBytes
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v3alpha.HttpGrpcAccessLogConfig")
	proto.RegisterType((*TcpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v3alpha.TcpGrpcAccessLogConfig")
	proto.RegisterType((*CommonGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v3alpha.CommonGrpcAccessLogConfig")
}

func init() {
	proto.RegisterFile("envoy/config/accesslog/v3alpha/als.proto", fileDescriptor_938e3858beb7bdc4)
}

var fileDescriptor_938e3858beb7bdc4 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xbf, 0x75, 0xfa, 0x41, 0xbb, 0x2d, 0x02, 0x8c, 0xa0, 0x69, 0x04, 0x26, 0xa4, 0x1c,
	0x02, 0x42, 0xb6, 0xd4, 0x9c, 0x38, 0x36, 0x45, 0x25, 0x45, 0x11, 0xaa, 0xdc, 0x80, 0xc4, 0xc9,
	0xda, 0x38, 0x9b, 0xcd, 0xa2, 0x8d, 0x67, 0xd9, 0x5d, 0x07, 0xd2, 0x33, 0x27, 0x1e, 0x85, 0x47,
	0xe8, 0x89, 0x23, 0x47, 0x1e, 0x01, 0xe5, 0xd6, 0xb7, 0x40, 0xde, 0x75, 0xd3, 0x84, 0xb6, 0x70,
	0xe4, 0x96, 0x68, 0x7e, 0xf3, 0x9b, 0xff, 0xd8, 0x63, 0xdc, 0xa4, 0xd9, 0x04, 0xa6, 0x51, 0x0a,
	0xd9, 0x90, 0xb3, 0x88, 0xa4, 0x29, 0xd5, 0x5a, 0x00, 0x8b, 0x26, 0x2d, 0x22, 0xe4, 0x88, 0x44,
	0x44, 0xe8, 0x50, 0x2a, 0x30, 0xe0, 0x07, 0x96, 0x0c, 0x1d, 0x19, 0xce, 0xc9, 0xb0, 0x24, 0x6b,
	0x4f, 0x9c, 0x89, 0x48, 0x3e, 0x6f, 0x4e, 0x41, 0xd1, 0x88, 0x29, 0x99, 0x26, 0x9a, 0xaa, 0x09,
	0x4f, 0xa9, 0x53, 0xd5, 0x02, 0x06, 0xc0, 0x04, 0x8d, 0xec, 0xbf, 0x7e, 0x3e, 0x8c, 0x06, 0xb9,
	0x22, 0x86, 0x43, 0x76, 0x55, 0xfd, 0xa3, 0x22, 0x52, 0x52, 0x55, 0x46, 0xa9, 0x6d, 0x4e, 0x88,
	0xe0, 0x03, 0x62, 0x68, 0x74, 0xf6, 0xc3, 0x15, 0x1a, 0xa7, 0x1e, 0xde, 0xec, 0x18, 0x23, 0x5f,
	0x2a, 0x99, 0xee, 0xda, 0x84, 0x5d, 0x60, 0x7b, 0x36, 0xb1, 0xff, 0x1e, 0xdf, 0x48, 0x61, 0x3c,
	0x86, 0x2c, 0x71, 0x2b, 0x54, 0x51, 0x1d, 0x35, 0xd7, 0x77, 0x9e, 0x87, 0x7f, 0xde, 0x2b, 0xdc,
	0xb3, 0x4d, 0x97, 0x18, 0xdb, 0xf8, 0xe4, 0xf4, 0x5b, 0xe5, 0xff, 0x2f, 0xc8, 0xbb, 0x85, 0xe2,
	0x0d, 0xe7, 0x2e, 0x67, 0x75, 0xf0, 0x23, 0x32, 0x18, 0xf0, 0x62, 0x25, 0x22, 0x12, 0x45, 0x3f,
	0xe4, 0x54, 0x9b, 0x64, 0x44, 0xc9, 0x80, 0x2a, 0x9d, 0x18, 0x48, 0x04, 0xb0, 0xaa, 0x57, 0xaf,
	0x34, 0xd7, 0xe2, 0x07, 0xe7, 0x60, 0xec, 0xb8, 0x8e, 0xc3, 0x7a, 0xd0, 0x05, 0xe6, 0xbf, 0xc2,
	0x8d, 0x25, 0x93, 0x96, 0x90, 0x69, 0xfa, 0xbb, 0xaa, 0x62, 0x55, 0xc1, 0xa2, 0xca, 0x81, 0x4b,
	0xae, 0x2e, 0xde, 0xbe, 0xcc, 0x65, 0x14, 0xe1, 0x62, 0x41, 0xb6, 0x62, 0x65, 0x0f, 0x2f, 0xca,
	0x7a, 0x25, 0x68, 0x6d, 0x8d, 0xcf, 0x08, 0xdf, 0xeb, 0xa5, 0xff, 0xfa, 0x51, 0x37, 0x4e, 0x3c,
	0xbc, 0x75, 0x65, 0x9f, 0xff, 0x18, 0xaf, 0x0a, 0x60, 0x49, 0x46, 0xc6, 0xd4, 0x86, 0x58, 0x6b,
	0xaf, 0x15, 0xa6, 0x15, 0xe5, 0xd5, 0x51, 0x7c, 0x5d, 0x00, 0x7b, 0x4d, 0xc6, 0xd4, 0x8f, 0xf1,
	0xc6, 0xe2, 0x95, 0x56, 0x3d, 0x1b, 0x77, 0xbb, 0x8c, 0x4b, 0x24, 0x9f, 0x27, 0x2c, 0x2e, 0x3a,
	0x2c, 0x06, 0x1d, 0x39, 0x74, 0x29, 0xd8, 0x3a, 0x3b, 0x2f, 0xf8, 0xef, 0xf0, 0xdd, 0x7e, 0x3e,
	0x1c, 0x52, 0x95, 0x0c, 0x45, 0xae, 0x47, 0x09, 0xcf, 0x0c, 0x55, 0x13, 0x22, 0xaa, 0x15, 0x2b,
	0xdf, 0x0a, 0xdd, 0x8d, 0x87, 0x67, 0x37, 0x1e, 0xbe, 0x28, 0xbf, 0x81, 0x52, 0xf9, 0x15, 0x79,
	0x4f, 0xff, 0x8b, 0xef, 0x38, 0xc7, 0x7e, 0xa1, 0x38, 0x28, 0x0d, 0x7e, 0x07, 0xdf, 0x2e, 0xd5,
	0x9a, 0x1f, 0xd3, 0xa4, 0x3f, 0x35, 0x54, 0x57, 0x57, 0xac, 0xf6, 0xfe, 0x05, 0xed, 0x9b, 0x83,
	0xcc, 0xb4, 0x76, 0xde, 0x12, 0x91, 0xd3, 0xf8, 0xa6, 0x6b, 0x3b, 0xe2, 0xc7, 0xb4, 0x5d, 0x34,
	0xb5, 0xf7, 0xbf, 0xcf, 0x02, 0xf4, 0x63, 0x16, 0xa0, 0x9f, 0xb3, 0x00, 0xe1, 0x67, 0x1c, 0xdc,
	0xca, 0x52, 0xc1, 0xa7, 0xe9, 0x5f, 0x5e, 0x56, 0x7b, 0x75, 0x57, 0xe8, 0xc3, 0x62, 0xca, 0x21,
	0xea, 0x5f, 0xb3, 0xe3, 0x5a, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x8b, 0xeb, 0x51, 0x4e,
	0x04, 0x00, 0x00,
}

func (m *HttpGrpcAccessLogConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpGrpcAccessLogConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpGrpcAccessLogConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.AdditionalResponseTrailersToLog) > 0 {
		for iNdEx := len(m.AdditionalResponseTrailersToLog) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AdditionalResponseTrailersToLog[iNdEx])
			copy(dAtA[i:], m.AdditionalResponseTrailersToLog[iNdEx])
			i = encodeVarintAls(dAtA, i, uint64(len(m.AdditionalResponseTrailersToLog[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AdditionalResponseHeadersToLog) > 0 {
		for iNdEx := len(m.AdditionalResponseHeadersToLog) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AdditionalResponseHeadersToLog[iNdEx])
			copy(dAtA[i:], m.AdditionalResponseHeadersToLog[iNdEx])
			i = encodeVarintAls(dAtA, i, uint64(len(m.AdditionalResponseHeadersToLog[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AdditionalRequestHeadersToLog) > 0 {
		for iNdEx := len(m.AdditionalRequestHeadersToLog) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AdditionalRequestHeadersToLog[iNdEx])
			copy(dAtA[i:], m.AdditionalRequestHeadersToLog[iNdEx])
			i = encodeVarintAls(dAtA, i, uint64(len(m.AdditionalRequestHeadersToLog[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CommonConfig != nil {
		{
			size, err := m.CommonConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAls(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TcpGrpcAccessLogConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TcpGrpcAccessLogConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TcpGrpcAccessLogConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CommonConfig != nil {
		{
			size, err := m.CommonConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAls(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommonGrpcAccessLogConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonGrpcAccessLogConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonGrpcAccessLogConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.BufferSizeBytes != nil {
		{
			size, err := m.BufferSizeBytes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAls(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.BufferFlushInterval != nil {
		{
			size, err := m.BufferFlushInterval.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAls(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.GrpcService != nil {
		{
			size, err := m.GrpcService.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAls(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.LogName) > 0 {
		i -= len(m.LogName)
		copy(dAtA[i:], m.LogName)
		i = encodeVarintAls(dAtA, i, uint64(len(m.LogName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAls(dAtA []byte, offset int, v uint64) int {
	offset -= sovAls(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HttpGrpcAccessLogConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommonConfig != nil {
		l = m.CommonConfig.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if len(m.AdditionalRequestHeadersToLog) > 0 {
		for _, s := range m.AdditionalRequestHeadersToLog {
			l = len(s)
			n += 1 + l + sovAls(uint64(l))
		}
	}
	if len(m.AdditionalResponseHeadersToLog) > 0 {
		for _, s := range m.AdditionalResponseHeadersToLog {
			l = len(s)
			n += 1 + l + sovAls(uint64(l))
		}
	}
	if len(m.AdditionalResponseTrailersToLog) > 0 {
		for _, s := range m.AdditionalResponseTrailersToLog {
			l = len(s)
			n += 1 + l + sovAls(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TcpGrpcAccessLogConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommonConfig != nil {
		l = m.CommonConfig.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CommonGrpcAccessLogConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LogName)
	if l > 0 {
		n += 1 + l + sovAls(uint64(l))
	}
	if m.GrpcService != nil {
		l = m.GrpcService.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if m.BufferFlushInterval != nil {
		l = m.BufferFlushInterval.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if m.BufferSizeBytes != nil {
		l = m.BufferSizeBytes.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAls(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAls(x uint64) (n int) {
	return sovAls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HttpGrpcAccessLogConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HttpGrpcAccessLogConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpGrpcAccessLogConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CommonConfig == nil {
				m.CommonConfig = &CommonGrpcAccessLogConfig{}
			}
			if err := m.CommonConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalRequestHeadersToLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalRequestHeadersToLog = append(m.AdditionalRequestHeadersToLog, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalResponseHeadersToLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalResponseHeadersToLog = append(m.AdditionalResponseHeadersToLog, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalResponseTrailersToLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalResponseTrailersToLog = append(m.AdditionalResponseTrailersToLog, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TcpGrpcAccessLogConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TcpGrpcAccessLogConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TcpGrpcAccessLogConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CommonConfig == nil {
				m.CommonConfig = &CommonGrpcAccessLogConfig{}
			}
			if err := m.CommonConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommonGrpcAccessLogConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommonGrpcAccessLogConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonGrpcAccessLogConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrpcService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GrpcService == nil {
				m.GrpcService = &core.GrpcService{}
			}
			if err := m.GrpcService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferFlushInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BufferFlushInterval == nil {
				m.BufferFlushInterval = &types.Duration{}
			}
			if err := m.BufferFlushInterval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferSizeBytes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BufferSizeBytes == nil {
				m.BufferSizeBytes = &types.UInt32Value{}
			}
			if err := m.BufferSizeBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAls
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAls
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAls
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAls
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAls
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAls
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAls(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAls
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAls = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAls   = fmt.Errorf("proto: integer overflow")
)
