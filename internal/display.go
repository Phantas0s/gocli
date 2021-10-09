package internal

import (
	"os"
)

func Display(flag bool) {
	os.Stdout.WriteString("Hello!")
	if flag {
		os.Stdout.WriteString("You've set a flag!")
	}
}
