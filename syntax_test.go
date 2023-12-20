// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode_test

// Import Go standard library package testing as well as lpcode, tserr and tsfio
import (
	"testing" // testing

	"github.com/thorstenrie/lpcode" // lpcode
	"github.com/thorstenrie/tserr"  // tserr
	"github.com/thorstenrie/tsfio"  // tsfio
)

var (
	testKey     string = "foo"
	testElem    string = "bar"
	testIdent   string = "dort"
	testExpr    string = "mund"
	testComment string = "Sumsemann hieß der dicke Maikaefer, der im Fruehling auf einer Kastanie im Garten von Peterchens Eltern hauste, nicht weit von der großen Wiese mit den vielen Sternblumen."
)

// TestTestVariables tests generated test variables by Testvariables. The test fails if
// the generated test variables do not match the contents of the golden file.
func TestTestVariables(t *testing.T) {
	// Configure Testvars to generate test variables of types string, error, int and float
	vars := &lpcode.Testvars{
		String: 1,
		Error:  1,
		Int:    1,
		Float:  1,
	}
	// Retrieve test variables with Testvariables
	c := lpcode.NewCode().Testvariables(vars)
	// Evaluate the retrieved source code
	if e := evalCode(c, "testvariables"); e != nil {
		// The test fails if the retrieved code does not match the contents of the golden file
		t.Error(e)
	}
}

// TestKeyedElement tests generated source code using the short variable declaration by ShortVarDecl,
// a keyed element by KeyedElement and function ending by FuncEnd. The test fails if the generated
// source code does not match the contents of the golden file.
func TestKeyedElement(t *testing.T) {
	// Retrieve the short variable declaration with ShortVarDecl
	c := lpcode.NewCode().ShortVarDecl(&lpcode.ShortVarDeclArgs{Ident: testIdent, Expr: testExpr + "{"})
	// Retrieve the keyed element with KeyedElement and the function ending with FuncEnd
	c.KeyedElement(&lpcode.KeyedElementArgs{Key: testKey, Elem: testElem}).FuncEnd()
	// Evaluate the retrieved source code
	if e := evalCode(c, "keyedelement"); e != nil {
		// The test fails if the generated source code does not match the contents of the golden file
		t.Error(e)
	}
}

// TestLineComment tests generated source code using a line comment by LineComment. The test fails
// if the generated source code does not match the contents of the golden file.
func TestLineComment(t *testing.T) {
	// Retrieve the line comment with LineComment
	c := lpcode.NewCode().LineComment(testComment)
	// Evaluate the retrieved source code
	if e := evalCode(c, "linecomment"); e != nil {
		// The test fails if the generated source code does not match the contents of the golden file
		t.Error(e)
	}
}

// TestTestVariablesEmpty tests Testvariables to return an empty string in case
// no test variables are configured. The test fails if the returned string is
// not empty.
func TestTestVariablesEmpty(t *testing.T) {
	// // Configure Testvars to generate no test variables
	vars := &lpcode.Testvars{}
	// Retrieve test variables with Testvariables
	c := lpcode.NewCode().Testvariables(vars)
	// The test fails in case the returned string is not empty
	if c.String() != "" {
		t.Error(tserr.NotEmpty("code"))
	}
}

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

// TestKeyedElementNil tests KeyedElement to return nil in case
// *Code is nil. The test fails if KeyedElement does not return nil.
func TestKeyedElementNil(t *testing.T) {
	// Declare c as type *Code and assign nil
	var c *lpcode.Code = nil
	// The test fails if KeyedElement does not return nil.
	if n := c.KeyedElement(&lpcode.KeyedElementArgs{}); n != nil {
		t.Error(tserr.NotNil("KeyedElementArgs"))
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

// TestFormatEmpty tests Format to return nil for an empty code.
// The test fails if Format returns an error.
func TestFormatEmpty(t *testing.T) {
	// Retrieve the return value of Format for an empty code
	if e := lpcode.NewCode().Format(); e != nil {
		// The test fails if Format returns an error
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Format", Fn: "empty source code", Err: e}))
	}
}

// TestFormatNoCode tests Format to return an error in case code contains a
// placeholder text which is not valid Go source code. The test fails if Format returns nil.
func TestFormatNoCode(t *testing.T) {
	// Declare testcase loremipsum
	tc := "loremipsum"
	// Retrieve golden file path for the testcase
	fn, err := tsfio.GoldenFilePath(tc)
	// The test fails if GoldenFilePath returns an error
	if err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "GoldenFilePath", Fn: tc, Err: err}))
	}
	// Retrieve the contents of the golden file for the test case
	li, err := tsfio.ReadFile(fn)
	// The test fails if ReadFile returns an error
	if err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(fn), Err: err}))
	}
	// Retrieve a new Code instance with the contents of the golden file
	c := lpcode.NewCode().Ident(string(li))
	// The test fails if Format returns nil
	if e := c.Format(); e == nil {
		t.Error(tserr.NilFailed("format code"))
	}
}
