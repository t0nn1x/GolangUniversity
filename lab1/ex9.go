package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)       // false
	fmt.Println("second = ", second)      // false
	fmt.Println("third  = ", third)       // true
	fmt.Println("fourth = ", fourth)      // false
	fmt.Println("fifth  = ", fifth, "\n") // true

	fmt.Println("!true  = ", !true)        // false
	fmt.Println("!false = ", !false, "\n") // true

	fmt.Println("true && true   = ", true && true)         // true
	fmt.Println("true && false  = ", true && false)        // false
	fmt.Println("false && false = ", false && false, "\n") // false

	fmt.Println("true || true   = ", true || true)         // true
	fmt.Println("true || false  = ", true || false)        // true
	fmt.Println("false || false = ", false || false, "\n") // false

	fmt.Println("2 < 3  = ", 2 < 3)        // true
	fmt.Println("2 > 3  = ", 2 > 3)        // false
	fmt.Println("3 < 3  = ", 3 < 3)        // false
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true
	fmt.Println("3 > 3  = ", 3 > 3)        // false
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true
	fmt.Println("2 == 3 = ", 2 == 3)       // false
	fmt.Println("3 == 3 = ", 3 == 3)       // true
	fmt.Println("2 != 3 = ", 2 != 3)       // true
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false

	//Задание.
	//1. Пояснить результаты операций

	/*
		Ініціалізація змінних
		first і second ініціалізуються як bool з типовим значенням false.
		third ініціалізується як true.
		fourth отримує протилежне значення до third, тобто !true, яке дорівнює false.
		fifth ініціалізується як true.
		Логічні операції
		!true — операція NOT, інвертує true до false.
		!false — інвертує false до true.
		Логічні операції AND (&&) та OR (||)
		true && true — дорівнює true, оскільки обидва операнди true.

		true && false — дорівнює false, оскільки один з операндів false.

		false && false — дорівнює false, оскільки обидва операнди false.

		true || true — дорівнює true, оскільки хоча б один з операндів true.

		true || false — дорівнює true, оскільки хоча б один з операндів true.

		false || false — дорівнює false, оскільки обидва операнди false.

		Порівняння чисел
		2 < 3 — true, оскільки 2 дійсно менше 3.
		2 > 3 — false, оскільки 2 не більше 3.
		3 < 3, 3 > 3 — обидва вирази false, оскільки числа однакові.
		3 <= 3, 3 >= 3 — обидва вирази true, оскільки числа однакові.
		2 == 3 — false, оскільки числа різні.
		3 == 3 — true, оскільки числа однакові.
		2 != 3 — true, оскільки числа різні.
		3 != 3 — false, оскільки числа однакові.
	*/
}
