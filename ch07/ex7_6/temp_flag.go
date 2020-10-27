package main

import (
	"flag"
	"fmt"
)

const unitC = "C"
const unitTempC = "℃"
const unitF = "F"
const unitTempF = "°F"
const unitK = "K"
const unitTempK = "°K"

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case unitC, unitTempC:
		f.Celsius = Celsius(value)
		return nil
	case unitF, unitTempF:
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case unitK, unitTempK:
		f.Celsius = KToC(Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.4g 度", c)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

var temp = CelsiusFlag("temp", 15.0, "temperature")

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
