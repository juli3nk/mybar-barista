package main

import (
	"time"

	"barista.run/bar"
	"barista.run/base/click"
	"barista.run/colors"
	"barista.run/outputs"
	"barista.run/pango"
)

func outputLocaltime(now time.Time) bar.Output {
	return outputs.Pango(
		pango.Text("󰃶").Color(colors.Scheme("dim-icon")),
		now.Format("Mon Jan 2 "),
		pango.Text("󱑏").Color(colors.Scheme("dim-icon")),
		now.Format("15:04:05"),
	).OnClick(click.RunLeft("gsimplecal"))
}
