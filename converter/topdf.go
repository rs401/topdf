package converter

import (
	"os/exec"
	"strings"
)

func Convtopdf(src string) (string, error) {
	cmd := exec.Command("lowriter", "--headless", "--convert-to", "pdf", src)

	bytes, err := cmd.Output()
	if err != nil {
		return "", err
	}
	sub := strings.Split(string(bytes), "-> ")
	sub2 := strings.Split(sub[1], " using filter :")
	return sub2[0], nil
}
