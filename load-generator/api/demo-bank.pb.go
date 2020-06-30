// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.2
// source: demo-bank.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CasaAccount_Status int32

const (
	CasaAccount_ACTIVE  CasaAccount_Status = 0
	CasaAccount_BLOCKED CasaAccount_Status = 1
	CasaAccount_DORMANT CasaAccount_Status = 2
)

// Enum value maps for CasaAccount_Status.
var (
	CasaAccount_Status_name = map[int32]string{
		0: "ACTIVE",
		1: "BLOCKED",
		2: "DORMANT",
	}
	CasaAccount_Status_value = map[string]int32{
		"ACTIVE":  0,
		"BLOCKED": 1,
		"DORMANT": 2,
	}
)

func (x CasaAccount_Status) Enum() *CasaAccount_Status {
	p := new(CasaAccount_Status)
	*p = x
	return p
}

func (x CasaAccount_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CasaAccount_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_demo_bank_proto_enumTypes[0].Descriptor()
}

func (CasaAccount_Status) Type() protoreflect.EnumType {
	return &file_demo_bank_proto_enumTypes[0]
}

func (x CasaAccount_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CasaAccount_Status.Descriptor instead.
func (CasaAccount_Status) EnumDescriptor() ([]byte, []int) {
	return file_demo_bank_proto_rawDescGZIP(), []int{2, 0}
}

type Balance_Type int32

const (
	Balance_CURRENT   Balance_Type = 0
	Balance_AVAILABLE Balance_Type = 1
)

// Enum value maps for Balance_Type.
var (
	Balance_Type_name = map[int32]string{
		0: "CURRENT",
		1: "AVAILABLE",
	}
	Balance_Type_value = map[string]int32{
		"CURRENT":   0,
		"AVAILABLE": 1,
	}
)

func (x Balance_Type) Enum() *Balance_Type {
	p := new(Balance_Type)
	*p = x
	return p
}

func (x Balance_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Balance_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_demo_bank_proto_enumTypes[1].Descriptor()
}

func (Balance_Type) Type() protoreflect.EnumType {
	return &file_demo_bank_proto_enumTypes[1]
}

func (x Balance_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Balance_Type.Descriptor instead.
func (Balance_Type) EnumDescriptor() ([]byte, []int) {
	return file_demo_bank_proto_rawDescGZIP(), []int{3, 0}
}

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LoginName  string `protobuf:"bytes,3,opt,name=login_name,json=loginName,proto3" json:"login_name,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_demo_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_demo_bank_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

type GetCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetCustomerRequest) Reset() {
	*x = GetCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRequest) ProtoMessage() {}

func (x *GetCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return file_demo_bank_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomerRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type CasaAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId         string               `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Nickname          string               `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	ProdCode          string               `protobuf:"bytes,4,opt,name=prod_code,json=prodCode,proto3" json:"prod_code,omitempty"`
	ProdName          string               `protobuf:"bytes,5,opt,name=prod_name,json=prodName,proto3" json:"prod_name,omitempty"`
	Currency          string               `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	Status            CasaAccount_Status   `protobuf:"varint,8,opt,name=status,proto3,enum=demobank.api.CasaAccount_Status" json:"status,omitempty"`
	StatusLastUpdated *timestamp.Timestamp `protobuf:"bytes,9,opt,name=status_last_updated,json=statusLastUpdated,proto3" json:"status_last_updated,omitempty"`
	Balances          []*Balance           `protobuf:"bytes,10,rep,name=balances,proto3" json:"balances,omitempty"`
}

func (x *CasaAccount) Reset() {
	*x = CasaAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CasaAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CasaAccount) ProtoMessage() {}

func (x *CasaAccount) ProtoReflect() protoreflect.Message {
	mi := &file_demo_bank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CasaAccount.ProtoReflect.Descriptor instead.
func (*CasaAccount) Descriptor() ([]byte, []int) {
	return file_demo_bank_proto_rawDescGZIP(), []int{2}
}

func (x *CasaAccount) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CasaAccount) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *CasaAccount) GetProdCode() string {
	if x != nil {
		return x.ProdCode
	}
	return ""
}

