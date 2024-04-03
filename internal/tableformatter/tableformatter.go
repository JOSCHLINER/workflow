package tableformatter

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

func NewTableFormatter(opts ...TableOptions) *TableFormatter {

	printer := getDefaultTableFormatter()
	for _, opt := range opts {
		opt(&printer.Options)
	}

	return printer
}
