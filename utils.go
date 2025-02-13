package main

import (
	"path/filepath"
	"time"

	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/colors"
	"github.com/barista-run/barista/modules/clock"
	"github.com/barista-run/barista/outputs"
	"github.com/barista-run/barista/pango"
	"github.com/juli3nk/go-utils/user"
)

func home(path ...string) string {
	usr := user.New()

	args := append([]string{usr.HomeDir}, path...)
	return filepath.Join(args...)
}

func makeTzClock(lbl, tzName string) bar.Module {
	c, err := clock.ZoneByName(tzName)
	if err != nil {
		panic(err)
	}
	return c.Output(time.Minute, func(now time.Time) bar.Output {
		return outputs.Pango(
			pango.Text(lbl).Smaller(),
			spacer,
			now.Format("15:04"),
		)
	})
}

func threshold(out *bar.Segment, urgent bool, color ...bool) *bar.Segment {
	if urgent {
		return out.Urgent(true)
	}
	colorKeys := []string{"critical", "degraded", "warning", "good"}
	for i, c := range colorKeys {
		if len(color) > i && color[i] {
			return out.Color(colors.Scheme(c))
		}
	}
	return out
}

func truncate(in string, l int) string {
	fromStart := false
	if l < 0 {
		fromStart = true
		l = -l
	}
	inLen := len([]rune(in))
	if inLen <= l {
		return in
	}
	if fromStart {
		return "⋯" + string([]rune(in)[inLen-l+1:])
	}
	return string([]rune(in)[:l-1]) + "⋯"
}
