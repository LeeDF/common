// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type EnumCommonCodes int32

const (
	EnumCommonCodes_Fail           EnumCommonCodes = 0
	EnumCommonCodes_Success        EnumCommonCodes = 1
	EnumCommonCodes_MysqlErr       EnumCommonCodes = 2
	EnumCommonCodes_ParamErr       EnumCommonCodes = 3
	EnumCommonCodes_DaoNotFound    EnumCommonCodes = 4
	EnumCommonCodes_DaoUpdateFail  EnumCommonCodes = 5
	EnumCommonCodes_DaoRecordExist EnumCommonCodes = 6
	EnumCommonCodes_DaoSaveFail    EnumCommonCodes = 7
	EnumCommonCodes_DaoDeleteFail  EnumCommonCodes = 8
	EnumCommonCodes_DaoErr         EnumCommonCodes = 9
	//userService
	EnumCommonCodes_NicknameErr  EnumCommonCodes = 100001
	EnumCommonCodes_UserNameErr  EnumCommonCodes = 100002
	EnumCommonCodes_FeedbackErr  EnumCommonCodes = 100003
	EnumCommonCodes_CheckUserErr EnumCommonCodes = 100004
	//verifyCodeService;
	EnumCommonCodes_CheckCodeErr EnumCommonCodes = 200001
	EnumCommonCodes_SendCodeErr  EnumCommonCodes = 200002
	//pushService
	EnumCommonCodes_PushServiceErr EnumCommonCodes = 300001
)

var EnumCommonCodes_name = map[int32]string{
	0:      "Fail",
	1:      "Success",
	2:      "MysqlErr",
	3:      "ParamErr",
	4:      "DaoNotFound",
	5:      "DaoUpdateFail",
	6:      "DaoRecordExist",
	7:      "DaoSaveFail",
	8:      "DaoDeleteFail",
	9:      "DaoErr",
	100001: "NicknameErr",
	100002: "UserNameErr",
	100003: "FeedbackErr",
	100004: "CheckUserErr",
	200001: "CheckCodeErr",
	200002: "SendCodeErr",
	300001: "PushServiceErr",
}

var EnumCommonCodes_value = map[string]int32{
	"Fail":           0,
	"Success":        1,
	"MysqlErr":       2,
	"ParamErr":       3,
	"DaoNotFound":    4,
	"DaoUpdateFail":  5,
	"DaoRecordExist": 6,
	"DaoSaveFail":    7,
	"DaoDeleteFail":  8,
	"DaoErr":         9,
	"NicknameErr":    100001,
	"UserNameErr":    100002,
	"FeedbackErr":    100003,
	"CheckUserErr":   100004,
	"CheckCodeErr":   200001,
	"SendCodeErr":    200002,
	"PushServiceErr": 300001,
}

func (x EnumCommonCodes) String() string {
	return proto.EnumName(EnumCommonCodes_name, int32(x))
}

