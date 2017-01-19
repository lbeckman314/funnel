// Code generated by protoc-gen-go.
// source: task_worker.proto
// DO NOT EDIT!

/*
Package ga4gh_task_ref is a generated protocol buffer package.

It is generated from these files:
	task_worker.proto

It has these top-level messages:
	LocalStorage
	S3Storage
	StorageConfig
	ServerConfig
	WorkerInfo
	JobRequest
	JobResponse
	UpdateStatusRequest
	QueuedTaskInfoRequest
	QueuedTaskInfo
*/
package ga4gh_task_ref

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ga4gh_task_exec "tes/ga4gh"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LocalStorage struct {
	AllowedDirs []string `protobuf:"bytes,1,rep,name=allowed_dirs,json=allowedDirs" json:"allowed_dirs,omitempty"`
}

func (m *LocalStorage) Reset()                    { *m = LocalStorage{} }
func (m *LocalStorage) String() string            { return proto.CompactTextString(m) }
func (*LocalStorage) ProtoMessage()               {}
func (*LocalStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LocalStorage) GetAllowedDirs() []string {
	if m != nil {
		return m.AllowedDirs
	}
	return nil
}

type S3Storage struct {
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	Key      string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Secret   string `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
}

func (m *S3Storage) Reset()                    { *m = S3Storage{} }
func (m *S3Storage) String() string            { return proto.CompactTextString(m) }
func (*S3Storage) ProtoMessage()               {}
func (*S3Storage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S3Storage) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *S3Storage) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *S3Storage) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type StorageConfig struct {
	Local *LocalStorage `protobuf:"bytes,1,opt,name=local" json:"local,omitempty"`
	S3    *S3Storage    `protobuf:"bytes,2,opt,name=s3" json:"s3,omitempty"`
}

func (m *StorageConfig) Reset()                    { *m = StorageConfig{} }
func (m *StorageConfig) String() string            { return proto.CompactTextString(m) }
func (*StorageConfig) ProtoMessage()               {}
func (*StorageConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StorageConfig) GetLocal() *LocalStorage {
	if m != nil {
		return m.Local
	}
	return nil
}

func (m *StorageConfig) GetS3() *S3Storage {
	if m != nil {
		return m.S3
	}
	return nil
}

type ServerConfig struct {
	Storage       []*StorageConfig `protobuf:"bytes,1,rep,name=storage" json:"storage,omitempty"`
	Secret        string           `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	ServerAddress string           `protobuf:"bytes,3,opt,name=server_address,json=serverAddress" json:"server_address,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ServerConfig) GetStorage() []*StorageConfig {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *ServerConfig) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *ServerConfig) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

// *
// Worker Info
type WorkerInfo struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	LastPing int64  `protobuf:"varint,2,opt,name=last_ping,json=lastPing" json:"last_ping,omitempty"`
	Hostname string `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
}

func (m *WorkerInfo) Reset()                    { *m = WorkerInfo{} }
func (m *WorkerInfo) String() string            { return proto.CompactTextString(m) }
func (*WorkerInfo) ProtoMessage()               {}
func (*WorkerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WorkerInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WorkerInfo) GetLastPing() int64 {
	if m != nil {
		return m.LastPing
	}
	return 0
}

func (m *WorkerInfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type JobRequest struct {
	Worker    *WorkerInfo                `protobuf:"bytes,1,opt,name=worker" json:"worker,omitempty"`
	Resources *ga4gh_task_exec.Resources `protobuf:"bytes,2,opt,name=resources" json:"resources,omitempty"`
}

func (m *JobRequest) Reset()                    { *m = JobRequest{} }
func (m *JobRequest) String() string            { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()               {}
func (*JobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *JobRequest) GetWorker() *WorkerInfo {
	if m != nil {
		return m.Worker
	}
	return nil
}

func (m *JobRequest) GetResources() *ga4gh_task_exec.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

type JobResponse struct {
	Job     *ga4gh_task_exec.Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	Auth    string               `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
	Storage []*StorageConfig     `protobuf:"bytes,3,rep,name=storage" json:"storage,omitempty"`
}

func (m *JobResponse) Reset()                    { *m = JobResponse{} }
func (m *JobResponse) String() string            { return proto.CompactTextString(m) }
func (*JobResponse) ProtoMessage()               {}
func (*JobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *JobResponse) GetJob() *ga4gh_task_exec.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *JobResponse) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *JobResponse) GetStorage() []*StorageConfig {
	if m != nil {
		return m.Storage
	}
	return nil
}

type UpdateStatusRequest struct {
	Id       string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Step     int64                   `protobuf:"varint,2,opt,name=step" json:"step,omitempty"`
	State    ga4gh_task_exec.State   `protobuf:"varint,3,opt,name=state,enum=ga4gh_task_exec.State" json:"state,omitempty"`
	Log      *ga4gh_task_exec.JobLog `protobuf:"bytes,4,opt,name=log" json:"log,omitempty"`
	WorkerId string                  `protobuf:"bytes,5,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
}

func (m *UpdateStatusRequest) Reset()                    { *m = UpdateStatusRequest{} }
func (m *UpdateStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateStatusRequest) ProtoMessage()               {}
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateStatusRequest) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *UpdateStatusRequest) GetState() ga4gh_task_exec.State {
	if m != nil {
		return m.State
	}
	return ga4gh_task_exec.State_Unknown
}

