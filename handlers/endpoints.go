package handlers

import (
	"net/http"
	"fmt"
	"github.com/ivanthescientist/tournament_service/dtos"
	"encoding/json"
)

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	fmt.Fprintln(response, "Tournament Application")
}

func FundHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();
	var queryMap = rawRequest.Form;

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

    var request = dtos.FundRequest {
		PlayerId: GetString(rawRequest, "playerId"),
		Points: GetInteger(rawRequest, "points"),
    }


    if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
    }

	// Actual work here
	fmt.Println(queryMap)

    response.WriteHeader(http.StatusOK)
}

func TakeHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();
	var queryMap = rawRequest.Form;

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var request = dtos.TakeRequest {
		PlayerId: GetString(rawRequest, "playerId"),
		Points: GetInteger(rawRequest, "points"),
	}


	if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actual work here
	fmt.Println(queryMap)

	response.WriteHeader(http.StatusOK)
}

func AnnounceTournamentHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();
	var queryMap = rawRequest.Form;

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var request = dtos.AnnounceTournamentRequest {
		TournamentId: GetString(rawRequest, "tournamentId"),
		Deposit: GetInteger(rawRequest, "deposit"),
	}


	if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actual work here
	fmt.Println(queryMap)

	response.WriteHeader(http.StatusOK)
}

func JoinTournamentHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();
	var queryMap = rawRequest.Form;

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var request = dtos.JoinTournamentRequest {
		PlayerId: GetString(rawRequest, "playerId"),
		BackerId: GetStringArray(rawRequest, "backerId"),
		TournamentId: GetString(rawRequest, "tournamentId"),
	}


	if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actual work here
	fmt.Println(queryMap)

	response.WriteHeader(http.StatusOK)
}

func ResultTournamentHandler(response http.ResponseWriter, rawRequest *http.Request) {
	var request = dtos.ResultTournamentRequest {}

	err := json.NewDecoder(rawRequest.Body).Decode(&request)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actual work here
	fmt.Println(request)

	response.WriteHeader(http.StatusOK)
}

func BalanceHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();
	var queryMap = rawRequest.Form;

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var request = dtos.PlayerBalanceRequest {
		PlayerId: GetString(rawRequest, "playerId"),
	}


	if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actual work here
	fmt.Println(queryMap)

	response.WriteHeader(http.StatusOK)
}

func ResetHandler(response http.ResponseWriter, request *http.Request) {

	// Actual work here
	response.WriteHeader(http.StatusOK)
}
