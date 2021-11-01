package ctystructpb

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"google.golang.org/protobuf/types/known/structpb"
)

// ToStructValue produces a struct value representation of the given value that
// can later be decoded into a value of the given type.
//
// A type is specified separately to allow for the given type to include
// cty.DynamicPseudoType to represent situations where any type is permitted
// and so type information must be included to allow recovery of the stored
// structure when decoding.
//
// The given type will also be used to attempt automatic conversions of any
// non-conformant types in the given value, although this will not always be
// possible. If the value cannot be made to be conformant then an error is
// returned, which may be a cty.PathError.
//
// Capsule-typed values can be encoded via their own MarshalJSON
// implementations, but with some caveats. Since capsule values are compared by
// pointer equality, it is impossible to recover a value that will compare
// equal to the original value. Additionally, it's not possible to
// JSON-serialize the capsule type itself, so it's not valid to use capsule
// types within parts of the value that are conformed to cty.DynamicPseudoType.
func ToStructValue(v cty.Value, ty cty.Type) (*structpb.Value, error) {
	src, err := ctyjson.Marshal(v, ty)
	if err != nil {
		return nil, fmt.Errorf("internal mapping to JSON failed: %w", err)
	}

	ret := &structpb.Value{}
	err = ret.UnmarshalJSON(src)
	return ret, err
}
