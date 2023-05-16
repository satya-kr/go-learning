package main

import "fmt"

type Numbers interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	fmt.Println("Int Non-Generic Sum=", IntSum(ints))
	fmt.Println("Float Non-Generic Sum=", FloatSum(floats))

	// fmt.Println("Int Generic Sum=", GenericIntFloatSum(ints))
	// fmt.Println("Float Generic Sum=", GenericIntFloatSum(floats))

	fmt.Println("Int Generic Sum=", OptmiseUsingInterfaceGenericIntFloatSum(ints))
	fmt.Println("Float Generic Sum=", OptmiseUsingInterfaceGenericIntFloatSum(floats))

}

func IntSum(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func FloatSum(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

//generic only work in v1.8 or later
func GenericIntFloatSum[k comparable, V int64 | float64](m map[k]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func OptmiseUsingInterfaceGenericIntFloatSum[k comparable, V Numbers](m map[k]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
