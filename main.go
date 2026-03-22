// Package main реализует игровой симулятор тренировки персонажа в ролевой игре.
// Программа позволяет пользователю выбрать класс персонажа (воитель, маг, лекарь),
// ввести имя и тренировать боевые навыки, включая атаку, защиту и специальные умения.
package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// Randint возвращает случайное целое число в диапазоне от min до max.
func Randint(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// Attack возвращает сообщение о нанесённом уроне в зависимости от класса персонажа.
func Attack(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+Randint(3, 5))
	case "mage":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+Randint(5, 10))
	case "healer":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+Randint(-3, -1))
	default:
		return "неизвестный класс персонажа"
	}
}

// Defence возвращает сообщение о заблокированном уроне в зависимости от класса персонажа.
func Defence(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+Randint(5, 10))
	case "mage":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+Randint(-2, 2))
	case "healer":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+Randint(2, 5))
	default:
		return "неизвестный класс персонажа"
	}
}

// Special возвращает сообщение о применении специального умения в зависимости от класса.
func Special(charName, charClass string) string {
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

// StartTraining выводит приветственное сообщение в зависимости от класса,
// а затем запускает цикл тренировки. Пользователь может вводить команды:
// attack, defence, special, skip. Цикл продолжается до ввода команды skip.
func StartTraining(charName, charClass string) string {
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
			fmt.Println(Attack(charName, charClass))
		case "defence":
			fmt.Println(Defence(charName, charClass))
		case "special":
			fmt.Println(Special(charName, charClass))
		}
	}

	return "тренировка окончена"
}

// ChooseCharacterClass предлагает пользователю выбрать класс персонажа,
// показывает описание выбранного класса и запрашивает подтверждение.
func ChooseCharacterClass() string {
	var confirm string
	var charClass string

	for confirm != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		if _, err := fmt.Scanf("%s\n", &charClass); err != nil {
			fmt.Println("Ошибка ввода:", err)
			return ""
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

// Main запускает игровой процесс: приветствие, ввод имени, выбор класса, тренировку.
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

	charClass := ChooseCharacterClass()
	fmt.Println(StartTraining(charName, charClass))
}