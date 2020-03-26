package handler

import (
	"log"
	"net/http"

	"github.com/crontb/random/pkg/config"
	"github.com/crontb/random/pkg/http/response"
	"github.com/crontb/random/pkg/service/card"
)

// Card HTTP GET /card endpoint handler
func Card(w http.ResponseWriter, r *http.Request) {
	res := &response.JSON{}
	numbers, err := card.GenerateNumbers(config.CardTotalNumbers, config.CardMaxNumber)
	if err != nil {
		message := "Cannot generate card numbers: " + err.Error()
		log.Println(message)
		res.SetError(http.StatusBadRequest, message)
	}

	if numbers != nil {
		res.SetResult(numbers)
	}
	res.Write(w)
}
