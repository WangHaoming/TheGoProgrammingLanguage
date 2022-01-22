package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

var Foo Celsius

const (
	AbsoluteZeroc Celsius = -273.15
	FreezinC      Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func init() {
	Foo = 25
}
