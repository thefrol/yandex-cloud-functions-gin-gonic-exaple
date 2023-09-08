# Пример создания Go-сервера на gin-gonic в Яндекс.Облаке(Функция + Api Gateway)

В этом примере я рассказываю, как настроить работу gin-gonic(сервера на языке go) в яндекс облаке. Это работа исследование.

Поможет вам начать разработку собственного сервека, используйте этот репозиторий как шаблон. Тут вы встретите основые способы использования. Поддерживает локальный запуск `go run .` для отладки

## Локальный запуск

Запустим сервер локально, и он станет доступен по адресу `http://localhost:8080`

```bash
go run .
```

Отправим ему запрос

```bash
$ curl curl --data "{\"name\":\"dima\"}" "http://localhost:8080/stuff"
```

и если все нормально, то получаем такой ответ. 

```bash
{"user":"dima","value":"its working"}
```

> Все команды запускаются в терминале


## Шаг 1. Загрузка в облако

Создайте функцию и загрузите в облако, например одним из способов в [официальном гайде](https://cloud.yandex.ru/docs/functions/quickstart/create-function/go-function-quickstart) или воспользуйтесь [powershell скриптом](https://github.com/thefrol/powershell-yandexcloud-function-uploader)

## Шаг 2. Настройки функции

Убедитесь что в поле `точка входа` установлено `main.Handler`

## Шаг N-1. Создание Api-Gateway

ссылка на гайд

## Шаг N. Настройка Api-Gateway

Просто функция не обладает способностью к маршрутизации. Если попробовать вызвать что-то вроде `https://functions.yandexcloud.net/<id функции>/stuff/23`, то мы увидим ошибку 

    {"errorCode":400,"errorMessage":"Invalid functionID: /<id функции>/stuff/23","errorType":"ProxyIntegrationError"}

Он считает, что `/stuff/23` это часть id функции. Много раз меня останавливала эта часть, но решение проблемы есть - достаточно создать `Api Gateway` в том же каталоге, с вот такой спецификацией

```swagger
openapi: 3.0.0
info:
  title: Sample API
  version: 1.0.0
paths:
  /{url+}:
    x-yc-apigateway-any-method:
      parameters:
      - explode: false
        in: path
        name: url
        required: false
        style: simple
      x-yc-apigateway-integration:
        function_id: d4ebpjlgj00sqsg8g2n4
        tag: $latest
        type: cloud_functions
```

важные параметры:
+ `function_id: <id функции>` - тут нужно указать айди функции
+ (Опционально) `service_account: <id аккаунта>` - если нужно чтобы функция запускалась с использованием аккаунта сервисного 

Можно так же настроить сервисный аккаунт с которым она будет запускаться

# Что дальше
В дальнейшем хотелось бы сделать утилиту, которая во-первых создает шаблон для `gin` сервера, и умеет его загружать в облако и самостоятельно конфигурировать api gateway и все прочее


# TODO

[ ] Добавить фотографий, например, где настройки
