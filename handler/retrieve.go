package handler

import (
	"VideoCutter/helper"
	"VideoCutter/video"
)

// cutter test.mp4 -r 00:12:00 00:14:00 outputPath
func Retrieve(startTime, endTime, filePath, outputPath string) {
	parsedStartTime := helper.TimeConverter(startTime)
	parsedEndTime := helper.TimeConverter(endTime)
	v, err := video.Load(filePath)
	if err != nil {
		panic(err)
	}
	if err := v.SubClip(parsedStartTime, parsedEndTime).Output(outputPath).Run(); err != nil {
		panic("Retrieve error")
	}
}
