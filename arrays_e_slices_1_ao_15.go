package main

import "fmt"

// q.1

func main() {
	var x int
	var y int
	var z int
	fmt.Print("Escolha um número.")
	fmt.Scan(&x)
	fmt.Print("Escolha um número.")
	fmt.Scan(&y)
	fmt.Print("Escolha um número.")
	fmt.Scan(&z)

	var numeros = [3]int{x, y, z}
	fmt.Println(numeros)
	fmt.Println(x + y + z)
}

// q.2

func main() {
	var numeros = []int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	numeros = append(numeros[:2], numeros[3:]...)
	fmt.Println(numeros)
}

// q.3

func main() {
	var w float64
	var x float64
	var y float64
	var z float64
	fmt.Print("Escolha um número.")
	fmt.Scan(&w)
	fmt.Print("Escolha um número.")
	fmt.Scan(&x)
	fmt.Print("Escolha um número.")
	fmt.Scan(&y)
	fmt.Print("Escolha um número.")
	fmt.Scan(&z)

	var numeros = [4]float64{w, x, y, z}
	fmt.Println(numeros)
	fmt.Println(w * x * y * z)
}

// q.4

func main() {
	var slice []int
	var soma, tamanho, x int

	fmt.Print("Digite o tamanho do Slice:")
	fmt.Scan(&tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)
		slice = append(slice, x)
		soma = soma + x
	}

	fmt.Println(slice)
	fmt.Println(soma)

}

// q.5

func main() {
	var matriz [3][2]int
	fmt.Println(matriz)

	var v1 int
	var v2 int
	var v3 int
	var v4 int
	var v5 int
	var v6 int

	fmt.Print("Escolha um número.")
	fmt.Scan(&v1)
	fmt.Print("Escolha um número.")
	fmt.Scan(&v2)
	fmt.Print("Escolha um número.")
	fmt.Scan(&v3)
	fmt.Print("Escolha um número.")
	fmt.Scan(&v4)
	fmt.Print("Escolha um número.")
	fmt.Scan(&v5)
	fmt.Print("Escolha um número.")
	fmt.Scan(&v6)

	matrizinicializada := [3][2]int{{v1, v2}, {v3, v4}, {v5, v6}}
	fmt.Println(matrizinicializada)

}

// q.6

func main() {
	var x int
	var y int
	var numeros = [10]int{1, 23, 7, 44, 5, 38, 0, 68, 91, 100}
	fmt.Print("Insira o valor que deseja procurar.")
	fmt.Scan(&x)

	for {
		y++
		if x == numeros[y] {
			fmt.Println("O valor que você procura já está inserido no array.")
		} else if len(numeros) >= 10 {
			fmt.Println("O valor que você procura não está presente no array.")
			break
		}
	}
}

// q.7

func main() {
	var numeros = []int{1, 2, 3, 4, 5}
	var x int

	fmt.Print("Escolha um número inteiro.")
	fmt.Scan(&x)

	if x >= 1 && x <= 5 {
		fmt.Println("Esse número já está na lista.")
		fmt.Println(numeros)
	} else {
		fmt.Println("Esse número não está na lista, mas será adicionado.")
		numeros = append(numeros, x)

		fmt.Println(numeros)
	}

}

// q.8

func main() {
	var numeros = []string{"Marta", "Daniel", "Gabriel", "Artur", "Luiza", "Gabriel", "Ana", "Beatriz"}
	var valor string
	var x int

	fmt.Print("Escreva um nome.")
	fmt.Scan(&valor)

	w := x + 1

	for len(numeros) > x {
		if valor == numeros[x] {
			numeros = append(numeros[:x], numeros[w:]...)
		}
		w++
		x++
	}

}

// q.9

func main() {
	var numeros = [6]float64{1.5, 2.4, 3.8, 4.7, 5.5, 6.9}
	var x float64
	var i int = 0
	fmt.Println(numeros)

	fmt.Print("Escolha um número para ser adicionado ao array em todas as posições.")
	fmt.Scan(&x)

	for i < len(numeros) {
		numeros[i] = numeros[i] + x
		i++
	}
	fmt.Println("O array final é: ", numeros)

}

// q.10

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	numMin := 10000000000000000
	numMax := -1000000000000000

	for _, n := range slice {
		if n > numMax {
			numMax = n
		}
		if n < numMin {
			numMin = n
		}
	}
	fmt.Println("Valor mínimo:", numMin)
	fmt.Println("Valor máximo:", numMax)
}

// q.11

func main() {
	var matriz [2][3]int
	fmt.Println(matriz)

	matrizinicializada := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matrizinicializada)

	var indxLinha int
	var indxColuna int

	fmt.Print("Escolha um número.")
	fmt.Scan(&indxLinha)
	fmt.Print("Escolha um número.")
	fmt.Scan(&indxColuna)

	if indxLinha == 0 && indxColuna == 0 {
		fmt.Println(matrizinicializada[0][0])
	} else if indxLinha == 0 && indxColuna == 1 {
		fmt.Println(matrizinicializada[0][1])
	} else if indxLinha == 0 && indxColuna == 2 {
		fmt.Println(matrizinicializada[0][2])
	} else if indxLinha == 1 && indxColuna == 0 {
		fmt.Println(matrizinicializada[1][0])
	} else if indxLinha == 1 && indxColuna == 1 {
		fmt.Println(matrizinicializada[1][1])
	} else if indxLinha == 1 && indxColuna == 2 {
		fmt.Println(matrizinicializada[1][2])
	} else {
		fmt.Println("Índices não existem nessa matriz, portanto não há valor.")
	}
}

// q.12

func main() {
	array := [5]int{3, 7, 12, 24, 25}
	sliceResultante := []int{}

	for i := 0; i < len(array); i++ {
		if array[i]%3 == 0 {
			sliceResultante = append(sliceResultante, array[i])
		}
	}
	fmt.Println("Os múltiplos de 3 dentro do array são:", sliceResultante)
}

// q.13

func main() {
	var numeros = []int{1, 4, 7, 13, 22, 46, 69}
	var valor_novo int
	fmt.Println(numeros)

	fmt.Print("Escolha um número para ser adicionado ao array.")
	fmt.Scan(&valor_novo)

	numeros = append([]int{valor_novo}, numeros...)
	fmt.Println(numeros)

	numeros = append(numeros, valor_novo)
	fmt.Println("O array resultante é: ", numeros)
}

// q.14

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80}
	var n1, n2 int

	fmt.Print("Escolha um número (0-7):")
	fmt.Scan(&n1)
	fmt.Print("Escolha um número (0-7):")
	fmt.Scan(&n2)

	slice[n1], slice[n2] = slice[n2], slice[n1]
	fmt.Println(slice)

}

// q.15

func main() {
	array := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceRes := []float64{}

	for i := 0; i < len(array); i++ {
		if array[i] > 5 {
			sliceRes = append(sliceRes, array[i])
		}
	}
	fmt.Println("Elementos que são maiores do que 5:", sliceRes)

}
