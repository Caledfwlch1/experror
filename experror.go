package experror

const (
	Inform = iota
	Warning
	Critical
)

type ExpError interface {
	Error() string
	String() string
	ErrorCod() int
	ErrorLevel() int
}

type expanderror struct {
	code  int
	level int
	short string
	long  string
}

var _ ExpError = (*expanderror)(nil)

var (
	// example
	FileNotFound = New(1, Critical, "file not found", "file not fount in this directory")
)

func New(code, lvl int, short, long string) ExpError {
	return &expanderror{code: code, level: lvl, short: short, long: long}
}

func (e *expanderror) Error() string {
	return e.short
}

func (e *expanderror) String() string {
	return e.long
}

func (e *expanderror) ErrorCod() int {
	return e.code
}

func (e *expanderror) ErrorLevel() int {
	return e.level
}
