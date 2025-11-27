package utils

import (
	"os"
	"os/user"
	"strings"
)

// Save user default name in ~/.greeterconfig
func SaveName(name string) {
	u, _ := user.Current()
	path := u.HomeDir + "/.greeterconfig"
	os.WriteFile(path, []byte(name), 0644)
}

// Load default name
func LoadDefaultName() string {
	u, _ := user.Current()
	path := u.HomeDir + "/.greeterconfig"

	data, err := os.ReadFile(path)
	if err != nil {
		return "friend"
	}

	return strings.TrimSpace(string(data))
}
