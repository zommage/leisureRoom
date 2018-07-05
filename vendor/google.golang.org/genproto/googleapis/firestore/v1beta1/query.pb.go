// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/v1beta1/query.proto

package firestore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf3 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A sort direction.
type StructuredQuery_Direction int32

const (
	// Unspecified.
	StructuredQuery_DIRECTION_UNSPECIFIED StructuredQuery_Direction = 0
	// Ascending.
	StructuredQuery_ASCENDING StructuredQuery_Direction = 1
	// Descending.
	StructuredQuery_DESCENDING StructuredQuery_Direction = 2
)

var StructuredQuery_Direction_name = map[int32]string{
	0: "DIRECTION_UNSPECIFIED",
	1: "ASCENDING",
	2: "DESCENDING",
}
var StructuredQuery_Direction_value = map[string]int32{
	"DIRECTION_UNSPECIFIED": 0,
	"ASCENDING":             1,
	"DESCENDING":            2,
}

func (x StructuredQuery_Direction) String() string {
	return proto.EnumName(StructuredQuery_Direction_name, int32(x))
}
func (StructuredQuery_Direction) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

// A composite filter operator.
type StructuredQuery_CompositeFilter_Operator int32

const (
	// Unspecified. This value must not be used.
	StructuredQuery_CompositeFilter_OPERATOR_UNSPECIFIED StructuredQuery_CompositeFilter_Operator = 0
	// The results are required to satisfy each of the combined filters.
	StructuredQuery_CompositeFilter_AND StructuredQuery_CompositeFilter_Operator = 1
)

var StructuredQuery_CompositeFilter_Operator_name = map[int32]string{
	0: "OPERATOR_UNSPECIFIED",
	1: "AND",
}
var StructuredQuery_CompositeFilter_Operator_value = map[string]int32{
	"OPERATOR_UNSPECIFIED": 0,
	"AND": 1,
}

func (x StructuredQuery_CompositeFilter_Operator) String() string {
	return proto.EnumName(StructuredQuery_CompositeFilter_Operator_name, int32(x))
}
func (StructuredQuery_CompositeFilter_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 2, 0}
}

// A field filter operator.
type StructuredQuery_FieldFilter_Operator int32

const (
	// Unspecified. This value must not be used.
	StructuredQuery_FieldFilter_OPERATOR_UNSPECIFIED StructuredQuery_FieldFilter_Operator = 0
	// Less than. Requires that the field come first in `order_by`.
	StructuredQuery_FieldFilter_LESS_THAN StructuredQuery_FieldFilter_Operator = 1
	// Less than or equal. Requires that the field come first in `order_by`.
	StructuredQuery_FieldFilter_LESS_THAN_OR_EQUAL StructuredQuery_FieldFilter_Operator = 2
	// Greater than. Requires that the field come first in `order_by`.
	StructuredQuery_FieldFilter_GREATER_THAN StructuredQuery_FieldFilter_Operator = 3
	// Greater than or equal. Requires that the field come first in
	// `order_by`.
	StructuredQuery_FieldFilter_GREATER_THAN_OR_EQUAL StructuredQuery_FieldFilter_Operator = 4
	// Equal.
	StructuredQuery_FieldFilter_EQUAL StructuredQuery_FieldFilter_Operator = 5
)

var StructuredQuery_FieldFilter_Operator_name = map[int32]string{
	0: "OPERATOR_UNSPECIFIED",
	1: "LESS_THAN",
	2: "LESS_THAN_OR_EQUAL",
	3: "GREATER_THAN",
	4: "GREATER_THAN_OR_EQUAL",
	5: "EQUAL",
}
var StructuredQuery_FieldFilter_Operator_value = map[string]int32{
	"OPERATOR_UNSPECIFIED":  0,
	"LESS_THAN":             1,
	"LESS_THAN_OR_EQUAL":    2,
	"GREATER_THAN":          3,
	"GREATER_THAN_OR_EQUAL": 4,
	"EQUAL":                 5,
}

func (x StructuredQuery_FieldFilter_Operator) String() string {
	return proto.EnumName(StructuredQuery_FieldFilter_Operator_name, int32(x))
}
func (StructuredQuery_FieldFilter_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 3, 0}
}

// A unary operator.
type StructuredQuery_UnaryFilter_Operator int32

const (
	// Unspecified. This value must not be used.
	StructuredQuery_UnaryFilter_OPERATOR_UNSPECIFIED StructuredQuery_UnaryFilter_Operator = 0
	// Test if a field is equal to NaN.
	StructuredQuery_UnaryFilter_IS_NAN StructuredQuery_UnaryFilter_Operator = 2
	// Test if an exprestion evaluates to Null.
	StructuredQuery_UnaryFilter_IS_NULL StructuredQuery_UnaryFilter_Operator = 3
)

