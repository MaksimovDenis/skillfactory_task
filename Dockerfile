# Указываем версию golang для большей предсказуемости и стабильности сборки
FROM golang:1.22 AS compiling_stage

# Создаем и устанавливаем рабочую директорию
WORKDIR /skillfactory_task

# Копируем файлы go.mod и go.sum перед кодом приложения для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код приложения
COPY . .

# Сборка приложения
RUN go build -o skillfactory_task ./cmd/main.go

# Используем минимальный образ для запуска контейнера
FROM alpine:latest

# Устанавливаем метаданные
LABEL version="1.0.0"

# Устанавливаем рабочую директорию
WORKDIR /root/

# Установите libc6-compat для совместимости с некоторыми Go бинарниками
RUN apk add --no-cache libc6-compat

# Копируем бинарный файл из предыдущего этапа
COPY --from=compiling_stage /skillfactory_task .

RUN chmod +x /root/skillfactory_task

# Указываем точку входа
ENTRYPOINT ["./skillfactory_task"]

# Открываем порт 8000
EXPOSE 8000