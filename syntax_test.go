// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode_test

import (
	"fmt"
	"testing"

	"github.com/thorstenrie/lpcode"
)

func TestTestVariables(t *testing.T) {
	vars := &lpcode.Testvars{
		String: 1,
		Error:  1,
		Int:    1,
		Float:  1,
	}
	c := lpcode.NewCode().Testvariables(vars)
	c.Format()
	fmt.Print(c.String())
}
