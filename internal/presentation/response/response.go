package response

import (
	"encoding/json"
	"net/http"
)

// JSON - функция, необходимая для формирования HTTP ответа в формате JSON.
//
// В качестве параметров получает:
// - w, куда записывается ответ
// - status, HTTP статус, с которым возвращается ответ
// - payload, тело ответа, приходит ввиде любого тип данных 
func JSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/JSON")
	w.WriteHeader(status)

	buf, _ := json.Marshal(payload)
	w.Write(buf)
}

// Error - функция, необходимая для отправки ошибки в формате JSON по HTTP.
// Параметры получает те же, что и функция JSON
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, message)
}