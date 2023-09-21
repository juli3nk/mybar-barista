package main

import (
	"barista.run/bar"
	"barista.run/outputs"
	"barista.run/pango"
)

func outputBrightness(value int) bar.Output {
	return outputs.Pango(
		pango.Text("󰳲").Alpha(0.8),
		pango.Textf("%d%%", value),
	)
}
