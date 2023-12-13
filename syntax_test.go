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
	// Format the retrieved code
	if e := c.Format(); e != nil {
		// The test fails if Format returns an error
		t.Error(e)
	}
	// Evaluate the retrieved code with the golden file
	if e := tsfio.EvalGoldenFile(&tsfio.Testcase{Name: "testvariables", Data: c.String()}); e != nil {
		// The test fails if the retrieved code does not match the contents of the golden file
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
