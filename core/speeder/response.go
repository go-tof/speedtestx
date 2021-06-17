package speeder

import (
	"encoding/json"
	"time"
)

type Ping struct {
	Jitter  float64 `json:"jitter"`
	Latency float64 `json:"latency"`
}

type Interface struct {
	InternalIP string `json:"internalIp"`
	Name       string `json:"name"`
	MacAddr    string `json:"macAddr"`
	IsVpn      bool   `json:"isVpn"`
	ExternalIP string `json:"externalIp"`
}

type Server struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Country  string `json:"country"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	IP       string `json:"ip"`
}

type Result struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// Speeder run command network check bandwidth output structure.
type Speeder struct {
	Type      string           `json:"type"`
	Timestamp time.Time        `json:"timestamp"`
	Ping      Ping             `json:"ping"`
	Download  BandwidthSummary `json:"download"`
	Upload    BandwidthSummary `json:"upload"`
	Isp       string           `json:"isp"`
	Interface Interface        `json:"interface"`
	Server    Server           `json:"server"`
	Result    Result           `json:"result"`
}

// NewSpeederWithBytes initialize speeder with bytes data.
func NewSpeederWithBytes(data []byte) (*Speeder, error) {
	var speeder *Speeder
	err := json.Unmarshal(data, &speeder)
	if err != nil {
		return nil, err
	}
	return speeder, nil
}

// DownloadSummary speedtest network check result for download summary data.
func (s *Speeder) DownloadSummary() BandwidthSummary {
	return s.Download
}

// DownloadBandwidth speedtest network check result for download bandwidth size(default unit is byte(B)).
func (s *Speeder) DownloadBandwidth() int {
	return s.Download.Bandwidth
}

// DownloadBandwidthKBs speedtest network check result for download bandwidth size with KB/s.
func (s *Speeder) DownloadBandwidthKBs() int {
	return s.Download.KBs()
}

// DownloadBandwidthMBs speedtest network check result for download bandwidth size with MB/s.
func (s *Speeder) DownloadBandwidthMBs() int {
	return s.Download.MBs()
}

// DownloadBandwidthMBs speedtest network check result for download bandwidth size with GB/s.
func (s *Speeder) DownloadBandwidthGBs() int {
	return s.Download.GBs()
}

// UploadSummary speedtest network check result for upload summary data.
func (s *Speeder) UploadSummary() BandwidthSummary {
	return s.Upload
}

// UploadBandwidth speedtest network check result for upload bandwidth size(default unit is byte(B)).
func (s *Speeder) UploadBandwidth() int {
	return s.Upload.Bandwidth
}

// UploadBandwidthKBs speedtest network check result for upload bandwidth size with KB/s.
func (s *Speeder) UploadBandwidthKBs() int {
	return s.Upload.KBs()
}

// UploadBandwidthMBs speedtest network check result for upload bandwidth size with MB/s.
func (s *Speeder) UploadBandwidthMBs() int {
	return s.Upload.MBs()
}

// UploadBandwidthGBs speedtest network check result for upload bandwidth size with GB/s.
func (s *Speeder) UploadBandwidthGBs() int {
	return s.Upload.GBs()
}

// String structure convert to string.
func (s *Speeder) String() ([]byte, error) {
	return json.Marshal(s)
}
