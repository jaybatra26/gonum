// Code generated by "go generate gonum.org/v1/gonum/unit/constant”; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constant

import (
	"fmt"

	"gonum.org/v1/gonum/unit"
)

// ElectricConstant is the electric constant (ε₀), the value of the absolute dielectric permittivity of classical vacuum.
// The dimensions of ElectricConstant are A^2 s^4 kg^-1 m^-3.
const ElectricConstant = electricConstantUnits(8.854187817620389e-12)

type electricConstantUnits float64

// Unit converts the electricConstantUnits to a *unit.Unit
func (cnst electricConstantUnits) Unit() *unit.Unit {
	return unit.New(float64(cnst), unit.Dimensions{
		unit.CurrentDim: 2,
		unit.TimeDim:    4,
		unit.MassDim:    -1,
		unit.LengthDim:  -3,
	})
}

func (cnst electricConstantUnits) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", cnst, float64(cnst))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), w, p, cnst.Unit())
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, cnst.Unit())
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), w, cnst.Unit())
		default:
			fmt.Fprintf(fs, "%"+string(c), cnst.Unit())
		}
	default:
		fmt.Fprintf(fs, "%%!"+string(c)+"(constant.electricConstantUnits=%v A^2 s^4 kg^-1 m^-3)", float64(cnst))
	}
}
