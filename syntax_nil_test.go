// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode_test

// Import Go standard library package testing as well as lpcode and tserr
import (
	"testing" // testing

	"github.com/thorstenrie/lpcode" // lpcode
	"github.com/thorstenrie/tserr"  // tserr
)

// TestTestVariablesNil tests Testvariables to return nil in case
// *Code is nil. The test fails if Testvariables does not return nil.
func TestTestVariablesNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Testvariables does not return nil.
	if n := c.Testvariables(&lpcode.Testvars{}); n != nil {
		t.Error(tserr.NotNil("Testvariables"))
	}
}

// TestShortVarDeclNil tests ShortVarDecl to return nil in case
// *Code is nil. The test fails if ShortVarDecl does not return nil.
func TestShortVarDeclNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if ShortVarDecl does not return nil.
	if n := c.ShortVarDecl(&lpcode.ShortVarDeclArgs{}); n != nil {
		t.Error(tserr.NotNil("ShortVarDecl"))
	}
}

// TestIdentNil tests Ident to return nil in case
// *Code is nil. The test fails if Ident does not return nil.
func TestIdentNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Ident does not return nil.
	if n := c.Ident(""); n != nil {
		t.Error(tserr.NotNil("Ident"))
	}
}

// TestListNil tests List to return nil in case
// *Code is nil. The test fails if List does not return nil.
func TestListNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if List does not return nil.
	if n := c.List(); n != nil {
		t.Error(tserr.NotNil("List"))
	}
}

// TestListlnNil tests Listln to return nil in case
// *Code is nil. The test fails if Listln does not return nil.
func TestListlnNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Listln does not return nil.
	if n := c.Listln(); n != nil {
		t.Error(tserr.NotNil("Listln"))
	}
}

// TestParamEndlnNil tests ParamEndln to return nil in case
// *Code is nil. The test fails if ParamEndln does not return nil.
func TestParamEndlnNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if ParamEndln does not return nil.
	if n := c.ParamEndln(); n != nil {
		t.Error(tserr.NotNil("ParamEndln"))
	}
}

// TestCallNil tests Call to return nil in case
// *Code is nil. The test fails if Call does not return nil.
func TestCallNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Call does not return nil.
	if n := c.Call(""); n != nil {
		t.Error(tserr.NotNil("Call"))
	}
}

// TestLineCommentNil tests LineComment to return nil in case
// *Code is nil. The test fails if LineComment does not return nil.
func TestLineCommentNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if LineComment does not return nil.
	if n := c.LineComment(""); n != nil {
		t.Error(tserr.NotNil("LineComment"))
	}
}

// TestTypeStructNil tests TypeStruct to return nil in case
// *Code is nil. The test fails if TypeStruct does not return nil.
func TestTypeStructNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if TypeStruct does not return nil.
	if n := c.TypeStruct(""); n != nil {
		t.Error(tserr.NotNil("TypeStruct"))
	}
}

// TestStringNil tests String to return an empty string in case
// *Code is nil. The test fails if String does not return an empty string.
func TestStringNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Ident does not return an empty string.
	if n := c.String(); n != "" {
		t.Error(tserr.NotNil("String"))
	}
}

// TestParamEndNil tests ParamEnd to return nil in case
// *Code is nil. The test fails if ParamEnd does not return nil.
func TestParamEndNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if ParamEnd does not return nil.
	if n := c.ParamEnd(); n != nil {
		t.Error(tserr.NotNil("ParamEnd"))
	}
}

// TestFuncEndNil tests FuncEnd to return nil in case
// *Code is nil. The test fails if FuncEnd does not return nil.
func TestFuncEndNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if FuncEnd does not return nil.
	if n := c.FuncEnd(); n != nil {
		t.Error(tserr.NotNil("FuncEnd"))
	}
}

// TestBlockEndNil tests BlockEnd to return nil in case
// *Code is nil. The test fails if BlockEnd does not return nil.
func TestBlockEndNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if BlockEnd does not return nil.
	if n := c.BlockEnd(); n != nil {
		t.Error(tserr.NotNil("BlockEnd"))
	}
}

// TestKeyedElementNil tests KeyedElement to return nil in case
// *Code is nil. The test fails if KeyedElement does not return nil.
func TestKeyedElementNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if KeyedElement does not return nil.
	if n := c.KeyedElement(&lpcode.KeyedElementArgs{}); n != nil {
		t.Error(tserr.NotNil("KeyedElement"))
	}
}

// TestVarSpecNil tests VarSpec to return nil in case
// *Code is nil. The test fails if VarSpec does not return nil.
func TestVarSpecNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if VarSpec does not return nil.
	if n := c.VarSpec(&lpcode.VarSpecArgs{}); n != nil {
		t.Error(tserr.NotNil("VarSpec"))
	}
}

// TestSelFieldNil tests SelField to return nil in case
// *Code is nil. The test fails if SelField does not return nil.
func TestSelFieldNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if SelField does not return nil.
	if n := c.SelField(&lpcode.SelArgs{}); n != nil {
		t.Error(tserr.NotNil("SelField"))
	}
}

// TestSelMethodNil tests SelMethod to return nil in case
// *Code is nil. The test fails if SelMethod does not return nil.
func TestSelMethodNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if SelField does not return nil.
	if n := c.SelMethod(&lpcode.SelArgs{}); n != nil {
		t.Error(tserr.NotNil("SelMethod"))
	}
}

// TestTestVariablesNil2 tests Testvariables to return nil in case
// t is nil. The test fails if Testvariables does not return nil.
func TestTestVariablesNil2(t *testing.T) {
	// The test fails if Testvariables does not return nil.
	if n := lpcode.NewCode().Testvariables(nil); n != nil {
		t.Error(tserr.NotNil("Testvariables"))
	}
}

// TestShortVarDeclNil2 tests ShortVarDecl to return nil in case
// a is nil. The test fails if ShortVarDecl dows not return nil.
func TestShortVarDeclNil2(t *testing.T) {
	// The test fails if ShortVarDecl does not return nil.
	if n := lpcode.NewCode().ShortVarDecl(nil); n != nil {
		t.Error(tserr.NotNil("ShortVarDecl"))
	}
}

// TestSelFieldNil2 tests SelField to return nil in case
// a is nil. The test fails if SelField dows not return nil.
func TestSelFieldNil2(t *testing.T) {
	// The test fails if SelField does not return nil.
	if n := lpcode.NewCode().SelField(nil); n != nil {
		t.Error(tserr.NotNil("SelField"))
	}
}

// TestSelMethodNil2 tests SelMethod to return nil in case
// a is nil. The test fails if SelMethod dows not return nil.
func TestSelMethodNil2(t *testing.T) {
	// The test fails if SelMethod does not return nil.
	if n := lpcode.NewCode().SelMethod(nil); n != nil {
		t.Error(tserr.NotNil("SelMethod"))
	}
}

// TestKeyedElementNil2 tests KeyedElement to return nil in case
// a is nil. The test fails if KeyedElement does not return nil.
func TestKeyedElementNil2(t *testing.T) {
	// The test fails if KeyedElement does not return nil.
	if n := lpcode.NewCode().KeyedElement(nil); n != nil {
		t.Error(tserr.NotNil("KeyedElement"))
	}
}

// TestFormatNil tests Format to return nil in case
// *Code is nil. The test fails if Format does not return nil.
func TestFormatNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if Format does not return nil.
	if n := c.Format(); n == nil {
		t.Error(tserr.NilFailed("Format"))
	}
}
