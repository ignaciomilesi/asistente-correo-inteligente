package ollama

import (
	"encoding/json"
	"fmt"
	"os"
)

func cargarConfig() (OllamaConfig, error) {

	config := OllamaConfig{}

	// Abrir archivo
	archivo, err := os.ReadFile("config\\config.json")
	if err != nil {
		return config, fmt.Errorf("error al leer el archivo, %v", err)
	}

	// Parsear JSON a struct
	err = json.Unmarshal(archivo, &config)
	if err != nil {
		return config, fmt.Errorf("error al parsear el json, %v", err)
	}

	return config, nil
}
