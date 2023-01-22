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
		pango.Icon("material-today").Color(colors.Scheme("dim-icon")),
		now.Format("Mon Jan 2 "),
		pango.Icon("material-access-time").Color(colors.Scheme("dim-icon")),
		now.Format("15:04:05"),
	).OnClick(click.RunLeft("gsimplecal"))
}
