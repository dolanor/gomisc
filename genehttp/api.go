package main

import (
	"strings"
)

type UppercaseRequest struct {
	Text string
}

type UppercaseResponse struct {
	Text string
}

func UpperCase(req UppercaseRequest) (UppercaseResponse, error) {
	upped := strings.ToUpper(req.Text)
	return UppercaseResponse{
		Text: upped,
	}, nil
}

type SnakeCaseRequest struct {
	Text string
}

type SnakeCaseResponse struct {
	Text string
}

func SnakeCase(req SnakeCaseRequest) (SnakeCaseResponse, error) {
	snaked := strings.ReplaceAll(req.Text, " ", "_")
	return SnakeCaseResponse{
		Text: snaked,
	}, nil
}
