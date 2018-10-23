package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {

	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) //如果小于0，则返回我们自定义的错误
	}

	return x, nil
}

func main() {
	fmt.Println(Sqrt(2))
	_, err := Sqrt(-2)
	if err != nil {
		switch err.(type) { //类型断言，获取错误码
		case ErrNegativeSqrt:
			fmt.Printf("ErrNegativeSqrt\n")
		default:

		}
	}
}
