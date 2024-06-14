package config

import "github.com/gdamore/tcell/v2"

const HOME_FOLDER = ".envolve"
const DEFAULT_EDITOR = "nano"

var EXCLUDED_FILES = []string{".DS_Store"}

type CommandType int32

const (
	SYNC    CommandType = 0
	SYNCALL CommandType = 1
	SHOW    CommandType = 2
	EDIT    CommandType = 3
)

// colors
const MAIN_COLOR = tcell.ColorLightGray
const FOLDER_COLOR = tcell.ColorDarkSlateBlue
const FILE_COLOR = tcell.ColorLightSalmon

// colors
const (
	RESET          = "\033[0m"
	RED            = "\033[31m"
	PASTEL_RED     = "\033[91m"
	PASTEL_GREEN   = "\033[92m"
	PASTEL_YELLOW  = "\033[93m"
	PASTEL_BLUE    = "\033[94m"
	PASTEL_MAGENTA = "\033[95m"
	PASTEL_CYAN    = "\033[96m"
	PASTEL_WHITE   = "\033[97m"
	PASTEL_GRAY    = "\033[37m"
	PASTEL_PURPLE  = "\033[35m"
	PASTEL_ORANGE  = "\033[38;5;214m" // Bir pastel turuncu tonu ekledim
)