func (x *CasaAccount) GetProdName() string {
	if x != nil {
		return x.ProdName
	}
	return ""
}

func (x *CasaAccount) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CasaAccount) GetStatus() CasaAccount_Status {
	if x != nil {
		return x.Status
	}
	return CasaAccount_ACTIVE
}

func (x *CasaAccount) GetStatusLastUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.StatusLastUpdated
	}
	return nil
}

func (x *CasaAccount) GetBalances() []*Balance {
	if x != nil {
		return x.Balances
	}
	return nil
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount      float64              `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Type        Balance_Type         `protobuf:"varint,2,opt,name=type,proto3,enum=demobank.api.Balance_Type" json:"type,omitempty"` // balance type
	CreditFlag  bool                 `protobuf:"varint,3,opt,name=credit_flag,json=creditFlag,proto3" json:"credit_flag,omitempty"`
	LastUpdated *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_bank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_demo_bank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_demo_bank_proto_rawDescGZIP(), []int{3}
}

func (x *Balance) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Balance) GetType() Balance_Type {
	if x != nil {
		return x.Type
	}
	return Balance_CURRENT
}

func (x *Balance) GetCreditFlag() bool {
	if x != nil {
		return x.CreditFlag
	}
	return false
}

func (x *Balance) GetLastUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type GetCasaAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *GetCasaAccountRequest) Reset() {
	*x = GetCasaAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_bank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCasaAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCasaAccountRequest) ProtoMessage() {}

func (x *GetCasaAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_bank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCasaAccountRequest.ProtoReflect.Descriptor instead.
func (*GetCasaAccountRequest) Descriptor() ([]byte, []int) {
	return file_demo_bank_proto_rawDescGZIP(), []int{4}
}

func (x *GetCasaAccountRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type Dashboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer            *Customer            `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	Casa                []*CasaAccount       `protobuf:"bytes,2,rep,name=casa,proto3" json:"casa,omitempty"`
	LastSuccessfulLogin *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_successful_login,json=lastSuccessfulLogin,proto3" json:"last_successful_login,omitempty"`
}

func (x *Dashboard) Reset() {
	*x = Dashboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_bank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dashboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dashboard) ProtoMessage() {}

func (x *Dashboard) ProtoReflect() protoreflect.Message {
	mi := &file_demo_bank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dashboard.ProtoReflect.Descriptor instead.
func (*Dashboard) Descriptor() ([]byte, []int) {
	return file_demo_bank_proto_rawDescGZIP(), []int{5}
}

func (x *Dashboard) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *Dashboard) GetCasa() []*CasaAccount {
	if x != nil {
		return x.Casa
	}
	return nil
}

func (x *Dashboard) GetLastSuccessfulLogin() *timestamp.Timestamp {
	if x != nil {
		return x.LastSuccessfulLogin
	}
	return nil
}

type GetDashboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginName string `protobuf:"bytes,1,opt,name=login_name,json=loginName,proto3" json:"login_name,omitempty"`
}

func (x *GetDashboardRequest) Reset() {
	*x = GetDashboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_bank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardRequest) ProtoMessage() {}

func (x *GetDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_bank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardRequest.ProtoReflect.Descriptor instead.
func (*GetDashboardRequest) Descriptor() ([]byte, []int) {
	return file_demo_bank_proto_rawDescGZIP(), []int{6}
}

func (x *GetDashboardRequest) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

var File_demo_bank_proto protoreflect.FileDescriptor

var file_demo_bank_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x65, 0x6d, 0x6f, 0x2d, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5e, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x87, 0x03, 0x0a, 0x0b, 0x43, 0x61, 0x73, 0x61,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x73, 0x61, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x4a, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x31,
	0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x22, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4c, 0x4f, 0x43, 0x4b,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x4f, 0x52, 0x4d, 0x41, 0x4e, 0x54, 0x10,
	0x02, 0x22, 0xd5, 0x01, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f,
	0x66, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x22, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56,
	0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x22, 0x36, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x73, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0xbe, 0x01, 0x0a, 0x09, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12,
	0x32, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x61, 0x73, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x63, 0x61,
	0x73, 0x61, 0x12, 0x4e, 0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x66, 0x75, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x6c,
	0x61, 0x73, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x22, 0x34, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x5a, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x32, 0x62, 0x0a, 0x12, 0x43, 0x61, 0x73, 0x61, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x61, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x64, 0x65, 0x6d, 0x6f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x73,
	0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x5e, 0x0a, 0x10, 0x44, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x21, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x1e, 0x0a, 0x09, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_bank_proto_rawDescOnce sync.Once
	file_demo_bank_proto_rawDescData = file_demo_bank_proto_rawDesc
)

