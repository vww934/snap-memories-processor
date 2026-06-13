package processor

func MergeVideoOverlay(
	videoPath string,
	overlayPath string,
	outputPath string,
) error {

	return RunFFmpeg(
		"-y",

		"-i",
		videoPath,

		"-i",
		overlayPath,

		"-filter_complex",
		"[1:v][0:v]scale2ref[ov][base];[base][ov]overlay=0:0",

		"-codec:a",
		"copy",

		outputPath,
	)
}
