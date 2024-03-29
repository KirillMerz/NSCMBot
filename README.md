# NSCMBot

Telegram-бот для NSCM с автоматическим обновлением результатов экзаменов.

Доступен для использования: [NSCMBot](https://t.me/NSCMBot)

---

# Развёртывание

Также вы можете развернуть своего бота, используя этот проект.

Системные требования: 
- Процессор: только 64-битный
- Оперативная память: как минимум 4 ГБ

## 0. Загрузка проекта

- используя HTTPS:
```
git clone https://github.com/KirillMerz/NSCMBot.git && cd NSCMBot/
```

- используя SSH:
```
git clone git@github.com:KirillMerz/NSCMBot.git && cd NSCMBot/
```

В дальнейших пунктах работа ведётся из корня проекта

## 1. Конфигурирование

Для конфигурирования бота используется файл ```.env```.
Его пример уже присутствует в загруженном проекте, достаточно скопировать его под нужным именем:
```
cp .env.example .env
```

Далее вам необходимо указать следующие переменные:

- ```TELEGRAM_BOT_TOKEN``` - токен телеграм-бота. Можно получить у [BotFather](https://t.me/BotFather)
- ```MONGODB_URI``` - URI базы данных MongoDB. [Документация](https://mongodb.com/docs/manual/reference/connection-string/)

Следующие параметры необходимы лишь для работы Webhook'ов и являются необязательными. Если их не указать, бот будет использовать long polling

- ```WEBHOOK_URL``` - адрес, по которому будут отправляться запросы
- ```WEBHOOK_PORT``` - порт, который будет прослушивать бот

Если вы хотите, чтобы соединение между сервером Telegram и вашим ботом было безопасным, необходимо выпустить SSL-сертификат.
Сделать это можно с помощью сервиса [Let's Encrypt](https://letsencrypt.org/getting-started/). Выпущенный сертификат (а именно файлы ```fullchain.pem``` и ```privkey.pem```) необходимо разместить в каталоге ```./data/certs/```

## 2. Сборка и запуск

Для развёртывания существуют два пути:
[использовать контейнеризацию (2.1)](https://github.com/KirillMerz/NSCMBot#21-docker-compose)
или
[развернуть всё в системе напрямую (2.2)](https://github.com/KirillMerz/NSCMBot#22-локальная-сборка).

### 2.1 Docker Compose

Нужно установить [Docker Engine](https://docs.docker.com/engine/install/) и [Docker Compose](https://docs.docker.com/compose/install/).

При первом запуске все необходимые образы будут скачаны, выполнится сборка бота:
```
docker-compose up
```

### 2.2 Локальная сборка

Нужно установить [Go](https://go.dev/dl/) и [MongoDB](https://mongodb.com/docs/manual/installation/).

При первой сборке бота Go установит необходимые зависимости, а в каталоге ```./bin/``` появится исполняемый файл ```NSCMBot```:
```
go build -o ./bin/NSCMBot ./cmd/main.go
```
Запускаться бот должен из того же каталога, в котором находится файл ```.env```:
```
./bin/NSCMBot
```
