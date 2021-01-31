package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/dealancer/validate.v2"
)

// Camera configuration
type Camera struct {
	Title          string   `json:"title" validate:"empty=false"`
	Address        string   `json:"address" validate:"empty=false"`
	PersistLoc     string   `json:"persist_location" validate:"empty=false"`
	SecondsPerClip int      `json:"seconds_per_clip" validate:"gte=1 & lte=3"`
	Disabled       bool     `json:"disabled"`
	Schedule       Schedule `json:"schedule"`
}

// Schedule contains each day of the week and it's off and on time entries
type Schedule struct {
	Everyday  OnOffTimes `json:"everyday"`
	Monday    OnOffTimes `json:"monday"`
	Tuesday   OnOffTimes `json:"tuesday"`
	Wednesday OnOffTimes `json:"wednesday"`
	Thursday  OnOffTimes `json:"thursday"`
	Friday    OnOffTimes `json:"friday"`
	Saturday  OnOffTimes `json:"saturday"`
	Sunday    OnOffTimes `json:"sunday"`
}

// OnOffTimes for loading up on off time entries
type OnOffTimes struct {
	Off string `json:"off"`
	On  string `json:"on"`
}

type Config struct {
	Address string `json:"address" validate:""`
}

func LoadConfig(stdlog, errlog *log.Logger) Config {
	configPath := os.Getenv("DRAGON_WEB_SERVER_CONFIG")
	if configPath == "" {
		configPath = "dws.config"
	}

	stdlog.Println("Loading web configuration:", configPath)
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("Error: %v\n", err)
	}

	stdlog.Println("Loaded web configuration...")

	cfg := Config{}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("Error parsing dws.config: %v\n", err)
	}

	err = validate.Validate(&cfg)
	if err != nil {
		log.Fatalf("Error validation dws.config content: %v\n", err)
	}

	if cfg.Address == "" {
		cfg.Address = "localhost:8080"
	}

	return cfg
}

// DragonDaemonConfig to keep track of each loaded camera's configuration
type DragonDaemonConfig struct {
	Debug   bool     `json:"debug"`
	Cameras []Camera `json:"cameras"`
}

// LoadDragonDaemonConfig parses configuration file and loads settings
func LoadDragonDaemonConfig(stdlog, errlog *log.Logger) DragonDaemonConfig {
	configPath := os.Getenv("DRAGON_DAEMON_CONFIG")
	if configPath == "" {
		configPath = "dd.config"
	}

	stdlog.Println("Loading daemon's configuration:", configPath)
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	stdlog.Println("Loaded daemon's configuration...")

	cfg := DragonDaemonConfig{}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("Error parsing dd.config: %v\n", err)
	}

	err = validate.Validate(&cfg)
	if err != nil {
		log.Fatalf("Error validating dd.config content: %v\n", err)
	}

	return cfg
}
