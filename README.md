# PKK

Проект предназначени для извлечения и храниния данных о земельных участках

*Возможные функции:* 
- Загразка кадастрового плана территории
- Получение кадастрового плана территории по кадастровому кварталу
- Получение координатного описания земельного участка по кадастровому номеру

## Окружение

### Требования 
- Docker 20.x.x
- Docker Compose 1.29.x
- Go 1.16+

### Переменные окружения

| Параметр                   | Описание                              | Default          |
|----------------------------|---------------------------------------|------------------|
| `ENV`                      | Режим окружения                       | `development`    |
| `X_API_KEY`                | Ключ API                              |                  |
| `HTTP_SERVER_PORT`         | Порт HTTP сервера                     | `8080`           |
| `HTTP_SERVER_NAME`         | Имя HTTP сервера                      | `pkk_app_web`    |
| `HTTP_SERVER_TIMEOUT`      | Тайм-аут HTTP сервера                 | `4s`             |
| `HTTP_SERVER_IDLE_TIMEOUT` | Idle тайм-аут HTTP сервера            | `60s`            |
| `DB_HOST`                  | Хост базы данных                      | `localhost`      |
| `DB_PORT`                  | Порт базы данных                      | `5432`           |
| `DB_NAME`                  | Имя базы данных                       | `pkk_db`         |
| `DB_USER_NAME`             | Имя пользователя базы данных          | `user`           |
| `DB_PASSWORD`              | Пароль пользователя базы данных       | `secret`         |
| `DB_SSL_MODE`              | Режим SSL для базы данных             | `disable`        |
| `DB_DRIVER_NAME`           | Имя драйвера базы данных              | `postgres`       |
| `DB_MAX_CONNS`             | Максимальное количество соединений БД | `20`             |
| `MINIO_NAME`               | Имя Minio                             | `minio`          |
| `MINIO_ACCESS_KEY`         | Ключ доступа Minio                    | `minioPkk`       |
| `MINIO_SECRET_KEY`         | Секретный ключ Minio                  | `secretMinioKey` |
| `MINIO_PORT`               | Порт Minio                            | `9001`           |
| `MINIO_WEB_PORT`           | Веб-порт Minio                        | `9002`           |
| `MINIO_ENDPOINT`           | Эндпоинт Minio                        | `localhost:9001` |
| `MINIO_REGION`             | Регион Minio                          | `eu-central-1`   |
| `MINIO_BUCKET`             | Название корзины Minio                | `pkk`            |
| `RETRY_ATTEMPTS`           | Количество попыток повтора            | `5`              |
| `RETRY_DELAY`              | Задержка между попытками              | `1s`             |
| `IP_INFO_TOKEN`            | Токен IPInfo                          |                  |
| `IP_INFO_LOCAL_MACHINE`    | Локальная машина IPInfo               | true             |
| `IP_INFO_COUNTRY`          | Страна IPInfo                         | RU               |

## Установка

**PROJECT**

- Создать новую директорию для проекта. В консоли перейти в созданную директорию и написать: git clone https://github.com/ShevelevEvgeniy/pkk.git

**DOCKER**

*Сборка:*

Скопировать файл .env.dist и переименовать в .env, настроить параметры окружения cp .env.dist .env
Для развертывания, запустите установку, выполнив команду ниже: make install

*Служебное:*
- make migrate-up - Запуск миграций 
- make migrate-down - Откат миграций
- make migrate-create name="$" - Создание новой миграции 

Если у вас возникли вопросы или проблемы, вы можете связаться со мной по адресу Z_shevelev@mail.ru