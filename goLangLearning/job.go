package main

import (
	"os"
	"log"
)

type Job struct {
	Command string
	*log.Logger //implicity composition
}

func main() {
	job := &Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Print("test")
}