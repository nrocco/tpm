package gpg

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func DecodeGpgString(encoded string) (string, error) {
	cmd := exec.Command("gpg", "--decrypt", "-q")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, encoded)
	}()

	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}
