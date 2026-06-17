# Домашнее задание №14 (глава 3.5)

Цель: Написать тесты для простого HTTP-сервера

Задание:
Реализуйте HTTP-сервер с одним эндпоинтом:
GET /hello – возвращает {"message": "hello, world!"} с кодом 200 OK.
Пример сервера (можно использовать свой или модифицировать этот):
package main

import (
"encoding/json"
"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(map[string]string{"message": "hello, world!"})
}

func main() {
http.HandleFunc("/hello", helloHandler)
http.ListenAndServe(":8080", nil)
}
Напишите тесты для этого сервера, проверяющие:
- Корректность кода ответа (200 OK).
- Корректность заголовка Content-Type (application/json).
- Корректность тела ответа (должен быть JSON с полем message).

Необходимо выложить программу на github и отправить ссылку для просмотра или отправить файл программы.