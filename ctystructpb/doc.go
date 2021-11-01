// Package ctystructpb supports bidirectional conversions between cty values
// and google.protobuf.Struct protobuf messages.
//
// Because google.protobuf.Struct is a mapping of the JSON infoset into the
// protobuf infoset, ctystructpb always makes the same decisions as cty's
// standard JSON serialization, but produces a protobuf-compatible result
// instead of a raw byte buffer.
//
// Note that cty's JSON support uses arbitrary-precision number values, but
// google.protobuf.Struct always uses float64 values and so numbers outside
// of that range will silently lose precision on encoding.
package ctystructpb
