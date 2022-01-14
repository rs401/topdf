package converter

import (
	"os/exec"
	"strings"
)

// func Convtopdf(src, dst string) (string, error) { // error {
func Convtopdf(src string) (string, error) {
	// cmd := exec.Command("pandoc", src, "-o", dst)
	// sed 's/.*-> \(.*\) using filter :.*/\1/'
	cmd := exec.Command("lowriter", "--headless", "--convert-to", "pdf", src)
	// err := cmd.Run()
	// if err != nil {
	// 	return err
	// }
	// return nil
	bytes, err := cmd.Output()
	if err != nil {
		return "", err
	}
	sub := strings.Split(string(bytes), "-> ")
	sub2 := strings.Split(sub[1], " using filter :")
	return sub2[0], nil
}
