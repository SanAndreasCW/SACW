package auth

import (
	"github.com/RahRow/omp"
	"strings"
)

func SplitMessage(input string, limit int) *[]string {
	var messages []string
	words := strings.Fields(input)
	var currentMessage string
	for _, word := range words {
		if len(currentMessage)+len(word)+1 > limit {
			messages = append(messages, currentMessage)
			currentMessage = word
		} else {
			if currentMessage != "" {
				currentMessage += " "
			}
			currentMessage += word
		}
	}
	if currentMessage != "" {
		messages = append(messages, currentMessage)
	}
	return &messages
}

func SendClientMessageToAll(message string, color omp.Color) {
	messages := SplitMessage(message, 143)
	go func() {
		for _, player := range PlayersI {
			for _, msg := range *messages {
				player.SendClientMessage(msg, color)
			}
		}
	}()
}
