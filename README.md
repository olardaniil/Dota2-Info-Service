# Dota2-Info-Service
> Тестовое задание на позицию "Cтажёр IT направления" в МТС

## Используемые технологии
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)

## Реализованные требования к проекту
Сервис имеет 2 апи-метода:
- `[get]../info/{hero}` - выдаёт информацию по заданному герою.
- `[get]../сounter/{hero}` - выдаёт информацию о персонажах, против которых заданный герой играет плохо.

### Другое
* Предоставлена спецификация на API в формате Swagger 2.0.
* Unit тесты. Создан файл с визуализацией метрик покрытия кода. `test_coverage_report.html`
* Dockerfile для сборки образа.
* docker-compose для запуска окружения с работающим приложением.

## Локальный запуск проекта

> 1. Запуск проекта на локальном сервере производиться командой 
```
docker-compose up -d
```
> 2. Swagger на локальном сервере доступен по URL:
```
http://127.0.0.1:8080/swagger/
```
