/**
 *
 * @author: echomuof
 * @created: 2020/11/18
 */
package main

import (
	"fmt"
	"github.com/xfrr/goffmpeg/transcoder"
)

var inputPath = "https://vod2.pipi.cn/fd217976vodcq1400253221/f80bfb765285890808837859031/f0.mp4"
var outputPath = "./testmp4.mp4"

func main() {
	// Create new instance of transcoder
	trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize(inputPath, outputPath)
	// Handle error...
	if err != nil {

	}

	// SET FRAME RATE TO MEDIAFILE
	trans.MediaFile().SetFrameRate(70)
	// SET ULTRAFAST PRESET TO MEDIAFILE
	trans.MediaFile().SetPreset("ultrafast")

	trans.MediaFile().SetVideoFilter("vf")

	// Start transcoder process to check progress
	done := trans.Run(true)

	// Returns a channel to get the transcoding progress
	progress := trans.Output()

	// Example of printing transcoding progress
	for msg := range progress {
		fmt.Println(msg)
	}

	// This channel is used to wait for the transcoding process to end
	err = <-done

}
