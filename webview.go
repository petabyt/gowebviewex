package webview

import (
	"strings"
)

// Turn arrays into null joined
// strings, which can be seperated
// once it has entered Java code.
func array(array []string) string {
	return strings.Join(array, "\000")
}

func Startup() string {
	settings := []string{
		"file:///android_asset/index.html",
		"// Nothin",
	}

	settingsString := array(settings)
	return settingsString
}

func Test() string {
	return "Hey Vsauce, Michael here"
}