func file_demo_bank_proto_rawDescGZIP() []byte {
	file_demo_bank_proto_rawDescOnce.Do(func() {
		file_demo_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_bank_proto_rawDescData)
	})
	return file_demo_bank_proto_rawDescData
}

var file_demo_bank_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_demo_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_demo_bank_proto_goTypes = []interface{}{
	(CasaAccount_Status)(0),       // 0: demobank.api.CasaAccount.Status
	(Balance_Type)(0),             // 1: demobank.api.Balance.Type
	(*Customer)(nil),              // 2: demobank.api.Customer
	(*GetCustomerRequest)(nil),    // 3: demobank.api.GetCustomerRequest
	(*CasaAccount)(nil),           // 4: demobank.api.CasaAccount
	(*Balance)(nil),               // 5: demobank.api.Balance
	(*GetCasaAccountRequest)(nil), // 6: demobank.api.GetCasaAccountRequest
	(*Dashboard)(nil),             // 7: demobank.api.Dashboard
	(*GetDashboardRequest)(nil),   // 8: demobank.api.GetDashboardRequest
	(*timestamp.Timestamp)(nil),   // 9: google.protobuf.Timestamp
}
var file_demo_bank_proto_depIdxs = []int32{
	0,  // 0: demobank.api.CasaAccount.status:type_name -> demobank.api.CasaAccount.Status
	9,  // 1: demobank.api.CasaAccount.status_last_updated:type_name -> google.protobuf.Timestamp
	5,  // 2: demobank.api.CasaAccount.balances:type_name -> demobank.api.Balance
	1,  // 3: demobank.api.Balance.type:type_name -> demobank.api.Balance.Type
	9,  // 4: demobank.api.Balance.last_updated:type_name -> google.protobuf.Timestamp
	2,  // 5: demobank.api.Dashboard.customer:type_name -> demobank.api.Customer
	4,  // 6: demobank.api.Dashboard.casa:type_name -> demobank.api.CasaAccount
	9,  // 7: demobank.api.Dashboard.last_successful_login:type_name -> google.protobuf.Timestamp
	3,  // 8: demobank.api.CustomerService.GetCustomer:input_type -> demobank.api.GetCustomerRequest
	6,  // 9: demobank.api.CasaAccountService.GetAccount:input_type -> demobank.api.GetCasaAccountRequest
	8,  // 10: demobank.api.DashboardService.GetDashboard:input_type -> demobank.api.GetDashboardRequest
	2,  // 11: demobank.api.CustomerService.GetCustomer:output_type -> demobank.api.Customer
	4,  // 12: demobank.api.CasaAccountService.GetAccount:output_type -> demobank.api.CasaAccount
	7,  // 13: demobank.api.DashboardService.GetDashboard:output_type -> demobank.api.Dashboard
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_demo_bank_proto_init() }
func file_demo_bank_proto_init() {
	if File_demo_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_bank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_demo_bank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerRequest); i {
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
		file_demo_bank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CasaAccount); i {
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
		file_demo_bank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_demo_bank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCasaAccountRequest); i {
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
		file_demo_bank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dashboard); i {
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
		file_demo_bank_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDashboardRequest); i {
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
			RawDescriptor: file_demo_bank_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_demo_bank_proto_goTypes,
		DependencyIndexes: file_demo_bank_proto_depIdxs,
		EnumInfos:         file_demo_bank_proto_enumTypes,
		MessageInfos:      file_demo_bank_proto_msgTypes,
	}.Build()
	File_demo_bank_proto = out.File
	file_demo_bank_proto_rawDesc = nil
	file_demo_bank_proto_goTypes = nil
	file_demo_bank_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerServiceClient interface {
	GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*Customer, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/demobank.api.CustomerService/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
type CustomerServiceServer interface {
	GetCustomer(context.Context, *GetCustomerRequest) (*Customer, error)
}

// UnimplementedCustomerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (*UnimplementedCustomerServiceServer) GetCustomer(context.Context, *GetCustomerRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}

func RegisterCustomerServiceServer(s *grpc.Server, srv CustomerServiceServer) {
	s.RegisterService(&_CustomerService_serviceDesc, srv)
}

func _CustomerService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demobank.api.CustomerService/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomer(ctx, req.(*GetCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demobank.api.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomer",
			Handler:    _CustomerService_GetCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo-bank.proto",
}

// CasaAccountServiceClient is the client API for CasaAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CasaAccountServiceClient interface {
	GetAccount(ctx context.Context, in *GetCasaAccountRequest, opts ...grpc.CallOption) (*CasaAccount, error)
}

type casaAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCasaAccountServiceClient(cc grpc.ClientConnInterface) CasaAccountServiceClient {
	return &casaAccountServiceClient{cc}
}

func (c *casaAccountServiceClient) GetAccount(ctx context.Context, in *GetCasaAccountRequest, opts ...grpc.CallOption) (*CasaAccount, error) {
	out := new(CasaAccount)
	err := c.cc.Invoke(ctx, "/demobank.api.CasaAccountService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CasaAccountServiceServer is the server API for CasaAccountService service.
type CasaAccountServiceServer interface {
	GetAccount(context.Context, *GetCasaAccountRequest) (*CasaAccount, error)
}

// UnimplementedCasaAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCasaAccountServiceServer struct {
}

func (*UnimplementedCasaAccountServiceServer) GetAccount(context.Context, *GetCasaAccountRequest) (*CasaAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}

func RegisterCasaAccountServiceServer(s *grpc.Server, srv CasaAccountServiceServer) {
	s.RegisterService(&_CasaAccountService_serviceDesc, srv)
}

func _CasaAccountService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCasaAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CasaAccountServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demobank.api.CasaAccountService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CasaAccountServiceServer).GetAccount(ctx, req.(*GetCasaAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CasaAccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demobank.api.CasaAccountService",
	HandlerType: (*CasaAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccount",
			Handler:    _CasaAccountService_GetAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo-bank.proto",
}

// DashboardServiceClient is the client API for DashboardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DashboardServiceClient interface {
	GetDashboard(ctx context.Context, in *GetDashboardRequest, opts ...grpc.CallOption) (*Dashboard, error)
}

type dashboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardServiceClient(cc grpc.ClientConnInterface) DashboardServiceClient {
	return &dashboardServiceClient{cc}
}

func (c *dashboardServiceClient) GetDashboard(ctx context.Context, in *GetDashboardRequest, opts ...grpc.CallOption) (*Dashboard, error) {
	out := new(Dashboard)
	err := c.cc.Invoke(ctx, "/demobank.api.DashboardService/GetDashboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardServiceServer is the server API for DashboardService service.
type DashboardServiceServer interface {
	GetDashboard(context.Context, *GetDashboardRequest) (*Dashboard, error)
}

// UnimplementedDashboardServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDashboardServiceServer struct {
}

func (*UnimplementedDashboardServiceServer) GetDashboard(context.Context, *GetDashboardRequest) (*Dashboard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDashboard not implemented")
}

func RegisterDashboardServiceServer(s *grpc.Server, srv DashboardServiceServer) {
	s.RegisterService(&_DashboardService_serviceDesc, srv)
}

func _DashboardService_GetDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).GetDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demobank.api.DashboardService/GetDashboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).GetDashboard(ctx, req.(*GetDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DashboardService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demobank.api.DashboardService",
	HandlerType: (*DashboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDashboard",
			Handler:    _DashboardService_GetDashboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo-bank.proto",
}