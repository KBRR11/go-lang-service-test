package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Question struct {
	QuestionID string   `json:"question_id"`
	Question   string   `json:"question"`
	Answers    []Answer `json:"answers"`
}

type Answer struct {
	AnswerID string `json:"answer_id"`
	Answer   string `json:"answer"`
}

type QuestionsResponse struct {
	Date string     `json:"date"`
	Data []Question `json:"data"`
}

type ReestructuredResponse struct {
	Titulo     string           `json:"titulo"`
	Dia        string           `json:"dia"`
	Info       []QuestionDetail `json:"info"`
	APIVersion int              `json:"api_version"`
}

type QuestionDetail struct {
	PreguntaID int    `json:"pregunta_id"`
	Pregunta   string `json:"pregunta"`
}

func main() {
	http.HandleFunc("/get-questions", func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{Timeout: 10 * time.Second}
		req, err := http.NewRequest("GET", "https://us-central1-teamcore-retail.cloudfunctions.net/test_mobile/api/questions", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2NzM0NzU4MTEsImV4cCI6MTcwNTAxMTgxMSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.9wqriO_2Q8Xfwc9VcgMpr-2c4WVdLRJ5G6NcNaXdpuY")
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

		var questionsResponse QuestionsResponse
		err = json.Unmarshal(body, &questionsResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var reestructuredResponse ReestructuredResponse
		reestructuredResponse.Titulo = "Preguntas del día"
		reestructuredResponse.Dia = time.Now().Format("02-01-2006")

		for _, question := range questionsResponse.Data {
			var questionDetail QuestionDetail
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
	})

	http.ListenAndServe(":8080", nil)
}
