// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ctrl.proto

package ctrl_pb

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

type ContentType int32

const (
	ContentType_Zero                           ContentType = 0
	ContentType_SessionRequestType             ContentType = 1000
	ContentType_SessionSuccessType             ContentType = 1001
	ContentType_DialType                       ContentType = 1002
	ContentType_LinkType                       ContentType = 1003
	ContentType_FaultType                      ContentType = 1004
	ContentType_RouteType                      ContentType = 1005
	ContentType_UnrouteType                    ContentType = 1006
	ContentType_MetricsType                    ContentType = 1007
	ContentType_TogglePipeTracesRequestType    ContentType = 1008
	ContentType_TraceEventType                 ContentType = 1010
	ContentType_CreateTerminatorRequestType    ContentType = 1011
	ContentType_RemoveTerminatorRequestType    ContentType = 1012
	ContentType_InspectRequestType             ContentType = 1013
	ContentType_InspectResponseType            ContentType = 1014
	ContentType_StartXgressType                ContentType = 1015
	ContentType_SessionFailedType              ContentType = 1016
	ContentType_ValidateTerminatorsRequestType ContentType = 1017
	ContentType_UpdateTerminatorRequestType    ContentType = 10018
)

var ContentType_name = map[int32]string{
	0:     "Zero",
	1000:  "SessionRequestType",
	1001:  "SessionSuccessType",
	1002:  "DialType",
	1003:  "LinkType",
	1004:  "FaultType",
	1005:  "RouteType",
	1006:  "UnrouteType",
	1007:  "MetricsType",
	1008:  "TogglePipeTracesRequestType",
	1010:  "TraceEventType",
	1011:  "CreateTerminatorRequestType",
	1012:  "RemoveTerminatorRequestType",
	1013:  "InspectRequestType",
	1014:  "InspectResponseType",
	1015:  "StartXgressType",
	1016:  "SessionFailedType",
	1017:  "ValidateTerminatorsRequestType",
	10018: "UpdateTerminatorRequestType",
}

var ContentType_value = map[string]int32{
	"Zero":                           0,
	"SessionRequestType":             1000,
	"SessionSuccessType":             1001,
	"DialType":                       1002,
	"LinkType":                       1003,
	"FaultType":                      1004,
	"RouteType":                      1005,
	"UnrouteType":                    1006,
	"MetricsType":                    1007,
	"TogglePipeTracesRequestType":    1008,
	"TraceEventType":                 1010,
	"CreateTerminatorRequestType":    1011,
	"RemoveTerminatorRequestType":    1012,
	"InspectRequestType":             1013,
	"InspectResponseType":            1014,
	"StartXgressType":                1015,
	"SessionFailedType":              1016,
	"ValidateTerminatorsRequestType": 1017,
	"UpdateTerminatorRequestType":    10018,
}

func (x ContentType) String() string {
	return proto.EnumName(ContentType_name, int32(x))
}

func (ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{0}
}

type TerminatorPrecedence int32

const (
	TerminatorPrecedence_Default  TerminatorPrecedence = 0
	TerminatorPrecedence_Required TerminatorPrecedence = 1
	TerminatorPrecedence_Failed   TerminatorPrecedence = 2
)

var TerminatorPrecedence_name = map[int32]string{
	0: "Default",
	1: "Required",
	2: "Failed",
}

var TerminatorPrecedence_value = map[string]int32{
	"Default":  0,
	"Required": 1,
	"Failed":   2,
}

func (x TerminatorPrecedence) String() string {
	return proto.EnumName(TerminatorPrecedence_name, int32(x))
}

func (TerminatorPrecedence) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{1}
}

type FaultSubject int32

const (
	FaultSubject_IngressFault FaultSubject = 0
	FaultSubject_EgressFault  FaultSubject = 1
	FaultSubject_LinkFault    FaultSubject = 2
)

var FaultSubject_name = map[int32]string{
	0: "IngressFault",
	1: "EgressFault",
	2: "LinkFault",
}

var FaultSubject_value = map[string]int32{
	"IngressFault": 0,
	"EgressFault":  1,
	"LinkFault":    2,
}

func (x FaultSubject) String() string {
	return proto.EnumName(FaultSubject_name, int32(x))
}

func (FaultSubject) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{2}
}

