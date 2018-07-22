// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-22
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
error exercise
==================================================================
error 练习 制造错误，走正常流程
------------------------------------------------------------------*/
package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", fmt.Sprint(float64(e)))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	return x, nil
}

func main() {
	fmt.Println(ErrNegativeSqrt(-2).Error())	// cannot Sqrt negative number: -2

	fmt.Println(Sqrt(2))		// 2 <nil>
	fmt.Println(Sqrt(-2))		// -2 cannot Sqrt negative number: -2

	if foo, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(foo)	// 2
	}

	if foo, err := Sqrt(-2); err != nil {
		fmt.Println(err)	// cannot Sqrt negative number: -2
	} else {
		fmt.Println(foo)
	}
}
