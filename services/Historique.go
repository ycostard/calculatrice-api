package services

import (
	"encoding/json"
	"os"
	"sync"
)

const historyFile = "historique.json"

type HistoryEntry struct {
	Expression string  `json:"expression"`
	Result     float64 `json:"result"`
}

var mu sync.Mutex

func LoadHistory() ([]HistoryEntry, error) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(historyFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var history []HistoryEntry
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&history)
	if err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return history, nil
}

func SaveHistory(history []HistoryEntry) error {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(historyFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(history)
}