type SessionRequest struct {
	IngressId            string            `protobuf:"bytes,1,opt,name=ingressId,proto3" json:"ingressId,omitempty"`
	ServiceId            string            `protobuf:"bytes,2,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	PeerData             map[uint32][]byte `protobuf:"bytes,3,rep,name=peerData,proto3" json:"peerData,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SessionRequest) Reset()         { *m = SessionRequest{} }
func (m *SessionRequest) String() string { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()    {}
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{0}
}

func (m *SessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRequest.Unmarshal(m, b)
}
func (m *SessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRequest.Marshal(b, m, deterministic)
}
func (m *SessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRequest.Merge(m, src)
}
func (m *SessionRequest) XXX_Size() int {
	return xxx_messageInfo_SessionRequest.Size(m)
}
func (m *SessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRequest proto.InternalMessageInfo

func (m *SessionRequest) GetIngressId() string {
	if m != nil {
		return m.IngressId
	}
	return ""
}

func (m *SessionRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *SessionRequest) GetPeerData() map[uint32][]byte {
	if m != nil {
		return m.PeerData
	}
	return nil
}

type CreateTerminatorRequest struct {
	ServiceId            string               `protobuf:"bytes,2,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	Binding              string               `protobuf:"bytes,3,opt,name=binding,proto3" json:"binding,omitempty"`
	Address              string               `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	PeerData             map[uint32][]byte    `protobuf:"bytes,5,rep,name=peerData,proto3" json:"peerData,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Cost                 uint32               `protobuf:"varint,6,opt,name=cost,proto3" json:"cost,omitempty"`
	Precedence           TerminatorPrecedence `protobuf:"varint,7,opt,name=precedence,proto3,enum=ctrl.pb.TerminatorPrecedence" json:"precedence,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateTerminatorRequest) Reset()         { *m = CreateTerminatorRequest{} }
func (m *CreateTerminatorRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTerminatorRequest) ProtoMessage()    {}
func (*CreateTerminatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{1}
}

func (m *CreateTerminatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTerminatorRequest.Unmarshal(m, b)
}
func (m *CreateTerminatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTerminatorRequest.Marshal(b, m, deterministic)
}
func (m *CreateTerminatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTerminatorRequest.Merge(m, src)
}
func (m *CreateTerminatorRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTerminatorRequest.Size(m)
}
func (m *CreateTerminatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTerminatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTerminatorRequest proto.InternalMessageInfo

func (m *CreateTerminatorRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *CreateTerminatorRequest) GetBinding() string {
	if m != nil {
		return m.Binding
	}
	return ""
}

func (m *CreateTerminatorRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CreateTerminatorRequest) GetPeerData() map[uint32][]byte {
	if m != nil {
		return m.PeerData
	}
	return nil
}

func (m *CreateTerminatorRequest) GetCost() uint32 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *CreateTerminatorRequest) GetPrecedence() TerminatorPrecedence {
	if m != nil {
		return m.Precedence
	}
	return TerminatorPrecedence_Default
}

type RemoveTerminatorRequest struct {
	TerminatorId         string   `protobuf:"bytes,1,opt,name=terminatorId,proto3" json:"terminatorId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTerminatorRequest) Reset()         { *m = RemoveTerminatorRequest{} }
func (m *RemoveTerminatorRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveTerminatorRequest) ProtoMessage()    {}
func (*RemoveTerminatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{2}
}

func (m *RemoveTerminatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTerminatorRequest.Unmarshal(m, b)
}
func (m *RemoveTerminatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTerminatorRequest.Marshal(b, m, deterministic)
}
func (m *RemoveTerminatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTerminatorRequest.Merge(m, src)
}
func (m *RemoveTerminatorRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveTerminatorRequest.Size(m)
}
func (m *RemoveTerminatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTerminatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTerminatorRequest proto.InternalMessageInfo

func (m *RemoveTerminatorRequest) GetTerminatorId() string {
	if m != nil {
		return m.TerminatorId
	}
	return ""
}

type Terminator struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Binding              string   `protobuf:"bytes,2,opt,name=binding,proto3" json:"binding,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Terminator) Reset()         { *m = Terminator{} }
func (m *Terminator) String() string { return proto.CompactTextString(m) }
func (*Terminator) ProtoMessage()    {}
func (*Terminator) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{3}
}

func (m *Terminator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Terminator.Unmarshal(m, b)
}
func (m *Terminator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Terminator.Marshal(b, m, deterministic)
}
func (m *Terminator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Terminator.Merge(m, src)
}
func (m *Terminator) XXX_Size() int {
	return xxx_messageInfo_Terminator.Size(m)
}
func (m *Terminator) XXX_DiscardUnknown() {
	xxx_messageInfo_Terminator.DiscardUnknown(m)
}

var xxx_messageInfo_Terminator proto.InternalMessageInfo

func (m *Terminator) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Terminator) GetBinding() string {
	if m != nil {
		return m.Binding
	}
	return ""
}

func (m *Terminator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ValidateTerminatorsRequest struct {
	Terminators          []*Terminator `protobuf:"bytes,1,rep,name=terminators,proto3" json:"terminators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ValidateTerminatorsRequest) Reset()         { *m = ValidateTerminatorsRequest{} }
func (m *ValidateTerminatorsRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateTerminatorsRequest) ProtoMessage()    {}
func (*ValidateTerminatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{4}
}

func (m *ValidateTerminatorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTerminatorsRequest.Unmarshal(m, b)
}
func (m *ValidateTerminatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTerminatorsRequest.Marshal(b, m, deterministic)
}
func (m *ValidateTerminatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTerminatorsRequest.Merge(m, src)
}
func (m *ValidateTerminatorsRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateTerminatorsRequest.Size(m)
}
func (m *ValidateTerminatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTerminatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTerminatorsRequest proto.InternalMessageInfo

func (m *ValidateTerminatorsRequest) GetTerminators() []*Terminator {
	if m != nil {
		return m.Terminators
	}
	return nil
}

type UpdateTerminatorRequest struct {
	TerminatorId         string               `protobuf:"bytes,1,opt,name=terminatorId,proto3" json:"terminatorId,omitempty"`
	UpdatePrecedence     bool                 `protobuf:"varint,2,opt,name=updatePrecedence,proto3" json:"updatePrecedence,omitempty"`
	UpdateCost           bool                 `protobuf:"varint,3,opt,name=updateCost,proto3" json:"updateCost,omitempty"`
	Precedence           TerminatorPrecedence `protobuf:"varint,4,opt,name=precedence,proto3,enum=ctrl.pb.TerminatorPrecedence" json:"precedence,omitempty"`
	Cost                 uint32               `protobuf:"varint,5,opt,name=cost,proto3" json:"cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdateTerminatorRequest) Reset()         { *m = UpdateTerminatorRequest{} }
func (m *UpdateTerminatorRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTerminatorRequest) ProtoMessage()    {}
func (*UpdateTerminatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{5}
}

func (m *UpdateTerminatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTerminatorRequest.Unmarshal(m, b)
}
func (m *UpdateTerminatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTerminatorRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTerminatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTerminatorRequest.Merge(m, src)
}
func (m *UpdateTerminatorRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTerminatorRequest.Size(m)
}
func (m *UpdateTerminatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTerminatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTerminatorRequest proto.InternalMessageInfo

func (m *UpdateTerminatorRequest) GetTerminatorId() string {
	if m != nil {
		return m.TerminatorId
	}
	return ""
}

func (m *UpdateTerminatorRequest) GetUpdatePrecedence() bool {
	if m != nil {
		return m.UpdatePrecedence
	}
	return false
}

func (m *UpdateTerminatorRequest) GetUpdateCost() bool {
	if m != nil {
		return m.UpdateCost
	}
	return false
}

func (m *UpdateTerminatorRequest) GetPrecedence() TerminatorPrecedence {
	if m != nil {
		return m.Precedence
	}
	return TerminatorPrecedence_Default
}

func (m *UpdateTerminatorRequest) GetCost() uint32 {
	if m != nil {
		return m.Cost
	}
	return 0
}

type Dial struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dial) Reset()         { *m = Dial{} }
func (m *Dial) String() string { return proto.CompactTextString(m) }
func (*Dial) ProtoMessage()    {}
func (*Dial) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{6}
}

func (m *Dial) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dial.Unmarshal(m, b)
}
func (m *Dial) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dial.Marshal(b, m, deterministic)
}
func (m *Dial) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dial.Merge(m, src)
}
func (m *Dial) XXX_Size() int {
	return xxx_messageInfo_Dial.Size(m)
}
func (m *Dial) XXX_DiscardUnknown() {
	xxx_messageInfo_Dial.DiscardUnknown(m)
}

