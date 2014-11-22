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
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/slack", commandHandler)
	http.HandleFunc("/slack_hook", commandHandler)

	http.Handle("/images/", http.StripPrefix("/images/",
		http.FileServer(http.Dir("images"))))

	startServer()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Whistler</h1><img src="/images/whistler.jpg">`))
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

		// Check if we got a command
		if command.Command == "" {
			log.Println("No command given", r.PostForm)
			return
		}

		if hook {
			c := strings.Split(command.Text, " ")

			if len(c) < 2 {
				log.Println("No command given")
				return
			}

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
	log.Printf("Listening on http://0.0.0.0:%s", robots.Config.Port)

	err := http.ListenAndServe(":"+robots.Config.Port, nil)
	if err != nil {
		log.Fatal("Server start error: ", err)
	}
}

func getRobot(command string) robots.Robot {
	if robotInitFunction, ok := robots.Robots[command]; ok {
		return robotInitFunction()
	}

	return nil
}
