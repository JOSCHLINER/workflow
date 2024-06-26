package tableformatter

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type TableFormatter struct {
	Options TableFormatterOptions
	Spacing []int
}

func getDefaultTableFormatter() *TableFormatter {
	return &TableFormatter{
		Options: TableFormatterOptions{
			onlyActive: false,
			padding:    1,
		},
	}
}

type TableFormatterOptions struct {
	onlyActive bool
	padding    int
}

type TableOptions func(*TableFormatterOptions)

func OnlyActiveTasks() TableOptions {
	return func(tfo *TableFormatterOptions) {
		tfo.onlyActive = true
	}
}

func SetPadding(padding int) TableOptions {
	return func(tfo *TableFormatterOptions) {
		tfo.padding = padding
	}
}

func NewTableFormatter(opts ...TableOptions) *TableFormatter {

	printer := getDefaultTableFormatter()
	for _, opt := range opts {
		opt(&printer.Options)
	}

	return printer
}

func (tf *TableFormatter) ConstructTable(items [][]string) string {
	tf.getSpacing(items)

	out := tf.constructDivider() + tf.constructRow(items[0]) + tf.constructDivider()
	for i := 1; i < len(items); i++ {
		out += tf.constructRow(items[i])
	}

	return out + tf.constructDivider()
}

func (tf *TableFormatter) getSpacing(items [][]string) {
	tf.Spacing = make([]int, len(items[0]))

	for _, row := range items {
		for c, column := range row {

			if n := utf8.RuneCountInString(column); n > tf.Spacing[c] {
				tf.Spacing[c] = n
			}

		}
	}
}

func (tf *TableFormatter) constructRow(items []string) string {

	padding, line := strings.Repeat(" ", tf.Options.padding), ""
	for i, item := range items {
		filler := strings.Repeat(" ", tf.Spacing[i]-utf8.RuneCountInString(item)+1)
		line += fmt.Sprintf("| %s%s%s", padding, item, filler)
	}

	return line + "|\n"
}

func (tf *TableFormatter) constructDivider() string {

	line := "+"
	for _, spacing := range tf.Spacing {
		line += strings.Repeat("-", tf.Options.padding+spacing+2) + "+"
	}

	return line + "\n"
}
