package main

import (
	"github.com/1Ginz/jump/internal/command"
	"github.com/1Ginz/jump/internal/config"
	"github.com/1Ginz/jump/internal/history"
	"github.com/1Ginz/jump/internal/interactive"
	"github.com/1Ginz/jump/internal/ssh"
	"github.com/1Ginz/jump/internal/tsh"
	"github.com/1Ginz/jump/internal/workspace"
	"os"
)

func main() {
	args := os.Args[1:]
	mode := config.TSHMode
	action, value := command.Which()
	switch action {
	case command.InteractiveHistory:
		args, mode = interactive.History()
	case command.InteractiveConfig:
		args, mode = interactive.Config("")
	case command.InteractiveConfigWithSearch:
		args, mode = interactive.Config(value)
	case command.ListHistory:
		history.Print()
		return
	case command.ListConfig:
		config.Print()
		return
	case command.ListWorkspace:
		workspace.Print()
		return
	case command.ActiveWorkspace:
		interactive.Active()
		return
	default:

	}
	history.AddHistoryFromArgs(args, mode)

	if mode == config.SSHMode {
		ssh.Run(args)
	} else {
		tsh.Run(args)
	}
}
