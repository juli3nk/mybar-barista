package main

import (
	"barista.run/bar"
	"barista.run/outputs"
	"barista.run/pango"
)

func outputBrightness(value int) bar.Output {
	return outputs.Pango(
		//pango.Icon("material-brightness-medium").Color(colors.Scheme("dim-icon")),
		pango.Icon("material-brightness-medium").Alpha(0.8),
		pango.Textf("%d%%", value),
	)
}
