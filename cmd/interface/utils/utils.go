package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearTerminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ConvertCurrencyToSimbol(currency string) string{
  switch currency{
  case "BRL":
    return "R$"
  case "USD":
    return "$"
  case "EUR":
    return "€"
  case "JPY":
    return "¥"
  default:
    return ""
  }
}
