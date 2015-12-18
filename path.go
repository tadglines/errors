// Copyright 2013, 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package errors

import (
	"runtime"
	"strings"
)

// prefixSize is used internally to trim the user specific path from the
// front of the returned filenames from the runtime call stack.
var prefixSize int

// goPath is the deduced path based on the location of this file as compiled.
var goPath string

func determineGoPath(path, knownSuffix string) string {
	if !strings.HasSuffix(path, knownSuffix) {
		return path
	}
	prefix := path[:len(path)-len(knownSuffix)]
	if strings.HasSuffix(prefix, "/vendor/src/") {
		return prefix[:len(prefix)-len("/vendor/src")]
	}
	if strings.HasSuffix(prefix, "/src/") {
		return prefix[:len(prefix)-len("/src")]
	}
	return prefix
}

func init() {
	_, file, _, ok := runtime.Caller(0)
	if ok {
		// We know that the end of the file should be:
		// github.com/tadglines/errors/path.go
		goPath = determineGoPath(file, "github.com/tadglines/errors/path.go")
		prefixSize = len(goPath)
	}
}

func trimGoPath(filename string) string {
	if strings.HasPrefix(filename, goPath) {
		suffix := filename[prefixSize:]
		if strings.HasPrefix(suffix, "src/") {
			return suffix[4:]
		}
		if strings.HasPrefix(suffix, "vendor/src/") {
			return suffix[11:]
		}
		return suffix
	}
	return filename
}
