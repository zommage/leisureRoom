// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/location/locations.proto

/*
Package location is a generated protocol buffer package.

It is generated from these files:
	google/cloud/location/locations.proto

It has these top-level messages:
	ListLocationsRequest
	ListLocationsResponse
	GetLocationRequest
	Location
*/
package location

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"

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

// The request message for [Locations.ListLocations][google.cloud.location.Locations.ListLocations].
type ListLocationsRequest struct {
	// The resource that owns the locations collection, if applicable.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The standard list filter.
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// The standard list page size.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// The standard list page token.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListLocationsRequest) Reset()                    { *m = ListLocationsRequest{} }
func (m *ListLocationsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLocationsRequest) ProtoMessage()               {}
func (*ListLocationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListLocationsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListLocationsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListLocationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListLocationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for [Locations.ListLocations][google.cloud.location.Locations.ListLocations].
type ListLocationsResponse struct {
	// A list of locations that matches the specified filter in the request.
	Locations []*Location `protobuf:"bytes,1,rep,name=locations" json:"locations,omitempty"`
	// The standard List next-page token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLocationsResponse) Reset()                    { *m = ListLocationsResponse{} }
func (m *ListLocationsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLocationsResponse) ProtoMessage()               {}
func (*ListLocationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListLocationsResponse) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *ListLocationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for [Locations.GetLocation][google.cloud.location.Locations.GetLocation].
type GetLocationRequest struct {
	// Resource name for the location.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetLocationRequest) Reset()                    { *m = GetLocationRequest{} }
func (m *GetLocationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLocationRequest) ProtoMessage()               {}
func (*GetLocationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetLocationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A resource that represents Google Cloud Platform location.
type Location struct {
	// Resource name for the location, which may vary between implementations.
	// For example: `"projects/example-project/locations/us-east1"`
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The canonical id for this location. For example: `"us-east1"`.
	LocationId string `protobuf:"bytes,4,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	// Cross-service attributes for the location. For example
	//
	//     {"cloud.googleapis.com/region": "us-east1"}
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Service-specific metadata. For example the available capacity at the given
	// location.
	Metadata *google_protobuf1.Any `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Location) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Location) GetLocationId() string {
	if m != nil {
		return m.LocationId
	}
	return ""
}

func (m *Location) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Location) GetMetadata() *google_protobuf1.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*ListLocationsRequest)(nil), "google.cloud.location.ListLocationsRequest")
	proto.RegisterType((*ListLocationsResponse)(nil), "google.cloud.location.ListLocationsResponse")
	proto.RegisterType((*GetLocationRequest)(nil), "google.cloud.location.GetLocationRequest")
	proto.RegisterType((*Location)(nil), "google.cloud.location.Location")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Locations service

type LocationsClient interface {
	// Lists information about the supported locations for this service.
	ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error)
	// Get information about a location.
	GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*Location, error)
}

type locationsClient struct {
	cc *grpc.ClientConn
}

func NewLocationsClient(cc *grpc.ClientConn) LocationsClient {
	return &locationsClient{cc}
}

