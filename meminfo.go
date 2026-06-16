package main

import (
	"log" // Обработка ошибок
	"os"  // Для работы с файлами (os.FileRead)
	// Для преобразования байтовых строк в человекочитаемые
)

func GetStatsReport() string {
	meminfoLocation := "/proc/meminfo" // Расположение читаемого файла в виде переменной
	// cpuinfoLocation := "/proc/cpuinfo" - позже добавить и другие файлы

	file, err := os.ReadFile(meminfoLocation) // Читаем файл
	if err != nil {
		log.Printf("Ошибка чтения файла: %s", err) // Выводим ошибку
	}

	return string(file)
}
