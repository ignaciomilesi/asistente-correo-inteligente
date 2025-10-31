package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type OllamaConfig struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

func Resumir(text string) (string, error) {
	client := &http.Client{}

	reqBody, err := cargarConfig()
	if err != nil {
		return "", fmt.Errorf("No se pudo cargar config, %s", err.Error())
	}

	reqBody.Prompt += "\n \n" + text // agregamos el texto a resumir al prompt

	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "http://localhost:11434/api/generate", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("no se pudo conectar con Ollama: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Error al decodificar JSON: %v", err)
	}

	return result["response"].(string), nil
}
