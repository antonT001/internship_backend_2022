package helpers

import (
	"encoding/json"
	"net/http"
	c "user_balance/service/internal/constants"
)

func HttpResponse(w http.ResponseWriter, out interface{}, statusCode int) {
	result, err := json.Marshal(out)
	if err != nil {
		w.Write([]byte(c.SYSTEM_ERROR + ": " + err.Error()))
		return
	}
	w.WriteHeader(statusCode)

	w.Write(result)
}
