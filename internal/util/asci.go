package util

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func TcellColorToAnsi(c tcell.Color) string {
	// Assuming 24-bit color support, we convert the tcell.Color to ANSI escape code.
	// This is a simplified approach assuming 24-bit color support.
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", (c>>16)&0xFF, (c>>8)&0xFF, c&0xFF)
}
