package main

import (
	"flag"
	"stresser/pkg/task"
)

func main() {
	// Parse the command line flags.
	concurrentCount := flag.Uint("c", 10, "Number of concurrent requests")
	totalCount := flag.Uint("t", 30, "Total number of requests")
	file := flag.String("f", "task.toml", "Path to the task definition file")
	flag.Parse()
	// init task
	t, err := task.ParseFile(*file)
	if err != nil {
		panic(err)
	}
	r, err := t.Execute(&task.Config{
		Concurrency: uint16(*concurrentCount),
		Total:       uint32(*totalCount),
	})
	if err != nil {
		panic(err)
	}

	r.Render()
}
