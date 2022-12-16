# Описание схемы работы

## Сущности

* Письмо;
* Шаблон;
* Запрос на отправку письма.

## Компоненты

### Сервер
Сервер принимает запросы на отправку письма от клиентов и добавляет их в очередь сообщений.

### Обработчик
Обработчик отправляет письмо при появлении сообщения в очереди.

### Очередь сообщений
Очередь сообщений получает сообщения от сервера и передает их обработчику.

## Модули сервера
```mermaid
sequenceDiagram
    participant client
    participant API
    participant domain
    participant request savier
    participant queue
    client->>API: http request
    API->>domain: request to send email
    domain->>request savier: request to send email
    request savier->>queue: message
    queue->>request savier: saved!
    request savier->>domain: saved!
    domain->>API: saved!
    API->>client: http response OK
```

### Домен
Модуль реализует бизнес-логику сервера. Принимает на вход **запрос на отправку письма** и отправляет его в **модуль записи в очередь сообщений**.

### API
Модуль обрабатывает HTTP запрос и передает на обработку **домену**.

### Модуль записи в очередь сообщений
Модуль получает сообщение и сохраняет ее в очередь сообщений.

## Модули обработчика

```mermaid
sequenceDiagram
    participant queue
    participant listner
    participant domain
    participant template storager
    participant mail builder
    participant mail sender
    participant mail server
    queue->>listner: message
    listner->>domain: request to send email
    domain->>template storager: template ID
    template storager->>domain: template
    domain->>mail builder: template, params
    mail builder->>domain: mail
    domain->>mail sender: mail
    mail sender->>mail server: ...
    mail server->>mail sender: sended!
    mail sender->>domain: sended!
    domain->>listner: ok!
    listner->>queue: message acknowleged!
```
### Слушатель сообщений
Модуль ждет сообщения в очереди и передает их на обработку в **домен**.

### Домен
Модуль реализует бизнес-логику обработчика. Принимает на вход **запрос на отправку письма**. Получает **шаблон** из **хранилища шаблонов**. Отправляет **шаблон** и данных для подстановки в шаблон на вход **сборщика писем**, получает сформированное **письмо**. Передает его в **отправщик писем** для отправки.

### Хранилище шаблонов
Модуль реализует метод получения шаблона по его идентификатору. Опционально — постраничное получение, добавление, редактирование и удаление шаблонов.

### Сборщик писем
Модуль формирует письмо из шаблона и данных.

### Отправщик писем
Модуль отправляет письмо, реализуя один из протоколов, например SMTP.
