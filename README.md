Этот проект представляет собой веб-сервис, который принимает арифметическое выражение через HTTP-запрос и возвращает результат вычисления.

Использование
Отправьте POST-запрос на `/api/v1/calculate` со следующим телом:
{
  "expression": "ваше арифметическое выражение"
}


Запуск проекта

Шаг 1: Клонирование репозитория
Откройте VS Code.
Откройте терминал в VS Code, нажав Ctrl+``.
В терминале выполните команду для клонирования репозитория:
git clone https://github.com/vonavitya/go_calc.git
Перейдите в директорию проекта:
cd go_calc

Шаг 2: Настройка окружения
В VS Code откройте директорию проекта (если она еще не открыта):
code .
Убедитесь, что все зависимости установлены. В терминале выполните команду:
go mod tidy

Шаг 3: Запуск сервера
В терминале в корневой директории проекта выполните команду для запуска сервера:
go run ./cmd/calc_service/...
Если появится запрос на разрешение запуска, дайте разрешение. Сервер будет запущен и будет слушать на порту 8080.

Шаг 4: Тестирование API
Откройте новый терминал в VS Code (Terminal -> New Terminal).
Отправьте запрос к API с помощью curl:
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'
Вы должны увидеть ответ от сервера с результатом вычислений:
{
  "result": "6.000000"
}




Примеры использования с curl

Успешное выполнение:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
Ответ:
{
  "result": "6.000000"
}

Ошибка 422 (некорректное выражение):
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*"
}'
Ответ:
{
  "error": "Expression is not valid"
}

Ошибка 500 (внутренняя ошибка сервера): В данном случае, чтобы симулировать внутреннюю ошибку сервера, вам нужно модифицировать код. Например, вы можете искусственно вызвать панику в функции обработки:
// В main.go добавьте функцию, вызывающую панику
func panicExampleHandler(w http.ResponseWriter, r *http.Request) {
    panic("internal server error example")
}
// И добавьте новый route в main
http.HandleFunc("/api/v1/panic", panicExampleHandler)
Затем отправьте запрос:
curl --location 'http://localhost:8080/api/v1/panic'
Ответ:
{
  "error": "Internal server error"
}
