package ctystructpb

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/zclconf/go-cty/cty"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/structpb"
)

var protoCmpOpt = protocmp.Transform()

func TestToStructValue(t *testing.T) {
	tests := []struct {
		v       cty.Value
		ty      cty.Type
		want    *structpb.Value
		wantErr string
	}{
		{
			cty.StringVal("hello"),
			cty.String,
			structpb.NewStringValue("hello"),
			``,
		},
		{
			cty.NumberIntVal(2),
			cty.Number,
			structpb.NewNumberValue(2),
			``,
		},
		{
			cty.EmptyTupleVal,
			cty.EmptyTuple,
			structpb.NewListValue(mustNewList(nil)),
			``,
		},
		{
			cty.StringVal("hello"),
			cty.DynamicPseudoType,
			structpb.NewStructValue(mustNewStruct(map[string]interface{}{
				"type":  "string",
				"value": "hello",
			})),
			``,
		},
	}

	for _, test := range tests {
		t.Run(test.v.GoString(), func(t *testing.T) {
			got, err := ToStructValue(test.v, test.ty)

			if test.wantErr != "" {
				if err == nil {
					t.Fatalf("unexpected success\nwant error: %s", test.wantErr)
				}
				if err.Error() != test.wantErr {
					t.Fatalf("wrong error\ngot error:  %s\nwant error: %s", err.Error(), test.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error\ngot error: %s", err.Error())
			}

			if diff := cmp.Diff(got, test.want, protoCmpOpt); diff != "" {
				t.Fatalf("wrong result\n%s", diff)
			}
		})
	}
}

func mustNewStruct(v map[string]interface{}) *structpb.Struct {
	ret, err := structpb.NewStruct(v)
	if err != nil {
		panic(err)
	}
	return ret
}

func mustNewList(v []interface{}) *structpb.ListValue {
	ret, err := structpb.NewList(v)
	if err != nil {
		panic(err)
	}
	return ret
}