var xxx_messageInfo_Dial proto.InternalMessageInfo

func (m *Dial) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Dial) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Link struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{7}
}

func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Fault struct {
	Subject              FaultSubject `protobuf:"varint,1,opt,name=subject,proto3,enum=ctrl.pb.FaultSubject" json:"subject,omitempty"`
	Id                   string       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Fault) Reset()         { *m = Fault{} }
func (m *Fault) String() string { return proto.CompactTextString(m) }
func (*Fault) ProtoMessage()    {}
func (*Fault) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{8}
}

func (m *Fault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fault.Unmarshal(m, b)
}
func (m *Fault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fault.Marshal(b, m, deterministic)
}
func (m *Fault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fault.Merge(m, src)
}
func (m *Fault) XXX_Size() int {
	return xxx_messageInfo_Fault.Size(m)
}
func (m *Fault) XXX_DiscardUnknown() {
	xxx_messageInfo_Fault.DiscardUnknown(m)
}

var xxx_messageInfo_Fault proto.InternalMessageInfo

func (m *Fault) GetSubject() FaultSubject {
	if m != nil {
		return m.Subject
	}
	return FaultSubject_IngressFault
}

func (m *Fault) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Route struct {
	SessionId            string           `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Egress               *Route_Egress    `protobuf:"bytes,2,opt,name=egress,proto3" json:"egress,omitempty"`
	Forwards             []*Route_Forward `protobuf:"bytes,3,rep,name=forwards,proto3" json:"forwards,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{9}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Route) GetEgress() *Route_Egress {
	if m != nil {
		return m.Egress
	}
	return nil
}

