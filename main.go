package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/athega/whistler/robots"
	"github.com/gorilla/schema"
)

func main() {
	http.HandleFunc("/slack", commandHandler)
	http.HandleFunc("/slack_hook", commandHandler)

	startServer()
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	hook := r.URL.Path == "/slack_hook"

	if err == nil {
		decoder := schema.NewDecoder()
		command := new(robots.SlashCommand)
		err := decoder.Decode(command, r.PostForm)

		if err != nil {
			log.Println("Couldn't parse post request:", err)
		}

		if hook {
			c := strings.Split(command.Text, " ")
			command.Command = c[1]
			command.Text = strings.Join(c[2:], " ")
		} else {
			command.Command = command.Command[1:]
		}

		robot := getRobot(command.Command)
		w.WriteHeader(http.StatusOK)

		respFn := plainResp
		if hook {
			respFn = jsonResp
		}

		if robot != nil {
			respFn(w, robot.Run(command))
		} else {
			respFn(w, "No robot for that command yet :(")
		}
	}
}

func jsonResp(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resp := map[string]string{"text": msg}
	r, err := json.Marshal(resp)

	if err != nil {
		log.Println("Couldn't marshal hook response:", err)
	} else {
		io.WriteString(w, string(r))
	}
}

func plainResp(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	io.WriteString(w, msg)
}

func startServer() {
	log.Printf("Starting HTTP server on %s", robots.Config.Port)

	err := http.ListenAndServe(":"+robots.Config.Port, nil)
	if err != nil {
		log.Fatal("Server start error: ", err)
	}
}

func getRobot(command string) robots.Robot {
	if RobotInitFunction, ok := robots.Robots[command]; ok {
		return RobotInitFunction()
	}
	return nil
}
