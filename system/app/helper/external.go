package helper

import "os/exec"

func OpenExpoler(path string) error {
	// cmd := exec.Command("explorer.exe", "/select,", path)
	cmd := exec.Command("explorer.exe", path)

	// Set the Stdout and Stderr to the os.Stdout and os.Stderr to see the command's output.
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
