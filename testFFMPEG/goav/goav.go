/**
 *
 * @author: echomuof
 * @created: 2020/11/18
 */
package main

import (
	"github.com/giorgisio/goav/avformat"
	"log"
)

func main() {

	filename := "https://vod2.pipi.cn/fd217976vodcq1400253221/f80bfb765285890808837859031/f0.mp4"

	// Register all formats and codecs
	avformat.AvRegisterAll()

	ctx := avformat.AvformatAllocContext()

	// Open video file
	if avformat.AvformatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AvformatCloseInput()
		return
	}

	//...

}