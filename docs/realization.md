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
    participant mail sender
    participant request savier
    participant queue
    client->>API: http request
    API->>mail sender: request to send email
    mail sender->>request savier: request to send email
    request savier->>queue: message
    queue->>request savier: saved!
    request savier->>mail sender: saved!
    mail sender->>API: saved!
    API->>client: http response OK
```

### Отправщик писем
Модуль реализует бизнес-логику сервера. Принимает на вход **запрос на отправку письма** и отправляет его в **модуль записи в очередь сообщений**.

### API
Модуль обрабатывает HTTP запрос и передает на обработку **отправщику писем**.

### Модуль записи в очередь сообщений
Модуль получает сообщение и сохраняет ее в очередь сообщений.

## Модули обработчика

Пример успешной обработки сообщения:

```mermaid
sequenceDiagram
    participant queue
    participant listner
    participant request processor
    participant template storager
    participant mail builder
    participant mail sender
    participant mail server
    queue->>listner: message
    listner->>request processor: request to send email
    request processor->>template storager: template ID
    template storager->>request processor: template
    request processor->>mail builder: template, params
    mail builder->>request processor: mail
    request processor->>mail sender: mail
    mail sender->>mail server: ...
    mail server->>mail sender: sended!
    mail sender->>request processor: sended!
    request processor->>listner: ok!
    listner->>queue: message acknowleged!
```

В случае ошибки обработки сообщения счетчик оставшихся попыток уменьшается, сообщение попадает в очередь queueDLX на время delay, после чего снова возвращается в очередь queue:

```mermaid
sequenceDiagram
    participant queue
    participant queueDLX
    participant listner
    participant request processor
    participant template storager
    participant mail builder
    participant mail sender
    participant mail server
    queue->>listner: message, retries: n
    listner->>request processor: request to send email
    request processor->>template storager: template ID
    template storager->>request processor: template
    request processor->>mail builder: template, params
    mail builder->>request processor: mail
    request processor->>mail sender: mail
    mail sender->>mail server: ...
    mail server->>mail sender: error!
    mail sender->>request processor: error!
    request processor->>listner: error!
    listner->>queueDLX: message, retries: n-1
    queueDLX->>queue: message, retries: n-1
```

### Слушатель сообщений
Модуль ждет сообщения в очереди и передает их на обработку в **обработчик запросов отправки писем**.

### Обработчик запросов отправки писем
Модуль реализует бизнес-логику обработчика. Принимает на вход **запрос на отправку письма**. Получает **шаблон** из **хранилища шаблонов**. Отправляет **шаблон** и данных для подстановки в шаблон на вход **сборщика писем**, получает сформированное **письмо**. Передает его в **отправщик писем** для отправки.

### Хранилище шаблонов
Модуль реализует метод получения шаблона по его идентификатору. Опционально — постраничное получение, добавление, редактирование и удаление шаблонов.

### Сборщик писем
Модуль формирует письмо из шаблона и данных.

### Отправщик писем
Модуль отправляет письмо, реализуя один из протоколов, например SMTP.
