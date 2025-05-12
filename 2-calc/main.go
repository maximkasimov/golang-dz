package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Добро пожаловать в консольный калькулятор")
	if userSelect() == 2 {
		fmt.Println("Сумма чисел массива = ", userSum(userInput()))
	} else if userSelect() == 1 {
		fmt.Println("Среднее массива чисел =", userSrednee(userInput()))
	} else {
		fmt.Println("Медиана =", userMedian(userInput()))
	}
}

func userInput() []int {
	var userData string
	fmt.Print("Введите массив чисел для вычисления (через запятую): ")
	fmt.Scanln(&userData)

	userDivide := strings.Split(userData, ",")
	userArray := make([]int, len(userDivide))

	for i, part := range userDivide {
		num, err := strconv.Atoi(strings.TrimSpace(part))

		if err != nil {
			fmt.Printf("Ошибка при преобразовании '%s': %v\n", part, err)
		}
		userArray[i] = num
	}
	return userArray
}

func userSelect() uint {
	var userChoise uint
	for {
		fmt.Println("Выберите операцию которую необходимо провести над массивом чисел")
		fmt.Println("1: Среднее\n2: Сумма\n3: Медиана")
		fmt.Print("Введите номер математической операции: ")
		fmt.Scan(&userChoise)
		if userChoise == 1 || userChoise == 2 || userChoise == 3 {
			break
		} else {
			fmt.Println("Ошибка ввода, повторите пожалуйста")
			continue
		}
	}
	return userChoise
}

func userSum(userArray []int) int {
	var userSum int = 0
	for i := 0; i < len(userArray); i++ {
		userSum = userSum + userArray[i]
	}
	return userSum
}

func userSrednee(userArray []int) float32 {
	userSrd := float32(userSum(userArray)) / float32(len(userArray))
	return userSrd

}

func userMedian(userArray []int) float64 {
	if len(userArray) == 0 {
		return 0
	}

	arr := make([]int, len(userArray))
	copy(arr, userArray)

	sort.Ints(arr)

	mid := len(arr) / 2

	// Если длина массива нечётная, возвращаем центральный элемент
	if len(arr)%2 != 0 {
		return float64(arr[mid])
	}

	// Если чётная, возвращаем среднее двух центральных элементов
	return float64(arr[mid-1]+arr[mid]) / 2.0
}