var StructuredQuery_UnaryFilter_Operator_name = map[int32]string{
	0: "OPERATOR_UNSPECIFIED",
	2: "IS_NAN",
	3: "IS_NULL",
}
var StructuredQuery_UnaryFilter_Operator_value = map[string]int32{
	"OPERATOR_UNSPECIFIED": 0,
	"IS_NAN":               2,
	"IS_NULL":              3,
}

func (x StructuredQuery_UnaryFilter_Operator) String() string {
	return proto.EnumName(StructuredQuery_UnaryFilter_Operator_name, int32(x))
}
func (StructuredQuery_UnaryFilter_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 4, 0}
}

// A Firestore query.
type StructuredQuery struct {
	// The projection to return.
	Select *StructuredQuery_Projection `protobuf:"bytes,1,opt,name=select" json:"select,omitempty"`
	// The collections to query.
	From []*StructuredQuery_CollectionSelector `protobuf:"bytes,2,rep,name=from" json:"from,omitempty"`
	// The filter to apply.
	Where *StructuredQuery_Filter `protobuf:"bytes,3,opt,name=where" json:"where,omitempty"`
	// The order to apply to the query results.
	//
	// Firestore guarantees a stable ordering through the following rules:
	//
	//  * Any field required to appear in `order_by`, that is not already
	//    specified in `order_by`, is appended to the order in field name order
	//    by default.
	//  * If an order on `__name__` is not specified, it is appended by default.
	//
	// Fields are appended with the same sort direction as the last order
	// specified, or 'ASCENDING' if no order was specified. For example:
	//
	//  * `SELECT * FROM Foo ORDER BY A` becomes
	//    `SELECT * FROM Foo ORDER BY A, __name__`
	//  * `SELECT * FROM Foo ORDER BY A DESC` becomes
	//    `SELECT * FROM Foo ORDER BY A DESC, __name__ DESC`
	//  * `SELECT * FROM Foo WHERE A > 1` becomes
	//    `SELECT * FROM Foo WHERE A > 1 ORDER BY A, __name__`
	OrderBy []*StructuredQuery_Order `protobuf:"bytes,4,rep,name=order_by,json=orderBy" json:"order_by,omitempty"`
	// A starting point for the query results.
	StartAt *Cursor `protobuf:"bytes,7,opt,name=start_at,json=startAt" json:"start_at,omitempty"`
	// A end point for the query results.
	EndAt *Cursor `protobuf:"bytes,8,opt,name=end_at,json=endAt" json:"end_at,omitempty"`
	// The number of results to skip.
	//
	// Applies before limit, but after all other constraints. Must be >= 0 if
	// specified.
	Offset int32 `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
	// The maximum number of results to return.
	//
	// Applies after all other constraints.
	// Must be >= 0 if specified.
	Limit *google_protobuf3.Int32Value `protobuf:"bytes,5,opt,name=limit" json:"limit,omitempty"`
}

func (m *StructuredQuery) Reset()                    { *m = StructuredQuery{} }
func (m *StructuredQuery) String() string            { return proto.CompactTextString(m) }
func (*StructuredQuery) ProtoMessage()               {}
func (*StructuredQuery) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *StructuredQuery) GetSelect() *StructuredQuery_Projection {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *StructuredQuery) GetFrom() []*StructuredQuery_CollectionSelector {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *StructuredQuery) GetWhere() *StructuredQuery_Filter {
	if m != nil {
		return m.Where
	}
	return nil
}

func (m *StructuredQuery) GetOrderBy() []*StructuredQuery_Order {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

func (m *StructuredQuery) GetStartAt() *Cursor {
	if m != nil {
		return m.StartAt
	}
	return nil
}

func (m *StructuredQuery) GetEndAt() *Cursor {
	if m != nil {
		return m.EndAt
	}
	return nil
}

func (m *StructuredQuery) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *StructuredQuery) GetLimit() *google_protobuf3.Int32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

// A selection of a collection, such as `messages as m1`.
type StructuredQuery_CollectionSelector struct {
	// The collection ID.
	// When set, selects only collections with this ID.
	CollectionId string `protobuf:"bytes,2,opt,name=collection_id,json=collectionId" json:"collection_id,omitempty"`
	// When false, selects only collections that are immediate children of
	// the `parent` specified in the containing `RunQueryRequest`.
	// When true, selects all descendant collections.
	AllDescendants bool `protobuf:"varint,3,opt,name=all_descendants,json=allDescendants" json:"all_descendants,omitempty"`
}

func (m *StructuredQuery_CollectionSelector) Reset()         { *m = StructuredQuery_CollectionSelector{} }
func (m *StructuredQuery_CollectionSelector) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery_CollectionSelector) ProtoMessage()    {}
func (*StructuredQuery_CollectionSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0}
}

func (m *StructuredQuery_CollectionSelector) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *StructuredQuery_CollectionSelector) GetAllDescendants() bool {
	if m != nil {
		return m.AllDescendants
	}
	return false
}

// A filter.
type StructuredQuery_Filter struct {
	// The type of filter.
	//
	// Types that are valid to be assigned to FilterType:
	//	*StructuredQuery_Filter_CompositeFilter
	//	*StructuredQuery_Filter_FieldFilter
	//	*StructuredQuery_Filter_UnaryFilter
	FilterType isStructuredQuery_Filter_FilterType `protobuf_oneof:"filter_type"`
}

func (m *StructuredQuery_Filter) Reset()                    { *m = StructuredQuery_Filter{} }
func (m *StructuredQuery_Filter) String() string            { return proto.CompactTextString(m) }
func (*StructuredQuery_Filter) ProtoMessage()               {}
func (*StructuredQuery_Filter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

type isStructuredQuery_Filter_FilterType interface {
	isStructuredQuery_Filter_FilterType()
}

type StructuredQuery_Filter_CompositeFilter struct {
	CompositeFilter *StructuredQuery_CompositeFilter `protobuf:"bytes,1,opt,name=composite_filter,json=compositeFilter,oneof"`
}
type StructuredQuery_Filter_FieldFilter struct {
	FieldFilter *StructuredQuery_FieldFilter `protobuf:"bytes,2,opt,name=field_filter,json=fieldFilter,oneof"`
}
type StructuredQuery_Filter_UnaryFilter struct {
	UnaryFilter *StructuredQuery_UnaryFilter `protobuf:"bytes,3,opt,name=unary_filter,json=unaryFilter,oneof"`
}

func (*StructuredQuery_Filter_CompositeFilter) isStructuredQuery_Filter_FilterType() {}
func (*StructuredQuery_Filter_FieldFilter) isStructuredQuery_Filter_FilterType()     {}
func (*StructuredQuery_Filter_UnaryFilter) isStructuredQuery_Filter_FilterType()     {}

func (m *StructuredQuery_Filter) GetFilterType() isStructuredQuery_Filter_FilterType {
	if m != nil {
		return m.FilterType
	}
	return nil
}

func (m *StructuredQuery_Filter) GetCompositeFilter() *StructuredQuery_CompositeFilter {
	if x, ok := m.GetFilterType().(*StructuredQuery_Filter_CompositeFilter); ok {
		return x.CompositeFilter
	}
	return nil
}

func (m *StructuredQuery_Filter) GetFieldFilter() *StructuredQuery_FieldFilter {
	if x, ok := m.GetFilterType().(*StructuredQuery_Filter_FieldFilter); ok {
		return x.FieldFilter
	}
	return nil
}

func (m *StructuredQuery_Filter) GetUnaryFilter() *StructuredQuery_UnaryFilter {
	if x, ok := m.GetFilterType().(*StructuredQuery_Filter_UnaryFilter); ok {
		return x.UnaryFilter
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StructuredQuery_Filter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StructuredQuery_Filter_OneofMarshaler, _StructuredQuery_Filter_OneofUnmarshaler, _StructuredQuery_Filter_OneofSizer, []interface{}{
		(*StructuredQuery_Filter_CompositeFilter)(nil),
		(*StructuredQuery_Filter_FieldFilter)(nil),
		(*StructuredQuery_Filter_UnaryFilter)(nil),
	}
}

func _StructuredQuery_Filter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StructuredQuery_Filter)
	// filter_type
	switch x := m.FilterType.(type) {
	case *StructuredQuery_Filter_CompositeFilter:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CompositeFilter); err != nil {
			return err
		}
	case *StructuredQuery_Filter_FieldFilter:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FieldFilter); err != nil {
			return err
		}
	case *StructuredQuery_Filter_UnaryFilter:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UnaryFilter); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StructuredQuery_Filter.FilterType has unexpected type %T", x)
	}
	return nil
}

func _StructuredQuery_Filter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StructuredQuery_Filter)
	switch tag {
	case 1: // filter_type.composite_filter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StructuredQuery_CompositeFilter)
		err := b.DecodeMessage(msg)
		m.FilterType = &StructuredQuery_Filter_CompositeFilter{msg}
		return true, err
	case 2: // filter_type.field_filter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StructuredQuery_FieldFilter)
		err := b.DecodeMessage(msg)
		m.FilterType = &StructuredQuery_Filter_FieldFilter{msg}
		return true, err
	case 3: // filter_type.unary_filter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StructuredQuery_UnaryFilter)
		err := b.DecodeMessage(msg)
		m.FilterType = &StructuredQuery_Filter_UnaryFilter{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StructuredQuery_Filter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StructuredQuery_Filter)
	// filter_type
	switch x := m.FilterType.(type) {
	case *StructuredQuery_Filter_CompositeFilter:
		s := proto.Size(x.CompositeFilter)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StructuredQuery_Filter_FieldFilter:
		s := proto.Size(x.FieldFilter)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StructuredQuery_Filter_UnaryFilter:
		s := proto.Size(x.UnaryFilter)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A filter that merges multiple other filters using the given operator.
type StructuredQuery_CompositeFilter struct {
	// The operator for combining multiple filters.
	Op StructuredQuery_CompositeFilter_Operator `protobuf:"varint,1,opt,name=op,enum=google.firestore.v1beta1.StructuredQuery_CompositeFilter_Operator" json:"op,omitempty"`
	// The list of filters to combine.
	// Must contain at least one filter.
	Filters []*StructuredQuery_Filter `protobuf:"bytes,2,rep,name=filters" json:"filters,omitempty"`
}

func (m *StructuredQuery_CompositeFilter) Reset()         { *m = StructuredQuery_CompositeFilter{} }
func (m *StructuredQuery_CompositeFilter) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery_CompositeFilter) ProtoMessage()    {}
func (*StructuredQuery_CompositeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 2}
}

func (m *StructuredQuery_CompositeFilter) GetOp() StructuredQuery_CompositeFilter_Operator {
	if m != nil {
		return m.Op
	}
	return StructuredQuery_CompositeFilter_OPERATOR_UNSPECIFIED
}

func (m *StructuredQuery_CompositeFilter) GetFilters() []*StructuredQuery_Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

// A filter on a specific field.
type StructuredQuery_FieldFilter struct {
	// The field to filter by.
	Field *StructuredQuery_FieldReference `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	// The operator to filter by.
	Op StructuredQuery_FieldFilter_Operator `protobuf:"varint,2,opt,name=op,enum=google.firestore.v1beta1.StructuredQuery_FieldFilter_Operator" json:"op,omitempty"`
	// The value to compare to.
	Value *Value `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *StructuredQuery_FieldFilter) Reset()                    { *m = StructuredQuery_FieldFilter{} }
func (m *StructuredQuery_FieldFilter) String() string            { return proto.CompactTextString(m) }
func (*StructuredQuery_FieldFilter) ProtoMessage()               {}
func (*StructuredQuery_FieldFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 3} }

func (m *StructuredQuery_FieldFilter) GetField() *StructuredQuery_FieldReference {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *StructuredQuery_FieldFilter) GetOp() StructuredQuery_FieldFilter_Operator {
	if m != nil {
		return m.Op
	}
	return StructuredQuery_FieldFilter_OPERATOR_UNSPECIFIED
}

func (m *StructuredQuery_FieldFilter) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

// A filter with a single operand.
type StructuredQuery_UnaryFilter struct {
	// The unary operator to apply.
	Op StructuredQuery_UnaryFilter_Operator `protobuf:"varint,1,opt,name=op,enum=google.firestore.v1beta1.StructuredQuery_UnaryFilter_Operator" json:"op,omitempty"`
	// The argument to the filter.
	//
	// Types that are valid to be assigned to OperandType:
	//	*StructuredQuery_UnaryFilter_Field
	OperandType isStructuredQuery_UnaryFilter_OperandType `protobuf_oneof:"operand_type"`
}

func (m *StructuredQuery_UnaryFilter) Reset()                    { *m = StructuredQuery_UnaryFilter{} }
func (m *StructuredQuery_UnaryFilter) String() string            { return proto.CompactTextString(m) }
func (*StructuredQuery_UnaryFilter) ProtoMessage()               {}
func (*StructuredQuery_UnaryFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 4} }

type isStructuredQuery_UnaryFilter_OperandType interface {
	isStructuredQuery_UnaryFilter_OperandType()
}

type StructuredQuery_UnaryFilter_Field struct {
	Field *StructuredQuery_FieldReference `protobuf:"bytes,2,opt,name=field,oneof"`
}

func (*StructuredQuery_UnaryFilter_Field) isStructuredQuery_UnaryFilter_OperandType() {}

func (m *StructuredQuery_UnaryFilter) GetOperandType() isStructuredQuery_UnaryFilter_OperandType {
	if m != nil {
		return m.OperandType
	}
	return nil
}

func (m *StructuredQuery_UnaryFilter) GetOp() StructuredQuery_UnaryFilter_Operator {
	if m != nil {
		return m.Op
	}
	return StructuredQuery_UnaryFilter_OPERATOR_UNSPECIFIED
}

func (m *StructuredQuery_UnaryFilter) GetField() *StructuredQuery_FieldReference {
	if x, ok := m.GetOperandType().(*StructuredQuery_UnaryFilter_Field); ok {
		return x.Field
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StructuredQuery_UnaryFilter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StructuredQuery_UnaryFilter_OneofMarshaler, _StructuredQuery_UnaryFilter_OneofUnmarshaler, _StructuredQuery_UnaryFilter_OneofSizer, []interface{}{
		(*StructuredQuery_UnaryFilter_Field)(nil),
	}
}

func _StructuredQuery_UnaryFilter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StructuredQuery_UnaryFilter)
	// operand_type
	switch x := m.OperandType.(type) {
	case *StructuredQuery_UnaryFilter_Field:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Field); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StructuredQuery_UnaryFilter.OperandType has unexpected type %T", x)
	}
	return nil
}

func _StructuredQuery_UnaryFilter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StructuredQuery_UnaryFilter)
	switch tag {
	case 2: // operand_type.field
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StructuredQuery_FieldReference)
		err := b.DecodeMessage(msg)
		m.OperandType = &StructuredQuery_UnaryFilter_Field{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StructuredQuery_UnaryFilter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StructuredQuery_UnaryFilter)
	// operand_type
	switch x := m.OperandType.(type) {
	case *StructuredQuery_UnaryFilter_Field:
		s := proto.Size(x.Field)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// An order on a field.
type StructuredQuery_Order struct {
	// The field to order by.
	Field *StructuredQuery_FieldReference `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	// The direction to order by. Defaults to `ASCENDING`.
	Direction StructuredQuery_Direction `protobuf:"varint,2,opt,name=direction,enum=google.firestore.v1beta1.StructuredQuery_Direction" json:"direction,omitempty"`
}

