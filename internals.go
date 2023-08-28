package main

import (
	"os/exec"
	"log"
	"strings"
)


type containerDetails struct {
        ID      string `json:"id"`
        Name    string `json:"name"`
        IP      string `json:"ip"`
        Tag     string `json:"tag"`
}

func Run(name, tag string) containerDetails {
        arg := []string{name, tag}
	args := strings.Join(arg, ":")
	log.Print(args)
	out, err := exec.Command("docker", "run", "-d", args).Output() 	

	if err != nil {
		log.Fatal(err)
	}

	_out := string(out)
	ip, err := exec.Command("docker", "inspect", "--format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'", _out).Output()
	detailedContainer := containerDetails {
		ID: _out[:12],
		Name: name,
		IP: string(ip),
		Tag: tag,
	}

	return detailedContainer
}
