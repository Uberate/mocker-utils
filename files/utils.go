package files

import "os"

// IsFileExists return true when file is exists
func IsFileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsPermission(err) {
			return true
		}
		return false
	}
	return true
}

// ReadFile return the file value.
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// WriteFile will write the data to specify file, and use os.ModePerm.
func WriteFile(path string, data []byte) error {
	err := os.WriteFile(path, data, os.ModePerm) // ignore_security_alert
	if err != nil {
		return err
	}
	return err
}