func (EnumCommonCodes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type EnumUserId int32

const (
	EnumUserId_Rid  EnumUserId = 0
	EnumUserId_Uqid EnumUserId = 1
	EnumUserId_Aid  EnumUserId = 2
)

var EnumUserId_name = map[int32]string{
	0: "Rid",
	1: "Uqid",
	2: "Aid",
}

var EnumUserId_value = map[string]int32{
	"Rid":  0,
	"Uqid": 1,
	"Aid":  2,
}

func (x EnumUserId) String() string {
	return proto.EnumName(EnumUserId_name, int32(x))
}

func (EnumUserId) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

type EnumYN int32

const (
	EnumYN_N EnumYN = 0
	EnumYN_Y EnumYN = 1
)

var EnumYN_name = map[int32]string{
	0: "N",
	1: "Y",
}

var EnumYN_value = map[string]int32{
	"N": 0,
	"Y": 1,
}

func (x EnumYN) String() string {
	return proto.EnumName(EnumYN_name, int32(x))
}

func (EnumYN) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

type EnumOnlineStatus int32

const (
	EnumOnlineStatus_Unknown EnumOnlineStatus = 0
	EnumOnlineStatus_Up      EnumOnlineStatus = 1
	EnumOnlineStatus_Down    EnumOnlineStatus = 2
)

var EnumOnlineStatus_name = map[int32]string{
	0: "Unknown",
	1: "Up",
	2: "Down",
}

var EnumOnlineStatus_value = map[string]int32{
	"Unknown": 0,
	"Up":      1,
	"Down":    2,
}

func (x EnumOnlineStatus) String() string {
	return proto.EnumName(EnumOnlineStatus_name, int32(x))
}

func (EnumOnlineStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

type EnumCommonSign int32

const (
	EnumCommonSign_Zero           EnumCommonSign = 0
	EnumCommonSign_Chid           EnumCommonSign = 1
	EnumCommonSign_Subchid        EnumCommonSign = 2
	EnumCommonSign_ApiLevel       EnumCommonSign = 3
	EnumCommonSign_Model          EnumCommonSign = 4
	EnumCommonSign_ABtest         EnumCommonSign = 5
	EnumCommonSign_ApkVCode       EnumCommonSign = 6
	EnumCommonSign_VipFlag        EnumCommonSign = 7
	EnumCommonSign_VMVCode        EnumCommonSign = 8
	EnumCommonSign_MainJarVCode   EnumCommonSign = 9
	EnumCommonSign_Brand          EnumCommonSign = 10
	EnumCommonSign_TargetSdkVCode EnumCommonSign = 11
)

var EnumCommonSign_name = map[int32]string{
	0:  "Zero",
	1:  "Chid",
	2:  "Subchid",
	3:  "ApiLevel",
	4:  "Model",
	5:  "ABtest",
	6:  "ApkVCode",
	7:  "VipFlag",
	8:  "VMVCode",
	9:  "MainJarVCode",
	10: "Brand",
	11: "TargetSdkVCode",
}

var EnumCommonSign_value = map[string]int32{
	"Zero":           0,
	"Chid":           1,
	"Subchid":        2,
	"ApiLevel":       3,
	"Model":          4,
	"ABtest":         5,
	"ApkVCode":       6,
	"VipFlag":        7,
	"VMVCode":        8,
	"MainJarVCode":   9,
	"Brand":          10,
	"TargetSdkVCode": 11,
}

func (x EnumCommonSign) String() string {
	return proto.EnumName(EnumCommonSign_name, int32(x))
}

func (EnumCommonSign) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{4}
}

type CommonResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResponse.Unmarshal(m, b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
}
func (m *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(m, src)
}
func (m *CommonResponse) XXX_Size() int {
	return xxx_messageInfo_CommonResponse.Size(m)
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CommonResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type CommonIn struct {
	Pkgname              string   `protobuf:"bytes,1,opt,name=pkgname,proto3" json:"pkgname,omitempty"`
	Chid                 int32    `protobuf:"varint,2,opt,name=chid,proto3" json:"chid,omitempty"`
	Subchid              int32    `protobuf:"varint,3,opt,name=subchid,proto3" json:"subchid,omitempty"`
	VersionCode          int32    `protobuf:"varint,4,opt,name=versionCode,proto3" json:"versionCode,omitempty"`
	VersionName          string   `protobuf:"bytes,5,opt,name=versionName,proto3" json:"versionName,omitempty"`
	Abtest               int32    `protobuf:"varint,6,opt,name=abtest,proto3" json:"abtest,omitempty"`
	Android              string   `protobuf:"bytes,7,opt,name=android,proto3" json:"android,omitempty"`
	Imsi                 string   `protobuf:"bytes,8,opt,name=imsi,proto3" json:"imsi,omitempty"`
	Imei                 string   `protobuf:"bytes,9,opt,name=imei,proto3" json:"imei,omitempty"`
	Brand                string   `protobuf:"bytes,10,opt,name=brand,proto3" json:"brand,omitempty"`
	Manufacturer         string   `protobuf:"bytes,11,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Model                string   `protobuf:"bytes,12,opt,name=model,proto3" json:"model,omitempty"`
	Product              string   `protobuf:"bytes,13,opt,name=product,proto3" json:"product,omitempty"`
	Compver              int32    `protobuf:"varint,14,opt,name=compver,proto3" json:"compver,omitempty"`
	Mainver              int32    `protobuf:"varint,15,opt,name=mainver,proto3" json:"mainver,omitempty"`
	Apiver               int32    `protobuf:"varint,16,opt,name=apiver,proto3" json:"apiver,omitempty"`
	Release              string   `protobuf:"bytes,17,opt,name=release,proto3" json:"release,omitempty"`
	Abi                  string   `protobuf:"bytes,18,opt,name=abi,proto3" json:"abi,omitempty"`
	Rid                  int64    `protobuf:"varint,19,opt,name=rid,proto3" json:"rid,omitempty"`
	Uid                  int64    `protobuf:"varint,20,opt,name=uid,proto3" json:"uid,omitempty"`
	Uqid                 string   `protobuf:"bytes,21,opt,name=uqid,proto3" json:"uqid,omitempty"`
	Cqid                 string   `protobuf:"bytes,22,opt,name=cqid,proto3" json:"cqid,omitempty"`
	Ip                   string   `protobuf:"bytes,23,opt,name=Ip,proto3" json:"Ip,omitempty"`
	Productid            int32    `protobuf:"varint,24,opt,name=productid,proto3" json:"productid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonIn) Reset()         { *m = CommonIn{} }
func (m *CommonIn) String() string { return proto.CompactTextString(m) }
func (*CommonIn) ProtoMessage()    {}
func (*CommonIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *CommonIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonIn.Unmarshal(m, b)
}
func (m *CommonIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonIn.Marshal(b, m, deterministic)
}
func (m *CommonIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonIn.Merge(m, src)
}
func (m *CommonIn) XXX_Size() int {
	return xxx_messageInfo_CommonIn.Size(m)
}
func (m *CommonIn) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonIn.DiscardUnknown(m)
}

var xxx_messageInfo_CommonIn proto.InternalMessageInfo

func (m *CommonIn) GetPkgname() string {
	if m != nil {
		return m.Pkgname
	}
	return ""
}

func (m *CommonIn) GetChid() int32 {
	if m != nil {
		return m.Chid
	}
	return 0
}

func (m *CommonIn) GetSubchid() int32 {
	if m != nil {
		return m.Subchid
	}
	return 0
}

func (m *CommonIn) GetVersionCode() int32 {
	if m != nil {
		return m.VersionCode
	}
	return 0
}

func (m *CommonIn) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *CommonIn) GetAbtest() int32 {
	if m != nil {
		return m.Abtest
	}
	return 0
}

func (m *CommonIn) GetAndroid() string {
	if m != nil {
		return m.Android
	}
	return ""
}

func (m *CommonIn) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *CommonIn) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *CommonIn) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *CommonIn) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *CommonIn) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *CommonIn) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *CommonIn) GetCompver() int32 {
	if m != nil {
		return m.Compver
	}
	return 0
}

