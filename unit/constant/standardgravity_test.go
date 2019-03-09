// Code generated by "go generate gonum.org/v1/gonum/unit/constant”; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constant

import (
	"fmt"
	"testing"
)

func TestStandardGravityFormat(t *testing.T) {
	for _, test := range []struct {
		format string
		want   string
	}{
		{"%v", "9.80665 m s^-2"},
		{"%.1v", "1e+01 m s^-2"},
		{"%50.1v", "                                      1e+01 m s^-2"},
		{"%50v", "                                    9.80665 m s^-2"},
		{"%1v", "9.80665 m s^-2"},
		{"%#v", "constant.standardGravityUnits(9.80665)"},
		{"%s", "%!s(constant.standardGravityUnits=9.80665 m s^-2)"},
	} {
		got := fmt.Sprintf(test.format, StandardGravity)
		if got != test.want {
			t.Errorf("Format %q: got: %q want: %q", test.format, got, test.want)
		}
	}
}
