package colorOutput

import "fmt"

const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
)

func Green(s string) {
	fmt.Printf("\x1b[%dm%s\x1b[0m", color_green, s)
}
func Yellow(s string) {
	fmt.Printf("\x1b[%dm%s\x1b[0m", color_yellow, s)
}
func Red(s string) {
	fmt.Printf("\x1b[%dm%s\x1b[0m", color_red, s)
}