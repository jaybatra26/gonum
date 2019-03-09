// Code generated by "go generate gonum.org/v1/gonum/unit/constant”; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constant

import (
	"fmt"
	"testing"
)

func TestLightSpeedInVacuumFormat(t *testing.T) {
	for _, test := range []struct {
		format string
		want   string
	}{
		{"%v", "2.99792458e+08 m s^-1"},
		{"%.1v", "3e+08 m s^-1"},
		{"%50.1v", "                                      3e+08 m s^-1"},
		{"%50v", "                             2.99792458e+08 m s^-1"},
		{"%1v", "2.99792458e+08 m s^-1"},
		{"%#v", "constant.lightSpeedInVacuumUnits(2.99792458e+08)"},
		{"%s", "%!s(constant.lightSpeedInVacuumUnits=2.99792458e+08 m s^-1)"},
	} {
		got := fmt.Sprintf(test.format, LightSpeedInVacuum)
		if got != test.want {
			t.Errorf("Format %q: got: %q want: %q", test.format, got, test.want)
		}
	}
}
