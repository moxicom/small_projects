package internal

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	// Очистка экрана для разных ОС
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintColoredText(text, colorCode string) {
	log.Printf("%s%s%s", colorCode, text, "\033[0m")
}
