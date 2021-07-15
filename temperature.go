package main

import "fmt" 

type celsius float64
type kelvin float64
type fahrenheit float64
	
func (f fahrenheit) celsius() celsius {
	return celsius((f-32.0)*5.0/9.0)
}

func (f kelvin) celsius() celsius {
	return celsius(f - 273.15)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func main() {

	var f fahrenheit = 45.0
	c := f.celsius()
	k := c.kelvin()
	fmt.Println(f, c, k )
}
