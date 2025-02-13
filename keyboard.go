package main

import (
	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/colors"
	"github.com/barista-run/barista/outputs"
	"github.com/barista-run/barista/pango"
)

func outputKeyboardLayout(layout string) bar.Output {
	return outputs.Pango(
		pango.Icon("mdi-keyboard-variant").Color(colors.Scheme("color10")),
		spacer,
		pango.Text(layout),
	)
}
