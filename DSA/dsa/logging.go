package dsa

//"Enumerated types (enums) are a special case of sum types" - https://gobyexample.com/enums
//There is no special enum type in golang

type LOG_LEVEL uint8

// "The special keyword iota generates successive constant values automatically; in this case 0, 1, 2 " - https://gobyexample.com/enums
const (
	ERROR LOG_LEVEL = iota
	CRITICAL
	WARNING
	DEBUG
	INFO
)

//By implementing the fmt.Stringer interface, values of LOG_LEVEL can be printed using fmt

var LogStrings = map[LOG_LEVEL]string{
	ERROR:    "error",
	CRITICAL: "critical",
	WARNING:  "warning",
	DEBUG:    "debug",
	INFO:     "info",
}

func (ll LOG_LEVEL) String() string {
	return LogStrings[ll]
}

//if there are many enum state we can automate the above implementation using go:generate and stringer
//https://pkg.go.dev/golang.org/x/tools/cmd/stringer
//https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate
