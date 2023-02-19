package main

import (
	"context"
	"fmt"
	"go/internal/service"
	"go/settings"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
		),
		fx.Invoke(
			func(s *settings.Settings) {
				port := fmt.Sprintf(":%s", s.Port)
				http.HandleFunc("/get-questions", func(w http.ResponseWriter, r *http.Request) {
					service.GetQuestions(w, r, s)
				})
				http.ListenAndServe((port), nil)
			},
		),
	)
	app.Run()
}
