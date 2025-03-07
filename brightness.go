package main

import (
	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/colors"
	"github.com/barista-run/barista/outputs"
	"github.com/barista-run/barista/pango"
)

func outputBrightness(value int) bar.Output {
	return outputs.Pango(
		pango.Icon("mdi-brightness-5").Color(colors.Scheme("color13")),
		spacer,
		pango.Textf("%d%%", value),
	)
}
