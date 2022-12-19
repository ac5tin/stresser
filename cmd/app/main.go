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
	export := flag.String("e", "", "Path to the file where the responses will be exported")
	exportErrOnly := flag.Bool("err", false, "Export errors only")
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
	// export
	{
		if *export != "" {
			if err := r.ExportResponsesToFile(*export, *exportErrOnly); err != nil {
				panic(err)
			}
		}
	}
}
