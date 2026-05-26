package cliprinter

const (
	_ formatType = iota
	json
	pretty

	defaultFormatType = pretty
)

type formatType int64

func strToFormatType(f string) (formatType, error) {
	_ = "STUB: not implemented"
	return *new(formatType), nil
}

func formatTypeToStr(f formatType) string { _ = "STUB: not implemented"; return "" }
