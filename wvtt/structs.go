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
	Id    int
	Start string
	End   string
	Position []Position
	Line  []string
}

type Position struct {
	Position int
	Align string
	Size int
}

type Style struct{
}
