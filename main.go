package main

import (
	"fmt"
	"os"

	"github.com/rosenhouse/bowl/application"
	"github.com/rosenhouse/bowl/model"
	"github.com/rosenhouse/bowl/ui"
)

func main() {
	app := application.App{
		Input:        os.Stdin,
		Output:       os.Stdout,
		Arguments:    os.Args,
		CommandNew:   &CommandNew{},
		CommandScore: &CommandScore{},
	}
	err := app.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
}

type CommandNew struct{}

func (c *CommandNew) Run() string {
	return ui.Template
}

type CommandScore struct{}

func (c *CommandScore) Run(userRecord string) (scoredBoard string, err error) {
	game, err := ui.ParseScorecard(userRecord)
	if err != nil {
		return "", fmt.Errorf("parsing: %s", err)
	}
	scoredGame, err := model.Score(game)
	if err != nil {
		return "", fmt.Errorf("scoring: %s", err)
	}
	scoredBoard = ui.Output(scoredGame)
	return scoredBoard, nil
}
