package webvtt

type WebVtt struct {
	Header string
	Styles []Style
	Subtitle struct {
		Content []Subtitle
	}
}

type Subtitle struct {
	Note bool
	Cue    string
	Start string
	End   string
	Position Position
	Line  []string
}

type Position struct {
	Vertical string
	Line int
	Position int
	Linepercent bool
	Align string
	Size int
}

type Style struct{
}
