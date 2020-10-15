package main

import "fmt"

func fibonacci(n uint) uint {
	if n < 2 {
		return n
	}
	return (fibonacci(n-1) + fibonacci(n-2))
}

func sum(args ...int) int {
	mayor := args[0]
	for _, v := range args {
		if v > mayor {
			mayor = v
		}
	}
	return mayor
}

func generadorPares() func() uint {
	i := uint(0)
	return func() uint {
		var par = i
		i += 2
		return par
	}
}

func intercambia(a *int, b *int) {
	tem := *a
	*a = *b
	*b = tem
}

func main() {
	fmt.Println(fibonacci(10))
	fmt.Println(sum(4, 5, 2, 7, 1, 9, 3, 6, 8))
	sigPar := generadorPares()
	fmt.Println(sigPar())
	fmt.Println(sigPar())
	fmt.Println(sigPar())
	fmt.Println(sigPar())
	a := 10
	b := 5
	intercambia(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
