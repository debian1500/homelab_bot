# Homelab Startup Notification Bot 🚀

[English](<#English>) [Русский](<#Russian>)

### English
Telegram bot on Go for monitoring. 
The bot automatically sends a notification to the administrator at system startup if the server's uptime (time since startup) is less than 60 seconds.

## 🛠 Technology stack
- **Language:** Go 1.2x
- **Framework:** `gopkg.in/telebot.v3`
- **Configuration:** `github.com/joho/godotenv`
- **OS Target:** Linux (uses `/proc/uptime` parsing)

## 📋 How it works
When starting a binary or service:
1. The bot reads environment variables from the '.env` file.
2. Accesses the Linux virtual file system and parses the first line `/proc/uptime'.
3. If the system's operating time is less than 60 seconds, the bot sends a message "Server Started!" to the specified `USER_ID`.
4. Long Polling starts to process incoming commands (for example, `/start`).

## ⚙️ Deployment and configuration

### 1. Environment preparation
Create a `.env` file in the root directory of the project:
```env
BOT_TOKEN=your_telegram_bot_token_here
USER_ID=your_telegram_user_id_here
```
### 2. Build and launch
Compile the binary file:
`go build . -o homelab_bot main.go`

Launch the bot: `./homelab_bot`

#### 🔄 Autorun at boot (systemd)
To make the bot run automatically with the server, create a systemd service for it.

Create a configuration file for the service:
`sudo nano /etc/systemd/system/homelab-bot.service`

Add the following content (replace the paths with your own):

```
[Unit]
Description=Homelab Telegram Notification Bot
After=network.target

[Service]
Type=simple
User="$USER"
WorkingDirectory="$HOME/homelab_bot"
ExecStart="$HOME/homelab_bot/homelab_bot"
Restart=on-failure

[Install]
WantedBy=multi-user.target
```
Restart the systemd daemon, turn on autorun, and run the service:
```bash
sudo systemctl daemon-reload
sudo systemctl enable homelab-bot.service
sudo systemctl start homelab-bot.service
```
#### Development Plans (Roadmap)
- Adding the `/poweroff` command to shut down the server
- System status output using `/proc` (`/proc/meminfo` for RAM, `/proc/loadavg` for medium load, etc.)



---
### Russian
Telegram-бот на Go для мониторинга. 
Бот автоматически отправляет уведомление администратору при старте системы, если аптайм (время с момента запуска) сервера составляет менее 60 секунд.

## 🛠 Технологический стек
- **Language:** Go 1.2x
- **Framework:** `gopkg.in/telebot.v3`
- **Configuration:** `github.com/joho/godotenv`
- **OS Target:** Linux (использует парсинг `/proc/uptime`)

## 📋 Как это работает
При запуске бинарника или службы:
1. Бот считывает переменные окружения из файла `.env`.
2. Обращается к виртуальной файловой системе Linux и парсит первую строку `/proc/uptime`.
3. Если значение времени работы системы `≤ 60` секунд, бот отправляет сообщение "Сервер запущен!" на указанный `USER_ID`.
4. Запускается Long Polling для обработки входящих команд (например, `/start`).

## ⚙️ Развёртывание и настройка

### 1. Подготовка окружения
Создайте файл `.env` в корневой директории проекта:
```env
BOT_TOKEN=your_telegram_bot_token_here
USER_ID=your_telegram_user_id_here
```
### 2. Сборка и запуск
Скомпилируйте бинарный файл:
`go build . -o homelab_bot main.go`

Запустите бота: `./homelab_bot`

#### 🔄 Автозапуск при загрузке (systemd)
Чтобы бот запускался автоматически с сервером, создайте для него службу systemd.

Создайте конфигурационный файл службы:
`sudo nano /etc/systemd/system/homelab-bot.service`

Добавьте следующее содержимое (замените пути на свои):

```
[Unit]
Description=Homelab Telegram Notification Bot
After=network.target

[Service]
Type=simple
User="$USER"
WorkingDirectory="$HOME/homelab_bot"
ExecStart="$HOME/homelab_bot/homelab_bot"
Restart=on-failure

[Install]
WantedBy=multi-user.target
```
Перезапустите демон systemd, включите автозапуск и запустите службу:
```bash
sudo systemctl daemon-reload
sudo systemctl enable homelab-bot.service
sudo systemctl start homelab-bot.service
```
#### Планы по развитию (Roadmap)
- Добавление команды `/poweroff` для выключения сервера
- Вывод статуса системы с использованием `/proc` (`/proc/meminfo` для RAM, `/proc/loadavg` для средней нагрузки и т.п.)



