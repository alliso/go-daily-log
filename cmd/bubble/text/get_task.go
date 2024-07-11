package text

import (
	"github.com/charmbracelet/bubbles/textinput"
)

var value string

type (
	errMsg error
)

func GetTask() string {
	ti := textinput.New()
	ti.Placeholder = "Celeritas event"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 20

	taskModel := model{
		textInput:   ti,
		err:         nil,
		description: "Task description:\n\n%s\n",
	}
	return getTextInput(taskModel)
}
