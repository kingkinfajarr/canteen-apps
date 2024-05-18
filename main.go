package main

import (
	"canteen-apps/services"
	"canteen-apps/utils"
	"os"
	"os/exec"
	"runtime"
)

func clearTerminal() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		return // unsupported platform
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	admin := services.NewAdmin()
	services.SeedAdminData(admin)

	clearTerminal()
	utils.ShowMenu(admin)
}