func (m *CommonIn) GetMainver() int32 {
	if m != nil {
		return m.Mainver
	}
	return 0
}

func (m *CommonIn) GetApiver() int32 {
	if m != nil {
		return m.Apiver
	}
	return 0
}

func (m *CommonIn) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *CommonIn) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *CommonIn) GetRid() int64 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *CommonIn) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *CommonIn) GetUqid() string {
	if m != nil {
		return m.Uqid
	}
	return ""
}

func (m *CommonIn) GetCqid() string {
	if m != nil {
		return m.Cqid
	}
	return ""
}

func (m *CommonIn) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CommonIn) GetProductid() int32 {
	if m != nil {
		return m.Productid
	}
	return 0
}

type VoidMsg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoidMsg) Reset()         { *m = VoidMsg{} }
func (m *VoidMsg) String() string { return proto.CompactTextString(m) }
func (*VoidMsg) ProtoMessage()    {}
func (*VoidMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

func (m *VoidMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoidMsg.Unmarshal(m, b)
}
func (m *VoidMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoidMsg.Marshal(b, m, deterministic)
}
func (m *VoidMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoidMsg.Merge(m, src)
}
func (m *VoidMsg) XXX_Size() int {
	return xxx_messageInfo_VoidMsg.Size(m)
}
func (m *VoidMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_VoidMsg.DiscardUnknown(m)
}

var xxx_messageInfo_VoidMsg proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("common.EnumCommonCodes", EnumCommonCodes_name, EnumCommonCodes_value)
	proto.RegisterEnum("common.EnumUserId", EnumUserId_name, EnumUserId_value)
	proto.RegisterEnum("common.EnumYN", EnumYN_name, EnumYN_value)
	proto.RegisterEnum("common.EnumOnlineStatus", EnumOnlineStatus_name, EnumOnlineStatus_value)
	proto.RegisterEnum("common.EnumCommonSign", EnumCommonSign_name, EnumCommonSign_value)
	proto.RegisterType((*CommonResponse)(nil), "common.CommonResponse")
	proto.RegisterType((*CommonIn)(nil), "common.CommonIn")
	proto.RegisterType((*VoidMsg)(nil), "common.VoidMsg")
}

