package classpath

import (
	"os"
	"strings"
)

// :(Linux/Unix) ;(Win)
const pathListSeparator  = string(os.PathListSeparator)

type Entry interface {

	ReadClass(className string) ([]byte, Entry, error)

	String() string
}

// Return different implementations depending on different file types
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {

		return newZipEntry(path)
	}

	return newDirEntry(path)
}
