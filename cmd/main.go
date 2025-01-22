package main

import (
	"fmt"
	"sort"
)

// простой поиск элемента
func simpleSearch(sl []int, n int) (int, string) {
	for i, e := range sl {
		if e == n {
			return i, fmt.Sprintf("Найден элемент %v под индексом %v при помощи простого поиска. \n\n", e, i)
		}
	}

	return -1, fmt.Sprintf("Не найден элемент %v в исходном слайсе при помощи простого поиска. \n\n", n)
}

// бинарный поиск элемента в слайсе
func binarSearch(sl []int, n int) (int, string) {
	if len(sl) == 0 {
		return -1, fmt.Sprintf("Не найден элемент %v в исходном слайсе при помощи бинарного поиска. \n\n", n)
	}

	midElement := len(sl) / 2

	if sl[midElement] == n {
		return midElement, fmt.Sprintf("Найден элемент %v под индексом %v при помощи бинарного поиска. \n\n", n, midElement)

	} else if sl[midElement] > n {
		lowSl := sl[:midElement]
		return binarSearch(lowSl, n)

	} else if sl[midElement] < n {
		res, text := binarSearch(sl[midElement+1:], n)
		return res + midElement + 1, text
	}

	return -1, fmt.Sprintf("Не найден элемент %v в исходном слайсе при помощи бинарного поиска. \n\n", n)
}

func main() {
	sl := []int{4, 5, 9, 1, 3, 2, 44, 31, 51}
	sort.Ints(sl) // отсортировали числа по возрастанию [1 2 3 4 5 9 31 44 51]

	fmt.Println(sl)

	var inputNum int
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("Введите число: ")
	fmt.Scan(&inputNum)

	res, text1 := simpleSearch(sl, inputNum)
	res2, text2 := binarSearch(sl, inputNum)

	/////////// Вывод в консоль \\\\\\\\\\\\
	fmt.Println("\n\nSimple Search:\n")
	fmt.Println("input num:", inputNum)
	fmt.Print("result: ", res, "\ntext: ", text1)

	fmt.Println("============================")

	fmt.Println("\nBinary Search:\n")
	fmt.Println("input num:", inputNum)
	fmt.Println("result:", res2, "\ntext:", text2)
	//////////////////////\\\\\\\\\\\\\\\\\\\\\\\\
}
