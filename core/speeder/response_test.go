package speeder

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	bytes, err := ioutil.ReadFile("./speedtest_fail.txt")
	assert.NoError(t, err)

	speeder, err := NewSpeederWithBytes(bytes)
	assert.NoError(t, err)

	downloadBandwidth := speeder.DownloadBandwidth()
	assert.Equal(t, 4005863, downloadBandwidth)

	downloadBandwidthKBs := speeder.DownloadBandwidthKBs()
	assert.Equal(t, 3911, downloadBandwidthKBs)

	downloadBandwidthMBs := speeder.DownloadBandwidthMBs()
	assert.Equal(t, 3, downloadBandwidthMBs)

	DownloadBandwidthGBs := speeder.DownloadBandwidthGBs()
	assert.Equal(t, 0, DownloadBandwidthGBs)
}

func TestUpload(t *testing.T) {
	bytes, err := ioutil.ReadFile("./speedtest_fail.txt")
	assert.NoError(t, err)

	speeder, err := NewSpeederWithBytes(bytes)
	assert.NoError(t, err)

	uploadBandwidth := speeder.UploadBandwidth()
	assert.Equal(t, 5960041, uploadBandwidth)

	uploadBandwidthKBs := speeder.UploadBandwidthKBs()
	assert.Equal(t, 5820, uploadBandwidthKBs)

	uploadBandwidthMBs := speeder.UploadBandwidthMBs()
	assert.Equal(t, 5, uploadBandwidthMBs)

	uploadBandwidthGBs := speeder.UploadBandwidthGBs()
	assert.Equal(t, 0, uploadBandwidthGBs)
}
