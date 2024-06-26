package main

import (
	"fmt"
	"go-basics/modules/calculator"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(1, 2))
	fmt.Println(calculator.Multiply(1, 2))

	fmt.Printf("%v\n", "---------")

	var a1 int = 3              // 明示的な代入を省略した場合は0値で初期化 0が3つの要素で初期化
	var a2 = [3]int{10, 20, 30} // 明示的な代入
	a3 := [...]int{10, 20}      // 要素数のところでスリードットをつけると自動的にカウントしてくれる
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v %v %v\n", a1, a2, a3)
	fmt.Printf("%v %v\n", len(a3), cap(a3))
	fmt.Printf("%T %T\n", a2, a3)

	var s1 []int
	s2 := []int{}
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
}
