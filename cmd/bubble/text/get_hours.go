package text

import (
	"github.com/charmbracelet/bubbles/textinput"
	"log"
	"strconv"
)

func GetHours() float64 {
	ti := textinput.New()
	ti.Placeholder = "1"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 20

	taskModel := model{
		textInput:   ti,
		err:         nil,
		description: "Hours:\n\n%s\n",
	}

	hours, err := strconv.ParseFloat(getTextInput(taskModel), 64)
	if err != nil {
		log.Fatal("unable to parse hours: ")
	}

	return hours
}
