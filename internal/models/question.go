package models

type Question struct {
	Key string
	
	Text string
	Options []string
	Next *Question
}

type Answers map[string]any