func (m *Route) GetForwards() []*Route_Forward {
	if m != nil {
		return m.Forwards
	}
	return nil
}

type Route_Egress struct {
	Binding              string            `protobuf:"bytes,1,opt,name=binding,proto3" json:"binding,omitempty"`
	Address              string            `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Destination          string            `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	PeerData             map[uint32][]byte `protobuf:"bytes,4,rep,name=peerData,proto3" json:"peerData,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Route_Egress) Reset()         { *m = Route_Egress{} }
func (m *Route_Egress) String() string { return proto.CompactTextString(m) }
func (*Route_Egress) ProtoMessage()    {}
func (*Route_Egress) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{9, 0}
}

func (m *Route_Egress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route_Egress.Unmarshal(m, b)
}
func (m *Route_Egress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route_Egress.Marshal(b, m, deterministic)
}
func (m *Route_Egress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route_Egress.Merge(m, src)
}
func (m *Route_Egress) XXX_Size() int {
	return xxx_messageInfo_Route_Egress.Size(m)
}
func (m *Route_Egress) XXX_DiscardUnknown() {
	xxx_messageInfo_Route_Egress.DiscardUnknown(m)
}

var xxx_messageInfo_Route_Egress proto.InternalMessageInfo

func (m *Route_Egress) GetBinding() string {
	if m != nil {
		return m.Binding
	}
	return ""
}

func (m *Route_Egress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Route_Egress) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *Route_Egress) GetPeerData() map[uint32][]byte {
	if m != nil {
		return m.PeerData
	}
	return nil
}

type Route_Forward struct {
	SrcAddress           string   `protobuf:"bytes,1,opt,name=srcAddress,proto3" json:"srcAddress,omitempty"`
	DstAddress           string   `protobuf:"bytes,2,opt,name=dstAddress,proto3" json:"dstAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route_Forward) Reset()         { *m = Route_Forward{} }
func (m *Route_Forward) String() string { return proto.CompactTextString(m) }
func (*Route_Forward) ProtoMessage()    {}
func (*Route_Forward) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{9, 1}
}

func (m *Route_Forward) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route_Forward.Unmarshal(m, b)
}
func (m *Route_Forward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route_Forward.Marshal(b, m, deterministic)
}
func (m *Route_Forward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route_Forward.Merge(m, src)
}
func (m *Route_Forward) XXX_Size() int {
	return xxx_messageInfo_Route_Forward.Size(m)
}
func (m *Route_Forward) XXX_DiscardUnknown() {
	xxx_messageInfo_Route_Forward.DiscardUnknown(m)
}

var xxx_messageInfo_Route_Forward proto.InternalMessageInfo

func (m *Route_Forward) GetSrcAddress() string {
	if m != nil {
		return m.SrcAddress
	}
	return ""
}

func (m *Route_Forward) GetDstAddress() string {
	if m != nil {
		return m.DstAddress
	}
	return ""
}

type Unroute struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Now                  bool     `protobuf:"varint,2,opt,name=now,proto3" json:"now,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Unroute) Reset()         { *m = Unroute{} }
func (m *Unroute) String() string { return proto.CompactTextString(m) }
func (*Unroute) ProtoMessage()    {}
func (*Unroute) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{10}
}

func (m *Unroute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unroute.Unmarshal(m, b)
}
func (m *Unroute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unroute.Marshal(b, m, deterministic)
}
func (m *Unroute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unroute.Merge(m, src)
}
func (m *Unroute) XXX_Size() int {
	return xxx_messageInfo_Unroute.Size(m)
}
func (m *Unroute) XXX_DiscardUnknown() {
	xxx_messageInfo_Unroute.DiscardUnknown(m)
}

var xxx_messageInfo_Unroute proto.InternalMessageInfo

func (m *Unroute) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Unroute) GetNow() bool {
	if m != nil {
		return m.Now
	}
	return false
}

type InspectRequest struct {
	RequestedValues      []string `protobuf:"bytes,1,rep,name=requestedValues,proto3" json:"requestedValues,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectRequest) Reset()         { *m = InspectRequest{} }
