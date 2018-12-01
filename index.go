package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		clear = "\033[2J"
	)

	for i := 10; i > 0; i-- {
		fmt.Fprint(w, clear, clear, clear, clear)

		f := figure.NewFigure(fmt.Sprintf("Left: %d", i), "", true)
		figure.Write(w, f)

		if fl, ok := w.(http.Flusher); ok {
			fl.Flush()
		}

		time.Sleep(time.Second)
	}
	fmt.Fprint(w, clear)

	f := figure.NewFigure("Done!", "", true)
	figure.Write(w, f)
}
