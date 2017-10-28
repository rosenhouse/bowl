package application

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

//go:generate counterfeiter -o fakes/command_new.go --fake-name CommandNew . commandNew
type commandNew interface {
	Run() (blankBoard string)
}

//go:generate counterfeiter -o fakes/command_score.go --fake-name CommandScore . commandScore
type commandScore interface {
	Run(userRecord string) (scoredBoard string, err error)
}

type App struct {
	Input     io.Reader
	Output    io.Writer
	Arguments []string

	CommandNew   commandNew
	CommandScore commandScore
}

var ErrUsage = errors.New("usage: expecting command 'new' or 'score'")

func (a *App) Run() error {
	if len(a.Arguments) <= 1 {
		return ErrUsage
	}
	switch a.Arguments[1] {
	case "new":
		fmt.Fprintf(a.Output, "%s", a.CommandNew.Run())
	case "score":
		inputBytes, _ := ioutil.ReadAll(a.Input)
		out, err := a.CommandScore.Run(string(inputBytes))
		if err != nil {
			return err
		}
		fmt.Fprintf(a.Output, "%s", out)
	default:
		return ErrUsage
	}
	return nil
}