func (m *InspectRequest) String() string { return proto.CompactTextString(m) }
func (*InspectRequest) ProtoMessage()    {}
func (*InspectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{11}
}

func (m *InspectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectRequest.Unmarshal(m, b)
}
func (m *InspectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectRequest.Marshal(b, m, deterministic)
}
func (m *InspectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectRequest.Merge(m, src)
}
func (m *InspectRequest) XXX_Size() int {
	return xxx_messageInfo_InspectRequest.Size(m)
}
func (m *InspectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InspectRequest proto.InternalMessageInfo

func (m *InspectRequest) GetRequestedValues() []string {
	if m != nil {
		return m.RequestedValues
	}
	return nil
}

type InspectResponse struct {
	Success              bool                            `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Errors               []string                        `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	Values               []*InspectResponse_InspectValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *InspectResponse) Reset()         { *m = InspectResponse{} }
func (m *InspectResponse) String() string { return proto.CompactTextString(m) }
func (*InspectResponse) ProtoMessage()    {}
func (*InspectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{12}
}

func (m *InspectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectResponse.Unmarshal(m, b)
}
func (m *InspectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectResponse.Marshal(b, m, deterministic)
}
func (m *InspectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectResponse.Merge(m, src)
}
func (m *InspectResponse) XXX_Size() int {
	return xxx_messageInfo_InspectResponse.Size(m)
}
func (m *InspectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InspectResponse proto.InternalMessageInfo

func (m *InspectResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *InspectResponse) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *InspectResponse) GetValues() []*InspectResponse_InspectValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type InspectResponse_InspectValue struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectResponse_InspectValue) Reset()         { *m = InspectResponse_InspectValue{} }
func (m *InspectResponse_InspectValue) String() string { return proto.CompactTextString(m) }
func (*InspectResponse_InspectValue) ProtoMessage()    {}
func (*InspectResponse_InspectValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{12, 0}
}

func (m *InspectResponse_InspectValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectResponse_InspectValue.Unmarshal(m, b)
}
func (m *InspectResponse_InspectValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectResponse_InspectValue.Marshal(b, m, deterministic)
}
func (m *InspectResponse_InspectValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectResponse_InspectValue.Merge(m, src)
}
func (m *InspectResponse_InspectValue) XXX_Size() int {
	return xxx_messageInfo_InspectResponse_InspectValue.Size(m)
}
func (m *InspectResponse_InspectValue) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectResponse_InspectValue.DiscardUnknown(m)
}

var xxx_messageInfo_InspectResponse_InspectValue proto.InternalMessageInfo

func (m *InspectResponse_InspectValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InspectResponse_InspectValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("ctrl.pb.ContentType", ContentType_name, ContentType_value)
	proto.RegisterEnum("ctrl.pb.TerminatorPrecedence", TerminatorPrecedence_name, TerminatorPrecedence_value)
	proto.RegisterEnum("ctrl.pb.FaultSubject", FaultSubject_name, FaultSubject_value)
	proto.RegisterType((*SessionRequest)(nil), "ctrl.pb.SessionRequest")
	proto.RegisterMapType((map[uint32][]byte)(nil), "ctrl.pb.SessionRequest.PeerDataEntry")
	proto.RegisterType((*CreateTerminatorRequest)(nil), "ctrl.pb.CreateTerminatorRequest")
	proto.RegisterMapType((map[uint32][]byte)(nil), "ctrl.pb.CreateTerminatorRequest.PeerDataEntry")
	proto.RegisterType((*RemoveTerminatorRequest)(nil), "ctrl.pb.RemoveTerminatorRequest")
	proto.RegisterType((*Terminator)(nil), "ctrl.pb.Terminator")
	proto.RegisterType((*ValidateTerminatorsRequest)(nil), "ctrl.pb.ValidateTerminatorsRequest")
	proto.RegisterType((*UpdateTerminatorRequest)(nil), "ctrl.pb.UpdateTerminatorRequest")
	proto.RegisterType((*Dial)(nil), "ctrl.pb.Dial")
	proto.RegisterType((*Link)(nil), "ctrl.pb.Link")
	proto.RegisterType((*Fault)(nil), "ctrl.pb.Fault")
	proto.RegisterType((*Route)(nil), "ctrl.pb.Route")
	proto.RegisterType((*Route_Egress)(nil), "ctrl.pb.Route.Egress")
	proto.RegisterMapType((map[uint32][]byte)(nil), "ctrl.pb.Route.Egress.PeerDataEntry")
	proto.RegisterType((*Route_Forward)(nil), "ctrl.pb.Route.Forward")
	proto.RegisterType((*Unroute)(nil), "ctrl.pb.Unroute")
	proto.RegisterType((*InspectRequest)(nil), "ctrl.pb.InspectRequest")
	proto.RegisterType((*InspectResponse)(nil), "ctrl.pb.InspectResponse")
	proto.RegisterType((*InspectResponse_InspectValue)(nil), "ctrl.pb.InspectResponse.InspectValue")
}

