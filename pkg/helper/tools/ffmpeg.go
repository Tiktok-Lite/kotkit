package tools

import (
	"bytes"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
)

func GetScreenshot(videoPath, screenshotPath string, frameNum int) (string, error) {
	logger := log.Logger()

	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n, %d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		logger.Errorf("get screenshot error: %v", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		logger.Errorf("decode image error: %v", err)
		return "", err
	}

	err = imaging.Save(img, screenshotPath+".jpg")
	if err != nil {
		logger.Errorf("save image error: %v", err)
		return "", err
	}

	imgPath := screenshotPath + ".jpg"

	return imgPath, nil
}

func GetScreenshotBuffer(playURL string, frameNum int) (*bytes.Buffer, error) {
	logger := log.Logger()

	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(playURL).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n, %d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		logger.Errorf("get screenshot error: %v", err)
		return nil, err
	}

	return buf, nil
}
