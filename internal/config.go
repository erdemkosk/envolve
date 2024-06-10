package config

import "github.com/gdamore/tcell/v2"

const HOME_FOLDER = ".envolve-go"
const DEFAULT_EDITOR = "nano"

var EXCLUDED_FILES = []string{".DS_Store"}

// colors
const MAIN_COLOR = tcell.ColorLightGray
const FOLDER_COLOR = tcell.ColorDarkSlateBlue
const FILE_COLOR = tcell.ColorLightSalmon
