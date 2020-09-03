package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func CtoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FtoK(f Fahrenheit) Kelvin {
	return Kelvin((f + 459.67) * 5/9)
}

func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func KtoF(k Kelvin) Fahrenheit {
	return Fahrenheit((k*9/5) - 459.67)
}
