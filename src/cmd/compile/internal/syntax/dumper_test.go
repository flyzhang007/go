// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

import (
	"os"
	"testing"
)

func TestDump(t *testing.T) {
	ast, err := ReadFile(*src, 0)
	if err != nil {
		t.Fatal(err)
	}
	Fdump(os.Stdout, ast)
}