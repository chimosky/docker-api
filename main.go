package main

import (
//	"net/http"
	"os/exec"
	"log"
	"fmt"

//	"github.com/gin-gonic/gin"
)

// Figure out the parameters you want each
// container to share and then add that here,
// typically it'd be the name and maybe base image.
// A default base image where there's none specified.
type Container struct {
	Name	string `json:"name"`
}

func main() {

	out, err := exec.Command("docker", "images").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%s\n", out)
}
