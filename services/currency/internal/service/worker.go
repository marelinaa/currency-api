package service

import (
	"encoding/json"
	"fmt"
	"log"
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
	log.Println("Start worker")
	if w.RetrieveImmediately {
		log.Println("Retrieve immediately")

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

	log.Printf("Next retrieve is scheduled for: %s", nextRun)

	if currentTime.After(nextRun) {
		nextRun = nextRun.Add(24 * time.Hour)
	}

	return nextRun
}

func (w *Worker) retrieveData() {
	url := fmt.Sprintf("%s", w.externalURL)
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)

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

	log.Println(response.Date)

	err = w.currencyService.SaveCurrencyData(response)
	if err != nil {
		return
	}

	log.Println("Fetched data and saved")
}

//query
