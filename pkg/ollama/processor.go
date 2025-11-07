package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ollamaConfig struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type Processor struct {
	ollamaRequestTemplate ollamaConfig
	key                   string
}

func (p *Processor) Iniciar() error {
	if err := p.cargarConfig(); err != nil {
		return fmt.Errorf("error al cargar config, %v", err)
	}

	if err := p.cargarKey(); err != nil {
		return fmt.Errorf("error al cargar la key de acceso, %v", err)
	}

	fmt.Println("Processor iniciado")
	return nil
}

func (p Processor) Resumir(text string) (string, error) {
	client := &http.Client{}

	reqBody := p.ollamaRequestTemplate

	reqBody.Prompt += "\n \n" + text // agregamos el texto a resumir al prompt

	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "https://ollama.com/api/generate", bytes.NewBuffer(body))
	//req, err := http.NewRequest("POST", "http://localhost:11434/api/generate", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+p.key)
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

func (p *Processor) cargarConfig() error {

	// Abrir archivo
	archivo, err := os.ReadFile("config.json")
	if err != nil {
		return fmt.Errorf("error al leer el archivo, %v", err)
	}

	// Parsear JSON a struct
	err = json.Unmarshal(archivo, &p.ollamaRequestTemplate)
	if err != nil {
		return fmt.Errorf("error al parsear el json, %v", err)
	}

	return nil
}

func (p *Processor) cargarKey() error {
	data, _ := os.ReadFile("key.json")

	var result map[string]string
	if err := json.Unmarshal(data, &result); err != nil {
		return fmt.Errorf("error al parsear el json, %v", err)
	}

	p.key = result["key"]

	return nil
}
