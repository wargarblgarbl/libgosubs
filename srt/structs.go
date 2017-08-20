package srt

//SubRip is the overall file descriptor.
//It provides a container for us to dump subtitles into.
type SubRip struct {
	Subtitle struct {
		Content []Subtitle
	}
}

//Subtitle struct provides all of the elements of an .srt subtitle
//with lines of subtitles being broken up into []strings
type Subtitle struct {
	Id    int
	Start string
	End   string
	Line  []string
}
