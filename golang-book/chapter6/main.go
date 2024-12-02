package main

import "fmt"

func main() {
	//fmt.Println(Q2(1))
	//fmt.Println(Q2(2))

	//fmt.Println(Q3(1.4, 2.6, 1.5))

	/*
		nextOdd := makeOddGenerator()
		fmt.Println(nextOdd())
		fmt.Println(nextOdd())
		fmt.Println(nextOdd())
	*/

	/*
		fmt.Println(fib(0))
		fmt.Println(fib(1))
		fmt.Println(fib(2))
		fmt.Println(fib(3))
		fmt.Println(fib(4))
		fmt.Println(fib(5))
		fmt.Println(fib(6))
		fmt.Println(fib(7))
	*/

	// fmt.Println(deferTest(1))

	/*
		x := 1
		y := 2
		swap(&x, &y)
		fmt.Println(x, y)
	*/

}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func deferTest(x int) int {
	defer defferedFunc()
	fmt.Println(x)
	return x + 1
}

func defferedFunc() {
	fmt.Println("Deffered Func Ran")
}

func fib(n uint) (ret uint) {
	switch n {
	case 0:
		ret = n
	case 1:
		ret = n
	default:
		ret = fib(n-1) + fib(n-2)
	}
	return
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func Q3(args ...float32) (biggest float32) {
	biggest = 0.0
	for _, v := range args {
		if v > biggest {
			biggest = v
		}
	}
	return
}

func Q2(x int) (half int, even bool) {
	half = x / 2
	even = (half%2 == 0)

	return
}
