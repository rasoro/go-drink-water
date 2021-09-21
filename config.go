package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
)

//go:embed gopher.png
var DEFAULT_ICON []byte

const (
	DEFAULT_ICON_PATH  = "/tmp/go-drink-water-icon.png"
	DEFAULT_SOUND_PATH = "/usr/share/sounds/freedesktop/stereo/complete.oga"
)

type Config struct {
	IconPath  string
	SoundPath string
}

var AppConfig *Config = nil

func GetConfig() *Config {
	if AppConfig == nil {
		AppConfig = &Config{
			IconPath:  DEFAULT_ICON_PATH,
			SoundPath: DEFAULT_SOUND_PATH,
		}
	}
	return AppConfig
}

func init() {
	log.Println("Welcome")
	iconImg := bytes.NewReader(DEFAULT_ICON)
	dest, err := os.Create(DEFAULT_ICON_PATH)
	if err != nil {
		fmt.Println("not today")
		log.Fatal(err)
	}
	defer dest.Close()
	io.Copy(dest, iconImg)
}
