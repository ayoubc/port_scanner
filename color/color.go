package color

import (
	"fmt"
)

var (
	OK      = "\033[92m" // GREEN
	WARNING = "\033[93m" // YELLOW
	FAIL    = "\033[91m" // RED
	RESET   = "\033[0m"  // RESET COLOR

)

var mapping = map[string]string{
	"green": OK,
	"red":   FAIL,
	"warn":  WARNING,
}

func Cprint(msg, color string) {
	fmt.Printf("%s%s%s\n", mapping[color], msg, RESET)
}
