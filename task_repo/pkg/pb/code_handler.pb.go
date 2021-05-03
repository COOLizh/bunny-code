// Code generated by protoc-gen-go. DO NOT EDIT.
// source: codeHandler.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CodeHandleRequest struct {
	Solution             []byte                        `protobuf:"bytes,1,opt,name=solution,proto3" json:"solution,omitempty"`
	MemoryLimit          int64                         `protobuf:"varint,2,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	TimeLimit            int64                         `protobuf:"varint,3,opt,name=time_limit,json=timeLimit,proto3" json:"time_limit,omitempty"`
	Language             string                        `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	TestCases            []*CodeHandleRequest_TestCase `protobuf:"bytes,5,rep,name=test_cases,json=testCases,proto3" json:"test_cases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CodeHandleRequest) Reset()         { *m = CodeHandleRequest{} }
func (m *CodeHandleRequest) String() string { return proto.CompactTextString(m) }
func (*CodeHandleRequest) ProtoMessage()    {}
func (*CodeHandleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad23e5995772978a, []int{0}
}

func (m *CodeHandleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeHandleRequest.Unmarshal(m, b)
}
func (m *CodeHandleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeHandleRequest.Marshal(b, m, deterministic)
}
func (m *CodeHandleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeHandleRequest.Merge(m, src)
}
func (m *CodeHandleRequest) XXX_Size() int {
	return xxx_messageInfo_CodeHandleRequest.Size(m)
}
func (m *CodeHandleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeHandleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CodeHandleRequest proto.InternalMessageInfo

func (m *CodeHandleRequest) GetSolution() []byte {
	if m != nil {
		return m.Solution
	}
	return nil
}

func (m *CodeHandleRequest) GetMemoryLimit() int64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *CodeHandleRequest) GetTimeLimit() int64 {
	if m != nil {
		return m.TimeLimit
	}
	return 0
}

func (m *CodeHandleRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *CodeHandleRequest) GetTestCases() []*CodeHandleRequest_TestCase {
	if m != nil {
		return m.TestCases
	}
	return nil
}

type CodeHandleRequest_TestCase struct {
	TestData             []byte   `protobuf:"bytes,1,opt,name=test_data,json=testData,proto3" json:"test_data,omitempty"`
	Answer               []byte   `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeHandleRequest_TestCase) Reset()         { *m = CodeHandleRequest_TestCase{} }
func (m *CodeHandleRequest_TestCase) String() string { return proto.CompactTextString(m) }
func (*CodeHandleRequest_TestCase) ProtoMessage()    {}
func (*CodeHandleRequest_TestCase) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad23e5995772978a, []int{0, 0}
}

func (m *CodeHandleRequest_TestCase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeHandleRequest_TestCase.Unmarshal(m, b)
}
func (m *CodeHandleRequest_TestCase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeHandleRequest_TestCase.Marshal(b, m, deterministic)
}
func (m *CodeHandleRequest_TestCase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeHandleRequest_TestCase.Merge(m, src)
}
func (m *CodeHandleRequest_TestCase) XXX_Size() int {
	return xxx_messageInfo_CodeHandleRequest_TestCase.Size(m)
}
func (m *CodeHandleRequest_TestCase) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeHandleRequest_TestCase.DiscardUnknown(m)
}

var xxx_messageInfo_CodeHandleRequest_TestCase proto.InternalMessageInfo

func (m *CodeHandleRequest_TestCase) GetTestData() []byte {
	if m != nil {
		return m.TestData
	}
	return nil
}

func (m *CodeHandleRequest_TestCase) GetAnswer() []byte {
	if m != nil {
		return m.Answer
	}
	return nil
}

type CodeHandleResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	JobCreated           bool     `protobuf:"varint,2,opt,name=job_created,json=jobCreated,proto3" json:"job_created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeHandleResponse) Reset()         { *m = CodeHandleResponse{} }
func (m *CodeHandleResponse) String() string { return proto.CompactTextString(m) }
func (*CodeHandleResponse) ProtoMessage()    {}
func (*CodeHandleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad23e5995772978a, []int{1}
}

func (m *CodeHandleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeHandleResponse.Unmarshal(m, b)
}
func (m *CodeHandleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeHandleResponse.Marshal(b, m, deterministic)
}
func (m *CodeHandleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeHandleResponse.Merge(m, src)
}
func (m *CodeHandleResponse) XXX_Size() int {
	return xxx_messageInfo_CodeHandleResponse.Size(m)
}
func (m *CodeHandleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeHandleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CodeHandleResponse proto.InternalMessageInfo

func (m *CodeHandleResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *CodeHandleResponse) GetJobCreated() bool {
	if m != nil {
		return m.JobCreated
	}
	return false
}

type StatusHandleRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusHandleRequest) Reset()         { *m = StatusHandleRequest{} }
func (m *StatusHandleRequest) String() string { return proto.CompactTextString(m) }
func (*StatusHandleRequest) ProtoMessage()    {}
func (*StatusHandleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad23e5995772978a, []int{2}
}

func (m *StatusHandleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusHandleRequest.Unmarshal(m, b)
}
func (m *StatusHandleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusHandleRequest.Marshal(b, m, deterministic)
}
func (m *StatusHandleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusHandleRequest.Merge(m, src)
}
func (m *StatusHandleRequest) XXX_Size() int {
	return xxx_messageInfo_StatusHandleRequest.Size(m)
}
func (m *StatusHandleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusHandleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusHandleRequest proto.InternalMessageInfo

func (m *StatusHandleRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type StatusHandleResponse struct {
	ID                   string                          `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Ready                bool                            `protobuf:"varint,2,opt,name=ready,proto3" json:"ready,omitempty"`
	TestsData            *StatusHandleResponse_TestsData `protobuf:"bytes,3,opt,name=tests_data,json=testsData,proto3" json:"tests_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *StatusHandleResponse) Reset()         { *m = StatusHandleResponse{} }
func (m *StatusHandleResponse) String() string { return proto.CompactTextString(m) }
func (*StatusHandleResponse) ProtoMessage()    {}
func (*StatusHandleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad23e5995772978a, []int{3}
}

func (m *StatusHandleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusHandleResponse.Unmarshal(m, b)
}
func (m *StatusHandleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusHandleResponse.Marshal(b, m, deterministic)
}
func (m *StatusHandleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusHandleResponse.Merge(m, src)
}
func (m *StatusHandleResponse) XXX_Size() int {
	return xxx_messageInfo_StatusHandleResponse.Size(m)
}
func (m *StatusHandleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusHandleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusHandleResponse proto.InternalMessageInfo

func (m *StatusHandleResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *StatusHandleResponse) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

func (m *StatusHandleResponse) GetTestsData() *StatusHandleResponse_TestsData {
	if m != nil {
		return m.TestsData
	}
	return nil
}

type StatusHandleResponse_TestsData struct {
	PassedTestsCount     int64                                        `protobuf:"varint,1,opt,name=passed_tests_count,json=passedTestsCount,proto3" json:"passed_tests_count,omitempty"`
	TestResults          []*StatusHandleResponse_TestsData_TestResult `protobuf:"bytes,2,rep,name=test_results,json=testResults,proto3" json:"test_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *StatusHandleResponse_TestsData) Reset()         { *m = StatusHandleResponse_TestsData{} }
func (m *StatusHandleResponse_TestsData) String() string { return proto.CompactTextString(m) }
func (*StatusHandleResponse_TestsData) ProtoMessage()    {}
func (*StatusHandleResponse_TestsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad23e5995772978a, []int{3, 0}
}

func (m *StatusHandleResponse_TestsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusHandleResponse_TestsData.Unmarshal(m, b)
}
func (m *StatusHandleResponse_TestsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusHandleResponse_TestsData.Marshal(b, m, deterministic)
}
func (m *StatusHandleResponse_TestsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusHandleResponse_TestsData.Merge(m, src)
}
func (m *StatusHandleResponse_TestsData) XXX_Size() int {
	return xxx_messageInfo_StatusHandleResponse_TestsData.Size(m)
}
func (m *StatusHandleResponse_TestsData) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusHandleResponse_TestsData.DiscardUnknown(m)
}

var xxx_messageInfo_StatusHandleResponse_TestsData proto.InternalMessageInfo

func (m *StatusHandleResponse_TestsData) GetPassedTestsCount() int64 {
	if m != nil {
		return m.PassedTestsCount
	}
	return 0
}

