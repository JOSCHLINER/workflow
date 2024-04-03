package tableformatter

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
