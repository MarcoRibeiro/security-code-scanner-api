package valueObjects

type Finding struct {
	Rule, File, Message string
	Line                int
}