func init() {
	proto.RegisterFile("ctrl.proto", fileDescriptor_a0572e205f89e843)
}

var fileDescriptor_a0572e205f89e843 = []byte{
	// 966 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0x5b, 0x6f, 0xe3, 0x44,
	0x14, 0xc6, 0xce, 0x65, 0x92, 0x93, 0x34, 0x0d, 0xd3, 0xa5, 0x8d, 0x02, 0xac, 0x2a, 0xaf, 0x90,
	0xaa, 0x4a, 0x04, 0x54, 0x84, 0xb4, 0x2c, 0xaa, 0xa0, 0xea, 0x45, 0x14, 0x01, 0xaa, 0x26, 0xdd,
	0x15, 0xe2, 0xcd, 0xb5, 0x67, 0x2b, 0xb3, 0xa9, 0x1d, 0xc6, 0x93, 0xac, 0xfa, 0x77, 0xf8, 0x09,
	0x3c, 0xf3, 0xce, 0x0b, 0xe2, 0x17, 0xf0, 0x03, 0xb8, 0x83, 0xb8, 0xf3, 0xc6, 0xcc, 0x99, 0x89,
	0x33, 0x0e, 0x09, 0x2a, 0x42, 0xfb, 0x36, 0xe7, 0xe6, 0x73, 0xbe, 0xf3, 0x9d, 0x33, 0x63, 0x80,
	0x48, 0x8a, 0xd1, 0x60, 0x2c, 0x32, 0x99, 0x51, 0x62, 0xce, 0x17, 0xc1, 0xe7, 0x1e, 0x74, 0x86,
	0x3c, 0xcf, 0x93, 0x2c, 0x65, 0xfc, 0xa3, 0x09, 0xcf, 0x25, 0x7d, 0x0e, 0x9a, 0x49, 0x7a, 0x29,
	0x94, 0xf2, 0x34, 0xee, 0x79, 0xdb, 0xde, 0x4e, 0x93, 0xcd, 0x15, 0xda, 0x9a, 0x73, 0x31, 0x4d,
	0x22, 0xae, 0xac, 0xbe, 0xb1, 0x16, 0x0a, 0x7a, 0x00, 0x8d, 0x31, 0xe7, 0xe2, 0x28, 0x94, 0x61,
	0xaf, 0xb2, 0x5d, 0xd9, 0x69, 0xed, 0xbd, 0x30, 0xb0, 0xa9, 0x06, 0xe5, 0x34, 0x83, 0x33, 0xeb,
	0x77, 0x9c, 0x4a, 0x71, 0xcd, 0x8a, 0xb0, 0xfe, 0xeb, 0xb0, 0x56, 0x32, 0xd1, 0x2e, 0x54, 0x1e,
	0xf1, 0x6b, 0xac, 0x64, 0x8d, 0xe9, 0x23, 0xbd, 0x05, 0xb5, 0x69, 0x38, 0x9a, 0x70, 0xcc, 0xdf,
	0x66, 0x46, 0xb8, 0xe7, 0xdf, 0xf5, 0x82, 0xcf, 0x7c, 0xd8, 0x3a, 0x14, 0x3c, 0x94, 0xfc, 0x9c,
	0x8b, 0xab, 0x24, 0x0d, 0x65, 0x26, 0x1c, 0x5c, 0xff, 0x52, 0x79, 0x0f, 0xc8, 0x45, 0x92, 0xc6,
	0x0a, 0xa8, 0x2a, 0x5c, 0xdb, 0x66, 0xa2, 0xb6, 0x84, 0x71, 0xac, 0xe1, 0xf7, 0xaa, 0xc6, 0x62,
	0x45, 0xfa, 0xb6, 0x83, 0xb6, 0x86, 0x68, 0x07, 0x05, 0xda, 0x15, 0x55, 0xac, 0x82, 0x4d, 0x29,
	0x54, 0xa3, 0x2c, 0x97, 0xbd, 0x3a, 0xc2, 0xc4, 0x33, 0xdd, 0x07, 0x18, 0x0b, 0x1e, 0xf1, 0x98,
	0xa7, 0x11, 0xef, 0x11, 0x65, 0xe9, 0xec, 0x3d, 0x5f, 0x64, 0x98, 0x7f, 0xfb, 0xac, 0x70, 0x62,
	0x4e, 0xc0, 0xff, 0xeb, 0xe4, 0x3e, 0x6c, 0x31, 0x7e, 0x95, 0x4d, 0x97, 0x34, 0x32, 0x80, 0xb6,
	0x2c, 0x94, 0xc5, 0x8c, 0x94, 0x74, 0xc1, 0x19, 0xc0, 0x3c, 0x90, 0x76, 0xc0, 0x4f, 0x66, 0x7e,
	0xea, 0xe4, 0x36, 0xdb, 0x5f, 0xd9, 0xec, 0x4a, 0xa9, 0xd9, 0xc1, 0x10, 0xfa, 0x0f, 0xc2, 0x51,
	0x12, 0x97, 0xba, 0x9a, 0xcf, 0x6a, 0x7a, 0x15, 0x5a, 0xf3, 0xfc, 0xb9, 0x4a, 0xa5, 0xd9, 0xd8,
	0x58, 0xd2, 0x2b, 0xe6, 0xfa, 0x05, 0x5f, 0x7a, 0xb0, 0x75, 0x7f, 0x1c, 0x2f, 0x9d, 0x97, 0x1b,
	0xc0, 0xa4, 0xbb, 0xd0, 0x9d, 0x60, 0xf8, 0x9c, 0x02, 0x44, 0xd4, 0x60, 0xff, 0xd0, 0xd3, 0xdb,
	0x00, 0x46, 0x77, 0xa8, 0x79, 0xae, 0xa0, 0x97, 0xa3, 0x59, 0x60, 0xbb, 0xfa, 0x1f, 0xd9, 0x2e,
	0x06, 0xa8, 0x36, 0x1f, 0xa0, 0xe0, 0x65, 0xa8, 0x1e, 0x25, 0xe1, 0x68, 0x59, 0xff, 0x67, 0x5d,
	0xf6, 0xcb, 0x5d, 0xde, 0x84, 0xea, 0x3b, 0x49, 0xfa, 0x68, 0x31, 0x22, 0x78, 0x0b, 0x6a, 0x27,
	0xe1, 0x64, 0x24, 0xe9, 0x4b, 0x40, 0xf2, 0xc9, 0xc5, 0x87, 0x3c, 0x92, 0x68, 0xed, 0xec, 0x3d,
	0x53, 0x94, 0x88, 0x0e, 0x43, 0x63, 0x64, 0x33, 0x2f, 0xfb, 0x25, 0xbf, 0xf8, 0xd2, 0x27, 0x15,
	0xa8, 0xb1, 0x6c, 0x22, 0xb9, 0x59, 0x48, 0xbc, 0x13, 0xe6, 0x17, 0x4d, 0xa1, 0xa0, 0x2f, 0x42,
	0x9d, 0x5f, 0x16, 0x25, 0xb6, 0x9c, 0x3c, 0x18, 0x3d, 0x38, 0x46, 0x23, 0xb3, 0x4e, 0x74, 0x0f,
	0x1a, 0x0f, 0x33, 0xf1, 0x38, 0x14, 0x71, 0x6e, 0x6f, 0x9e, 0xcd, 0x85, 0x80, 0x13, 0x63, 0x66,
	0x85, 0x5f, 0x5f, 0xb1, 0x5f, 0x37, 0x9f, 0x71, 0x27, 0xd2, 0x5b, 0x39, 0x91, 0xe5, 0x5e, 0xd1,
	0x6d, 0x68, 0xc5, 0x6a, 0x50, 0x34, 0x2b, 0xaa, 0x64, 0x3b, 0xaf, 0xae, 0x8a, 0xbe, 0xe1, 0x5c,
	0x10, 0x55, 0x2c, 0xea, 0xce, 0x52, 0x14, 0x4f, 0xe4, 0x32, 0xec, 0x9f, 0x02, 0xb1, 0x98, 0xf5,
	0xec, 0xe5, 0x22, 0x3a, 0xb0, 0x38, 0x0c, 0x42, 0x47, 0xa3, 0xed, 0x71, 0x2e, 0x0f, 0x4a, 0x38,
	0x1d, 0x4d, 0xf0, 0x1a, 0x90, 0xfb, 0xa9, 0xb8, 0x01, 0x6b, 0xaa, 0xbe, 0x34, 0x7b, 0x6c, 0x77,
	0x40, 0x1f, 0x83, 0x7b, 0xd0, 0x39, 0x4d, 0xf3, 0xb1, 0x9e, 0x09, 0xbb, 0x58, 0x3b, 0xb0, 0x2e,
	0xcc, 0x91, 0xc7, 0x0f, 0x74, 0xb5, 0x66, 0x5f, 0x9b, 0x6c, 0x51, 0x1d, 0x7c, 0xea, 0xc1, 0x7a,
	0x11, 0x9c, 0x8f, 0xb3, 0x34, 0xe7, 0x9a, 0x8f, 0x7c, 0x12, 0x45, 0x33, 0x1c, 0x0d, 0x36, 0x13,
	0xe9, 0xa6, 0x9a, 0x18, 0x21, 0xf4, 0xfa, 0xfb, 0xf8, 0x39, 0x2b, 0xa9, 0xc5, 0xaa, 0x4f, 0x4d,
	0x9a, 0xc5, 0x27, 0x69, 0xe1, 0xdb, 0x33, 0x19, 0xb3, 0x33, 0x1b, 0xd4, 0xbf, 0x0b, 0x6d, 0x57,
	0xaf, 0x17, 0x2d, 0x0d, 0xaf, 0xb8, 0xc5, 0x8e, 0xe7, 0x32, 0x09, 0x4d, 0x4b, 0xc2, 0xee, 0x17,
	0x15, 0x68, 0x1d, 0x66, 0xa9, 0xe4, 0xa9, 0x3c, 0xbf, 0x1e, 0x73, 0xda, 0x80, 0xea, 0x07, 0x5c,
	0x64, 0xdd, 0xa7, 0xe8, 0x16, 0xd0, 0xf2, 0x73, 0xa8, 0xed, 0xdd, 0xaf, 0x88, 0x63, 0x18, 0x1a,
	0x54, 0x68, 0xf8, 0x9a, 0xd0, 0x35, 0x68, 0xe8, 0x55, 0x46, 0xf1, 0x1b, 0x14, 0xf5, 0x9e, 0xa2,
	0xf8, 0x2d, 0x51, 0x4b, 0xd6, 0xc4, 0xed, 0x43, 0xf9, 0x3b, 0x94, 0x71, 0xbe, 0x50, 0xfe, 0x9e,
	0x28, 0x5a, 0x5a, 0x96, 0x3f, 0xd4, 0xfc, 0x80, 0x9a, 0x77, 0xb9, 0x14, 0x49, 0x64, 0x32, 0xfc,
	0x48, 0xd4, 0x38, 0x3f, 0x7b, 0x9e, 0x5d, 0x5e, 0x8e, 0xf8, 0x59, 0x32, 0xe6, 0xe7, 0x22, 0x54,
	0xd9, 0xdd, 0xe2, 0x7e, 0x22, 0x74, 0x03, 0x3a, 0xa8, 0x3f, 0x9e, 0x5a, 0x44, 0xdd, 0x9f, 0x31,
	0x6c, 0xc5, 0x5b, 0x87, 0x1e, 0xbf, 0xa0, 0xc7, 0x8a, 0xa7, 0x04, 0x3d, 0x7e, 0x45, 0xd4, 0xe5,
	0x19, 0x41, 0xc3, 0x6f, 0x44, 0x91, 0xbd, 0xb1, 0xc0, 0x11, 0x5a, 0x7e, 0x27, 0xaa, 0xe3, 0xeb,
	0x43, 0x19, 0x0a, 0xf9, 0x3e, 0x2e, 0x10, 0x6a, 0xff, 0x20, 0x6a, 0x04, 0x9e, 0xb6, 0xed, 0x3b,
	0x09, 0x93, 0x11, 0x8f, 0x51, 0xff, 0x27, 0xa1, 0x77, 0xe0, 0xf6, 0xea, 0xc7, 0x03, 0x9d, 0xfe,
	0xc2, 0x3a, 0x57, 0xbc, 0x05, 0xe8, 0xf1, 0xf1, 0x7b, 0xbb, 0xfb, 0x70, 0x6b, 0xd9, 0x3d, 0x4c,
	0x5b, 0x40, 0x8e, 0xf8, 0x43, 0x4d, 0x80, 0xe2, 0xb6, 0x0d, 0x0d, 0x1d, 0x96, 0x08, 0x1e, 0x77,
	0x3d, 0x0a, 0x50, 0x37, 0xa5, 0x74, 0xfd, 0xdd, 0x37, 0xa1, 0xed, 0xde, 0x91, 0x8a, 0x03, 0x35,
	0x59, 0x58, 0xff, 0x89, 0x8d, 0x5d, 0x87, 0xd6, 0xb1, 0xa3, 0xf0, 0x14, 0xcf, 0x4d, 0xcd, 0xb3,
	0x11, 0xfd, 0x8b, 0x3a, 0xfe, 0xbe, 0xbd, 0xf2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x89,
	0x60, 0xe3, 0xcc, 0x09, 0x00, 0x00,
}
