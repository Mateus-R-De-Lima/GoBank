package utils

import (
	"encoding/json"
	"fmt"
	"gobank/internal/validador"
	"net/http"
)

func CodificarJson[T any](w http.ResponseWriter, r *http.Request, statusCode int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("Falha ao codificar o json %w", err)
	}
	return nil
}

func DecodificarJson[T validador.Validador](r *http.Request) (T, map[string]string, error) {
	var data T

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, nil, fmt.Errorf("Erro ao decodificar o json %w", err)
	}

	if problemas := data.Valid(r.Context()); len(problemas) > 0 {
		return data, problemas, fmt.Errorf("Invalido %T:%d problemas", data, len(problemas))
	}

	return data, nil, nil
}
