package classpath

import "os"

// :(Linux/Unix) ;(Win)
const pathListSeparator  = string(os.PathListSeparator)

type Entry interface {

	ReadClass(className string) ([]byte, Entry, error)

	String() string
}

// Return different implementations depending on different file types
func newEntry(path string) Entry {
	return nil
}
