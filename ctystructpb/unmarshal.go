package ctystructpb

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"google.golang.org/protobuf/types/known/structpb"
)

// FromStructValue decodes a JSON representation of the given value into a
// cty Value conforming to the given type.
//
// While decoding, type conversions will be made where possible to make the
// result conformant even if the types given in JSON are not exactly correct.
// If conversion isn't possible then an error is returned, which may be a cty.PathError.
func FromStructValue(sv *structpb.Value, ty cty.Type) (cty.Value, error) {
	src, err := sv.MarshalJSON()
	if err != nil {
		return cty.NilVal, fmt.Errorf("internal mapping to JSON failed: %w", err)
	}

	return ctyjson.Unmarshal(src, ty)
}
