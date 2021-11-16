package main

import "fmt"

func main() {
	out := f()
	fmt.Printf("out = %d\n", out)
}

func f() (r int) {
	/* return文を含んでいないのに、ゼロ値ではない値を返す関数 */
	defer func() {
		if p := recover(); p != nil {
			r = p.(int)
		}
	}()
	panic(1)
}
