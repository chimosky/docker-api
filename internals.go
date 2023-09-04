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

	id := runAndReturnID(args)
	ip := getIP(id)

	detailedContainer := containerDetails {
		ID: id,
		Name: name,
		IP: ip,
		Tag: tag,
	}

	return detailedContainer
}

func runAndReturnID(args string) string {
	out, err := exec.Command("docker", "run", "-d", args).Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)[:12]
}

func getIP(id string) string {
	format := "--format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'"
	ip, err := exec.Command("docker", "inspect", format, id).Output()

	if err != nil {
		log.Fatal(err)
	}

	return string(ip)
}
