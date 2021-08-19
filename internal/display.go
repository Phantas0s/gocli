package internal

import "os"

func Display() {
	os.Stdout.WriteString("Hello!")
}
