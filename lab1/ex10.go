package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	//Задание.
	//1. Вывести украинскую букву 'Ї'
	//2. Пояснить назначение типа "rune"

	var ukrainianChar rune = 'Ї'
	fmt.Printf("Code '%c' - %d\n", ukrainianChar, ukrainianChar)
}
