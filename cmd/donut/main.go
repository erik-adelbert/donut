package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/erik-adelbert/donut/donut"
	"golang.org/x/term"
)

func main() {
	w, h, err := term.GetSize(int(os.Stdin.Fd()))

	if err != nil {
		fatal("Could not get terminal size:", err)
	}

	h = max(1, h) // ensure the dimensions are strictly positive
	w = max(1, w)
	p := tea.NewProgram(donut.NewModel(h, w))

	if _, err := p.Run(); err != nil {
		fatal("Error running program:", err)
	}
}

func fatal(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
