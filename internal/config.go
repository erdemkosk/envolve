package config

import "github.com/gdamore/tcell/v2"

const HOME_FOLDER = ".envolve-go"
const DEFAULT_EDITOR = "nano"

var EXCLUDED_FILES = []string{".DS_Store"}

type CommandType int32

const (
	SYNC    CommandType = 0
	SYNCALL CommandType = 1
	SHOW    CommandType = 2
)

// colors
const MAIN_COLOR = tcell.ColorLightGray
const FOLDER_COLOR = tcell.ColorDarkSlateBlue
const FILE_COLOR = tcell.ColorLightSalmon
