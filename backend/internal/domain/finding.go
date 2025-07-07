package domain

type Finding struct {
	Rule    string
	File    string
	Message string
	Line    int
}