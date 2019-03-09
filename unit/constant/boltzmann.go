// Code generated by "go generate gonum.org/v1/gonum/unit/constant”; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constant

import (
	"fmt"

	"gonum.org/v1/gonum/unit"
)

// Boltzmann is the Boltzmann constant (k), it relates the average relative kinetic energy of particles in a gas with the temperature of the gas.
// The dimensions of Boltzmann are kg m^2 K^-1 s^-2.
const Boltzmann = boltzmannUnits(1.380649e-23)

type boltzmannUnits float64

// Unit converts the boltzmannUnits to a *unit.Unit
func (cnst boltzmannUnits) Unit() *unit.Unit {
	return unit.New(float64(cnst), unit.Dimensions{
		unit.MassDim:        1,
		unit.LengthDim:      2,
		unit.TimeDim:        -2,
		unit.TemperatureDim: -1,
	})
}

func (cnst boltzmannUnits) Format(fs fmt.State, c rune) {
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
		fmt.Fprintf(fs, "%%!"+string(c)+"(constant.boltzmannUnits=%v kg m^2 K^-1 s^-2)", float64(cnst))
	}
}