func (m *StructuredQuery_Order) Reset()                    { *m = StructuredQuery_Order{} }
func (m *StructuredQuery_Order) String() string            { return proto.CompactTextString(m) }
func (*StructuredQuery_Order) ProtoMessage()               {}
func (*StructuredQuery_Order) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 5} }

func (m *StructuredQuery_Order) GetField() *StructuredQuery_FieldReference {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *StructuredQuery_Order) GetDirection() StructuredQuery_Direction {
	if m != nil {
		return m.Direction
	}
	return StructuredQuery_DIRECTION_UNSPECIFIED
}

// A reference to a field, such as `max(messages.time) as max_time`.
type StructuredQuery_FieldReference struct {
	FieldPath string `protobuf:"bytes,2,opt,name=field_path,json=fieldPath" json:"field_path,omitempty"`
}

func (m *StructuredQuery_FieldReference) Reset()         { *m = StructuredQuery_FieldReference{} }
func (m *StructuredQuery_FieldReference) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery_FieldReference) ProtoMessage()    {}
func (*StructuredQuery_FieldReference) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 6}
}

func (m *StructuredQuery_FieldReference) GetFieldPath() string {
	if m != nil {
		return m.FieldPath
	}
	return ""
}

// The projection of document's fields to return.
type StructuredQuery_Projection struct {
	// The fields to return.
	//
	// If empty, all fields are returned. To only return the name
	// of the document, use `['__name__']`.
	Fields []*StructuredQuery_FieldReference `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
}

func (m *StructuredQuery_Projection) Reset()                    { *m = StructuredQuery_Projection{} }
func (m *StructuredQuery_Projection) String() string            { return proto.CompactTextString(m) }
func (*StructuredQuery_Projection) ProtoMessage()               {}
func (*StructuredQuery_Projection) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 7} }

func (m *StructuredQuery_Projection) GetFields() []*StructuredQuery_FieldReference {
	if m != nil {
		return m.Fields
	}
	return nil
}

// A position in a query result set.
type Cursor struct {
	// The values that represent a position, in the order they appear in
	// the order by clause of a query.
	//
	// Can contain fewer values than specified in the order by clause.
	Values []*Value `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// If the position is just before or just after the given values, relative
	// to the sort order defined by the query.
	Before bool `protobuf:"varint,2,opt,name=before" json:"before,omitempty"`
}

