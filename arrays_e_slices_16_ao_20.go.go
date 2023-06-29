package main

import "fmt"

// q.16

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceRes := []int{}

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			sliceRes = append(sliceRes, array[i])
		}
	}

	fmt.Println("Elementos pares do array original:", sliceRes)
}

// q.17

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	soma := array[0] + array[2] + array[4] + array[6] + array[8]
	fmt.Println(soma)
}

// q.18

func main() {
	var n int

	fmt.Print("Digite o valor de n: ")
	fmt.Scan(&n)

	var primos []int

	for i := 2; i < n; i++ {
		primo := true
		for x := 2; x < i; x++ {
			if i%x == 0 {
				primo = false
				break
			}
		}
		if primo {
			primos = append(primos, i)
		}
	}
}

// q.19

func main() {
	var n int
	fmt.Print("Digite o tamanho dos arrays: ")
	fmt.Scanln(&n)

	array1 := make([]int, n)
	array2 := make([]int, n)

	fmt.Println("Digite os elementos do primeiro array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i)
		fmt.Scanln(&array1[i])
	}

	fmt.Println("Digite os elementos do segundo array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i)
		fmt.Scanln(&array2[i])
	}

	arraySoma := make([]int, n)
	for i := 0; i < n; i++ {
		arraySoma[i] = array1[i] + array2[i]
	}
	fmt.Println("Array soma:", arraySoma)
}

// q.20

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	crescente := true
	for i := 1; i < len(slice); i++ {
		atual := slice[i]
		anterior := slice[i-1]

		if anterior >= atual {
			crescente = false
			break
		}
	}

	if crescente {
		fmt.Println()
	}
}
