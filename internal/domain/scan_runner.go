package domain

type ScanRunner interface {
	RunScan(scan *Scan, analyzers []Analyzer)
}