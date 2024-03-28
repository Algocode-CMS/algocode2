package common

type JudgeData struct {
	EjudgeUrl       string
	CodeforcesUrl   string `gorm:"default:'https://codeforces.com'"`
	PcmsUrl         string
	ExternalGroupID string
}
