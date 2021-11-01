package ctystructpb

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"google.golang.org/protobuf/types/known/structpb"
)

// ImpliedType returns the cty Type implied by the structure of the given
// structpb value. This function implements the default type mapping behavior
// used when decoding arbitrary values without explicit cty Type information.
//
// The mapping rules are the same as for cty's JSON ImpliedType function.
func ImpliedType(sv *structpb.Value) (cty.Type, error) {
	src, err := sv.MarshalJSON()
	if err != nil {
		return cty.NilType, fmt.Errorf("internal JSON mapping failed: %w", err)
	}

	return ctyjson.ImpliedType(src)
}
