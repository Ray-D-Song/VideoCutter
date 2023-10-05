This is a command-line-based video cropping tool. You can divide a video into n segments based on time units, or you can input a time axis to obtain a specific segment you desire.

> The time units for segmentation are in `hours:minutes:seconds`.

# How to use
## Divide into multiple segments.
```bash
cutter test.mp4 -d 00:12:00 outputPath
```
The video will be divided into several 12-minute short clips, named in sequential order as xxx-1.mp4.
## Retrieve a specific segment.
```bash
cutter test.mp4 -r 00:12:00 00:14:00 outputPath/outputName
```
The input time range will be segmented and named as outputName.mp4.