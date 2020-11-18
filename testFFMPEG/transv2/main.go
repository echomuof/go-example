package main

import (
	"github.com/floostack/transcoder/ffmpeg"
	"log"
)

var inputPath = "https://vod2.pipi.cn/fd217976vodcq1400253221/f80bfb765285890808837859031/f0.mp4"
var outputPath = "./testmp4.mp4"

func main() {
	format := "mp4"
	overwrite := true
	fileterOption := "scale=iw:ih,pad=iw:iw/720*1280:0:(iw/720*1280-ih)/2:black,scale=720:1280,delogo=x=570:455:w=140:h=50:show=0"

	opts := ffmpeg.Options{
		OutputFormat: &format,
		Overwrite:    &overwrite,
		VideoFilter:  &fileterOption,
	}

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   "/usr/local/bin/ffmpeg",
		FfprobeBinPath:  "/usr/local/bin/ffprobe",
		ProgressEnabled: true,
	}

	progress, err := ffmpeg.
		New(ffmpegConf).
		Input(inputPath).
		Output(outputPath).
		WithOptions(opts).
		Start(opts)

	if err != nil {
		log.Fatal(err)
	}

	for msg := range progress {
		log.Printf("%+v", msg)
	}
}
