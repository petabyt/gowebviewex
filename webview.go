package gowebviewex

import (
	"strings"
)

// Turn arrays into null joined
// strings, which can be seperated
// once it has entered Java code.
func array(array []string) string {
	return strings.Join(array, "\000")
}

// Called from Java always at startup
func Startup() string {
	settings := []string{
		"file:///android_asset/index.html", // File the webview will go to
		"// Nothin", // Startup code once the webview has loaded
	}

	settingsString := array(settings)
	return settingsString
}

// Called from Java
func Test() string {
	return "Hey Vsauce, Michael here"
}
