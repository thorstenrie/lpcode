// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode_test

// Import Go packages lpcode, tserr and tsfio
import (
	"github.com/thorstenrie/lpcode" // lpcode
	"github.com/thorstenrie/tserr"  // tserr
	"github.com/thorstenrie/tsfio"  // tsfio
)

// evalCode evaluates the source code of c by comparing the source code
// with the associated golden file for test case tc. The function returns
// an error if the source code does not match the contents of the golden file.
func evalCode(
	c *lpcode.Code,
	tc string,
) error {
	// Return an error if c is nil
	if c == nil {
		return tserr.NilPtr()
	}
	// Format the retrieved code
	if e := c.Format(); e != nil {
		// The test fails if Format returns an error
		return e
	}
	// Evaluate the retrieved code with the golden file
	if e := tsfio.EvalGoldenFile(&tsfio.Testcase{Name: tc, Data: c.String()}); e != nil {
		// The test fails if the retrieved code does not match the contents of the golden file
		return e
	}
	// Return nil
	return nil
}
