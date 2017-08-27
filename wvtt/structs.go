package webvtt

//WebVtt is the container struct for a WebVtt subtitle
type WebVtt struct {
	Header   string
	Styles   []Style
	Subtitle struct {
		Content []Subtitle
	}
}

//Subtitle is the struct that stores subtitles and notes for WebVtt.
//The Note bool determines whether the Subtitle is actually a note or not.
type Subtitle struct {
	Note     bool
	Cue      string
	Start    string
	End      string
	Haspos   bool
	Position Position
	Line     []string
}

//Position is a struct that stores the positional data on a per-line level
type Position struct {
	Vertical    string
	Line        int
	Posstring   string
	Position    int
	Linepercent bool
	Align       string
	Size        int
}

//Style struct contains the style header information
// and a map of the contained values
type Style struct {
	Header string
	Value  map[string]interface{}
}
