package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/marelinaa/currency-api/currency/internal/config"
	"github.com/marelinaa/currency-api/currency/internal/domain"
)

type Worker struct {
	currencyService     *CurrencyService
	externalURL         string
	RetrieveImmediately bool
	RetrieveTime        time.Time
}

func NewWorker(currencyService *CurrencyService, workerConfig config.WorkerConfig) *Worker {
	return &Worker{
		currencyService:     currencyService,
		externalURL:         workerConfig.ApiURL,
		RetrieveImmediately: workerConfig.RunFetchingOnStart,
		RetrieveTime:        workerConfig.RunTime,
	}
}

func (w *Worker) Start() {
	if w.RetrieveImmediately {
		go w.retrieveData()
	}

	go func() {
		for {
			nextRetrieve := w.calculateNextRetrieveTime(time.Now())
			time.Sleep(time.Until(nextRetrieve))

			go w.retrieveData()
		}
	}()
}

func (w *Worker) calculateNextRetrieveTime(currentTime time.Time) time.Time {
	nextRun := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), w.RetrieveTime.Hour(), w.RetrieveTime.Minute(), 0, 0, currentTime.Location())

	if currentTime.After(nextRun) {
		nextRun = nextRun.Add(24 * time.Hour)
	}

	return nextRun
}

func (w *Worker) retrieveData() {
	url := fmt.Sprintf("%s", w.externalURL)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	var response domain.CurrencyResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return
	}

	fmt.Println(response.Date)

	err = w.currencyService.SaveCurrencyData(response) //todo: check this
	if err != nil {
		return
	}
}
