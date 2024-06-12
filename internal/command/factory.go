package command

import (
	config "github.com/erdemkosk/envolve-go/internal"
)

func CommandFactory(commandType config.CommandType, path string) ICommand {
	if commandType == config.SYNC {
		return &SyncCommand{path: path}
	}
	if commandType == config.SHOW {
		return &ShowCommand{}
	}
	if commandType == config.SYNCALL {
		return &SyncAllCommand{path: path}
	}

	return nil
}
