package service

import (
	"encoding/json"
	"fmt"
	"go/internal/entity"
	"go/settings"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetQuestions(w http.ResponseWriter, r *http.Request, s *settings.Settings) {
	token := fmt.Sprintf("Bearer %s", s.External.Token)
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", s.External.Url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("unexpected status code %d", resp.StatusCode), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var questionsResponse entity.QuestionsResponse
	err = json.Unmarshal(body, &questionsResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var reestructuredResponse entity.ReestructuredResponse
	reestructuredResponse.Titulo = "Preguntas del día"
	reestructuredResponse.Dia = time.Now().Format("02-01-2006")

	for _, question := range questionsResponse.Data {
		var questionDetail entity.QuestionDetail
		questionDetail.PreguntaID, err = strconv.Atoi(question.QuestionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		questionDetail.Pregunta = question.Question
		reestructuredResponse.Info = append(reestructuredResponse.Info, questionDetail)
	}

	reestructuredResponse.APIVersion = 1

	response, err := json.Marshal(reestructuredResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
