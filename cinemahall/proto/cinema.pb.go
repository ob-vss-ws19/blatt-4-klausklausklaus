// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cinemahall/proto/cinema.proto

package cinemahallproto

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

type SizeRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SizeRequest) Reset()         { *m = SizeRequest{} }
func (m *SizeRequest) String() string { return proto.CompactTextString(m) }
func (*SizeRequest) ProtoMessage()    {}
func (*SizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{0}
}

func (m *SizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SizeRequest.Unmarshal(m, b)
}
func (m *SizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SizeRequest.Marshal(b, m, deterministic)
}
func (m *SizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SizeRequest.Merge(m, src)
}
func (m *SizeRequest) XXX_Size() int {
	return xxx_messageInfo_SizeRequest.Size(m)
}
func (m *SizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SizeRequest proto.InternalMessageInfo

func (m *SizeRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SizeResponse struct {
	Row                  int32    `protobuf:"varint,1,opt,name=row,proto3" json:"row,omitempty"`
	Column               int32    `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SizeResponse) Reset()         { *m = SizeResponse{} }
func (m *SizeResponse) String() string { return proto.CompactTextString(m) }
func (*SizeResponse) ProtoMessage()    {}
func (*SizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{1}
}

func (m *SizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SizeResponse.Unmarshal(m, b)
}
func (m *SizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SizeResponse.Marshal(b, m, deterministic)
}
func (m *SizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SizeResponse.Merge(m, src)
}
func (m *SizeResponse) XXX_Size() int {
	return xxx_messageInfo_SizeResponse.Size(m)
}
func (m *SizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SizeResponse proto.InternalMessageInfo

func (m *SizeResponse) GetRow() int32 {
	if m != nil {
		return m.Row
	}
	return 0
}

func (m *SizeResponse) GetColumn() int32 {
	if m != nil {
		return m.Column
	}
	return 0
}

type SeatMessage struct {
	Row                  int32    `protobuf:"varint,1,opt,name=row,proto3" json:"row,omitempty"`
	Column               int32    `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SeatMessage) Reset()         { *m = SeatMessage{} }
func (m *SeatMessage) String() string { return proto.CompactTextString(m) }
func (*SeatMessage) ProtoMessage()    {}
func (*SeatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{2}
}

func (m *SeatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SeatMessage.Unmarshal(m, b)
}
func (m *SeatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SeatMessage.Marshal(b, m, deterministic)
}
func (m *SeatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeatMessage.Merge(m, src)
}
func (m *SeatMessage) XXX_Size() int {
	return xxx_messageInfo_SeatMessage.Size(m)
}
func (m *SeatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SeatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SeatMessage proto.InternalMessageInfo

func (m *SeatMessage) GetRow() int32 {
	if m != nil {
		return m.Row
	}
	return 0
}

func (m *SeatMessage) GetColumn() int32 {
	if m != nil {
		return m.Column
	}
	return 0
}

type CreateCinemaRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Row                  int32    `protobuf:"varint,2,opt,name=row,proto3" json:"row,omitempty"`
	Column               int32    `protobuf:"varint,3,opt,name=column,proto3" json:"column,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCinemaRequest) Reset()         { *m = CreateCinemaRequest{} }
func (m *CreateCinemaRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCinemaRequest) ProtoMessage()    {}
func (*CreateCinemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{3}
}

func (m *CreateCinemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCinemaRequest.Unmarshal(m, b)
}
func (m *CreateCinemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCinemaRequest.Marshal(b, m, deterministic)
}
func (m *CreateCinemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCinemaRequest.Merge(m, src)
}
func (m *CreateCinemaRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCinemaRequest.Size(m)
}
func (m *CreateCinemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCinemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCinemaRequest proto.InternalMessageInfo

func (m *CreateCinemaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCinemaRequest) GetRow() int32 {
	if m != nil {
		return m.Row
	}
	return 0
}

func (m *CreateCinemaRequest) GetColumn() int32 {
	if m != nil {
		return m.Column
	}
	return 0
}

type CreateCinemaResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCinemaResponse) Reset()         { *m = CreateCinemaResponse{} }
func (m *CreateCinemaResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCinemaResponse) ProtoMessage()    {}
func (*CreateCinemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{4}
}

func (m *CreateCinemaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCinemaResponse.Unmarshal(m, b)
}
func (m *CreateCinemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCinemaResponse.Marshal(b, m, deterministic)
}
func (m *CreateCinemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCinemaResponse.Merge(m, src)
}
func (m *CreateCinemaResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCinemaResponse.Size(m)
}
func (m *CreateCinemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCinemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCinemaResponse proto.InternalMessageInfo

func (m *CreateCinemaResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCinemaResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteCinemaRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCinemaRequest) Reset()         { *m = DeleteCinemaRequest{} }
func (m *DeleteCinemaRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCinemaRequest) ProtoMessage()    {}
func (*DeleteCinemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{5}
}

func (m *DeleteCinemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCinemaRequest.Unmarshal(m, b)
}
func (m *DeleteCinemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCinemaRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCinemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCinemaRequest.Merge(m, src)
}
func (m *DeleteCinemaRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCinemaRequest.Size(m)
}
func (m *DeleteCinemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCinemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCinemaRequest proto.InternalMessageInfo

func (m *DeleteCinemaRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteCinemaResponse struct {
	Answer               bool     `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCinemaResponse) Reset()         { *m = DeleteCinemaResponse{} }
func (m *DeleteCinemaResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCinemaResponse) ProtoMessage()    {}
func (*DeleteCinemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{6}
}

func (m *DeleteCinemaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCinemaResponse.Unmarshal(m, b)
}
func (m *DeleteCinemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCinemaResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCinemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCinemaResponse.Merge(m, src)
}
func (m *DeleteCinemaResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCinemaResponse.Size(m)
}
func (m *DeleteCinemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCinemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCinemaResponse proto.InternalMessageInfo

func (m *DeleteCinemaResponse) GetAnswer() bool {
	if m != nil {
		return m.Answer
	}
	return false
}

type ReservationRequest struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Seatreservation      []*SeatMessage `protobuf:"bytes,2,rep,name=seatreservation,proto3" json:"seatreservation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReservationRequest) Reset()         { *m = ReservationRequest{} }
func (m *ReservationRequest) String() string { return proto.CompactTextString(m) }
func (*ReservationRequest) ProtoMessage()    {}
func (*ReservationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{7}
}

func (m *ReservationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReservationRequest.Unmarshal(m, b)
}
func (m *ReservationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReservationRequest.Marshal(b, m, deterministic)
}
func (m *ReservationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReservationRequest.Merge(m, src)
}
func (m *ReservationRequest) XXX_Size() int {
	return xxx_messageInfo_ReservationRequest.Size(m)
}
func (m *ReservationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReservationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReservationRequest proto.InternalMessageInfo

func (m *ReservationRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReservationRequest) GetSeatreservation() []*SeatMessage {
	if m != nil {
		return m.Seatreservation
	}
	return nil
}

type ReservationResponse struct {
	Answer               bool     `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReservationResponse) Reset()         { *m = ReservationResponse{} }
func (m *ReservationResponse) String() string { return proto.CompactTextString(m) }
func (*ReservationResponse) ProtoMessage()    {}
func (*ReservationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{8}
}

func (m *ReservationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReservationResponse.Unmarshal(m, b)
}
func (m *ReservationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReservationResponse.Marshal(b, m, deterministic)
}
func (m *ReservationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReservationResponse.Merge(m, src)
}
func (m *ReservationResponse) XXX_Size() int {
	return xxx_messageInfo_ReservationResponse.Size(m)
}
func (m *ReservationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReservationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReservationResponse proto.InternalMessageInfo

func (m *ReservationResponse) GetAnswer() bool {
	if m != nil {
		return m.Answer
	}
	return false
}

type StornoRequest struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Seatstorno           []*SeatMessage `protobuf:"bytes,2,rep,name=seatstorno,proto3" json:"seatstorno,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StornoRequest) Reset()         { *m = StornoRequest{} }
func (m *StornoRequest) String() string { return proto.CompactTextString(m) }
func (*StornoRequest) ProtoMessage()    {}
func (*StornoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{9}
}

func (m *StornoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StornoRequest.Unmarshal(m, b)
}
func (m *StornoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StornoRequest.Marshal(b, m, deterministic)
}
func (m *StornoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StornoRequest.Merge(m, src)
}
func (m *StornoRequest) XXX_Size() int {
	return xxx_messageInfo_StornoRequest.Size(m)
}
func (m *StornoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StornoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StornoRequest proto.InternalMessageInfo

func (m *StornoRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StornoRequest) GetSeatstorno() []*SeatMessage {
	if m != nil {
		return m.Seatstorno
	}
	return nil
}

type StornoResponse struct {
	Answer               bool     `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StornoResponse) Reset()         { *m = StornoResponse{} }
func (m *StornoResponse) String() string { return proto.CompactTextString(m) }
func (*StornoResponse) ProtoMessage()    {}
func (*StornoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{10}
}

func (m *StornoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StornoResponse.Unmarshal(m, b)
}
func (m *StornoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StornoResponse.Marshal(b, m, deterministic)
}
func (m *StornoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StornoResponse.Merge(m, src)
}
func (m *StornoResponse) XXX_Size() int {
	return xxx_messageInfo_StornoResponse.Size(m)
}
func (m *StornoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StornoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StornoResponse proto.InternalMessageInfo

func (m *StornoResponse) GetAnswer() bool {
	if m != nil {
		return m.Answer
	}
	return false
}

type CheckSeatsRequest struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Seatcheck            []*SeatMessage `protobuf:"bytes,2,rep,name=seatcheck,proto3" json:"seatcheck,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CheckSeatsRequest) Reset()         { *m = CheckSeatsRequest{} }
func (m *CheckSeatsRequest) String() string { return proto.CompactTextString(m) }
func (*CheckSeatsRequest) ProtoMessage()    {}
func (*CheckSeatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{11}
}

func (m *CheckSeatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSeatsRequest.Unmarshal(m, b)
}
func (m *CheckSeatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSeatsRequest.Marshal(b, m, deterministic)
}
func (m *CheckSeatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSeatsRequest.Merge(m, src)
}
func (m *CheckSeatsRequest) XXX_Size() int {
	return xxx_messageInfo_CheckSeatsRequest.Size(m)
}
func (m *CheckSeatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSeatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSeatsRequest proto.InternalMessageInfo

func (m *CheckSeatsRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CheckSeatsRequest) GetSeatcheck() []*SeatMessage {
	if m != nil {
		return m.Seatcheck
	}
	return nil
}

type CheckSeatsResponse struct {
	Answer               bool     `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckSeatsResponse) Reset()         { *m = CheckSeatsResponse{} }
func (m *CheckSeatsResponse) String() string { return proto.CompactTextString(m) }
func (*CheckSeatsResponse) ProtoMessage()    {}
func (*CheckSeatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{12}
}

func (m *CheckSeatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSeatsResponse.Unmarshal(m, b)
}
func (m *CheckSeatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSeatsResponse.Marshal(b, m, deterministic)
}
func (m *CheckSeatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSeatsResponse.Merge(m, src)
}
func (m *CheckSeatsResponse) XXX_Size() int {
	return xxx_messageInfo_CheckSeatsResponse.Size(m)
}
func (m *CheckSeatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSeatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSeatsResponse proto.InternalMessageInfo

func (m *CheckSeatsResponse) GetAnswer() bool {
	if m != nil {
		return m.Answer
	}
	return false
}

type FreeSeatsRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FreeSeatsRequest) Reset()         { *m = FreeSeatsRequest{} }
func (m *FreeSeatsRequest) String() string { return proto.CompactTextString(m) }
func (*FreeSeatsRequest) ProtoMessage()    {}
func (*FreeSeatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{13}
}

func (m *FreeSeatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FreeSeatsRequest.Unmarshal(m, b)
}
func (m *FreeSeatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FreeSeatsRequest.Marshal(b, m, deterministic)
}
func (m *FreeSeatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FreeSeatsRequest.Merge(m, src)
}
func (m *FreeSeatsRequest) XXX_Size() int {
	return xxx_messageInfo_FreeSeatsRequest.Size(m)
}
func (m *FreeSeatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FreeSeatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FreeSeatsRequest proto.InternalMessageInfo

func (m *FreeSeatsRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FreeSeatsResponse struct {
	Freeseats            []*SeatMessage `protobuf:"bytes,1,rep,name=freeseats,proto3" json:"freeseats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FreeSeatsResponse) Reset()         { *m = FreeSeatsResponse{} }
func (m *FreeSeatsResponse) String() string { return proto.CompactTextString(m) }
func (*FreeSeatsResponse) ProtoMessage()    {}
func (*FreeSeatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6a30189af5a9b5, []int{14}
}

func (m *FreeSeatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FreeSeatsResponse.Unmarshal(m, b)
}
func (m *FreeSeatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FreeSeatsResponse.Marshal(b, m, deterministic)
}
func (m *FreeSeatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FreeSeatsResponse.Merge(m, src)
}
func (m *FreeSeatsResponse) XXX_Size() int {
	return xxx_messageInfo_FreeSeatsResponse.Size(m)
}
func (m *FreeSeatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FreeSeatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FreeSeatsResponse proto.InternalMessageInfo

func (m *FreeSeatsResponse) GetFreeseats() []*SeatMessage {
	if m != nil {
		return m.Freeseats
	}
	return nil
}

func init() {
	proto.RegisterType((*SizeRequest)(nil), "cinemahallproto.SizeRequest")
	proto.RegisterType((*SizeResponse)(nil), "cinemahallproto.SizeResponse")
	proto.RegisterType((*SeatMessage)(nil), "cinemahallproto.SeatMessage")
	proto.RegisterType((*CreateCinemaRequest)(nil), "cinemahallproto.CreateCinemaRequest")
	proto.RegisterType((*CreateCinemaResponse)(nil), "cinemahallproto.CreateCinemaResponse")
	proto.RegisterType((*DeleteCinemaRequest)(nil), "cinemahallproto.DeleteCinemaRequest")
	proto.RegisterType((*DeleteCinemaResponse)(nil), "cinemahallproto.DeleteCinemaResponse")
	proto.RegisterType((*ReservationRequest)(nil), "cinemahallproto.ReservationRequest")
	proto.RegisterType((*ReservationResponse)(nil), "cinemahallproto.ReservationResponse")
	proto.RegisterType((*StornoRequest)(nil), "cinemahallproto.StornoRequest")
	proto.RegisterType((*StornoResponse)(nil), "cinemahallproto.StornoResponse")
	proto.RegisterType((*CheckSeatsRequest)(nil), "cinemahallproto.CheckSeatsRequest")
	proto.RegisterType((*CheckSeatsResponse)(nil), "cinemahallproto.CheckSeatsResponse")
	proto.RegisterType((*FreeSeatsRequest)(nil), "cinemahallproto.FreeSeatsRequest")
	proto.RegisterType((*FreeSeatsResponse)(nil), "cinemahallproto.FreeSeatsResponse")
}

func init() { proto.RegisterFile("cinemahall/proto/cinema.proto", fileDescriptor_8e6a30189af5a9b5) }

var fileDescriptor_8e6a30189af5a9b5 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x61, 0x6b, 0xd3, 0x40,
	0x1c, 0xc6, 0xd7, 0x74, 0x06, 0xfb, 0x54, 0xd7, 0xed, 0x3a, 0x46, 0x09, 0x56, 0xe7, 0x6d, 0x83,
	0xbe, 0xd0, 0x0e, 0xe6, 0x0b, 0xa5, 0xf8, 0xae, 0x32, 0x5f, 0x88, 0x4c, 0x52, 0x41, 0x10, 0x44,
	0xce, 0xee, 0x3f, 0x17, 0x4c, 0x93, 0x79, 0x97, 0x39, 0xf0, 0xf3, 0xfa, 0x41, 0x24, 0x97, 0x34,
	0x49, 0x93, 0x4b, 0xd3, 0xbd, 0xcb, 0xdd, 0x3d, 0xff, 0xdf, 0x3d, 0x7d, 0xfa, 0x24, 0x18, 0xce,
	0xbd, 0x80, 0x16, 0xe2, 0x5a, 0xf8, 0xfe, 0xe9, 0x8d, 0x0c, 0xa3, 0xf0, 0x34, 0xd9, 0x18, 0xeb,
	0x05, 0xeb, 0xe5, 0xc7, 0x7a, 0x83, 0x0f, 0xd1, 0x9d, 0x79, 0x7f, 0xc9, 0xa5, 0xdf, 0xb7, 0xa4,
	0x22, 0xb6, 0x03, 0xcb, 0xbb, 0x1c, 0xb4, 0x0e, 0x5b, 0xa3, 0x07, 0xae, 0xe5, 0x5d, 0xf2, 0x37,
	0x78, 0x94, 0x1c, 0xab, 0x9b, 0x30, 0x50, 0xc4, 0x76, 0xd1, 0x96, 0xe1, 0x5d, 0x2a, 0x88, 0x1f,
	0xd9, 0x01, 0xec, 0x79, 0xe8, 0xdf, 0x2e, 0x82, 0x81, 0xa5, 0x37, 0xd3, 0x15, 0x7f, 0x8d, 0xee,
	0x8c, 0x44, 0xf4, 0x91, 0x94, 0x12, 0x3f, 0xef, 0x33, 0x38, 0x43, 0x7f, 0x2a, 0x49, 0x44, 0x34,
	0xd5, 0x56, 0x97, 0xce, 0x18, 0xb6, 0x03, 0xb1, 0x20, 0x4d, 0xe8, 0xb8, 0xfa, 0x79, 0x09, 0xb5,
	0x4c, 0xd0, 0xf6, 0x0a, 0x74, 0x82, 0xfd, 0x55, 0x68, 0xfa, 0x7b, 0x4c, 0xd4, 0x24, 0x03, 0x2b,
	0xcb, 0xe0, 0x04, 0xfd, 0x77, 0xe4, 0x53, 0xd9, 0x50, 0x39, 0xaa, 0x31, 0xf6, 0x57, 0x65, 0xe9,
	0x15, 0x07, 0xb0, 0x45, 0xa0, 0xee, 0x48, 0x6a, 0xed, 0x43, 0x37, 0x5d, 0x71, 0x1f, 0xcc, 0x25,
	0x45, 0xf2, 0x8f, 0x88, 0xbc, 0x30, 0xa8, 0xa1, 0xb2, 0x73, 0xf4, 0x14, 0x89, 0x48, 0xe6, 0xca,
	0x81, 0x75, 0xd8, 0x1e, 0x75, 0xcf, 0x9e, 0x8c, 0x4b, 0x7f, 0xe5, 0xb8, 0x10, 0xb7, 0x5b, 0x1e,
	0xe2, 0x2f, 0xd1, 0x5f, 0xb9, 0xad, 0xc1, 0xdc, 0x37, 0x3c, 0x9e, 0x45, 0xa1, 0x0c, 0xc2, 0x3a,
	0x5f, 0x6f, 0x81, 0xf8, 0x0a, 0xa5, 0x45, 0x1b, 0x59, 0x2a, 0xe8, 0xf9, 0x08, 0x3b, 0x4b, 0x7c,
	0x83, 0x91, 0xef, 0xd8, 0x9b, 0x5e, 0xd3, 0xfc, 0x57, 0x4c, 0x52, 0x75, 0x66, 0x26, 0xe8, 0xc4,
	0xf0, 0x79, 0x2c, 0xdc, 0xc8, 0x4b, 0x2e, 0xe7, 0x2f, 0xc0, 0x8a, 0x17, 0x34, 0xd8, 0xe1, 0xd8,
	0x3d, 0x97, 0x44, 0xeb, 0xdc, 0xf0, 0x0b, 0xec, 0x15, 0x34, 0x29, 0x70, 0x82, 0xce, 0x95, 0x24,
	0xd2, 0x19, 0x0c, 0x5a, 0x9b, 0x58, 0xcc, 0xe4, 0x67, 0xff, 0xb6, 0x61, 0x27, 0xa5, 0x62, 0x5f,
	0x60, 0x27, 0x3d, 0x66, 0xc7, 0x95, 0x69, 0xc3, 0x5b, 0xe3, 0x9c, 0x34, 0xa8, 0x12, 0x77, 0x7c,
	0x2b, 0x06, 0x27, 0xed, 0x35, 0x80, 0x0d, 0xed, 0x37, 0x80, 0x4d, 0xe5, 0xe7, 0x5b, 0xec, 0x2b,
	0xba, 0x85, 0xe2, 0xb1, 0xa3, 0xca, 0x5c, 0xf5, 0x25, 0x70, 0x8e, 0xd7, 0x8b, 0x32, 0xf6, 0x07,
	0xd8, 0x49, 0x8d, 0xd8, 0xd3, 0x6a, 0x96, 0xc5, 0xfa, 0x3a, 0xcf, 0x6a, 0xcf, 0x0b, 0x09, 0x20,
	0x2f, 0x02, 0xe3, 0xd5, 0xe0, 0xca, 0x35, 0x74, 0x8e, 0xd6, 0x6a, 0x32, 0xf0, 0x67, 0x74, 0xb2,
	0x3e, 0xb0, 0xe7, 0x95, 0x99, 0x72, 0x9f, 0x1c, 0xbe, 0x4e, 0x92, 0x51, 0x3f, 0xa1, 0xf7, 0x9e,
	0xa2, 0xf8, 0xe3, 0x7c, 0x71, 0x95, 0x96, 0xc3, 0x50, 0xa8, 0xfc, 0xd3, 0xee, 0x0c, 0x6b, 0x4e,
	0x97, 0xc4, 0x1f, 0xb6, 0xde, 0x7d, 0xf5, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x9e, 0xf9, 0x77,
	0x43, 0x06, 0x00, 0x00,
}
