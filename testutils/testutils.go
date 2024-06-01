package testutils

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

func Testutils() {

	url := `https://www.youtube.com/watch?v=Wz0H8HFkI9U`
	progress := `--progress`
	progressDelta := `--progress-delta 5` //seconds
	filepath := `--print after_move:filepath`
	channel := `--print after_move:channel`
	likes := `--print after_move:like_count`

	command := `./yt-dlp_x86.exe`
	space := " "
	var args []string
	args = append(args, url)
	args = append(args, progressDelta)
	args = append(args, filepath)
	args = append(args, channel)
	args = append(args, likes)
	args = append(args, progress)

	arguments := strings.Join(args, space)
	cmd := exec.Command(command)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.CmdLine = command + space + arguments

	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(cmd.String())

	if err = cmd.Start(); err != nil {
		fmt.Print(err)
	}
	for {
		tmp := make([]byte, 1024)
		n, err := stdout.Read(tmp)
		fmt.Println("\r", string(tmp[:n]))
		if err != nil {
			break
		}
	}
}
