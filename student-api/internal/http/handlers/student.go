package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/AlNomanCSE/learngo/internal/types"
	respon "github.com/AlNomanCSE/learngo/internal/utils"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			if errors.Is(err, io.EOF) {
				respon.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "empty request body"})
				return
			}
			respon.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}
		// Validate required fields
		if student.Name == "" || student.Email == "" || student.Age == 0 {
			respon.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "missing required fields: name, email, and age are required"})
			return
		}

		slog.Info("Creating a student", "name", student.Name)
		respon.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