func (m *Cursor) Reset()                    { *m = Cursor{} }
func (m *Cursor) String() string            { return proto.CompactTextString(m) }
func (*Cursor) ProtoMessage()               {}
func (*Cursor) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Cursor) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Cursor) GetBefore() bool {
	if m != nil {
		return m.Before
	}
	return false
}

func init() {
	proto.RegisterType((*StructuredQuery)(nil), "google.firestore.v1beta1.StructuredQuery")
	proto.RegisterType((*StructuredQuery_CollectionSelector)(nil), "google.firestore.v1beta1.StructuredQuery.CollectionSelector")
	proto.RegisterType((*StructuredQuery_Filter)(nil), "google.firestore.v1beta1.StructuredQuery.Filter")
	proto.RegisterType((*StructuredQuery_CompositeFilter)(nil), "google.firestore.v1beta1.StructuredQuery.CompositeFilter")
	proto.RegisterType((*StructuredQuery_FieldFilter)(nil), "google.firestore.v1beta1.StructuredQuery.FieldFilter")
	proto.RegisterType((*StructuredQuery_UnaryFilter)(nil), "google.firestore.v1beta1.StructuredQuery.UnaryFilter")
	proto.RegisterType((*StructuredQuery_Order)(nil), "google.firestore.v1beta1.StructuredQuery.Order")
	proto.RegisterType((*StructuredQuery_FieldReference)(nil), "google.firestore.v1beta1.StructuredQuery.FieldReference")
	proto.RegisterType((*StructuredQuery_Projection)(nil), "google.firestore.v1beta1.StructuredQuery.Projection")
	proto.RegisterType((*Cursor)(nil), "google.firestore.v1beta1.Cursor")
	proto.RegisterEnum("google.firestore.v1beta1.StructuredQuery_Direction", StructuredQuery_Direction_name, StructuredQuery_Direction_value)
	proto.RegisterEnum("google.firestore.v1beta1.StructuredQuery_CompositeFilter_Operator", StructuredQuery_CompositeFilter_Operator_name, StructuredQuery_CompositeFilter_Operator_value)
	proto.RegisterEnum("google.firestore.v1beta1.StructuredQuery_FieldFilter_Operator", StructuredQuery_FieldFilter_Operator_name, StructuredQuery_FieldFilter_Operator_value)
	proto.RegisterEnum("google.firestore.v1beta1.StructuredQuery_UnaryFilter_Operator", StructuredQuery_UnaryFilter_Operator_name, StructuredQuery_UnaryFilter_Operator_value)
}

