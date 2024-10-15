package service

import (
	"encoding/json"
	"fmt"
	"github.com/marelinaa/currency-api/services/currency/internal/config"
	"github.com/marelinaa/currency-api/services/currency/internal/domain"
	"net/http"
	"time"
)

type Worker struct {
	currencyService *CurrencyService
	externalURL     string
	RunImmediately  bool
	RunTime         time.Time
}

func NewWorker(currencyService *CurrencyService, workerConfig config.WorkerConfig) *Worker {
	return &Worker{
		currencyService: currencyService,
		externalURL:     workerConfig.ApiURL,
		RunImmediately:  workerConfig.RunFetchingOnStart,
		RunTime:         workerConfig.RunTime,
	}
}

func (w *Worker) Start() {
	if w.RunImmediately {
		go w.fetchData()
	}

	go func() {
		for {
			nextRun := w.calculateNextRunTime(time.Now())
			time.Sleep(time.Until(nextRun))

			go w.fetchData()
		}
	}()
}

func (w *Worker) calculateNextRunTime(currentTime time.Time) time.Time {
	nextRun := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), w.RunTime.Hour(), w.RunTime.Minute(), 0, 0, currentTime.Location())

	if currentTime.After(nextRun) {
		nextRun = nextRun.Add(24 * time.Hour)
	}

	return nextRun
}

func (w *Worker) fetchData() {
	url := fmt.Sprintf("%s", w.externalURL)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	var currencyData domain.CurrencyData
	if err := json.NewDecoder(resp.Body).Decode(&currencyData); err != nil {
		return
	}

	err = w.currencyService.SaveCurrencyData(currencyData) //todo: check this
	if err != nil {
		return
	}
}
