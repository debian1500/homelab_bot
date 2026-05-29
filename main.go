package main

import (
	"log"     // Для обработки ошибок
	"os"      // Для парсинга /proc/uptime (os.FileRead)
	"strconv" // Для преобразования секунд в float
	"strings" // Для преобразования прочитанного файла в человеко-читаемый вид
	"time"    // Для константы time.Second

	tele "gopkg.in/telebot.v3"
	"github.com/joho/godotenv" // Для работы с переменными

)

func main() {


	godotenv.Load()

	// Токен бота из переменной (env)
	BotToken := os.Getenv("BOT_TOKEN")
	if BotToken == "" {
		log.Fatal("Ошибка: Переменная BOT_TOKEN не найдена!")
	}

	// userID из переменной (env) 
	UserIDRaw := os.Getenv("USER_ID")
        if UserIDRaw == "" {
                log.Fatal("Ошибка: Переменная USER_ID не найдена!")
	}


	// Преобразование userID в int64

	UserID, err := strconv.ParseInt(UserIDRaw, 10, 64)
	if err != nil {
		log.Fatalf("Ошибка: не удалось преобразовать USER_ID в число: %v", err)
	}


	pref := tele.Settings{
		Token:  BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)

	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %v", err)
	}

	// Проверка времени с момента запуска

	uptimeLocation := "/proc/uptime"              // Выносим расположение файла в переменную
	rawUptime, err := os.ReadFile(uptimeLocation) // Читаем файл
	if err != nil {                               // Если есть ошибка - выводим сообщение (букв. Если ошибка не пуста)
		log.Printf("Не удалось прочитать файл /proc/uptime: %v", err)
		return
	}

	uptimeSlice := strings.Fields(string(rawUptime))
	/* Преобразования:
	1. string - преобразуем в строку (без этого при чтении отображается слайс из ASCII-байтов)
	2. strings.Fields - преобразуем в слайс (аналог массива - list в Python)
	*/

	uptimeSeconds, err := strconv.ParseFloat(uptimeSlice[0], 64)
	/*
		1. Берём только первую строку из слайса (индекс 0)
		В /proc/uptime первая строка - время с момента запуска.
		Вторая нам не нужна - сумма времени запуска для всех ядер ЦП.

		2. Преобразуем первую строку в float64 (обычно значение в виде десятичной дроби, а не целого числа)
	*/

	if uptimeSeconds <= 60 {
		bot.Send(tele.ChatID(UserID), "Сервер запущен!")
	}

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Привет! Бот работает.")
	})

	bot.Start()
}
