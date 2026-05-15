// +build windows

package spinner

import (
	"fmt"

	"github.com/leaanthony/wincursor"
)

func showCursor() {
	wincursor.Show()
}

func hideCursor() {
	wincursor.Hide()
}

func (s *Spinner) clearCurrentLine() {
	// *shudder*
	fmt.Fprint(s.writer, "\r")

	// Get the current line length
	var length = len(s.GetMessage()) + len(s.getCurrentSpinnerFrame()) + 1

	for i := 0; i < length; i++ {
		fmt.Fprint(s.writer, " ")
	}
	fmt.Fprint(s.writer, "\r")
}
