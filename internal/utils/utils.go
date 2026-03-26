// Package utils provides a few simple funcs for logging and error checking.
package utils

import (
	"fmt"
)

// Check panics if an error is not nil.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// Must panics if an error is not nil.
func Must(err error) {
	Check(err)
}

// Log simply prints the provided item.
func Log(item any) {
	fmt.Println(item)
}

// Logf prints the provided item using a format string.
func Logf(fmtStr string, item any) {
	fmt.Println(fmt.Sprintf(fmtStr, item))
}

// Logs prints the provided items.
func Logs(items ...any) {
	fmt.Println(items...)
}

// Logsf prints the provided items using a format string.
func Logsf(fmtStr string, items ...any) {
	fmt.Println(fmt.Sprintf(fmtStr, items...))
}
