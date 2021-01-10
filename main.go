package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/takama/daemon"
	"github.com/tauraamui/dragonmediaserver/db"
	"github.com/tauraamui/dragonmediaserver/web"
)

const (
	name        = "dragon_media_server"
	description = "Dragon media server service which hosts access to the clip manifest and video files"
)

var stdlog, errlog *log.Logger

type Service struct {
	daemon.Daemon
}

func (service *Service) Manage() (string, error) {
	usage := "Usage: dragonms install | remove | start | stop | status"

	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	stdlog.Println("Starting dragon media server...")

	dbConn, err := db.Connect()
	if err != nil {
		errlog.Printf("Unable to open DB: %v\n", err)
		os.Exit(1)
	}

	err = db.Setup(dbConn)
	if err != nil {
		errlog.Printf("Unable to setup database: %v\n", err)
		os.Exit(1)
	}

	server := web.NewServer(dbConn)
	go http.ListenAndServe("localhost:8080", server)

	killSignal := <-interrupt
	stdlog.Println("Received signal:", killSignal)

	return "Shutdown successful... BYE! 👋", nil
}

func init() {
	stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

func main() {
	daemonType := daemon.SystemDaemon
	if runtime.GOOS == "darwin" {
		daemonType = daemon.UserAgent
	}

	srv, err := daemon.New(name, description, daemonType)
	if err != nil {
		errlog.Println("Error:", err)
		os.Exit(1)
	}

	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errlog.Println(status, "\nError:", err)
		os.Exit(1)
	}

	fmt.Println(status)

	os.Exit(0)
}
