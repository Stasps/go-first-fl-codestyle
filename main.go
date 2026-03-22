package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// randint возвращает случайное целое число в диапазоне от min до max
func randint(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// attack возвращает сообщение о нанесённом уроне в зависимости от класса персонажа.
func attack(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(3, 5))
	case "mage":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(5, 10))
	case "healer":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(-3, -1))
	default:
		return "неизвестный класс персонажа"
	}
}

// defence возвращает сообщение о заблокированном уроне в зависимости от класса персонажа.
func defence(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(5, 10))
	case "mage":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(-2, 2))
	case "healer":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(2, 5))
	default:
		return "неизвестный класс персонажа"
	}
}

// special возвращает сообщение о применении специального умения в зависимости от класса.
func special(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", charName, 105)
	case "mage":
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", charName, 45)
	case "healer":
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", charName, 40)
	default:
		return "неизвестный класс персонажа"
	}
}

// startTraining выводит приветственное сообщение в зависимости от класса,
// а затем запускает цикл тренировки. Пользователь может вводить команды:
// attack, defence, special, skip. Цикл продолжается до ввода команды skip.
func startTraining(charName, charClass string) string {
	// Приветствие в зависимости от класса
	switch charClass {
	case "warrior":
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", charName)
	case "mage":
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", charName)
	case "healer":
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", charName)
	}

	// Инструкция
	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		if _, err := fmt.Scanf("%s\n", &cmd); err != nil {
			fmt.Println("Ошибка ввода:", err)
			return "тренировка прервана из-за ошибки"
		}

		switch cmd {
		case "attack":
			fmt.Println(attack(charName, charClass))
		case "defence":
			fmt.Println(defence(charName, charClass))
		case "special":
			fmt.Println(special(charName, charClass))
		}
	}

	return "тренировка окончена"
}

// chooseCharacterClass предлагает пользователю выбрать класс персонажа,
// показывает описание выбранного класса и запрашивает подтверждение.
func chooseCharacterClass() string {
	var confirm string
	var charClass string

	for confirm != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		if _, err := fmt.Scanf("%s\n", &charClass); err != nil {
			fmt.Println("Ошибка ввода:", err)
			return "" // возвращаем пустую строку при ошибке
		}

		// Описание выбранного класса
		switch charClass {
		case "warrior":
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		case "mage":
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		case "healer":
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		if _, err := fmt.Scanf("%s\n", &confirm); err != nil {
			fmt.Println("Ошибка ввода:", err)
			return ""
		}
		confirm = strings.ToLower(confirm)
	}

	return charClass
}

// main запускает игровой процесс: приветствие, ввод имени, выбор класса, тренировку.
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	if _, err := fmt.Scanf("%s\n", &charName); err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	charClass := chooseCharacterClass()
	fmt.Println(startTraining(charName, charClass))
}
