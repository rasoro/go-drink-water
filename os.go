package main

import (
	"log"
	"os/exec"
)

func OsNotify(title string, text string, icon string) {
	exec.Command("notify-send", "-i", icon, title, text, "-t", "4000").Run()
}

func OsSound(sound string) {
	if err := exec.Command("paplay", sound).Run(); err != nil {
		log.Println("error playing sound", err)
	}
}
