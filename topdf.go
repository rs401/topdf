package main

import (
	"os/exec"
)

func convtopdf(src, dst string) error {
	cmd := exec.Command("pandoc", src, "-o", dst)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
