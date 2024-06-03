package testutils

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func Testutils() {

	url := `https://www.youtube.com/watch?v=GW2g-5WALrc`
	progress := `--progress`
	progressDelta := `--progress-delta 1` //seconds
	filepath := `--print after_move:filepath`
	channel := `--print after_move:channel`
	title := `--print after_move:title`

	ext := `--print after_move:ext`
	duration := `--print after_move:duration`
	webpage_url_domain := `--print after_move:webpage_url_domain`
	original_url := `--print after_move:original_url`
	n_entries := `--print after_move:n_entries`
	playlist_title := `--print after_move:playlist_title`
	playlist_count := `--print after_move:playlist_count`
	playlist_index := `--print after_move:playlist_index`
	thumbnail := `--write-thumbnail`
	format := `--print format`
	epoch := `--print epoch`

	//infojson := `--write-info-json`
	skip_download := `--skip-download`

	command := `./testutils/yt-dlp_x86.exe`
	space := " "
	var args []string
	args = append(args, url)
	args = append(args, filepath)
	args = append(args, channel)
	args = append(args, title)
	//args = append(args, description)
	args = append(args, ext)
	args = append(args, duration)
	args = append(args, webpage_url_domain)
	args = append(args, original_url)
	args = append(args, n_entries)      //for playlists
	args = append(args, playlist_title) //for playlists
	args = append(args, playlist_count) //for playlists
	args = append(args, playlist_index) //for playlists
	args = append(args, skip_download)  //for metadata only
	args = append(args, thumbnail)
	args = append(args, format)
	args = append(args, epoch) // for timestamp time.Now().Unix() GMT TIME

	fmt.Println("Anant here", time.Now().Unix())

	//args = append(args, infojson)
	args = append(args, progressDelta)
	args = append(args, progress)
	args = append(args, "--write-auto-subs")

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
