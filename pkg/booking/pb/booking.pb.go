// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: pkg/booking/pb/booking.proto

package pb

import (
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

//Ticket booking
type CheckoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Compartmentid string `protobuf:"bytes,1,opt,name=compartmentid,proto3" json:"compartmentid,omitempty"`
	TrainId       string `protobuf:"bytes,2,opt,name=trainId,proto3" json:"trainId,omitempty"`
	Userid        int64  `protobuf:"varint,3,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *CheckoutRequest) Reset() {
	*x = CheckoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_booking_pb_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutRequest) ProtoMessage() {}

func (x *CheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_booking_pb_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutRequest.ProtoReflect.Descriptor instead.
func (*CheckoutRequest) Descriptor() ([]byte, []int) {
	return file_pkg_booking_pb_booking_proto_rawDescGZIP(), []int{0}
}

func (x *CheckoutRequest) GetCompartmentid() string {
	if x != nil {
		return x.Compartmentid
	}
	return ""
}

func (x *CheckoutRequest) GetTrainId() string {
	if x != nil {
		return x.TrainId
	}
	return ""
}

func (x *CheckoutRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

type CheckoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status             int64        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error              string       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	TrainName          string       `protobuf:"bytes,3,opt,name=trainName,proto3" json:"trainName,omitempty"`
	Trainnumber        int64        `protobuf:"varint,4,opt,name=trainnumber,proto3" json:"trainnumber,omitempty"`
	SoureceStation     string       `protobuf:"bytes,5,opt,name=soureceStation,proto3" json:"soureceStation,omitempty"`
	DestinationStation string       `protobuf:"bytes,6,opt,name=destinationStation,proto3" json:"destinationStation,omitempty"`
	Username           string       `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	Travelers          []*Travelers `protobuf:"bytes,8,rep,name=Travelers,proto3" json:"Travelers,omitempty"`
}

func (x *CheckoutResponse) Reset() {
	*x = CheckoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_booking_pb_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutResponse) ProtoMessage() {}

func (x *CheckoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_booking_pb_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutResponse.ProtoReflect.Descriptor instead.
func (*CheckoutResponse) Descriptor() ([]byte, []int) {
	return file_pkg_booking_pb_booking_proto_rawDescGZIP(), []int{1}
}

func (x *CheckoutResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CheckoutResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CheckoutResponse) GetTrainName() string {
	if x != nil {
		return x.TrainName
	}
	return ""
}

func (x *CheckoutResponse) GetTrainnumber() int64 {
	if x != nil {
		return x.Trainnumber
	}
	return 0
}

func (x *CheckoutResponse) GetSoureceStation() string {
	if x != nil {
		return x.SoureceStation
	}
	return ""
}

func (x *CheckoutResponse) GetDestinationStation() string {
	if x != nil {
		return x.DestinationStation
	}
	return ""
}

func (x *CheckoutResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CheckoutResponse) GetTravelers() []*Travelers {
	if x != nil {
		return x.Travelers
	}
	return nil
}

type Travelers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Travelername string `protobuf:"bytes,1,opt,name=Travelername,proto3" json:"Travelername,omitempty"`
}

func (x *Travelers) Reset() {
	*x = Travelers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_booking_pb_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Travelers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Travelers) ProtoMessage() {}

func (x *Travelers) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_booking_pb_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Travelers.ProtoReflect.Descriptor instead.
func (*Travelers) Descriptor() ([]byte, []int) {
	return file_pkg_booking_pb_booking_proto_rawDescGZIP(), []int{2}
}

func (x *Travelers) GetTravelername() string {
	if x != nil {
		return x.Travelername
	}
	return ""
}

//Booking
type SearchCompartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trainid string `protobuf:"bytes,1,opt,name=trainid,proto3" json:"trainid,omitempty"`
}

func (x *SearchCompartmentRequest) Reset() {
	*x = SearchCompartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_booking_pb_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCompartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCompartmentRequest) ProtoMessage() {}

func (x *SearchCompartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_booking_pb_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCompartmentRequest.ProtoReflect.Descriptor instead.
func (*SearchCompartmentRequest) Descriptor() ([]byte, []int) {
	return file_pkg_booking_pb_booking_proto_rawDescGZIP(), []int{3}
}

func (x *SearchCompartmentRequest) GetTrainid() string {
	if x != nil {
		return x.Trainid
	}
	return ""
}

type SearchCompartmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Compartment []*Compartment `protobuf:"bytes,1,rep,name=compartment,proto3" json:"compartment,omitempty"`
}

func (x *SearchCompartmentResponse) Reset() {
	*x = SearchCompartmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_booking_pb_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCompartmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCompartmentResponse) ProtoMessage() {}