func init() { proto.RegisterFile("google/firestore/v1beta1/query.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 970 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0xd7, 0x4e, 0xf3, 0x75, 0xd2, 0x0f, 0x6b, 0x04, 0x2b, 0x13, 0x96, 0xa5, 0x0a, 0x48,
	0xdb, 0x1b, 0x1c, 0xda, 0xb2, 0x02, 0xb4, 0x80, 0xe4, 0x24, 0x6e, 0x9b, 0x55, 0xe5, 0xa4, 0x93,
	0xb6, 0x12, 0xab, 0x0a, 0xcb, 0xb1, 0xc7, 0xa9, 0x91, 0xeb, 0x31, 0xe3, 0xf1, 0xae, 0x7a, 0xcd,
	0x9b, 0x70, 0xb9, 0x2f, 0x00, 0xcf, 0xc0, 0x53, 0x70, 0xcd, 0x23, 0x70, 0x81, 0x90, 0xc7, 0xe3,
	0xa4, 0xdd, 0xaa, 0x22, 0x29, 0xdc, 0xe5, 0x1c, 0x9f, 0xf3, 0x9b, 0xe3, 0xff, 0x39, 0xc7, 0x13,
	0xf8, 0x74, 0x46, 0xe9, 0x2c, 0x22, 0xdd, 0x20, 0x64, 0x24, 0xe5, 0x94, 0x91, 0xee, 0xeb, 0xdd,
	0x29, 0xe1, 0xee, 0x6e, 0xf7, 0xa7, 0x8c, 0xb0, 0x6b, 0x23, 0x61, 0x94, 0x53, 0xa4, 0x17, 0x51,
	0xc6, 0x3c, 0xca, 0x90, 0x51, 0xed, 0x27, 0x32, 0xdf, 0x4d, 0xc2, 0xae, 0x1b, 0xc7, 0x94, 0xbb,
	0x3c, 0xa4, 0x71, 0x5a, 0xe4, 0xb5, 0x9f, 0xdd, 0x4b, 0xf7, 0xa9, 0x97, 0x5d, 0x91, 0x98, 0xcb,
	0xc0, 0xa7, 0x32, 0x50, 0x58, 0xd3, 0x2c, 0xe8, 0xbe, 0x61, 0x6e, 0x92, 0x10, 0x26, 0x41, 0x9d,
	0xbf, 0x34, 0xd8, 0x9a, 0x70, 0x96, 0x79, 0x3c, 0x63, 0xc4, 0x3f, 0xc9, 0x4b, 0x43, 0xc7, 0x50,
	0x4b, 0x49, 0x44, 0x3c, 0xae, 0x2b, 0xdb, 0xca, 0x4e, 0x6b, 0xef, 0x0b, 0xe3, 0xbe, 0x2a, 0x8d,
	0x77, 0x52, 0x8d, 0x31, 0xa3, 0x3f, 0x12, 0x2f, 0xaf, 0x14, 0x4b, 0x06, 0x1a, 0xc3, 0x5a, 0xc0,
	0xe8, 0x95, 0xae, 0x6e, 0x57, 0x76, 0x5a, 0x7b, 0xdf, 0x2c, 0xcf, 0xea, 0xd3, 0x28, 0x2a, 0x58,
	0x13, 0x41, 0xa2, 0x0c, 0x0b, 0x12, 0x3a, 0x80, 0xea, 0x9b, 0x4b, 0xc2, 0x88, 0x5e, 0x11, 0xe5,
	0x7d, 0xbe, 0x3c, 0xf2, 0x20, 0x8c, 0x38, 0x61, 0xb8, 0x48, 0x47, 0x2f, 0xa1, 0x41, 0x99, 0x4f,
	0x98, 0x33, 0xbd, 0xd6, 0xd7, 0x44, 0x75, 0xdd, 0xe5, 0x51, 0xa3, 0x3c, 0x13, 0xd7, 0x05, 0xa0,
	0x77, 0x8d, 0x5e, 0x40, 0x23, 0xe5, 0x2e, 0xe3, 0x8e, 0xcb, 0xf5, 0xba, 0x28, 0x6b, 0xfb, 0x7e,
	0x56, 0x3f, 0x63, 0x29, 0x65, 0xb8, 0x2e, 0x32, 0x4c, 0x8e, 0xbe, 0x84, 0x1a, 0x89, 0xfd, 0x3c,
	0xb5, 0xb1, 0x64, 0x6a, 0x95, 0xc4, 0xbe, 0xc9, 0xd1, 0x63, 0xa8, 0xd1, 0x20, 0x48, 0x09, 0xd7,
	0x6b, 0xdb, 0xca, 0x4e, 0x15, 0x4b, 0x0b, 0xed, 0x42, 0x35, 0x0a, 0xaf, 0x42, 0xae, 0x57, 0x05,
	0xef, 0xc3, 0x92, 0x57, 0x4e, 0x81, 0x31, 0x8c, 0xf9, 0xfe, 0xde, 0xb9, 0x1b, 0x65, 0x04, 0x17,
	0x91, 0xed, 0x29, 0xa0, 0xbb, 0x82, 0xa3, 0x4f, 0x60, 0xc3, 0x9b, 0x7b, 0x9d, 0xd0, 0xd7, 0xd5,
	0x6d, 0x65, 0xa7, 0x89, 0xd7, 0x17, 0xce, 0xa1, 0x8f, 0x9e, 0xc1, 0x96, 0x1b, 0x45, 0x8e, 0x4f,
	0x52, 0x8f, 0xc4, 0xbe, 0x1b, 0xf3, 0x54, 0x74, 0xa6, 0x81, 0x37, 0xdd, 0x28, 0x1a, 0x2c, 0xbc,
	0xed, 0x5f, 0x55, 0xa8, 0x15, 0x2d, 0x40, 0x01, 0x68, 0x1e, 0xbd, 0x4a, 0x68, 0x1a, 0x72, 0xe2,
	0x04, 0xc2, 0x27, 0xa7, 0xed, 0xeb, 0x55, 0x26, 0x44, 0x12, 0x0a, 0xe8, 0xd1, 0x23, 0xbc, 0xe5,
	0xdd, 0x76, 0xa1, 0x57, 0xb0, 0x1e, 0x84, 0x24, 0xf2, 0xcb, 0x33, 0x54, 0x71, 0xc6, 0xf3, 0x55,
	0x46, 0x86, 0x44, 0xfe, 0x9c, 0xdf, 0x0a, 0x16, 0x66, 0xce, 0xce, 0x62, 0x97, 0x5d, 0x97, 0xec,
	0xca, 0xaa, 0xec, 0xb3, 0x3c, 0x7b, 0xc1, 0xce, 0x16, 0x66, 0x6f, 0x03, 0x5a, 0x05, 0xd5, 0xe1,
	0xd7, 0x09, 0x69, 0xff, 0xa1, 0xc0, 0xd6, 0x3b, 0x6f, 0x8b, 0x30, 0xa8, 0x34, 0x11, 0xa2, 0x6d,
	0xee, 0xf5, 0x1e, 0x2c, 0x9a, 0x31, 0x4a, 0x08, 0x73, 0xf3, 0xe5, 0x52, 0x69, 0x82, 0x5e, 0x42,
	0xbd, 0x38, 0x36, 0x95, 0xfb, 0xba, 0xfa, 0x72, 0x95, 0x80, 0xce, 0x67, 0xd0, 0x28, 0xd9, 0x48,
	0x87, 0xf7, 0x46, 0x63, 0x0b, 0x9b, 0xa7, 0x23, 0xec, 0x9c, 0xd9, 0x93, 0xb1, 0xd5, 0x1f, 0x1e,
	0x0c, 0xad, 0x81, 0xf6, 0x08, 0xd5, 0xa1, 0x62, 0xda, 0x03, 0x4d, 0x69, 0xff, 0xa9, 0x42, 0xeb,
	0x86, 0xd8, 0xc8, 0x86, 0xaa, 0x10, 0x5b, 0x8e, 0xc5, 0x57, 0x2b, 0xb6, 0x0c, 0x93, 0x80, 0x30,
	0x12, 0x7b, 0x04, 0x17, 0x18, 0x64, 0x0b, 0xb9, 0x54, 0x21, 0xd7, 0x77, 0x0f, 0xea, 0xff, 0x6d,
	0xa9, 0x9e, 0x43, 0xf5, 0x75, 0xbe, 0x40, 0xb2, 0xed, 0x1f, 0xdf, 0x8f, 0x94, 0x7b, 0x26, 0xa2,
	0x3b, 0x3f, 0x2b, 0x4b, 0xc9, 0xb2, 0x01, 0xcd, 0x63, 0x6b, 0x32, 0x71, 0x4e, 0x8f, 0x4c, 0x5b,
	0x53, 0xd0, 0x63, 0x40, 0x73, 0xd3, 0x19, 0x61, 0xc7, 0x3a, 0x39, 0x33, 0x8f, 0x35, 0x15, 0x69,
	0xb0, 0x7e, 0x88, 0x2d, 0xf3, 0xd4, 0xc2, 0x45, 0x64, 0x05, 0x7d, 0x00, 0xef, 0xdf, 0xf4, 0x2c,
	0x82, 0xd7, 0x50, 0x13, 0xaa, 0xc5, 0xcf, 0x6a, 0xfb, 0x6f, 0x05, 0x5a, 0x37, 0xa6, 0x4f, 0x8a,
	0xa3, 0xac, 0x2a, 0xce, 0x0d, 0xc4, 0x6d, 0x71, 0xc6, 0x65, 0xf3, 0xd4, 0xff, 0xd6, 0xbc, 0xa3,
	0x47, 0xb2, 0x7d, 0x9d, 0x6f, 0x97, 0x92, 0x0d, 0xa0, 0x36, 0x9c, 0x38, 0xb6, 0x69, 0x6b, 0x2a,
	0x6a, 0x41, 0x3d, 0xff, 0x7d, 0x76, 0x7c, 0xac, 0x55, 0x7a, 0x9b, 0xb0, 0x4e, 0xf3, 0xf4, 0xd8,
	0x2f, 0x16, 0xea, 0xad, 0x02, 0x55, 0xf1, 0x09, 0xff, 0xdf, 0xe7, 0xec, 0x04, 0x9a, 0x7e, 0xc8,
	0x8a, 0x8f, 0xa3, 0x1c, 0xb7, 0xfd, 0xe5, 0x99, 0x83, 0x32, 0x15, 0x2f, 0x28, 0xed, 0x2e, 0x6c,
	0xde, 0x3e, 0x0b, 0x7d, 0x04, 0x50, 0x7c, 0xd6, 0x12, 0x97, 0x5f, 0xca, 0x8f, 0x72, 0x53, 0x78,
	0xc6, 0x2e, 0xbf, 0x6c, 0xff, 0x00, 0xb0, 0xb8, 0x89, 0xd1, 0x18, 0x6a, 0xe2, 0x51, 0xb9, 0xd3,
	0x0f, 0x7f, 0x45, 0xc9, 0xe9, 0x58, 0xd0, 0x9c, 0x17, 0x9a, 0x4f, 0xdc, 0x60, 0x88, 0xad, 0xfe,
	0xe9, 0x70, 0x64, 0xdf, 0x9d, 0x62, 0x73, 0xd2, 0xb7, 0xec, 0xc1, 0xd0, 0x3e, 0xd4, 0x14, 0xb4,
	0x09, 0x30, 0xb0, 0xe6, 0xb6, 0xda, 0xf9, 0x1e, 0x6a, 0xc5, 0x7d, 0x96, 0xdf, 0x80, 0x62, 0x3d,
	0x52, 0x5d, 0x11, 0x25, 0xfe, 0xeb, 0x36, 0xc9, 0xf0, 0xfc, 0x06, 0x9c, 0x92, 0x80, 0x32, 0x22,
	0x44, 0x68, 0x60, 0x69, 0xf5, 0x7e, 0x53, 0xe0, 0x89, 0x47, 0xaf, 0xee, 0xc5, 0xf4, 0x40, 0xbc,
	0xe0, 0x38, 0xbf, 0x10, 0xc7, 0xca, 0x2b, 0x53, 0xc6, 0xcd, 0x68, 0xe4, 0xc6, 0x33, 0x83, 0xb2,
	0x59, 0x77, 0x46, 0x62, 0x71, 0x5d, 0x76, 0x8b, 0x47, 0x6e, 0x12, 0xa6, 0x77, 0xff, 0x6e, 0xbd,
	0x98, 0x7b, 0x7e, 0x51, 0xd7, 0x0e, 0xfb, 0x07, 0x93, 0xb7, 0xea, 0xd3, 0xc3, 0x02, 0xd5, 0x8f,
	0x68, 0xe6, 0x1b, 0x07, 0xf3, 0x83, 0xcf, 0x77, 0x7b, 0x79, 0xc6, 0xef, 0x65, 0xc0, 0x85, 0x08,
	0xb8, 0x98, 0x07, 0x5c, 0x9c, 0x17, 0xc8, 0x69, 0x4d, 0x1c, 0xbb, 0xff, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4c, 0xc9, 0x73, 0x9b, 0x42, 0x0a, 0x00, 0x00,
}
