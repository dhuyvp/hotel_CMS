package utils

import (
	"fmt"
	"os"
)

func ConnectionURLBuilder(n string) (string, error) {
	var url string

	switch n {
	case "mysql":
		url = fmt.Sprintf("%s",
			os.Getenv("MYSQL_URL"),
		)
	default:
		return "", fmt.Errorf("Connection name '%v' is not supported", n)
	}

	return url, nil
}
