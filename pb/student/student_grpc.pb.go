// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: student/student.proto

package student

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StudentService_Insert_FullMethodName                        = "/student.StudentService/Insert"
	StudentService_Login_FullMethodName                         = "/student.StudentService/Login"
	StudentService_CheckStudent_FullMethodName                  = "/student.StudentService/CheckStudent"
	StudentService_StudentList_FullMethodName                   = "/student.StudentService/StudentList"
	StudentService_GetClassList_FullMethodName                  = "/student.StudentService/GetClassList"
	StudentService_InsertAcademicSession_FullMethodName         = "/student.StudentService/InsertAcademicSession"
	StudentService_GetStudentListBySession_FullMethodName       = "/student.StudentService/GetStudentListBySession"
	StudentService_InsertAddress_FullMethodName                 = "/student.StudentService/InsertAddress"
	StudentService_GetStudentProfileById_FullMethodName         = "/student.StudentService/GetStudentProfileById"
	StudentService_GetDistrictList_FullMethodName               = "/student.StudentService/GetDistrictList"
	StudentService_GetUpazillaListByDistrictId_FullMethodName   = "/student.StudentService/GetUpazillaListByDistrictId"
	StudentService_GetUnionListByUpazillaId_FullMethodName      = "/student.StudentService/GetUnionListByUpazillaId"
	StudentService_GetVillageOrRoadListByUnionId_FullMethodName = "/student.StudentService/GetVillageOrRoadListByUnionId"
	StudentService_SearchStudent_FullMethodName                 = "/student.StudentService/SearchStudent"
)

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// rpc Validate(ValidateRequest) returns (ValidateResponse) {}
	CheckStudent(ctx context.Context, in *CheckStudentRequest, opts ...grpc.CallOption) (*CheckStudentResponse, error)
	StudentList(ctx context.Context, in *StudentListRequest, opts ...grpc.CallOption) (*StudentListResponse, error)
	GetClassList(ctx context.Context, in *GetClassListRequest, opts ...grpc.CallOption) (*GetClassListResponse, error)
	InsertAcademicSession(ctx context.Context, in *InsertAcademicSessionRequest, opts ...grpc.CallOption) (*InsertAcademicSessionResponse, error)
	GetStudentListBySession(ctx context.Context, in *GetStudentListBySessionRequest, opts ...grpc.CallOption) (*GetStudentListBySessionResponse, error)
	InsertAddress(ctx context.Context, in *InsertAddressRequest, opts ...grpc.CallOption) (*InsertAddressResponse, error)
	GetStudentProfileById(ctx context.Context, in *GetStudentProfileByIdRequest, opts ...grpc.CallOption) (*GetStudentProfileByIdResponse, error)
	GetDistrictList(ctx context.Context, in *GetDistrictListRequest, opts ...grpc.CallOption) (*GetDistrictListResponse, error)
	GetUpazillaListByDistrictId(ctx context.Context, in *GetUpazillaListByDistrictIdRequest, opts ...grpc.CallOption) (*GetUpazillaListByDistrictIdResponse, error)
	GetUnionListByUpazillaId(ctx context.Context, in *GetUnionListByUpazillaIdRequest, opts ...grpc.CallOption) (*GetUnionListByUpazillaIdResponse, error)
	GetVillageOrRoadListByUnionId(ctx context.Context, in *GetVillageOrRoadListByUnionIdRequest, opts ...grpc.CallOption) (*GetVillageOrRoadListByUnionIdResponse, error)
	SearchStudent(ctx context.Context, in *SearchStudentRequest, opts ...grpc.CallOption) (*SearchStudentResponse, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, StudentService_Insert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, StudentService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) CheckStudent(ctx context.Context, in *CheckStudentRequest, opts ...grpc.CallOption) (*CheckStudentResponse, error) {
	out := new(CheckStudentResponse)
	err := c.cc.Invoke(ctx, StudentService_CheckStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) StudentList(ctx context.Context, in *StudentListRequest, opts ...grpc.CallOption) (*StudentListResponse, error) {
	out := new(StudentListResponse)
	err := c.cc.Invoke(ctx, StudentService_StudentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetClassList(ctx context.Context, in *GetClassListRequest, opts ...grpc.CallOption) (*GetClassListResponse, error) {
	out := new(GetClassListResponse)
	err := c.cc.Invoke(ctx, StudentService_GetClassList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) InsertAcademicSession(ctx context.Context, in *InsertAcademicSessionRequest, opts ...grpc.CallOption) (*InsertAcademicSessionResponse, error) {
	out := new(InsertAcademicSessionResponse)
	err := c.cc.Invoke(ctx, StudentService_InsertAcademicSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetStudentListBySession(ctx context.Context, in *GetStudentListBySessionRequest, opts ...grpc.CallOption) (*GetStudentListBySessionResponse, error) {
	out := new(GetStudentListBySessionResponse)
	err := c.cc.Invoke(ctx, StudentService_GetStudentListBySession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) InsertAddress(ctx context.Context, in *InsertAddressRequest, opts ...grpc.CallOption) (*InsertAddressResponse, error) {
	out := new(InsertAddressResponse)
	err := c.cc.Invoke(ctx, StudentService_InsertAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetStudentProfileById(ctx context.Context, in *GetStudentProfileByIdRequest, opts ...grpc.CallOption) (*GetStudentProfileByIdResponse, error) {
	out := new(GetStudentProfileByIdResponse)
	err := c.cc.Invoke(ctx, StudentService_GetStudentProfileById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetDistrictList(ctx context.Context, in *GetDistrictListRequest, opts ...grpc.CallOption) (*GetDistrictListResponse, error) {
	out := new(GetDistrictListResponse)
	err := c.cc.Invoke(ctx, StudentService_GetDistrictList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetUpazillaListByDistrictId(ctx context.Context, in *GetUpazillaListByDistrictIdRequest, opts ...grpc.CallOption) (*GetUpazillaListByDistrictIdResponse, error) {
	out := new(GetUpazillaListByDistrictIdResponse)
	err := c.cc.Invoke(ctx, StudentService_GetUpazillaListByDistrictId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetUnionListByUpazillaId(ctx context.Context, in *GetUnionListByUpazillaIdRequest, opts ...grpc.CallOption) (*GetUnionListByUpazillaIdResponse, error) {
	out := new(GetUnionListByUpazillaIdResponse)
	err := c.cc.Invoke(ctx, StudentService_GetUnionListByUpazillaId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetVillageOrRoadListByUnionId(ctx context.Context, in *GetVillageOrRoadListByUnionIdRequest, opts ...grpc.CallOption) (*GetVillageOrRoadListByUnionIdResponse, error) {
	out := new(GetVillageOrRoadListByUnionIdResponse)
	err := c.cc.Invoke(ctx, StudentService_GetVillageOrRoadListByUnionId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) SearchStudent(ctx context.Context, in *SearchStudentRequest, opts ...grpc.CallOption) (*SearchStudentResponse, error) {
	out := new(SearchStudentResponse)
	err := c.cc.Invoke(ctx, StudentService_SearchStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	Insert(context.Context, *InsertRequest) (*InsertResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// rpc Validate(ValidateRequest) returns (ValidateResponse) {}
	CheckStudent(context.Context, *CheckStudentRequest) (*CheckStudentResponse, error)
	StudentList(context.Context, *StudentListRequest) (*StudentListResponse, error)
	GetClassList(context.Context, *GetClassListRequest) (*GetClassListResponse, error)
	InsertAcademicSession(context.Context, *InsertAcademicSessionRequest) (*InsertAcademicSessionResponse, error)
	GetStudentListBySession(context.Context, *GetStudentListBySessionRequest) (*GetStudentListBySessionResponse, error)
	InsertAddress(context.Context, *InsertAddressRequest) (*InsertAddressResponse, error)
	GetStudentProfileById(context.Context, *GetStudentProfileByIdRequest) (*GetStudentProfileByIdResponse, error)
	GetDistrictList(context.Context, *GetDistrictListRequest) (*GetDistrictListResponse, error)
	GetUpazillaListByDistrictId(context.Context, *GetUpazillaListByDistrictIdRequest) (*GetUpazillaListByDistrictIdResponse, error)
	GetUnionListByUpazillaId(context.Context, *GetUnionListByUpazillaIdRequest) (*GetUnionListByUpazillaIdResponse, error)
	GetVillageOrRoadListByUnionId(context.Context, *GetVillageOrRoadListByUnionIdRequest) (*GetVillageOrRoadListByUnionIdResponse, error)
	SearchStudent(context.Context, *SearchStudentRequest) (*SearchStudentResponse, error)
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) Insert(context.Context, *InsertRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedStudentServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedStudentServiceServer) CheckStudent(context.Context, *CheckStudentRequest) (*CheckStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStudent not implemented")
}
func (UnimplementedStudentServiceServer) StudentList(context.Context, *StudentListRequest) (*StudentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StudentList not implemented")
}
func (UnimplementedStudentServiceServer) GetClassList(context.Context, *GetClassListRequest) (*GetClassListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassList not implemented")
}
func (UnimplementedStudentServiceServer) InsertAcademicSession(context.Context, *InsertAcademicSessionRequest) (*InsertAcademicSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertAcademicSession not implemented")
}
func (UnimplementedStudentServiceServer) GetStudentListBySession(context.Context, *GetStudentListBySessionRequest) (*GetStudentListBySessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentListBySession not implemented")
}
func (UnimplementedStudentServiceServer) InsertAddress(context.Context, *InsertAddressRequest) (*InsertAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertAddress not implemented")
}
func (UnimplementedStudentServiceServer) GetStudentProfileById(context.Context, *GetStudentProfileByIdRequest) (*GetStudentProfileByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentProfileById not implemented")
}
func (UnimplementedStudentServiceServer) GetDistrictList(context.Context, *GetDistrictListRequest) (*GetDistrictListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDistrictList not implemented")
}
func (UnimplementedStudentServiceServer) GetUpazillaListByDistrictId(context.Context, *GetUpazillaListByDistrictIdRequest) (*GetUpazillaListByDistrictIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpazillaListByDistrictId not implemented")
}
func (UnimplementedStudentServiceServer) GetUnionListByUpazillaId(context.Context, *GetUnionListByUpazillaIdRequest) (*GetUnionListByUpazillaIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnionListByUpazillaId not implemented")
}
func (UnimplementedStudentServiceServer) GetVillageOrRoadListByUnionId(context.Context, *GetVillageOrRoadListByUnionIdRequest) (*GetVillageOrRoadListByUnionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVillageOrRoadListByUnionId not implemented")
}
func (UnimplementedStudentServiceServer) SearchStudent(context.Context, *SearchStudentRequest) (*SearchStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchStudent not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_Insert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_CheckStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).CheckStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_CheckStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).CheckStudent(ctx, req.(*CheckStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_StudentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).StudentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_StudentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).StudentList(ctx, req.(*StudentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetClassList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetClassList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetClassList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetClassList(ctx, req.(*GetClassListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_InsertAcademicSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertAcademicSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).InsertAcademicSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_InsertAcademicSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).InsertAcademicSession(ctx, req.(*InsertAcademicSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetStudentListBySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentListBySessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudentListBySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetStudentListBySession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudentListBySession(ctx, req.(*GetStudentListBySessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_InsertAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).InsertAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_InsertAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).InsertAddress(ctx, req.(*InsertAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetStudentProfileById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentProfileByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudentProfileById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetStudentProfileById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudentProfileById(ctx, req.(*GetStudentProfileByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetDistrictList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDistrictListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetDistrictList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetDistrictList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetDistrictList(ctx, req.(*GetDistrictListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetUpazillaListByDistrictId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpazillaListByDistrictIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetUpazillaListByDistrictId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetUpazillaListByDistrictId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetUpazillaListByDistrictId(ctx, req.(*GetUpazillaListByDistrictIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetUnionListByUpazillaId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnionListByUpazillaIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetUnionListByUpazillaId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetUnionListByUpazillaId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetUnionListByUpazillaId(ctx, req.(*GetUnionListByUpazillaIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetVillageOrRoadListByUnionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVillageOrRoadListByUnionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetVillageOrRoadListByUnionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetVillageOrRoadListByUnionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetVillageOrRoadListByUnionId(ctx, req.(*GetVillageOrRoadListByUnionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_SearchStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).SearchStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_SearchStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).SearchStudent(ctx, req.(*SearchStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _StudentService_Insert_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _StudentService_Login_Handler,
		},
		{
			MethodName: "CheckStudent",
			Handler:    _StudentService_CheckStudent_Handler,
		},
		{
			MethodName: "StudentList",
			Handler:    _StudentService_StudentList_Handler,
		},
		{
			MethodName: "GetClassList",
			Handler:    _StudentService_GetClassList_Handler,
		},
		{
			MethodName: "InsertAcademicSession",
			Handler:    _StudentService_InsertAcademicSession_Handler,
		},
		{
			MethodName: "GetStudentListBySession",
			Handler:    _StudentService_GetStudentListBySession_Handler,
		},
		{
			MethodName: "InsertAddress",
			Handler:    _StudentService_InsertAddress_Handler,
		},
		{
			MethodName: "GetStudentProfileById",
			Handler:    _StudentService_GetStudentProfileById_Handler,
		},
		{
			MethodName: "GetDistrictList",
			Handler:    _StudentService_GetDistrictList_Handler,
		},
		{
			MethodName: "GetUpazillaListByDistrictId",
			Handler:    _StudentService_GetUpazillaListByDistrictId_Handler,
		},
		{
			MethodName: "GetUnionListByUpazillaId",
			Handler:    _StudentService_GetUnionListByUpazillaId_Handler,
		},
		{
			MethodName: "GetVillageOrRoadListByUnionId",
			Handler:    _StudentService_GetVillageOrRoadListByUnionId_Handler,
		},
		{
			MethodName: "SearchStudent",
			Handler:    _StudentService_SearchStudent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student/student.proto",
}
