package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// util functions
func stringToInt(strNum string) int {
	num, _ := strconv.Atoi(strNum)
	return num
}

func stringToFloat64(strNum string) float64 {
	newFloat, _ := strconv.ParseFloat(strings.TrimSpace(strNum), 64)
	return newFloat
}

// Сумма цифр числа
func task_1_1() {
	var num string
	fmt.Print("Введите целое 4х-значное число: ")
	fmt.Scanf("%s\n", &num)

	sum := int(num[0]-'0') + int(num[1]-'0') + int(num[2]-'0') + int(num[3]-'0')

	fmt.Printf("Сумма цифр: %d\n", sum)
}

// Преобразование температуры
func task_1_2() {
	fmt.Println("Введите температуру с указанием единицы (например, 25 (Celsius) или 77 (Fahrenheit)): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		fmt.Println("Ошибка: неверный формат ввода.")
		return
	}

	valueStr := parts[0]
	value := stringToFloat64(valueStr)

	unit := strings.ToLower(parts[1])

	switch unit {
	case "celsius":
		fahrenheit := (value * 9 / 5) + 32
		fmt.Printf("%.2f°C = %.2f°F\n", value, fahrenheit)
	case "fahrenheit":
		celsius := (value - 32) * 5 / 9
		fmt.Printf("%.2f°F = %.2f°C\n", value, celsius)
	default:
		fmt.Println("Ошибка, задана некорректная единица измерения.")
	}
}

// Удвоение каждого элемента массива
func task_1_3() {
	var strArr string
	fmt.Println("Введите массив из 4х элементов через ',' без пробелов:")
	fmt.Scanf("%s\n", &strArr)

	stringArray := strings.Split(strArr, ",")

	fmt.Println("Удвоенный массив массив:")
	fmt.Printf("%d,", stringToInt(stringArray[0])*2)
	fmt.Printf("%d,", stringToInt(stringArray[1])*2)
	fmt.Printf("%d,", stringToInt(stringArray[2])*2)
	fmt.Println(stringToInt(stringArray[3]) * 2)
}

// Объединение строк
func task_1_4() {
	var strArr string
	fmt.Println("Введите строки через ',' без пробелов:")
	fmt.Scanf("%s\n", &strArr)

	stringArray := strings.Replace(strArr, ",", " ", -1)
	fmt.Println(stringArray)
}

// Расчет расстояния между двумя точками
func task_1_5() {
	var x1, y1, x2, y2 float64
	fmt.Println("Введите координыта первой точки (x1 y1)")
	fmt.Scan(&x1, &y1)
	fmt.Println("Введите координыта второй точки (x2 y2)")
	fmt.Scan(&x2, &y2)

	d := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
	fmt.Printf("Расстояние между точками: %.2f\n", d)
}

// Проверка на четность/нечетность
func task_2_1() {
	var x int

	fmt.Println("Введите целое число:")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

// Проверка високосного года
func task_2_2() {
	var year int

	fmt.Println("Введите год")
	fmt.Scan(&year)

	var leap bool

	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				leap = true
			} else {
				leap = false
			}
		} else {
			leap = true
		}
	}

	if leap {
		fmt.Printf("%d год - високосный год.\n", year)
	} else {
		fmt.Printf("%d год - не високосный год.\n", year)
	}
}

// Определение наибольшего из трех чисел
func task_2_3() {
	var a, b, c int

	fmt.Println("Введите три числа через пробел")
	fmt.Scan(&a, &b, &c)
	if a < b {
		a = b
	}
	if a > c {
		fmt.Printf("Наибольшее число из трех - %d.\n", a)
	} else {
		fmt.Printf("Наибольшее число из трех - %d.\n", c)
	}
}

// Категория возраста
func task_2_4() {
	var age int

	fmt.Println("Введите ваш возраст для определения категории")
	fmt.Println("[0 - 11]-ребенок, [12 - 19]-подросток, [20 - 60]-взрослый, [61 - 122]-пожилой")
	fmt.Scan(&age)
	if age < 0 {
		fmt.Println("Вы еще не рождены, возраст отрицательный")
		return
	}
	switch {
	case age <= 11:
		fmt.Println("ребенок")
	case age <= 19:
		fmt.Println("подросток")
	case age <= 60:
		fmt.Println("взрослый")
	case age <= 122:
		fmt.Println("пожилой")
	default:
		fmt.Println("Инопланетянин")
	}
}

// Проверка делимости на 3 и 5
func task_2_5() {
	var a int

	fmt.Println("Введите число, чтобы проверить делимостьна 3 и 5")
	fmt.Scan(&a)
	if a%15 == 0 {
		fmt.Println("Делится")
	} else {
		fmt.Println("Не делится")
	}
}

// Факториал числа
func task_3_1() {
	var a int

	fmt.Println("Введите число, чтобы посчитать его факториал")
	fmt.Scan(&a)
	result := 1
	for i := 1; i <= a; i++ {
		result *= i
	}
	fmt.Printf("Факториал = %d\n", result)
}

// Числа Фибоначчи
func task_3_2() {
	var n int
	fmt.Println("Введите количество чисел Фибоначчи, которые хотите увидеть")
	fmt.Scan(&n)
	fmt.Print(0)
	if n == 1 {
		return
	}
	fmt.Print(", 1")
	if n == 2 {
		return
	}
	a := 0
	b := 1
	var tmp int
	for i := 0; i < n-2; i++ {
		tmp = b
		b += a
		a = tmp
		fmt.Printf(", %d", b)
	}
	fmt.Println()
}

// Реверс массива
func task_3_3() {
	var strArr string
	fmt.Println("Введите массив элементов через ',' без пробелов:")
	fmt.Scanf("%s\n", &strArr)

	stringArray := strings.Split(strArr, ",")

	fmt.Println("Перевертыш:")
	for i := len(stringArray) - 1; i >= 0; i-- {
		fmt.Print(stringArray[i])
		if i != 0 {
			fmt.Print(",")
		} else {
			fmt.Print("\n")
		}
	}
}

// Поиск простых чисел
func task_3_4() {
	var n int
	fmt.Println("Введите число, чтобы найти простые числа до него")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		k := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				k++
			}
		}
		if k == 2 {
			fmt.Printf("%d, ", i)
		}
	}
	fmt.Println()
}

// Сумма чисел в массиве
func task_3_5() {
	var strArr string
	fmt.Println("Введите массив элементов через ',' без пробелов:")
	fmt.Scanf("%s\n", &strArr)

	stringArray := strings.Split(strArr, ",")

	sum := 0

	for i := 0; i < len(stringArray); i++ {
		sum += stringToInt(stringArray[i])
	}
	fmt.Printf("Сумма чисел в массиве=%d\n", sum)
}

func main() {
	// task_1_1()
	// task_1_2()
	// task_1_3()
	// task_1_4()
	// task_1_5()

	// task_2_1()
	// task_2_2()
	// task_2_3()
	// task_2_4()
	// task_2_5()

	// task_3_1()
	// task_3_2()
	// task_3_3()
	// task_3_4()
	task_3_5()

}
