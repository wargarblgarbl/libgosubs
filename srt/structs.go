package srt

//Subrip is the overall file descriptor.
//It provides a container for us to dump subtitles into.

type SubRip struct {
	subtitle struct {
			 content []Subtitle
		 }
}

//Subtitle struct provides all of the elements of an .srt subtitle
//with lines of subtitles being broken up into []strings
type Subtitle struct {
	id int
	start string
	end string
	line []string
}

