package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func sum(values ...int) (result int) {
	for _, value := range values {
		result += value
	}
	return
}

func sub(values ...int) *int {
	result := 0
	for _, value := range values {
		result -= value
	}
	return &result
}

func divide(x, y float64) (float64, error) {
	if y == 0.0 {
		return 0.0, fmt.Errorf("Can't divide by zero")
	}
	return x / y, nil
}

func main() {
	fmt.Println(getSum(3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(*sub(1, 2, 3, 4, 5))
	res, err := divide(1, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)

}
