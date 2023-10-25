package main

import (
	"fmt"
	"math"
)

func integral(x float64) float64 {

	y := x * x
	fmt.Println("y = ", y)
	return (math.Round(y * 1000)) / 1000
}

func main() {

	var f_point, l_point, step float64

	fmt.Println("Введите начало отрезка (x)")
	fmt.Scanf("%f\n", &f_point)

	fmt.Println("Введите конец отрезка (x)")
	fmt.Scanf("%f\n", &l_point)

	fmt.Println("Введите шаг")
	fmt.Scanf("%f\n", &step)

	s_f := 0.

	for i := f_point + step; i <= l_point; i += step {

		s_f += integral(i-step) + integral(i)

	}

	fmt.Println("Площадь = ", s_f*step*0.5)

}
