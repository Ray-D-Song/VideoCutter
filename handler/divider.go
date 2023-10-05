package handler

import (
	"VideoCutter/helper"
	"VideoCutter/video"
	"context"
	"os"
	"strconv"
	"sync"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

type Segment struct {
	startTime float64
	endTime   float64
}

func getVideoDuration(filePath string) float64 {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	fileReader, err := os.Open(filePath)
	if err != nil {
		panic("Read file error")
	}
	data, err := ffprobe.ProbeReader(ctx, fileReader)
	if err != nil {
		panic("Get file info error")
	}
	return data.Format.DurationSeconds
}

func generateSegments(start, end float64, segmentSize float64) []Segment {
	var segments []Segment

	for i := start; i <= end; i += segmentSize {
		if end-i <= segmentSize {
			segments = append(segments, Segment{startTime: i, endTime: end})
		} else {
			segments = append(segments, Segment{startTime: i, endTime: i + segmentSize})
		}
	}

	return segments
}

func Divider(wg *sync.WaitGroup, segmentDuration, filePath, outputPath string) {
	parsedDuration := helper.TimeConverter(segmentDuration)
	videoDuration := getVideoDuration(filePath)

	segments := generateSegments(0, videoDuration, parsedDuration)
	helper.ColorLog(segments)
	helper.ColorLog(videoDuration)
	v, err := video.Load(filePath)
	if err != nil {
		panic(err)
	}
	for i := range segments {
		val := segments[i]
		copy := i
		(*wg).Add(1)
		go func() {
			defer (*wg).Done()
			n, e := helper.GetFileNameAndExtension(filePath)
			o := outputPath + "/" + n
			if err := v.SubClip(val.startTime, val.endTime).Output(o + "-" + strconv.Itoa(copy) + "." + e).Run(); err != nil {
				panic("cut failed")
			}
		}()
	}
}
