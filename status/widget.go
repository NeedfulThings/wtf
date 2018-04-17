package status

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget

	Current int
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" 🎉 Status ", "status"),
		Current:    0,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	// FIXME: Use two calls to wtf.RightAlign here and get rid of this code duplication
	_, _, w, _ := widget.View.GetInnerRect()

	widget.View.Clear()
	fmt.Fprintf(
		widget.View,
		fmt.Sprintf("\n%%%ds", w-1),
		widget.timezones(),
	)

	widget.RefreshedAt = time.Now()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) animation() string {
	icons := []string{"👍", "🤜", "🤙", "🤜", "🤘", "🤜", "✊", "🤜", "👌", "🤜"}
	next := icons[widget.Current]

	widget.Current = widget.Current + 1
	if widget.Current == len(icons) {
		widget.Current = 0
	}

	return next
}

func (widget *Widget) timezones() string {
	timezones := Timezones(Config.UMap("wtf.mods.status.timezones"))

	if len(timezones) == 0 {
		return ""
	}

	// All this is to display the time entries in alphabetical order
	labels := []string{}
	for label, _ := range timezones {
		labels = append(labels, label)
	}

	sort.Strings(labels)

	tzs := []string{}
	for _, label := range labels {
		zoneStr := fmt.Sprintf("%s %s", label, timezones[label].Format(wtf.TimeFormat))
		tzs = append(tzs, zoneStr)
	}

	return strings.Join(tzs, " ◦ ")
}
