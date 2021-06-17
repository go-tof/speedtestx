package speeder

// UnitSize compute storage unit size.
const UnitSize = 1024

// BandwidthSummary bandwidth summary structure.
type BandwidthSummary struct {
	Bandwidth int `json:"bandwidth"` // default unit is byte(B).
	Bytes     int `json:"bytes"`
	Elapsed   int `json:"elapsed"`
}

// KBs bandwidth size for kB/s.
func (bw BandwidthSummary) KBs() int {
	return bw.Bandwidth / UnitSize
}

// MBs bandwidth size for MB/s.
func (bw BandwidthSummary) MBs() int {
	return bw.KBs() / UnitSize
}

// GBs bandwidth size for GB/s.
func (bw BandwidthSummary) GBs() int {
	return bw.MBs() / UnitSize
}
