//handles conversions between Celsius and Fahrenheit
package go_temp_conv

//create custom types to prevent erroneous
//arithmetic between Celsius and Fahrenheit values
//as they are not equal in base
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZero Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func main(){

}