func init() {
	proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6)
}

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 764 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x94, 0xcd, 0x8e, 0x22, 0x37,
	0x10, 0xc7, 0xa7, 0x1b, 0x68, 0xe8, 0x82, 0x65, 0x3c, 0xde, 0xc9, 0xc4, 0x87, 0x1c, 0x46, 0x9c,
	0x10, 0x87, 0x44, 0x51, 0xa4, 0xdc, 0x67, 0x61, 0x90, 0x88, 0x02, 0x59, 0x41, 0x18, 0x69, 0x73,
	0x33, 0x6d, 0x07, 0x2c, 0xba, 0xed, 0x1e, 0xbb, 0x9b, 0x24, 0x0f, 0xb1, 0x2f, 0xb0, 0xb7, 0x7c,
	0x48, 0x39, 0xe6, 0x96, 0x7b, 0x72, 0xce, 0x4b, 0xe4, 0x4d, 0xa2, 0xb2, 0x1b, 0x0d, 0x39, 0x51,
	0xff, 0x5f, 0x7d, 0xd9, 0x55, 0x6d, 0xe0, 0x75, 0x66, 0x8a, 0xc2, 0xe8, 0xcf, 0xc2, 0xcf, 0xa7,
	0xa5, 0x35, 0x95, 0xa1, 0x49, 0x50, 0xa3, 0x2f, 0x61, 0x38, 0xf5, 0xd6, 0x5a, 0xba, 0xd2, 0x68,
	0x27, 0x29, 0x85, 0x76, 0x66, 0x84, 0x64, 0xd1, 0x7d, 0x34, 0xee, 0xac, 0xbd, 0x4d, 0x09, 0xb4,
	0x0a, 0xb7, 0x67, 0xf1, 0x7d, 0x34, 0x4e, 0xd7, 0x68, 0x8e, 0xfe, 0x69, 0x43, 0x2f, 0x24, 0x2e,
	0x34, 0x65, 0xd0, 0x2d, 0x8f, 0x7b, 0xcd, 0x8b, 0x90, 0x95, 0xae, 0xcf, 0xd2, 0x17, 0x3b, 0x28,
	0xe1, 0x33, 0xb1, 0xd8, 0x41, 0x09, 0x8c, 0x76, 0xf5, 0xce, 0xe3, 0x96, 0xc7, 0x67, 0x49, 0xef,
	0xa1, 0x7f, 0x92, 0xd6, 0x29, 0xa3, 0xa7, 0x78, 0x82, 0xb6, 0xf7, 0x5e, 0xa2, 0x8b, 0x88, 0x15,
	0x76, 0xeb, 0xf8, 0x6e, 0x97, 0x88, 0xde, 0x41, 0xc2, 0x77, 0x95, 0x74, 0x15, 0x4b, 0x7c, 0x7a,
	0xa3, 0xb0, 0x2b, 0xd7, 0xc2, 0x1a, 0x25, 0x58, 0x37, 0x9c, 0xb1, 0x91, 0x78, 0x46, 0x55, 0x38,
	0xc5, 0x7a, 0x1e, 0x7b, 0x3b, 0x30, 0xa9, 0x58, 0x7a, 0x66, 0x52, 0xd1, 0x5b, 0xe8, 0xec, 0x2c,
	0xd7, 0x82, 0x81, 0x87, 0x41, 0xd0, 0x11, 0x0c, 0x0a, 0xae, 0xeb, 0xef, 0x79, 0x56, 0xd5, 0x56,
	0x5a, 0xd6, 0xf7, 0xce, 0xff, 0x31, 0xcc, 0x2c, 0x8c, 0x90, 0x39, 0x1b, 0x84, 0x4c, 0x2f, 0xfc,
	0xd4, 0xac, 0x11, 0x75, 0x56, 0xb1, 0x57, 0xcd, 0xd4, 0x82, 0x44, 0x4f, 0x66, 0x8a, 0xf2, 0x24,
	0x2d, 0x1b, 0x86, 0x09, 0x35, 0x12, 0x3d, 0x05, 0x57, 0x1a, 0x3d, 0xd7, 0xc1, 0xd3, 0x48, 0x7f,
	0xef, 0x52, 0xa1, 0x83, 0x34, 0xf7, 0xf6, 0x0a, 0x33, 0xac, 0xcc, 0x25, 0x77, 0x92, 0xdd, 0x84,
	0x2e, 0x8d, 0xc4, 0xa5, 0xf2, 0x9d, 0x62, 0x34, 0x2c, 0x95, 0xef, 0x14, 0x12, 0xab, 0x04, 0x7b,
	0x7d, 0x1f, 0x8d, 0x5b, 0x6b, 0x34, 0x91, 0xd4, 0x4a, 0xb0, 0xdb, 0x40, 0xea, 0x30, 0xad, 0xfa,
	0x59, 0x09, 0xf6, 0x51, 0x98, 0x0c, 0xda, 0x7e, 0xcb, 0xc8, 0xee, 0x02, 0x43, 0x9b, 0x0e, 0x21,
	0x5e, 0x94, 0xec, 0x63, 0x4f, 0xe2, 0x45, 0x49, 0x3f, 0x81, 0xb4, 0xb9, 0x9e, 0x12, 0x8c, 0xf9,
	0x23, 0xbe, 0x80, 0x51, 0x0a, 0xdd, 0x27, 0xa3, 0xc4, 0xd2, 0xed, 0x27, 0x7f, 0xc6, 0x70, 0xfd,
	0xa8, 0xeb, 0x22, 0x7c, 0x5d, 0xb8, 0x75, 0x47, 0x7b, 0xd0, 0x9e, 0x73, 0x95, 0x93, 0x2b, 0xda,
	0x87, 0xee, 0xa6, 0xce, 0x32, 0xe9, 0x1c, 0x89, 0xe8, 0x00, 0x7a, 0xcb, 0x9f, 0xdc, 0x73, 0xfe,
	0x68, 0x2d, 0x89, 0x51, 0xbd, 0xe5, 0x96, 0x17, 0xa8, 0x5a, 0xf4, 0x1a, 0xfa, 0x33, 0x6e, 0x56,
	0xa6, 0x9a, 0x9b, 0x5a, 0x0b, 0xd2, 0xa6, 0x37, 0xf0, 0x6a, 0xc6, 0xcd, 0xb6, 0x14, 0xbc, 0x92,
	0xbe, 0x58, 0x87, 0x52, 0x18, 0xce, 0xb8, 0x59, 0xcb, 0xcc, 0x58, 0xf1, 0xf8, 0xa3, 0x72, 0x15,
	0x49, 0x9a, 0xbc, 0x0d, 0x3f, 0x85, 0xa0, 0x6e, 0x93, 0x37, 0x93, 0xb9, 0x6c, 0xf2, 0x7a, 0x14,
	0x20, 0x99, 0x71, 0x83, 0x7d, 0x52, 0x7a, 0x03, 0xfd, 0x95, 0xca, 0x8e, 0xf8, 0xb5, 0x23, 0xf8,
	0xf9, 0x7d, 0x82, 0x68, 0xeb, 0xa4, 0x5d, 0x35, 0xe8, 0x97, 0x80, 0xe6, 0x52, 0x8a, 0x1d, 0xcf,
	0x8e, 0x88, 0x7e, 0x7d, 0x9f, 0x50, 0x0a, 0x83, 0xe9, 0x41, 0x66, 0x47, 0x0c, 0x45, 0xf6, 0xdb,
	0x05, 0xc3, 0x5b, 0x23, 0xfb, 0xeb, 0xc3, 0x00, 0x53, 0x37, 0x52, 0x8b, 0x33, 0xfa, 0xfb, 0xc3,
	0x80, 0xde, 0xc2, 0xf0, 0x6d, 0xed, 0x0e, 0x1b, 0x69, 0x4f, 0x2a, 0xf3, 0xf4, 0xdf, 0xdf, 0xe9,
	0x64, 0x0c, 0x80, 0x73, 0xc3, 0x7a, 0x0b, 0x41, 0xbb, 0xd0, 0x5a, 0x2b, 0x41, 0xae, 0x70, 0x76,
	0xdb, 0x67, 0x25, 0x48, 0x84, 0xe8, 0x41, 0x09, 0x12, 0x4f, 0xee, 0x20, 0xc1, 0xc8, 0x77, 0x2b,
	0xda, 0x81, 0x68, 0x45, 0xae, 0xf0, 0xe7, 0x1d, 0x89, 0x26, 0x9f, 0x03, 0x41, 0xfe, 0x8d, 0xce,
	0x95, 0x96, 0x9b, 0x8a, 0x57, 0xb5, 0xc3, 0x81, 0x6f, 0xf5, 0x51, 0x9b, 0x1f, 0x34, 0xb9, 0xa2,
	0x09, 0xc4, 0xdb, 0x92, 0x44, 0x58, 0x73, 0x86, 0x24, 0x9e, 0xfc, 0x11, 0xc1, 0xf0, 0x65, 0x5b,
	0x1b, 0xb5, 0xd7, 0xe8, 0xfc, 0x4e, 0x5a, 0x13, 0x5a, 0x4f, 0x0f, 0xbe, 0xb5, 0x5f, 0x9b, 0x7f,
	0xe4, 0x61, 0x51, 0x0f, 0xa5, 0xfa, 0x5a, 0x9e, 0x64, 0x4e, 0x5a, 0x34, 0x85, 0xce, 0x12, 0xdf,
	0x03, 0x69, 0xe3, 0x5c, 0x1f, 0xde, 0xe0, 0x6b, 0x25, 0x9d, 0x10, 0x74, 0x7c, 0xc2, 0x6b, 0x93,
	0x04, 0xf3, 0x9f, 0x54, 0x39, 0xcf, 0xf9, 0x9e, 0x74, 0xbd, 0x58, 0x06, 0x4f, 0x8f, 0x12, 0x18,
	0x2c, 0xb9, 0xd2, 0x5f, 0x71, 0x1b, 0x48, 0x8a, 0x05, 0xdf, 0xe0, 0xd3, 0x24, 0x80, 0x0b, 0xfe,
	0x96, 0xdb, 0xbd, 0xac, 0x36, 0xa2, 0x29, 0xd5, 0xdf, 0x25, 0xfe, 0x0f, 0xf0, 0x8b, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xaf, 0x88, 0x79, 0x49, 0x17, 0x05, 0x00, 0x00,
}
