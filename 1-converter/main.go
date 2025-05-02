package main

import "fmt"

var userInCurrency string
var userOutCurrency string
var userAmmount float32

func main() {

	fmt.Println("Добро пожаловать в консольный конвертер валют")
	fmt.Println("Введите исходную валюту (валюту ИЗ которой будет проходить конвертация): ")
	userInCurrency = chooseInCurrency()
	fmt.Println(userInCurrency)

	userAmmount = userAmmountInput()
	fmt.Println(userAmmount)

	fmt.Println("Введите целевую валюту (валюту В которую будет проходить конвертация)")
	userOutCurrency = chooseOutCurrency()
	fmt.Println(userOutCurrency)
}

func chooseInCurrency() string {
	for {
		fmt.Print("Возможные варианты исходной валюты - eur/Eur/EUR, rub/Rub/RUB, usd/Usd/USD: ")
		fmt.Scan(&userInCurrency)

		if userInCurrency == "eur" || userInCurrency == "Eur" || userInCurrency == "EUR" {
			fmt.Println("Выполняем конвертацию из Евро")
			userInCurrency = "eur"
			break
		} else if userInCurrency == "usd" || userInCurrency == "Usd" || userInCurrency == "USD" {
			fmt.Println("Выполняем конвертацию из Доллара США")
			userInCurrency = "usd"
			break
		} else if userInCurrency == "rub" || userInCurrency == "Rub" || userInCurrency == "RUB" {
			fmt.Println("Выполняем конвертацию из Рубля")
			userInCurrency = "rub"
			break
		} else {
			fmt.Println("Неправильный ввод исходной валюты. Повторите ввод")
			continue
		}
	}
	return userInCurrency
}

func chooseOutCurrency() string {
	for {
		switch userInCurrency {
		case "eur":
			fmt.Print("Возможные варианты целевой валюты - rub/Rub/RUB, usd/Usd/USD: ")
			fmt.Scan(&userOutCurrency)
			if userOutCurrency == "rub" || userOutCurrency == "Rub" || userOutCurrency == "RUB" {
				fmt.Println("Выполняем конвертацию в Рубль")
				userOutCurrency = "rub"
				break
			} else if userOutCurrency == "usd" || userOutCurrency == "Uub" || userOutCurrency == "USD" {
				fmt.Println("Выполняем конвертацию в доллар")
				userOutCurrency = "usd"
				break
			} else {
				fmt.Println("Неправильная целевая валюта. Повторите ввод")
				continue
			}
		case "usd":
			fmt.Print("Возможные варианты целевой валюты - rub/Rub/RUB, eur/Eur/EUR: ")
			fmt.Scan(&userOutCurrency)
			if userOutCurrency == "rub" || userOutCurrency == "Rub" || userOutCurrency == "RUB" {
				fmt.Println("Выполняем конвертацию в рубль")
				userOutCurrency = "rub"
				break
			} else if userOutCurrency == "eur" || userOutCurrency == "Eur" || userOutCurrency == "EUR" {
				fmt.Println("Выполняем конвертацию в евро")
				userOutCurrency = "eur"
				break
			} else {
				fmt.Println("Неправильная целевая валюта. Повторите ввод")
				continue
			}
		case "rub":
			fmt.Print("Возможные варианты целевой валюты - usd/Usd/USD, eur/Eur/EUR: ")
			fmt.Scan(&userOutCurrency)
			if userOutCurrency == "usd" || userOutCurrency == "Usd" || userOutCurrency == "USD" {
				fmt.Println("Выполняем конвертацию в доллар")
				userOutCurrency = "usd"
				break
			} else if userOutCurrency == "eur" || userOutCurrency == "Eur" || userOutCurrency == "EUR" {
				fmt.Println("Выполняем конвертацию в евро")
				userOutCurrency = "eur"
				break
			} else {
				fmt.Println("Неправильная целевая валюта. Повторите ввод")
				continue
			}
		}
		return userOutCurrency
	}
}

func userAmmountInput() float32 {
	for {
		fmt.Print("Введите объем исходной валюты (", userInCurrency, "): ")
		_, err := fmt.Scanf("%f", &userAmmount)

		if err == nil {
			return userAmmount
		}
		fmt.Println("Неправильный ввод числа, повторите ввод")

		//   очищаем строку чтобы не было зацикливания - запомнить прием
		var strDiscard string
		fmt.Scanln(&strDiscard)
		continue
	}
}

// func userCounting(userAmmount float64, userInCurrency string, userOutCurrency string) float64 {
// 	const usd2Eur = 0.88
// 	const usd2Rub = 83.5

// }
