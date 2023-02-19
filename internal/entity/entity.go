package entity

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