func (c *locationsClient) ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error) {
	out := new(ListLocationsResponse)
	err := grpc.Invoke(ctx, "/google.cloud.location.Locations/ListLocations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationsClient) GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := grpc.Invoke(ctx, "/google.cloud.location.Locations/GetLocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Locations service

type LocationsServer interface {
	// Lists information about the supported locations for this service.
	ListLocations(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error)
	// Get information about a location.
	GetLocation(context.Context, *GetLocationRequest) (*Location, error)
}

func RegisterLocationsServer(s *grpc.Server, srv LocationsServer) {
	s.RegisterService(&_Locations_serviceDesc, srv)
}

func _Locations_ListLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).ListLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.location.Locations/ListLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).ListLocations(ctx, req.(*ListLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Locations_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.location.Locations/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).GetLocation(ctx, req.(*GetLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Locations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.location.Locations",
	HandlerType: (*LocationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLocations",
			Handler:    _Locations_ListLocations_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _Locations_GetLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/location/locations.proto",
}

func init() { proto.RegisterFile("google/cloud/location/locations.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0x3a, 0x6d, 0x94, 0x8c, 0x55, 0x40, 0xa3, 0x14, 0xb9, 0x06, 0x94, 0xd4, 0x08, 0x48,
	0x0b, 0xf2, 0x42, 0xb8, 0xf0, 0xa3, 0x1c, 0x28, 0x42, 0x08, 0x29, 0x87, 0xc8, 0x70, 0xe2, 0x12,
	0x6d, 0x92, 0xad, 0x65, 0xea, 0xec, 0x1a, 0x7b, 0x53, 0x91, 0xa2, 0xf6, 0x80, 0x78, 0x83, 0xbe,
	0x04, 0xef, 0xc3, 0x2b, 0xf0, 0x02, 0xdc, 0x38, 0x22, 0xaf, 0x7f, 0x12, 0x8a, 0x4b, 0xb8, 0xcd,
	0xce, 0x7c, 0xb3, 0xdf, 0x7c, 0xf3, 0xad, 0x0d, 0x77, 0x7c, 0x29, 0xfd, 0x90, 0xd3, 0x49, 0x28,
	0xe7, 0x53, 0x1a, 0xca, 0x09, 0x53, 0x81, 0x14, 0x65, 0x90, 0xb8, 0x51, 0x2c, 0x95, 0xc4, 0xed,
	0x0c, 0xe6, 0x6a, 0x98, 0x5b, 0x54, 0xed, 0x9b, 0x79, 0x37, 0x8b, 0x02, 0xca, 0x84, 0x90, 0x6a,
	0xb5, 0xc9, 0xde, 0xc9, 0xab, 0xfa, 0x34, 0x9e, 0x1f, 0x52, 0x26, 0x16, 0x59, 0xc9, 0x39, 0x83,
	0xd6, 0x20, 0x48, 0xd4, 0xa0, 0xa0, 0xf1, 0xf8, 0xc7, 0x39, 0x4f, 0x14, 0x22, 0x6c, 0x08, 0x36,
	0xe3, 0x16, 0xe9, 0x90, 0x6e, 0xd3, 0xd3, 0x31, 0x5e, 0x87, 0xfa, 0x61, 0x10, 0x2a, 0x1e, 0x5b,
	0x86, 0xce, 0xe6, 0x27, 0xbc, 0x01, 0xcd, 0x88, 0xf9, 0x7c, 0x94, 0x04, 0x27, 0xdc, 0xaa, 0x75,
	0x48, 0x77, 0xd3, 0x6b, 0xa4, 0x89, 0xb7, 0xc1, 0x09, 0xc7, 0x5b, 0x00, 0xba, 0xa8, 0xe4, 0x11,
	0x17, 0xd6, 0x86, 0x6e, 0xd4, 0xf0, 0x77, 0x69, 0xc2, 0x39, 0x83, 0xed, 0x0b, 0xfc, 0x49, 0x24,
	0x45, 0xc2, 0xb1, 0x0f, 0xcd, 0x52, 0xbb, 0x45, 0x3a, 0xb5, 0xae, 0xd9, 0x6b, 0xbb, 0x95, 0xe2,
	0xdd, 0xa2, 0xd9, 0x5b, 0x76, 0xe0, 0x5d, 0xb8, 0x2a, 0xf8, 0x27, 0x35, 0x5a, 0xe1, 0xce, 0x86,
	0xde, 0x4a, 0xd3, 0xc3, 0x92, 0xbf, 0x0b, 0xf8, 0x9a, 0x97, 0xf4, 0xff, 0x50, 0xef, 0xfc, 0x24,
	0xd0, 0x28, 0x70, 0x95, 0xeb, 0x69, 0x83, 0x59, 0xf0, 0x8f, 0x82, 0x69, 0x2e, 0x15, 0x8a, 0xd4,
	0x9b, 0x29, 0xbe, 0x84, 0x7a, 0xc8, 0xc6, 0x3c, 0x4c, 0x2c, 0x43, 0xeb, 0xb9, 0xbf, 0x46, 0x8f,
	0x3b, 0xd0, 0xe8, 0x57, 0x42, 0xc5, 0x0b, 0x2f, 0x6f, 0xc5, 0x87, 0xd0, 0x98, 0x71, 0xc5, 0xa6,
	0x4c, 0x31, 0xbd, 0x6b, 0xb3, 0xd7, 0x2a, 0xae, 0x29, 0xec, 0x75, 0x5f, 0x88, 0x85, 0x57, 0xa2,
	0xec, 0xa7, 0x60, 0xae, 0x5c, 0x84, 0xd7, 0xa0, 0x76, 0xc4, 0x17, 0xf9, 0xe4, 0x69, 0x88, 0x2d,
	0xd8, 0x3c, 0x66, 0xe1, 0x9c, 0xe7, 0x1b, 0xca, 0x0e, 0xcf, 0x8c, 0x27, 0xa4, 0xf7, 0xcd, 0x80,
	0x66, 0x69, 0x0d, 0x9e, 0x13, 0xd8, 0xfa, 0xc3, 0x2c, 0xbc, 0x54, 0x41, 0xc5, 0x93, 0xb2, 0x1f,
	0xfc, 0x1f, 0x38, 0xf3, 0xdf, 0xb9, 0xf7, 0xe5, 0xfb, 0x8f, 0x73, 0x63, 0x17, 0xdb, 0xf4, 0xf8,
	0x11, 0xfd, 0x9c, 0x2e, 0xb8, 0x1f, 0xc5, 0xf2, 0x03, 0x9f, 0xa8, 0x84, 0xee, 0x9f, 0x2e, 0xbf,
	0x0b, 0xfc, 0x4a, 0xc0, 0x5c, 0xb1, 0x10, 0xf7, 0x2e, 0xa1, 0xf9, 0xdb, 0x66, 0x7b, 0xdd, 0x83,
	0x72, 0xf6, 0xf4, 0x10, 0xb7, 0x71, 0xb7, 0x6a, 0x88, 0xe5, 0x0c, 0x74, 0xff, 0xf4, 0x40, 0xc2,
	0xce, 0x44, 0xce, 0xaa, 0x2f, 0x3c, 0xb8, 0x52, 0xea, 0x1b, 0xa6, 0x1e, 0x0d, 0xc9, 0xfb, 0x7e,
	0x0e, 0xf4, 0x65, 0xc8, 0x84, 0xef, 0xca, 0xd8, 0xa7, 0x3e, 0x17, 0xda, 0x41, 0x9a, 0x95, 0x58,
	0x14, 0x24, 0x17, 0xfe, 0x06, 0xcf, 0x8b, 0xe0, 0x17, 0x21, 0xe3, 0xba, 0x06, 0x3f, 0xfe, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x0d, 0x57, 0x9e, 0xa7, 0x39, 0x04, 0x00, 0x00,
}
