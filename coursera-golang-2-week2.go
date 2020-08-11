package main

import (
	"fmt"
)


func main() {

	fmt.Print("input acceleration: ")
	_acceleration:= 0.0
	fmt.Scanf("%f\n",&_acceleration)

	fmt.Print("input initial velocity: ")
	_initialVelocity:= 0.0
	fmt.Scanf("%f\n",&_initialVelocity)

	fmt.Print("input initial displacement: ")
	_initialDisplacement:= 0.0
	fmt.Scanf("%f\n",&_initialDisplacement)
	
	fmt.Print("input time: ")
	_time:=0.0
	fmt.Scanf("%f\n",&_time)

	fn := GenDisplaceFn(_acceleration, _initialVelocity, _initialDisplacement)

	fmt.Println("displacement after", _time, "seconds is", fn(_time))
}

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}