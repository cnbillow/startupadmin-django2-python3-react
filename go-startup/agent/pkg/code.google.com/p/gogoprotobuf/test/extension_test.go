// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://code.google.com/p/gogoprotobuf/gogoproto
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package test

import (
	"code.google.com/p/gogoprotobuf/proto"
	"math"
	math_rand "math/rand"
	"testing"
	"time"
)

//func SetRawExtension(base extendableProto, id int32, b []byte) {
//func HasExtension(pb extendableProto, extension *ExtensionDesc) bool {
//func ClearExtension(pb extendableProto, extension *ExtensionDesc) {
//func GetExtension(pb extendableProto, extension *ExtensionDesc) (interface{}, error) {
//func GetExtensions(pb Message, es []*ExtensionDesc) (extensions []interface{}, err error) {
//func SetExtension(pb extendableProto, extension *ExtensionDesc, value interface{}) error {

type extendable interface {
	proto.Message
	ExtensionRangeArray() []proto.ExtensionRange
}

func check(t *testing.T, m extendable, fieldA float64, ext *proto.ExtensionDesc) {
	if !proto.HasExtension(m, ext) {
		t.Fatalf("expected extension to be set")
	}
	fieldA2Interface, err := proto.GetExtension(m, ext)
	if err != nil {
		panic(err)
	}
	fieldA2 := fieldA2Interface.(*float64)
	if fieldA != *fieldA2 {
		t.Fatalf("Expected %f got %f", fieldA, *fieldA2)
	}
	proto.ClearExtension(m, ext)
	if proto.HasExtension(m, ext) {
		t.Fatalf("expected extension to be cleared")
	}
}

var fieldA float64
var fieldABytes []byte
var extr = math_rand.New(math_rand.NewSource(time.Now().UnixNano()))

func init() {
	fieldA = float64(1.1)
	fieldABits := math.Float64bits(fieldA)
	x := uint64(uint32(100)<<3 | uint32(proto.WireFixed64))
	fieldABytes = encodeVarintPopulateThetest(nil, x)
	fieldABytes = append(fieldABytes, uint8(fieldABits))
	fieldABytes = append(fieldABytes, uint8(fieldABits>>8))
	fieldABytes = append(fieldABytes, uint8(fieldABits>>16))
	fieldABytes = append(fieldABytes, uint8(fieldABits>>24))
	fieldABytes = append(fieldABytes, uint8(fieldABits>>32))
	fieldABytes = append(fieldABytes, uint8(fieldABits>>40))
	fieldABytes = append(fieldABytes, uint8(fieldABits>>48))
	fieldABytes = append(fieldABytes, uint8(fieldABits>>56))
}

func TestExtensionsMyExtendable(t *testing.T) {
	m := NewPopulatedMyExtendable(extr, false)
	err := proto.SetExtension(m, E_FieldA, &fieldA)
	if err != nil {
		panic(err)
	}
	check(t, m, fieldA, E_FieldA)
	proto.SetRawExtension(m, 100, fieldABytes)
	check(t, m, fieldA, E_FieldA)
}

func TestExtensionsNoExtensionsMapSetExtension(t *testing.T) {
	m := NewPopulatedNoExtensionsMap(extr, false)
	err := proto.SetExtension(m, E_FieldA1, &fieldA)
	if err != nil {
		panic(err)
	}
	check(t, m, fieldA, E_FieldA1)
}

func TestExtensionsNoExtensionsMapSetRawExtension(t *testing.T) {
	m := NewPopulatedNoExtensionsMap(extr, false)
	proto.SetRawExtension(m, 100, fieldABytes)
	check(t, m, fieldA, E_FieldA1)
}
