package concepts

import "fmt"

func LazyLoad() {
	defer fmt.Println("Lazy Printing...")
	for i := 0; i < 10; i++ {
		fmt.Println("Synchronous printing")
	}
}

func ReversePrint() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func printNumber(i int) {
	defer fmt.Println(i)
	fmt.Print("Printing i: ")
}

func ReversePrintWithFunction() {
	for i := 0; i < 10; i++ {
		defer printNumber(i)
	}
}
