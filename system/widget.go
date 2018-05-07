package system

import (
	"fmt"
	"time"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget

	BuiltAt string
	Version string
}

func NewWidget(builtAt, version string) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Build ", "system", false),
		BuiltAt:    builtAt,
		Version:    version,
	}

	return &widget
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.View.Clear()

	fmt.Fprintf(
		widget.View,
		"%6s: %s\n%6s: %s",
		"Built",
		widget.prettyBuiltAt(),
		"Vers",
		widget.Version,
	)

	widget.RefreshedAt = time.Now()
}

func (widget *Widget) prettyBuiltAt() string {
	str, err := time.Parse(wtf.TimestampFormat, widget.BuiltAt)
	if err != nil {
		return err.Error()
	} else {
		return str.Format("Jan _2, 15:04")
	}
}
