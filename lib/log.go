package lib

import (
	"fmt"
	"time"

	color "github.com/fatih/color"
)

type Logger struct {
	timezone string
}

func (l Logger) ink(c *color.Color, text string) string {
	return c.Sprintf(text)
}

func (l Logger) timestamp() string {
	const layout = "15:04:05"

	// FIXME
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)

	c := color.New(color.FgHiMagenta, color.Bold)
	return c.Sprintf(l.ink(c, now.Format(layout)))
}

func (l Logger) logging(c *color.Color, text string) {
	fmt.Printf("[%s] %s\n", l.timestamp(), l.ink(c, text))
}

func (l Logger) Warn(text string) {
	c := color.New(color.FgYellow, color.Bold)
	l.logging(c, text)
}

func (l Logger) Error(text string) {
	c := color.New(color.FgRed, color.Bold)
	l.logging(c, text)
}

func (l Logger) Success(text string) {
	c := color.New(color.FgGreen, color.Bold)
	l.logging(c, text)
}

func (l Logger) Info(text string) {
	c := color.New(color.FgWhite, color.Bold)
	l.logging(c, text)
}