func (x *SearchCompartmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_booking_pb_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCompartmentResponse.ProtoReflect.Descriptor instead.
func (*SearchCompartmentResponse) Descriptor() ([]byte, []int) {
	return file_pkg_booking_pb_booking_proto_rawDescGZIP(), []int{4}
}

func (x *SearchCompartmentResponse) GetCompartment() []*Compartment {
	if x != nil {
		return x.Compartment
	}
	return nil
}

type Compartment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Compartmentid     string `protobuf:"bytes,1,opt,name=compartmentid,proto3" json:"compartmentid,omitempty"`
	Price             string `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	Typeofseat        string `protobuf:"bytes,3,opt,name=typeofseat,proto3" json:"typeofseat,omitempty"`
	CompartmentName   string `protobuf:"bytes,4,opt,name=CompartmentName,proto3" json:"CompartmentName,omitempty"`
	Availablitystatus string `protobuf:"bytes,5,opt,name=availablitystatus,proto3" json:"availablitystatus,omitempty"` // repeated seatDetails seatDetails =5;
}

func (x *Compartment) Reset() {
	*x = Compartment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_booking_pb_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Compartment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Compartment) ProtoMessage() {}

func (x *Compartment) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_booking_pb_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Compartment.ProtoReflect.Descriptor instead.
func (*Compartment) Descriptor() ([]byte, []int) {
	return file_pkg_booking_pb_booking_proto_rawDescGZIP(), []int{5}
}

func (x *Compartment) GetCompartmentid() string {
	if x != nil {
		return x.Compartmentid
	}
	return ""
}

func (x *Compartment) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Compartment) GetTypeofseat() string {
	if x != nil {
		return x.Typeofseat
	}
	return ""
}

func (x *Compartment) GetCompartmentName() string {
	if x != nil {
		return x.CompartmentName
	}
	return ""
}

func (x *Compartment) GetAvailablitystatus() string {
	if x != nil {
		return x.Availablitystatus
	}
	return ""
}

// message seatDetails{
//     string price = 1;
//     string isreserved = 2;
//     string seattype = 3;
//     int64 seatnumber = 4;
// }
//Search train
type SearchTrainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date                 string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Sourcestationid      string `protobuf:"bytes,2,opt,name=sourcestationid,proto3" json:"sourcestationid,omitempty"`
	Destinationstationid string `protobuf:"bytes,3,opt,name=destinationstationid,proto3" json:"destinationstationid,omitempty"`
}

func (x *SearchTrainRequest) Reset() {
	*x = SearchTrainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_booking_pb_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTrainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTrainRequest) ProtoMessage() {}

func (x *SearchTrainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_booking_pb_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTrainRequest.ProtoReflect.Descriptor instead.
func (*SearchTrainRequest) Descriptor() ([]byte, []int) {
	return file_pkg_booking_pb_booking_proto_rawDescGZIP(), []int{6}
}

func (x *SearchTrainRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *SearchTrainRequest) GetSourcestationid() string {
	if x != nil {
		return x.Sourcestationid
	}
	return ""
}

func (x *SearchTrainRequest) GetDestinationstationid() string {
	if x != nil {
		return x.Destinationstationid
	}
	return ""
}

type SearchTrainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    int64        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error     string       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Traindata []*TrainData `protobuf:"bytes,3,rep,name=traindata,proto3" json:"traindata,omitempty"`
}

func (x *SearchTrainResponse) Reset() {
	*x = SearchTrainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_booking_pb_booking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTrainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTrainResponse) ProtoMessage() {}

func (x *SearchTrainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_booking_pb_booking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTrainResponse.ProtoReflect.Descriptor instead.
func (*SearchTrainResponse) Descriptor() ([]byte, []int) {
	return file_pkg_booking_pb_booking_proto_rawDescGZIP(), []int{7}
}

func (x *SearchTrainResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SearchTrainResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SearchTrainResponse) GetTraindata() []*TrainData {
	if x != nil {
		return x.Traindata
	}
	return nil
}

type TrainData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trainname    string  `protobuf:"bytes,1,opt,name=trainname,proto3" json:"trainname,omitempty"`
	TrainNumber  int64   `protobuf:"varint,2,opt,name=trainNumber,proto3" json:"trainNumber,omitempty"`
	StartingTime string  `protobuf:"bytes,3,opt,name=startingTime,proto3" json:"startingTime,omitempty"`
	Endingtime   string  `protobuf:"bytes,4,opt,name=endingtime,proto3" json:"endingtime,omitempty"`
	Distance     float32 `protobuf:"fixed32,5,opt,name=distance,proto3" json:"distance,omitempty"`
	Trainid      string  `protobuf:"bytes,6,opt,name=trainid,proto3" json:"trainid,omitempty"`
}

func (x *TrainData) Reset() {
	*x = TrainData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_booking_pb_booking_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainData) ProtoMessage() {}

func (x *TrainData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_booking_pb_booking_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainData.ProtoReflect.Descriptor instead.
func (*TrainData) Descriptor() ([]byte, []int) {
	return file_pkg_booking_pb_booking_proto_rawDescGZIP(), []int{8}
}

func (x *TrainData) GetTrainname() string {
	if x != nil {
		return x.Trainname
	}
	return ""
}

func (x *TrainData) GetTrainNumber() int64 {
	if x != nil {
		return x.TrainNumber
	}
	return 0
}

func (x *TrainData) GetStartingTime() string {
	if x != nil {
		return x.StartingTime
	}
	return ""
}

func (x *TrainData) GetEndingtime() string {
	if x != nil {
		return x.Endingtime
	}
	return ""
}

func (x *TrainData) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *TrainData) GetTrainid() string {
	if x != nil {
		return x.Trainid
	}
	return ""
}

var File_pkg_booking_pb_booking_proto protoreflect.FileDescriptor

var file_pkg_booking_pb_booking_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x62,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x69, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x64, 0x22, 0xa6, 0x02, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x65, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x6f, 0x75, 0x72, 0x65, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x12, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x54, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73,
	0x52, 0x09, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x22, 0x2f, 0x0a, 0x09, 0x54,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x76,
	0x65, 0x6c, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x18,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x64, 0x22, 0x53, 0x0a, 0x19, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xc1, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x6f, 0x66, 0x73, 0x65, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x6f, 0x66, 0x73,
	0x65, 0x61, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x69, 0x74, 0x79, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x69, 0x74, 0x79, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x12,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x64,
	0x12, 0x32, 0x0a, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x69, 0x64, 0x22, 0x75, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x72,
	0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x30, 0x0a, 0x09, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc5, 0x01, 0x0a, 0x09,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x64, 0x32, 0x80, 0x02, 0x0a, 0x11, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x0b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12,
	0x18, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_booking_pb_booking_proto_rawDescOnce sync.Once
	file_pkg_booking_pb_booking_proto_rawDescData = file_pkg_booking_pb_booking_proto_rawDesc
)

func file_pkg_booking_pb_booking_proto_rawDescGZIP() []byte {
	file_pkg_booking_pb_booking_proto_rawDescOnce.Do(func() {
		file_pkg_booking_pb_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_booking_pb_booking_proto_rawDescData)
	})
	return file_pkg_booking_pb_booking_proto_rawDescData
}

var file_pkg_booking_pb_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_booking_pb_booking_proto_goTypes = []interface{}{
	(*CheckoutRequest)(nil),           // 0: Booking.CheckoutRequest
	(*CheckoutResponse)(nil),          // 1: Booking.CheckoutResponse
	(*Travelers)(nil),                 // 2: Booking.Travelers
	(*SearchCompartmentRequest)(nil),  // 3: Booking.SearchCompartmentRequest
	(*SearchCompartmentResponse)(nil), // 4: Booking.SearchCompartmentResponse
	(*Compartment)(nil),               // 5: Booking.compartment
	(*SearchTrainRequest)(nil),        // 6: Booking.SearchTrainRequest
	(*SearchTrainResponse)(nil),       // 7: Booking.SearchTrainResponse
	(*TrainData)(nil),                 // 8: Booking.TrainData
}
var file_pkg_booking_pb_booking_proto_depIdxs = []int32{
	2, // 0: Booking.CheckoutResponse.Travelers:type_name -> Booking.Travelers
	5, // 1: Booking.SearchCompartmentResponse.compartment:type_name -> Booking.compartment
	8, // 2: Booking.SearchTrainResponse.traindata:type_name -> Booking.TrainData
	6, // 3: Booking.BookingManagement.SearchTrain:input_type -> Booking.SearchTrainRequest
	3, // 4: Booking.BookingManagement.SearchCompartment:input_type -> Booking.SearchCompartmentRequest
	0, // 5: Booking.BookingManagement.Checkout:input_type -> Booking.CheckoutRequest
	7, // 6: Booking.BookingManagement.SearchTrain:output_type -> Booking.SearchTrainResponse
	4, // 7: Booking.BookingManagement.SearchCompartment:output_type -> Booking.SearchCompartmentResponse
	1, // 8: Booking.BookingManagement.Checkout:output_type -> Booking.CheckoutResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_booking_pb_booking_proto_init() }
func file_pkg_booking_pb_booking_proto_init() {
	if File_pkg_booking_pb_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_booking_pb_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutRequest); i {
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
		file_pkg_booking_pb_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutResponse); i {
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
		file_pkg_booking_pb_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Travelers); i {
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
		file_pkg_booking_pb_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCompartmentRequest); i {
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
		file_pkg_booking_pb_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCompartmentResponse); i {
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
		file_pkg_booking_pb_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Compartment); i {
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
		file_pkg_booking_pb_booking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTrainRequest); i {
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
		file_pkg_booking_pb_booking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTrainResponse); i {
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
		file_pkg_booking_pb_booking_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainData); i {
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
			RawDescriptor: file_pkg_booking_pb_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_booking_pb_booking_proto_goTypes,
		DependencyIndexes: file_pkg_booking_pb_booking_proto_depIdxs,
		MessageInfos:      file_pkg_booking_pb_booking_proto_msgTypes,
	}.Build()
	File_pkg_booking_pb_booking_proto = out.File
	file_pkg_booking_pb_booking_proto_rawDesc = nil
	file_pkg_booking_pb_booking_proto_goTypes = nil
	file_pkg_booking_pb_booking_proto_depIdxs = nil
}
