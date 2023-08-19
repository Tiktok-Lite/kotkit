package tools

import (
	"github.com/Tiktok-Lite/kotkit/pkg/helper/tools"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestScreenshot(t *testing.T) {
	_, err := tools.GetScreenshotBuffer("/Users/century/Desktop/circle.mp4", 1)
	assert.Equal(t, err, nil)
}