func (m *StatusHandleResponse_TestsData) GetTestResults() []*StatusHandleResponse_TestsData_TestResult {
	if m != nil {
		return m.TestResults
	}
	return nil
}

type StatusHandleResponse_TestsData_TestResult struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	TimeSpent            int64    `protobuf:"varint,2,opt,name=time_spent,json=timeSpent,proto3" json:"time_spent,omitempty"`
	MemorySpent          int64    `protobuf:"varint,3,opt,name=memory_spent,json=memorySpent,proto3" json:"memory_spent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusHandleResponse_TestsData_TestResult) Reset() {
	*m = StatusHandleResponse_TestsData_TestResult{}
}
func (m *StatusHandleResponse_TestsData_TestResult) String() string {
	return proto.CompactTextString(m)
}
func (*StatusHandleResponse_TestsData_TestResult) ProtoMessage() {}
func (*StatusHandleResponse_TestsData_TestResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad23e5995772978a, []int{3, 0, 0}
}

func (m *StatusHandleResponse_TestsData_TestResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusHandleResponse_TestsData_TestResult.Unmarshal(m, b)
}
func (m *StatusHandleResponse_TestsData_TestResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusHandleResponse_TestsData_TestResult.Marshal(b, m, deterministic)
}
func (m *StatusHandleResponse_TestsData_TestResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusHandleResponse_TestsData_TestResult.Merge(m, src)
}
func (m *StatusHandleResponse_TestsData_TestResult) XXX_Size() int {
	return xxx_messageInfo_StatusHandleResponse_TestsData_TestResult.Size(m)
}
func (m *StatusHandleResponse_TestsData_TestResult) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusHandleResponse_TestsData_TestResult.DiscardUnknown(m)
}

var xxx_messageInfo_StatusHandleResponse_TestsData_TestResult proto.InternalMessageInfo

func (m *StatusHandleResponse_TestsData_TestResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *StatusHandleResponse_TestsData_TestResult) GetTimeSpent() int64 {
	if m != nil {
		return m.TimeSpent
	}
	return 0
}

func (m *StatusHandleResponse_TestsData_TestResult) GetMemorySpent() int64 {
	if m != nil {
		return m.MemorySpent
	}
	return 0
}

func init() {
	proto.RegisterType((*CodeHandleRequest)(nil), "pb.CodeHandleRequest")
	proto.RegisterType((*CodeHandleRequest_TestCase)(nil), "pb.CodeHandleRequest.TestCase")
	proto.RegisterType((*CodeHandleResponse)(nil), "pb.CodeHandleResponse")
	proto.RegisterType((*StatusHandleRequest)(nil), "pb.StatusHandleRequest")
	proto.RegisterType((*StatusHandleResponse)(nil), "pb.StatusHandleResponse")
	proto.RegisterType((*StatusHandleResponse_TestsData)(nil), "pb.StatusHandleResponse.TestsData")
	proto.RegisterType((*StatusHandleResponse_TestsData_TestResult)(nil), "pb.StatusHandleResponse.TestsData.TestResult")
}

func init() { proto.RegisterFile("codeHandler.proto", fileDescriptor_ad23e5995772978a) }

var fileDescriptor_ad23e5995772978a = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x6d, 0x5a, 0xd9, 0xe3, 0x80, 0xe8, 0x52, 0x8a, 0x65, 0x04, 0x04, 0x4b, 0x48, 0x39,
	0x40, 0x90, 0xc2, 0x19, 0x21, 0xea, 0x20, 0x51, 0xa9, 0x07, 0xb4, 0xe1, 0xc4, 0xc5, 0x5a, 0xdb,
	0x43, 0x48, 0xeb, 0x78, 0x8d, 0x67, 0x2d, 0xd4, 0x1f, 0xc1, 0xaf, 0x45, 0xdc, 0xd1, 0xee, 0x3a,
	0x4e, 0xfa, 0x25, 0x6e, 0x33, 0x6f, 0xc6, 0x6f, 0x67, 0xde, 0x3c, 0xc3, 0x61, 0x21, 0x4b, 0xfc,
	0x2c, 0xea, 0xb2, 0xc2, 0x76, 0xda, 0xb4, 0x52, 0x49, 0xe6, 0x36, 0x79, 0xf2, 0xdb, 0x85, 0xc3,
	0x74, 0xa8, 0x70, 0xfc, 0xd9, 0x21, 0x29, 0x16, 0x83, 0x4f, 0xb2, 0xea, 0xd4, 0x4a, 0xd6, 0x91,
	0x33, 0x76, 0x26, 0x23, 0x3e, 0xe4, 0xec, 0x25, 0x8c, 0xd6, 0xb8, 0x96, 0xed, 0x65, 0x56, 0xad,
	0xd6, 0x2b, 0x15, 0xb9, 0x63, 0x67, 0xe2, 0xf1, 0xd0, 0x62, 0x67, 0x1a, 0x62, 0xcf, 0x00, 0xd4,
	0x6a, 0x8d, 0x7d, 0x83, 0x67, 0x1a, 0x02, 0x8d, 0xd8, 0x72, 0x0c, 0x7e, 0x25, 0xea, 0x65, 0x27,
	0x96, 0x18, 0xdd, 0x1b, 0x3b, 0x93, 0x80, 0x0f, 0x39, 0x7b, 0x0f, 0xa0, 0x90, 0x54, 0x56, 0x08,
	0x42, 0x8a, 0xf6, 0xc7, 0xde, 0x24, 0x9c, 0x3d, 0x9f, 0x36, 0xf9, 0xf4, 0xc6, 0x90, 0xd3, 0xaf,
	0x48, 0x2a, 0x15, 0x84, 0x3c, 0x50, 0x7d, 0x44, 0xf1, 0x07, 0xf0, 0x37, 0x30, 0x7b, 0x0a, 0xa6,
	0x90, 0x95, 0x42, 0x89, 0xcd, 0x16, 0x1a, 0x98, 0x0b, 0x25, 0xd8, 0x31, 0x1c, 0x88, 0x9a, 0x7e,
	0x61, 0x6b, 0xe6, 0x1f, 0xf1, 0x3e, 0x4b, 0x3e, 0x01, 0xdb, 0x7d, 0x89, 0x1a, 0x59, 0x13, 0xb2,
	0x07, 0xe0, 0x9e, 0xce, 0x0d, 0x47, 0xc0, 0xdd, 0xd3, 0x39, 0x7b, 0x01, 0xe1, 0xb9, 0xcc, 0xb3,
	0xa2, 0x45, 0xa1, 0xb0, 0x34, 0x14, 0x3e, 0x87, 0x73, 0x99, 0xa7, 0x16, 0x49, 0x5e, 0xc1, 0xa3,
	0x85, 0x12, 0xaa, 0xa3, 0xab, 0xba, 0x5e, 0xe3, 0x49, 0xfe, 0xba, 0x70, 0x74, 0xb5, 0xef, 0x8e,
	0x07, 0x8f, 0x60, 0xbf, 0x45, 0x51, 0x5e, 0xf6, 0x4f, 0xd9, 0x84, 0x7d, 0xb4, 0x62, 0x91, 0x5d,
	0x51, 0xeb, 0x1c, 0xce, 0x12, 0x2d, 0xd6, 0x6d, 0x9c, 0x46, 0x2f, 0xd2, 0xcb, 0x5b, 0xc1, 0x4c,
	0x18, 0xff, 0x71, 0x20, 0x18, 0x0a, 0xec, 0x35, 0xb0, 0x46, 0x10, 0x61, 0x99, 0x59, 0xde, 0x42,
	0x76, 0xb5, 0x32, 0x63, 0x78, 0xfc, 0xa1, 0xad, 0x98, 0xe6, 0x54, 0xe3, 0xec, 0x0b, 0x8c, 0x8c,
	0xc0, 0x2d, 0x52, 0x57, 0x29, 0x8a, 0x5c, 0x73, 0xad, 0x37, 0xff, 0x1f, 0xc0, 0x44, 0xdc, 0x7c,
	0xc5, 0x43, 0x35, 0xc4, 0x14, 0x7f, 0x07, 0xd8, 0x96, 0xf4, 0x8d, 0x2c, 0x75, 0x2f, 0x44, 0x9f,
	0x0d, 0xf6, 0xa2, 0x06, 0xeb, 0x8d, 0xff, 0x8c, 0xbd, 0x16, 0x1a, 0xd8, 0x31, 0xa8, 0x6d, 0xf0,
	0x76, 0x0d, 0x6a, 0x5a, 0x66, 0x67, 0x10, 0x6e, 0xaf, 0xdc, 0x6a, 0xd3, 0x6d, 0x53, 0xf6, 0xf8,
	0x56, 0xbb, 0xc5, 0xc7, 0xd7, 0x61, 0xbb, 0x55, 0xb2, 0x37, 0x5b, 0xc0, 0xfd, 0xdd, 0x7d, 0x5b,
	0x76, 0x02, 0xa1, 0x05, 0xd2, 0x1f, 0x58, 0x5c, 0xb0, 0x27, 0x37, 0x15, 0xb1, 0x94, 0xd1, 0x5d,
	0x52, 0x25, 0x7b, 0x27, 0xfe, 0xb7, 0x83, 0xe6, 0x62, 0xf9, 0xb6, 0xc9, 0xf3, 0x03, 0xf3, 0xb7,
	0xbe, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x6b, 0x84, 0xa4, 0xc2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CodeHandlerClient is the client API for CodeHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CodeHandlerClient interface {
	CodeHandle(ctx context.Context, in *CodeHandleRequest, opts ...grpc.CallOption) (*CodeHandleResponse, error)
}

type codeHandlerClient struct {
	cc *grpc.ClientConn
}

func NewCodeHandlerClient(cc *grpc.ClientConn) CodeHandlerClient {
	return &codeHandlerClient{cc}
}

func (c *codeHandlerClient) CodeHandle(ctx context.Context, in *CodeHandleRequest, opts ...grpc.CallOption) (*CodeHandleResponse, error) {
	out := new(CodeHandleResponse)
	err := c.cc.Invoke(ctx, "/pb.CodeHandler/CodeHandle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeHandlerServer is the server API for CodeHandler service.
type CodeHandlerServer interface {
	CodeHandle(context.Context, *CodeHandleRequest) (*CodeHandleResponse, error)
}

// UnimplementedCodeHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedCodeHandlerServer struct {
}

func (*UnimplementedCodeHandlerServer) CodeHandle(ctx context.Context, req *CodeHandleRequest) (*CodeHandleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeHandle not implemented")
}

func RegisterCodeHandlerServer(s *grpc.Server, srv CodeHandlerServer) {
	s.RegisterService(&_CodeHandler_serviceDesc, srv)
}

func _CodeHandler_CodeHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeHandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeHandlerServer).CodeHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CodeHandler/CodeHandle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeHandlerServer).CodeHandle(ctx, req.(*CodeHandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CodeHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CodeHandler",
	HandlerType: (*CodeHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CodeHandle",
			Handler:    _CodeHandler_CodeHandle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "code_handler.proto",
}

// StatusHandlerClient is the client API for StatusHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatusHandlerClient interface {
	StatusCheck(ctx context.Context, in *StatusHandleRequest, opts ...grpc.CallOption) (*StatusHandleResponse, error)
}

type statusHandlerClient struct {
	cc *grpc.ClientConn
}

func NewStatusHandlerClient(cc *grpc.ClientConn) StatusHandlerClient {
	return &statusHandlerClient{cc}
}

func (c *statusHandlerClient) StatusCheck(ctx context.Context, in *StatusHandleRequest, opts ...grpc.CallOption) (*StatusHandleResponse, error) {
	out := new(StatusHandleResponse)
	err := c.cc.Invoke(ctx, "/pb.StatusHandler/StatusCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatusHandlerServer is the server API for StatusHandler service.
type StatusHandlerServer interface {
	StatusCheck(context.Context, *StatusHandleRequest) (*StatusHandleResponse, error)
}

// UnimplementedStatusHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedStatusHandlerServer struct {
}

func (*UnimplementedStatusHandlerServer) StatusCheck(ctx context.Context, req *StatusHandleRequest) (*StatusHandleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusCheck not implemented")
}

func RegisterStatusHandlerServer(s *grpc.Server, srv StatusHandlerServer) {
	s.RegisterService(&_StatusHandler_serviceDesc, srv)
}

func _StatusHandler_StatusCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusHandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusHandlerServer).StatusCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StatusHandler/StatusCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusHandlerServer).StatusCheck(ctx, req.(*StatusHandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatusHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StatusHandler",
	HandlerType: (*StatusHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StatusCheck",
			Handler:    _StatusHandler_StatusCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "code_handler.proto",
}
