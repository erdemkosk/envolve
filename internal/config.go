package config

import "github.com/gdamore/tcell/v2"

const HOME_FOLDER = ".envolve-go"
const DEFAULT_EDITOR = "nano"

var EXCLUDED_FILES = []string{".DS_Store"}

type CommandType int32

const (
	SYNC CommandType = 0
	GET  CommandType = 1
)

// colors
const MAIN_COLOR = tcell.ColorLightGray
const FOLDER_COLOR = tcell.ColorDarkSlateBlue
const FILE_COLOR = tcell.ColorLightSalmon
