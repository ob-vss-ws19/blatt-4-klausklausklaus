// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/show.proto

package showproto

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

type ShowMessage struct {
	CinemaId             int32    `protobuf:"varint,1,opt,name=cinemaId,proto3" json:"cinemaId,omitempty"`
	MovieId              int32    `protobuf:"varint,2,opt,name=movieId,proto3" json:"movieId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowMessage) Reset()         { *m = ShowMessage{} }
func (m *ShowMessage) String() string { return proto.CompactTextString(m) }
func (*ShowMessage) ProtoMessage()    {}
func (*ShowMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{0}
}

func (m *ShowMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowMessage.Unmarshal(m, b)
}
func (m *ShowMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowMessage.Marshal(b, m, deterministic)
}
func (m *ShowMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowMessage.Merge(m, src)
}
func (m *ShowMessage) XXX_Size() int {
	return xxx_messageInfo_ShowMessage.Size(m)
}
func (m *ShowMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ShowMessage proto.InternalMessageInfo

func (m *ShowMessage) GetCinemaId() int32 {
	if m != nil {
		return m.CinemaId
	}
	return 0
}

func (m *ShowMessage) GetMovieId() int32 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

type CreateShowRequest struct {
	CreateData           *ShowMessage `protobuf:"bytes,1,opt,name=createData,proto3" json:"createData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateShowRequest) Reset()         { *m = CreateShowRequest{} }
func (m *CreateShowRequest) String() string { return proto.CompactTextString(m) }
func (*CreateShowRequest) ProtoMessage()    {}
func (*CreateShowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{1}
}

func (m *CreateShowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShowRequest.Unmarshal(m, b)
}
func (m *CreateShowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShowRequest.Marshal(b, m, deterministic)
}
func (m *CreateShowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShowRequest.Merge(m, src)
}
func (m *CreateShowRequest) XXX_Size() int {
	return xxx_messageInfo_CreateShowRequest.Size(m)
}
func (m *CreateShowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShowRequest proto.InternalMessageInfo

func (m *CreateShowRequest) GetCreateData() *ShowMessage {
	if m != nil {
		return m.CreateData
	}
	return nil
}

type CreateShowResponse struct {
	CreateShowId         int32    `protobuf:"varint,1,opt,name=createShowId,proto3" json:"createShowId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShowResponse) Reset()         { *m = CreateShowResponse{} }
func (m *CreateShowResponse) String() string { return proto.CompactTextString(m) }
func (*CreateShowResponse) ProtoMessage()    {}
func (*CreateShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{2}
}

func (m *CreateShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShowResponse.Unmarshal(m, b)
}
func (m *CreateShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShowResponse.Marshal(b, m, deterministic)
}
func (m *CreateShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShowResponse.Merge(m, src)
}
func (m *CreateShowResponse) XXX_Size() int {
	return xxx_messageInfo_CreateShowResponse.Size(m)
}
func (m *CreateShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShowResponse proto.InternalMessageInfo

func (m *CreateShowResponse) GetCreateShowId() int32 {
	if m != nil {
		return m.CreateShowId
	}
	return 0
}

type DeleteShowRequest struct {
	DeleteShowId         int32    `protobuf:"varint,1,opt,name=deleteShowId,proto3" json:"deleteShowId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteShowRequest) Reset()         { *m = DeleteShowRequest{} }
func (m *DeleteShowRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteShowRequest) ProtoMessage()    {}
func (*DeleteShowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{3}
}

func (m *DeleteShowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteShowRequest.Unmarshal(m, b)
}
func (m *DeleteShowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteShowRequest.Marshal(b, m, deterministic)
}
func (m *DeleteShowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteShowRequest.Merge(m, src)
}
func (m *DeleteShowRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteShowRequest.Size(m)
}
func (m *DeleteShowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteShowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteShowRequest proto.InternalMessageInfo

func (m *DeleteShowRequest) GetDeleteShowId() int32 {
	if m != nil {
		return m.DeleteShowId
	}
	return 0
}

type DeleteShowResponse struct {
	Answer               bool     `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteShowResponse) Reset()         { *m = DeleteShowResponse{} }
func (m *DeleteShowResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteShowResponse) ProtoMessage()    {}
func (*DeleteShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{4}
}

func (m *DeleteShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteShowResponse.Unmarshal(m, b)
}
func (m *DeleteShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteShowResponse.Marshal(b, m, deterministic)
}
func (m *DeleteShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteShowResponse.Merge(m, src)
}
func (m *DeleteShowResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteShowResponse.Size(m)
}
func (m *DeleteShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteShowResponse proto.InternalMessageInfo

func (m *DeleteShowResponse) GetAnswer() bool {
	if m != nil {
		return m.Answer
	}
	return false
}

type DeleteShowCinemaRequest struct {
	CinemaId             int32    `protobuf:"varint,1,opt,name=cinemaId,proto3" json:"cinemaId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteShowCinemaRequest) Reset()         { *m = DeleteShowCinemaRequest{} }
func (m *DeleteShowCinemaRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteShowCinemaRequest) ProtoMessage()    {}
func (*DeleteShowCinemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{5}
}

func (m *DeleteShowCinemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteShowCinemaRequest.Unmarshal(m, b)
}
func (m *DeleteShowCinemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteShowCinemaRequest.Marshal(b, m, deterministic)
}
func (m *DeleteShowCinemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteShowCinemaRequest.Merge(m, src)
}
func (m *DeleteShowCinemaRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteShowCinemaRequest.Size(m)
}
func (m *DeleteShowCinemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteShowCinemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteShowCinemaRequest proto.InternalMessageInfo

func (m *DeleteShowCinemaRequest) GetCinemaId() int32 {
	if m != nil {
		return m.CinemaId
	}
	return 0
}

type DeleteShowCinemaResponse struct {
	Answer               bool     `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteShowCinemaResponse) Reset()         { *m = DeleteShowCinemaResponse{} }
func (m *DeleteShowCinemaResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteShowCinemaResponse) ProtoMessage()    {}
func (*DeleteShowCinemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{6}
}

func (m *DeleteShowCinemaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteShowCinemaResponse.Unmarshal(m, b)
}
func (m *DeleteShowCinemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteShowCinemaResponse.Marshal(b, m, deterministic)
}
func (m *DeleteShowCinemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteShowCinemaResponse.Merge(m, src)
}
func (m *DeleteShowCinemaResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteShowCinemaResponse.Size(m)
}
func (m *DeleteShowCinemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteShowCinemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteShowCinemaResponse proto.InternalMessageInfo

func (m *DeleteShowCinemaResponse) GetAnswer() bool {
	if m != nil {
		return m.Answer
	}
	return false
}

type DeleteShowConnectedMovieRequest struct {
	MovieId              int32    `protobuf:"varint,1,opt,name=movieId,proto3" json:"movieId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteShowConnectedMovieRequest) Reset()         { *m = DeleteShowConnectedMovieRequest{} }
func (m *DeleteShowConnectedMovieRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteShowConnectedMovieRequest) ProtoMessage()    {}
func (*DeleteShowConnectedMovieRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{7}
}

func (m *DeleteShowConnectedMovieRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteShowConnectedMovieRequest.Unmarshal(m, b)
}
func (m *DeleteShowConnectedMovieRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteShowConnectedMovieRequest.Marshal(b, m, deterministic)
}
func (m *DeleteShowConnectedMovieRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteShowConnectedMovieRequest.Merge(m, src)
}
func (m *DeleteShowConnectedMovieRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteShowConnectedMovieRequest.Size(m)
}
func (m *DeleteShowConnectedMovieRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteShowConnectedMovieRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteShowConnectedMovieRequest proto.InternalMessageInfo

func (m *DeleteShowConnectedMovieRequest) GetMovieId() int32 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

type DeleteShowConnectedMovieResponse struct {
	Answer               bool     `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteShowConnectedMovieResponse) Reset()         { *m = DeleteShowConnectedMovieResponse{} }
func (m *DeleteShowConnectedMovieResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteShowConnectedMovieResponse) ProtoMessage()    {}
func (*DeleteShowConnectedMovieResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{8}
}

func (m *DeleteShowConnectedMovieResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteShowConnectedMovieResponse.Unmarshal(m, b)
}
func (m *DeleteShowConnectedMovieResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteShowConnectedMovieResponse.Marshal(b, m, deterministic)
}
func (m *DeleteShowConnectedMovieResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteShowConnectedMovieResponse.Merge(m, src)
}
func (m *DeleteShowConnectedMovieResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteShowConnectedMovieResponse.Size(m)
}
func (m *DeleteShowConnectedMovieResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteShowConnectedMovieResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteShowConnectedMovieResponse proto.InternalMessageInfo

func (m *DeleteShowConnectedMovieResponse) GetAnswer() bool {
	if m != nil {
		return m.Answer
	}
	return false
}

type ListShowRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListShowRequest) Reset()         { *m = ListShowRequest{} }
func (m *ListShowRequest) String() string { return proto.CompactTextString(m) }
func (*ListShowRequest) ProtoMessage()    {}
func (*ListShowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{9}
}

func (m *ListShowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListShowRequest.Unmarshal(m, b)
}
func (m *ListShowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListShowRequest.Marshal(b, m, deterministic)
}
func (m *ListShowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListShowRequest.Merge(m, src)
}
func (m *ListShowRequest) XXX_Size() int {
	return xxx_messageInfo_ListShowRequest.Size(m)
}
func (m *ListShowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListShowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListShowRequest proto.InternalMessageInfo

type ListShowResponse struct {
	ShowId               []int32        `protobuf:"varint,1,rep,packed,name=showId,proto3" json:"showId,omitempty"`
	AllShowsData         []*ShowMessage `protobuf:"bytes,2,rep,name=allShowsData,proto3" json:"allShowsData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListShowResponse) Reset()         { *m = ListShowResponse{} }
func (m *ListShowResponse) String() string { return proto.CompactTextString(m) }
func (*ListShowResponse) ProtoMessage()    {}
func (*ListShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{10}
}

func (m *ListShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListShowResponse.Unmarshal(m, b)
}
func (m *ListShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListShowResponse.Marshal(b, m, deterministic)
}
func (m *ListShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListShowResponse.Merge(m, src)
}
func (m *ListShowResponse) XXX_Size() int {
	return xxx_messageInfo_ListShowResponse.Size(m)
}
func (m *ListShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListShowResponse proto.InternalMessageInfo

func (m *ListShowResponse) GetShowId() []int32 {
	if m != nil {
		return m.ShowId
	}
	return nil
}

func (m *ListShowResponse) GetAllShowsData() []*ShowMessage {
	if m != nil {
		return m.AllShowsData
	}
	return nil
}

type FindShowConnectedMovieRequest struct {
	MovieId              int32    `protobuf:"varint,1,opt,name=movieId,proto3" json:"movieId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindShowConnectedMovieRequest) Reset()         { *m = FindShowConnectedMovieRequest{} }
func (m *FindShowConnectedMovieRequest) String() string { return proto.CompactTextString(m) }
func (*FindShowConnectedMovieRequest) ProtoMessage()    {}
func (*FindShowConnectedMovieRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{11}
}

func (m *FindShowConnectedMovieRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindShowConnectedMovieRequest.Unmarshal(m, b)
}
func (m *FindShowConnectedMovieRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindShowConnectedMovieRequest.Marshal(b, m, deterministic)
}
func (m *FindShowConnectedMovieRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindShowConnectedMovieRequest.Merge(m, src)
}
func (m *FindShowConnectedMovieRequest) XXX_Size() int {
	return xxx_messageInfo_FindShowConnectedMovieRequest.Size(m)
}
func (m *FindShowConnectedMovieRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindShowConnectedMovieRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindShowConnectedMovieRequest proto.InternalMessageInfo

func (m *FindShowConnectedMovieRequest) GetMovieId() int32 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

type FindShowConnectedMovieResponse struct {
	MovieData            []*ShowMessage `protobuf:"bytes,1,rep,name=movieData,proto3" json:"movieData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FindShowConnectedMovieResponse) Reset()         { *m = FindShowConnectedMovieResponse{} }
func (m *FindShowConnectedMovieResponse) String() string { return proto.CompactTextString(m) }
func (*FindShowConnectedMovieResponse) ProtoMessage()    {}
func (*FindShowConnectedMovieResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{12}
}

func (m *FindShowConnectedMovieResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindShowConnectedMovieResponse.Unmarshal(m, b)
}
func (m *FindShowConnectedMovieResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindShowConnectedMovieResponse.Marshal(b, m, deterministic)
}
func (m *FindShowConnectedMovieResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindShowConnectedMovieResponse.Merge(m, src)
}
func (m *FindShowConnectedMovieResponse) XXX_Size() int {
	return xxx_messageInfo_FindShowConnectedMovieResponse.Size(m)
}
func (m *FindShowConnectedMovieResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindShowConnectedMovieResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindShowConnectedMovieResponse proto.InternalMessageInfo

func (m *FindShowConnectedMovieResponse) GetMovieData() []*ShowMessage {
	if m != nil {
		return m.MovieData
	}
	return nil
}

type FindShowConnectedCinemaRequest struct {
	CinemaId             int32    `protobuf:"varint,1,opt,name=cinemaId,proto3" json:"cinemaId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindShowConnectedCinemaRequest) Reset()         { *m = FindShowConnectedCinemaRequest{} }
func (m *FindShowConnectedCinemaRequest) String() string { return proto.CompactTextString(m) }
func (*FindShowConnectedCinemaRequest) ProtoMessage()    {}
func (*FindShowConnectedCinemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{13}
}

func (m *FindShowConnectedCinemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindShowConnectedCinemaRequest.Unmarshal(m, b)
}
func (m *FindShowConnectedCinemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindShowConnectedCinemaRequest.Marshal(b, m, deterministic)
}
func (m *FindShowConnectedCinemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindShowConnectedCinemaRequest.Merge(m, src)
}
func (m *FindShowConnectedCinemaRequest) XXX_Size() int {
	return xxx_messageInfo_FindShowConnectedCinemaRequest.Size(m)
}
func (m *FindShowConnectedCinemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindShowConnectedCinemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindShowConnectedCinemaRequest proto.InternalMessageInfo

func (m *FindShowConnectedCinemaRequest) GetCinemaId() int32 {
	if m != nil {
		return m.CinemaId
	}
	return 0
}

type FindShowConnectedCinemaResponse struct {
	Ids                  []int32        `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	CinemaData           []*ShowMessage `protobuf:"bytes,2,rep,name=cinemaData,proto3" json:"cinemaData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FindShowConnectedCinemaResponse) Reset()         { *m = FindShowConnectedCinemaResponse{} }
func (m *FindShowConnectedCinemaResponse) String() string { return proto.CompactTextString(m) }
func (*FindShowConnectedCinemaResponse) ProtoMessage()    {}
func (*FindShowConnectedCinemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e145b6e6608d39, []int{14}
}

func (m *FindShowConnectedCinemaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindShowConnectedCinemaResponse.Unmarshal(m, b)
}
func (m *FindShowConnectedCinemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindShowConnectedCinemaResponse.Marshal(b, m, deterministic)
}
func (m *FindShowConnectedCinemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindShowConnectedCinemaResponse.Merge(m, src)
}
func (m *FindShowConnectedCinemaResponse) XXX_Size() int {
	return xxx_messageInfo_FindShowConnectedCinemaResponse.Size(m)
}
func (m *FindShowConnectedCinemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindShowConnectedCinemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindShowConnectedCinemaResponse proto.InternalMessageInfo

func (m *FindShowConnectedCinemaResponse) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *FindShowConnectedCinemaResponse) GetCinemaData() []*ShowMessage {
	if m != nil {
		return m.CinemaData
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowMessage)(nil), "showproto.ShowMessage")
	proto.RegisterType((*CreateShowRequest)(nil), "showproto.CreateShowRequest")
	proto.RegisterType((*CreateShowResponse)(nil), "showproto.CreateShowResponse")
	proto.RegisterType((*DeleteShowRequest)(nil), "showproto.DeleteShowRequest")
	proto.RegisterType((*DeleteShowResponse)(nil), "showproto.DeleteShowResponse")
	proto.RegisterType((*DeleteShowCinemaRequest)(nil), "showproto.DeleteShowCinemaRequest")
	proto.RegisterType((*DeleteShowCinemaResponse)(nil), "showproto.DeleteShowCinemaResponse")
	proto.RegisterType((*DeleteShowConnectedMovieRequest)(nil), "showproto.DeleteShowConnectedMovieRequest")
	proto.RegisterType((*DeleteShowConnectedMovieResponse)(nil), "showproto.DeleteShowConnectedMovieResponse")
	proto.RegisterType((*ListShowRequest)(nil), "showproto.ListShowRequest")
	proto.RegisterType((*ListShowResponse)(nil), "showproto.ListShowResponse")
	proto.RegisterType((*FindShowConnectedMovieRequest)(nil), "showproto.FindShowConnectedMovieRequest")
	proto.RegisterType((*FindShowConnectedMovieResponse)(nil), "showproto.FindShowConnectedMovieResponse")
	proto.RegisterType((*FindShowConnectedCinemaRequest)(nil), "showproto.FindShowConnectedCinemaRequest")
	proto.RegisterType((*FindShowConnectedCinemaResponse)(nil), "showproto.FindShowConnectedCinemaResponse")
}

func init() { proto.RegisterFile("proto/show.proto", fileDescriptor_67e145b6e6608d39) }

var fileDescriptor_67e145b6e6608d39 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5b, 0x6f, 0xd3, 0x30,
	0x14, 0x5e, 0x57, 0x3a, 0xba, 0x6f, 0x95, 0x68, 0xcf, 0x43, 0x29, 0x81, 0xb1, 0xe2, 0xa7, 0xee,
	0xa2, 0x22, 0x0d, 0xc4, 0x65, 0x43, 0xbc, 0x64, 0x20, 0x4d, 0x63, 0x2f, 0x41, 0xe2, 0x3d, 0x24,
	0x86, 0x45, 0x74, 0xf1, 0x88, 0x33, 0xca, 0x9f, 0xe3, 0xbf, 0xa1, 0x18, 0xc7, 0x71, 0x68, 0x92,
	0x46, 0x7b, 0xb3, 0x8f, 0xbf, 0xcb, 0xc9, 0xc9, 0xf9, 0x30, 0xbc, 0x49, 0x44, 0x2a, 0x9e, 0xcb,
	0x2b, 0xb1, 0x9c, 0xab, 0x23, 0x6d, 0x67, 0x67, 0x75, 0x64, 0x2e, 0x76, 0x3e, 0x5f, 0x89, 0xe5,
	0x25, 0x97, 0xd2, 0xff, 0xce, 0xc9, 0x41, 0x3f, 0x88, 0x62, 0x7e, 0xed, 0x9f, 0x87, 0x93, 0xce,
	0xb4, 0x33, 0xeb, 0x79, 0xe6, 0x4e, 0x13, 0xdc, 0xbf, 0x16, 0xbf, 0x22, 0x7e, 0x1e, 0x4e, 0x36,
	0xd5, 0x53, 0x7e, 0x65, 0x17, 0x18, 0xb9, 0x09, 0xf7, 0x53, 0x9e, 0x49, 0x79, 0xfc, 0xe7, 0x2d,
	0x97, 0x29, 0xbd, 0x02, 0x02, 0x55, 0x3c, 0xf3, 0x53, 0x5f, 0x89, 0xed, 0x1c, 0x8f, 0xe7, 0xc6,
	0x79, 0x6e, 0xd9, 0x7a, 0x16, 0x92, 0xbd, 0x01, 0xd9, 0x62, 0xf2, 0x46, 0xc4, 0x92, 0x13, 0xc3,
	0x20, 0x30, 0x55, 0xd3, 0x5c, 0xa9, 0xc6, 0x5e, 0x63, 0x74, 0xc6, 0x17, 0xbc, 0xdc, 0x06, 0xc3,
	0x20, 0x34, 0xc5, 0x82, 0x68, 0xd7, 0xd8, 0x11, 0xc8, 0x26, 0x6a, 0xcb, 0x31, 0xb6, 0xfc, 0x58,
	0x2e, 0x79, 0xa2, 0x38, 0x7d, 0x4f, 0xdf, 0xd8, 0x7b, 0x4c, 0x0b, 0xb4, 0x2b, 0xe2, 0x98, 0x07,
	0x29, 0x0f, 0x5d, 0x35, 0xa6, 0xdc, 0xb5, 0x61, 0x8e, 0xec, 0x14, 0xcf, 0x1a, 0xf8, 0x6b, 0xcc,
	0x4f, 0xb1, 0x57, 0x41, 0xbe, 0xcc, 0x7e, 0x44, 0xee, 0x6d, 0xfd, 0xa7, 0x4e, 0xf9, 0x3f, 0x9d,
	0x54, 0x76, 0xae, 0xc9, 0x6b, 0x8c, 0x47, 0x78, 0xf0, 0x29, 0x92, 0xa9, 0x35, 0x5a, 0xf6, 0x0d,
	0xc3, 0xa2, 0x54, 0xd0, 0x65, 0x3e, 0xe8, 0xee, 0xac, 0xe7, 0xe9, 0x1b, 0x9d, 0x60, 0xe0, 0x2f,
	0x16, 0x19, 0x54, 0xaa, 0x7d, 0xd8, 0x9c, 0x76, 0x1b, 0xf6, 0xa1, 0x84, 0x65, 0x6f, 0xb1, 0xfb,
	0x31, 0x8a, 0xc3, 0xbb, 0x7c, 0xf1, 0x17, 0x3c, 0xad, 0xa3, 0xea, 0x86, 0x5f, 0x62, 0x5b, 0x81,
	0xf5, 0x96, 0x36, 0x75, 0x55, 0x00, 0xd9, 0xbb, 0x0a, 0xdd, 0xf6, 0x1b, 0xf0, 0x03, 0x7b, 0xb5,
	0x6c, 0xdd, 0xd6, 0x10, 0xdd, 0x28, 0x94, 0x7a, 0x88, 0xd9, 0x51, 0xe5, 0x49, 0x61, 0x5a, 0xcc,
	0xcf, 0x42, 0x1e, 0xff, 0xe9, 0xe1, 0x5e, 0xf6, 0x46, 0x17, 0x40, 0x11, 0x2c, 0x7a, 0x62, 0x51,
	0x57, 0xc2, 0xeb, 0xec, 0xd6, 0xbc, 0xfe, 0xeb, 0x8e, 0x6d, 0x64, 0x62, 0xc5, 0x2a, 0x95, 0xc4,
	0x56, 0x22, 0x58, 0x12, 0x5b, 0xcd, 0x19, 0xdb, 0xa0, 0xdf, 0x78, 0x54, 0x9b, 0x08, 0x3a, 0xac,
	0x64, 0x57, 0x4f, 0xdd, 0x39, 0x6a, 0x07, 0x36, 0xce, 0xb7, 0x98, 0xd4, 0x25, 0x82, 0x0e, 0x9a,
	0xb5, 0xec, 0x0d, 0x74, 0x0e, 0x5b, 0x61, 0x8d, 0xed, 0x07, 0xf4, 0xf3, 0xe4, 0x90, 0x63, 0x51,
	0xff, 0x4b, 0x98, 0xf3, 0xb8, 0xf2, 0xcd, 0xc8, 0x24, 0x78, 0x58, 0xb3, 0x47, 0xb4, 0x6f, 0x31,
	0x9b, 0x37, 0xd5, 0x39, 0x68, 0x03, 0x35, 0x9e, 0x02, 0xe3, 0xea, 0x44, 0xd1, 0xac, 0x49, 0xa7,
	0x34, 0xad, 0xfd, 0x16, 0xc8, 0xdc, 0xf0, 0xeb, 0x96, 0xc2, 0xbd, 0xf8, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xd6, 0x71, 0xed, 0xd1, 0xc7, 0x06, 0x00, 0x00,
}
