package main

import (
	"bufio"   // Парсинг
	"log"     // Обработка ошибок
	"os"      // Чтение файла (os.Open)
	"strings" // Преобразование байтов в читаемые строки
)

// Главная функция этого файла, которую вызывает main.go по команде /stats в боте
func GetStatsReport() string {
	meminfoLocation := "/proc/meminfo"

	// Читаем файл
	file, err := os.Open(meminfoLocation)

	// Обработка ошибки
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	// ------------- //

	// Создаём объект scanner, который будет работать с прочитанным meminfo
	scanner := bufio.NewScanner(file)

	// Переменная, в которую будем добавлять отфильтрованные строки
	var memInfoCollected strings.Builder

	// Запускаем цикл со Scanner
	for scanner.Scan() {

		// Переменная с содержимым файла
		line := scanner.Text()

		// Проходимся циклом по массиву (слайсу) с искомыми строками
		for _, memInfoString := range []string{"MemTotal", "MemFree", "MemAvailable"} {

			// Ищем, существует ли искомая строка
			if strings.Contains(line, memInfoString) {

				// Добавляем найденную строку + символ переноса строки (аналог append в Python)
				memInfoCollected.WriteString(line + "\n")
			}

		}
	}

	// Обработка ошибки scanner'а
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return "Ошибка парсинга /proc/meminfo"
	}

	return memInfoCollected.String() // Преобразуем переменную из сырых байтов в строку и возвращаем

}