func (m *UpdateStatusRequest) GetLog() *ga4gh_task_exec.JobLog {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *UpdateStatusRequest) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

type QueuedTaskInfoRequest struct {
	MaxTasks int32 `protobuf:"varint,1,opt,name=max_tasks,json=maxTasks" json:"max_tasks,omitempty"`
}

func (m *QueuedTaskInfoRequest) Reset()                    { *m = QueuedTaskInfoRequest{} }
func (m *QueuedTaskInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*QueuedTaskInfoRequest) ProtoMessage()               {}
func (*QueuedTaskInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueuedTaskInfoRequest) GetMaxTasks() int32 {
	if m != nil {
		return m.MaxTasks
	}
	return 0
}

type QueuedTaskInfo struct {
	Inputs    []string                   `protobuf:"bytes,1,rep,name=inputs" json:"inputs,omitempty"`
	Resources *ga4gh_task_exec.Resources `protobuf:"bytes,2,opt,name=resources" json:"resources,omitempty"`
}

func (m *QueuedTaskInfo) Reset()                    { *m = QueuedTaskInfo{} }
func (m *QueuedTaskInfo) String() string            { return proto.CompactTextString(m) }
func (*QueuedTaskInfo) ProtoMessage()               {}
func (*QueuedTaskInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *QueuedTaskInfo) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *QueuedTaskInfo) GetResources() *ga4gh_task_exec.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterType((*LocalStorage)(nil), "ga4gh_task_ref.LocalStorage")
	proto.RegisterType((*S3Storage)(nil), "ga4gh_task_ref.S3Storage")
	proto.RegisterType((*StorageConfig)(nil), "ga4gh_task_ref.StorageConfig")
	proto.RegisterType((*ServerConfig)(nil), "ga4gh_task_ref.ServerConfig")
	proto.RegisterType((*WorkerInfo)(nil), "ga4gh_task_ref.WorkerInfo")
	proto.RegisterType((*JobRequest)(nil), "ga4gh_task_ref.JobRequest")
	proto.RegisterType((*JobResponse)(nil), "ga4gh_task_ref.JobResponse")
	proto.RegisterType((*UpdateStatusRequest)(nil), "ga4gh_task_ref.UpdateStatusRequest")
	proto.RegisterType((*QueuedTaskInfoRequest)(nil), "ga4gh_task_ref.QueuedTaskInfoRequest")
	proto.RegisterType((*QueuedTaskInfo)(nil), "ga4gh_task_ref.QueuedTaskInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Scheduler service

type SchedulerClient interface {
	GetServerConfig(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*ServerConfig, error)
	GetQueueInfo(ctx context.Context, in *QueuedTaskInfoRequest, opts ...grpc.CallOption) (Scheduler_GetQueueInfoClient, error)
	GetJobToRun(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error)
	UpdateJobStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*ga4gh_task_exec.JobID, error)
	WorkerPing(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*WorkerInfo, error)
	GetJobState(ctx context.Context, in *ga4gh_task_exec.JobID, opts ...grpc.CallOption) (*ga4gh_task_exec.JobDesc, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) GetServerConfig(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*ServerConfig, error) {
	out := new(ServerConfig)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/GetServerConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) GetQueueInfo(ctx context.Context, in *QueuedTaskInfoRequest, opts ...grpc.CallOption) (Scheduler_GetQueueInfoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Scheduler_serviceDesc.Streams[0], c.cc, "/ga4gh_task_ref.Scheduler/GetQueueInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerGetQueueInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Scheduler_GetQueueInfoClient interface {
	Recv() (*QueuedTaskInfo, error)
	grpc.ClientStream
}

type schedulerGetQueueInfoClient struct {
	grpc.ClientStream
}

func (x *schedulerGetQueueInfoClient) Recv() (*QueuedTaskInfo, error) {
	m := new(QueuedTaskInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerClient) GetJobToRun(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/GetJobToRun", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) UpdateJobStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*ga4gh_task_exec.JobID, error) {
	out := new(ga4gh_task_exec.JobID)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/UpdateJobStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) WorkerPing(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*WorkerInfo, error) {
	out := new(WorkerInfo)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/WorkerPing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) GetJobState(ctx context.Context, in *ga4gh_task_exec.JobID, opts ...grpc.CallOption) (*ga4gh_task_exec.JobDesc, error) {
	out := new(ga4gh_task_exec.JobDesc)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/GetJobState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	GetServerConfig(context.Context, *WorkerInfo) (*ServerConfig, error)
	GetQueueInfo(*QueuedTaskInfoRequest, Scheduler_GetQueueInfoServer) error
	GetJobToRun(context.Context, *JobRequest) (*JobResponse, error)
	UpdateJobStatus(context.Context, *UpdateStatusRequest) (*ga4gh_task_exec.JobID, error)
	WorkerPing(context.Context, *WorkerInfo) (*WorkerInfo, error)
	GetJobState(context.Context, *ga4gh_task_exec.JobID) (*ga4gh_task_exec.JobDesc, error)
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_GetServerConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetServerConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/GetServerConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetServerConfig(ctx, req.(*WorkerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_GetQueueInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueuedTaskInfoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchedulerServer).GetQueueInfo(m, &schedulerGetQueueInfoServer{stream})
}

type Scheduler_GetQueueInfoServer interface {
	Send(*QueuedTaskInfo) error
	grpc.ServerStream
}

type schedulerGetQueueInfoServer struct {
	grpc.ServerStream
}

func (x *schedulerGetQueueInfoServer) Send(m *QueuedTaskInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _Scheduler_GetJobToRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetJobToRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/GetJobToRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetJobToRun(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_UpdateJobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).UpdateJobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/UpdateJobStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).UpdateJobStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_WorkerPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).WorkerPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/WorkerPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).WorkerPing(ctx, req.(*WorkerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_GetJobState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ga4gh_task_exec.JobID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetJobState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/GetJobState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetJobState(ctx, req.(*ga4gh_task_exec.JobID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ga4gh_task_ref.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerConfig",
			Handler:    _Scheduler_GetServerConfig_Handler,
		},
		{
			MethodName: "GetJobToRun",
			Handler:    _Scheduler_GetJobToRun_Handler,
		},
		{
			MethodName: "UpdateJobStatus",
			Handler:    _Scheduler_UpdateJobStatus_Handler,
		},
		{
			MethodName: "WorkerPing",
			Handler:    _Scheduler_WorkerPing_Handler,
		},
		{
			MethodName: "GetJobState",
			Handler:    _Scheduler_GetJobState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetQueueInfo",
			Handler:       _Scheduler_GetQueueInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "task_worker.proto",
}

func init() { proto.RegisterFile("task_worker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x5d, 0x6f, 0x13, 0x3b,
	0x10, 0x4d, 0xb2, 0x4d, 0x6f, 0x33, 0x49, 0xd3, 0x7b, 0x7d, 0x7b, 0x7b, 0xc3, 0x16, 0x50, 0x31,
	0x2a, 0x6a, 0x25, 0x14, 0x41, 0x52, 0x09, 0x5e, 0x2b, 0x2a, 0x4a, 0xab, 0x22, 0xb5, 0x4e, 0x2b,
	0xc4, 0x53, 0xe4, 0x64, 0xa7, 0x9b, 0x25, 0xe9, 0x7a, 0xb1, 0xbd, 0xb4, 0xf0, 0xc8, 0x03, 0x7f,
	0x87, 0x3f, 0xc4, 0x8f, 0x41, 0x6b, 0x7b, 0xdb, 0x7c, 0x11, 0x09, 0xde, 0xd6, 0xe3, 0xe3, 0x99,
	0x39, 0x67, 0xce, 0x68, 0xe1, 0x1f, 0xcd, 0xd5, 0xb0, 0x7b, 0x2d, 0xe4, 0x10, 0x65, 0x33, 0x91,
	0x42, 0x0b, 0x52, 0x0f, 0xf9, 0x5e, 0x38, 0xe8, 0x9a, 0x0b, 0x89, 0x97, 0xfe, 0xba, 0xf9, 0xc2,
	0x1b, 0xec, 0xa7, 0x3a, 0x12, 0xb1, 0x45, 0xd1, 0xe7, 0x50, 0x3b, 0x11, 0x7d, 0x3e, 0xea, 0x68,
	0x21, 0x79, 0x88, 0xe4, 0x11, 0xd4, 0xf8, 0x68, 0x24, 0xae, 0x31, 0xe8, 0x06, 0x91, 0x54, 0x8d,
	0xe2, 0x96, 0xb7, 0x53, 0x61, 0x55, 0x17, 0x3b, 0x88, 0xa4, 0xa2, 0x67, 0x50, 0xe9, 0xb4, 0x73,
	0xbc, 0x0f, 0x2b, 0x18, 0x07, 0x89, 0x88, 0x62, 0xdd, 0x28, 0x6e, 0x15, 0x77, 0x2a, 0xec, 0xf6,
	0x4c, 0xfe, 0x06, 0x6f, 0x88, 0x9f, 0x1b, 0x25, 0x13, 0xce, 0x3e, 0xc9, 0x06, 0x2c, 0x2b, 0xec,
	0x4b, 0xd4, 0x0d, 0xcf, 0x04, 0xdd, 0x89, 0xc6, 0xb0, 0xea, 0x12, 0xbe, 0x12, 0xf1, 0x65, 0x14,
	0x92, 0x16, 0x94, 0x47, 0x59, 0x5b, 0x26, 0x67, 0xb5, 0x75, 0xbf, 0x39, 0x49, 0xa6, 0x39, 0xde,
	0x33, 0xb3, 0x50, 0xb2, 0x0b, 0x25, 0xd5, 0x36, 0xd5, 0xaa, 0xad, 0x7b, 0xd3, 0x0f, 0x6e, 0x3b,
	0x66, 0x25, 0xd5, 0xa6, 0xdf, 0x8a, 0x50, 0xeb, 0xa0, 0xfc, 0x84, 0xd2, 0xd5, 0x7b, 0x01, 0x7f,
	0x29, 0x7b, 0x6f, 0x18, 0x57, 0x5b, 0x0f, 0x66, 0x12, 0x8c, 0xf7, 0xc7, 0x72, 0xf4, 0x18, 0xa3,
	0xd2, 0x38, 0x23, 0xb2, 0x0d, 0x75, 0x65, 0x0a, 0x74, 0x79, 0x10, 0x48, 0x54, 0xca, 0x31, 0x5e,
	0xb5, 0xd1, 0x7d, 0x1b, 0xa4, 0x17, 0x00, 0xef, 0xcc, 0xd0, 0x8e, 0xe2, 0x4b, 0x41, 0xea, 0x50,
	0x8a, 0x02, 0x27, 0x63, 0x29, 0x0a, 0xc8, 0x26, 0x54, 0x46, 0x5c, 0xe9, 0x6e, 0x12, 0xc5, 0xa1,
	0xc9, 0xef, 0xb1, 0x95, 0x2c, 0x70, 0x1a, 0xc5, 0x61, 0xa6, 0xfc, 0x40, 0x28, 0x1d, 0xf3, 0x2b,
	0x74, 0xb9, 0x6f, 0xcf, 0xf4, 0x0b, 0xc0, 0xb1, 0xe8, 0x31, 0xfc, 0x98, 0xa2, 0xd2, 0xa4, 0x05,
	0xcb, 0xd6, 0x19, 0x4e, 0x4d, 0x7f, 0x9a, 0xdb, 0x5d, 0x0b, 0xcc, 0x21, 0xc9, 0x4b, 0xa8, 0x48,
	0x54, 0x22, 0x95, 0x7d, 0x54, 0x4e, 0xd3, 0x89, 0x67, 0x99, 0x8f, 0x9a, 0x2c, 0x47, 0xb0, 0x3b,
	0x30, 0xfd, 0x5a, 0x84, 0xaa, 0x29, 0xae, 0x12, 0x11, 0x2b, 0x24, 0x4f, 0xc0, 0xfb, 0x20, 0x7a,
	0xae, 0xf4, 0xfa, 0x4c, 0x8e, 0x0c, 0x9a, 0x01, 0x08, 0x81, 0x25, 0x9e, 0xea, 0x81, 0xd3, 0xd1,
	0x7c, 0x8f, 0x8f, 0xc5, 0xfb, 0x9d, 0xb1, 0xd0, 0xef, 0x45, 0xf8, 0xf7, 0x22, 0x09, 0xb8, 0xc6,
	0x8e, 0xe6, 0x3a, 0x55, 0xb9, 0x14, 0xd3, 0x0a, 0x13, 0x58, 0x52, 0x1a, 0x13, 0x27, 0xae, 0xf9,
	0x26, 0x4f, 0xa1, 0xac, 0x34, 0xd7, 0x56, 0xd5, 0x7a, 0x6b, 0x63, 0xa6, 0xe5, 0x2c, 0x25, 0x32,
	0x0b, 0x22, 0xbb, 0xe0, 0x8d, 0x44, 0xd8, 0x58, 0x32, 0xf4, 0xfe, 0x9f, 0x47, 0xef, 0x44, 0x84,
	0x2c, 0xc3, 0x64, 0xe3, 0xb4, 0xea, 0x76, 0xa3, 0xa0, 0x51, 0xb6, 0x23, 0xb3, 0x81, 0xa3, 0x80,
	0xee, 0xc1, 0x7f, 0x67, 0x29, 0xa6, 0x18, 0x9c, 0x73, 0x35, 0x34, 0xa3, 0x70, 0x2d, 0x6f, 0x42,
	0xe5, 0x8a, 0xdf, 0x98, 0x94, 0xca, 0x74, 0x5e, 0x66, 0x2b, 0x57, 0xfc, 0x26, 0x83, 0x29, 0xda,
	0x83, 0xfa, 0xe4, 0xab, 0xcc, 0x90, 0x51, 0x9c, 0xa4, 0x3a, 0x5f, 0x5d, 0x77, 0xfa, 0xf3, 0x81,
	0xb6, 0x7e, 0x78, 0x50, 0xe9, 0xf4, 0x07, 0x18, 0xa4, 0x23, 0x94, 0xe4, 0x2d, 0xac, 0x1d, 0xa2,
	0x9e, 0x58, 0x9e, 0x05, 0x7e, 0xf2, 0x67, 0x36, 0x77, 0xfc, 0x25, 0x2d, 0x90, 0xf7, 0x50, 0x3b,
	0x44, 0x6d, 0x38, 0x98, 0xf6, 0xb7, 0xa7, 0xf1, 0x73, 0x45, 0xf1, 0x1f, 0x2e, 0x86, 0xd1, 0xc2,
	0xb3, 0x22, 0x79, 0x03, 0xd5, 0x43, 0xd4, 0xc7, 0xa2, 0x77, 0x2e, 0x58, 0x1a, 0xcf, 0x76, 0x79,
	0xb7, 0x21, 0xfe, 0xe6, 0xdc, 0x3b, 0x6b, 0x60, 0x5a, 0x20, 0xa7, 0xb0, 0x66, 0xcd, 0x74, 0x2c,
	0x7a, 0xd6, 0x4f, 0xe4, 0xf1, 0xf4, 0x8b, 0x39, 0x6e, 0xf3, 0x37, 0xe6, 0xd9, 0xe1, 0xe8, 0x80,
	0x16, 0xc8, 0xeb, 0x7c, 0xef, 0xed, 0x2a, 0x2f, 0x10, 0x70, 0xc1, 0x1d, 0x2d, 0x90, 0xfd, 0x9c,
	0xa3, 0xf1, 0x24, 0xf9, 0x45, 0x41, 0xbf, 0x31, 0x2f, 0x7e, 0x80, 0xaa, 0x4f, 0x0b, 0xbd, 0x65,
	0xf3, 0x23, 0x68, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x49, 0x50, 0xb2, 0xa7, 0x43, 0x06, 0x00,
	0x00,
}
