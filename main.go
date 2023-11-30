// Написать программу, которая считывает из файла список математических выражений, считает результат и записывает в другой файл.
// (Только сложение и вычитание для положительных целых чисел от 0 до 99)

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func main() {

	var fileR, fileW string

	fmt.Println("Введите название входного файла:")
	_, err := fmt.Scanf("%s\n", &fileR) // Файл чтения выражений
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Введите название выходного файла для вывода результатов:")
	_, err = fmt.Scanf("%s\n", &fileW) // Файл записи результатов
	if err != nil {
		fmt.Println(err)
	}

	mathCalc(fileR, fileW)

}

// Функция считывает из файла список математических выражений, считает результат и записывает в другой файл
func mathCalc(fileR, fileW string) {

	re := regexp.MustCompile(`[0-9]{1,2}[\+,\-][0-9]{1,2}[=][?]`)

	b, err := ioutil.ReadFile(fileR) // Считываем содержимое файла
	if err != nil {
		panic(err)
	}

	testStr := string(b) // Преобразуем слайс байт в строку

	textLines := re.FindAllString(testStr, -1) // Находим все математические выражения и записываем в слайс строк

	w, err := os.Create(fileW) // Создаем файл для записи, если он не создан или переписываем старый
	if err != nil {
		panic(err)
	}

	defer w.Close() // Отложенный вызов функции закрытия файла

	buffer := bufio.NewWriter(w) // Создаем буфер для хранения строк с результатами

	// Для каждой найденной подстроки считаем результат математического выражения
	for j := 0; j < len(textLines); j++ {

		num1B := make([]byte, 0) // первое число
		num2B := make([]byte, 0) // второе число
		var znak, equal string   // знак сложения или вычитания и знак равенства
		var znakFlag bool        // флаг нахождения знака сложения или вычитания в выражении
		str := textLines[j]      // подстрока из файла (математическое выражение)

		for i := 0; i < len(textLines[j])-1; i++ {

			if znakFlag {

				if string(str[i]) != "=" {

					num2B = append(num2B, str[i])
				} else {
					equal = string(str[i])
				}

			}

			if !znakFlag {

				if string(str[i]) != "+" && string(str[i]) != "-" {

					num1B = append(num1B, str[i])
				} else {
					znak = string(str[i])
					znakFlag = true
				}
			}

		}

		num1, err := strconv.Atoi(string(num1B)) // Преобразование набора байт в целое число
		if err != nil {
			panic(err)
		}

		num2, err := strconv.Atoi(string(num2B)) // Преобразование набора байт в целое число
		if err != nil {
			panic(err)
		}

		var result int // результат выражения
		switch znak {
		case "+":
			result = num1 + num2

		case "-":
			result = num1 - num2
		}

		res := strconv.Itoa(result) // Преобразование результата в строку

		otv := string(num1B) + znak + string(num2B) + equal + res + "\n"

		buffer.WriteString(otv) // Запись резултата в буфер обмена
	}

	buffer.Flush() // Запись результата в файл

}
