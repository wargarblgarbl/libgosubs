package webvtt

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//CreateSubtitle creates a VTT formatted subtitle
func CreateSubtitle(note bool, cue string, start string, end string, text []string, pos Position) *Subtitle {
	return &Subtitle{
		Note:     note,
		Cue:      cue,
		Start:    start,
		End:      end,
		Position: pos,
		Line:     text,
	}
}

//CreatePosition creates a VTT subtiotle position
func CreatePosition(vertical string, line int, pos int, lpercent bool, align string, size int) *Position {
	return &Position{
		Vertical:    vertical,
		Line:        line,
		Position:    pos,
		Linepercent: lpercent,
		Align:       align,
		Size:        size,
	}
}

/*HELPER FUNCTIONS*/
//a quick function to int strings
func intit(in string) (out int) {
	if in == "" {
		out = 0
	} else {
		outa, err := strconv.Atoi(in)
		if err != nil {
			panic(err)
		}
		out = outa
	}
	return
}

//a quick function to parse the timecode
func parsetimecode(tc string) (start string, end string, pos bool) {
	pos = false
	split := strings.Split(tc, " ")
	start = split[0]
	end = split[2]
	if len(split) > 2 {
		pos = true
	}
	return
}

//a function to parse the positioning part of a timecode
func parsepos(tc string) (vertical string, line int, pos int, lpercent bool, align string, size int) {
	split := strings.Split(tc, " ")
	for _, a := range split {
		b := strings.Split(a, ":")
		switch b[0] {
		case "line":
			if strings.Contains(b[1], "%") {
				lpercent = true
				line = intit(strings.Replace(b[1], "%", "", -1))
			} else {
				line = intit(b[1])
			}
		case "position":
			if strings.Contains(b[1], "%") {
				pos = intit(strings.Replace(b[1], "%", "", -1))
			}
		case "size":
			if strings.Contains(b[1], "%") {
				size = intit(strings.Replace(b[1], "%", "", -1))
			}
		case "align":
			align = b[1]
		case "vertical":
			vertical = b[1]
		}
	}
	return
}

//check line - checks line to see if it's a header, a note, or a subtitle
func checkline(line []string) (out int) {
	for _, i := range line {
		if strings.Contains(i, "WEBVTT") {
			return 0
		} else if strings.Contains(i, "NOTE") {
			return 1
		} else if strings.Contains(i, "-->") {
			return 2
		}
	}
	return
}

//LoadWebVtt loads a WebVtt file
func LoadWebVtt(v *WebVtt, filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	//scanner.Split(bufio.ScanLines)
	var file [][]string
	var lines []string
	for scanner.Scan() {
		if scanner.Text() != "" {
			lines = append(lines, scanner.Text())
		} else {
			file = append(file, lines)
			lines = nil
		}
	}
	file = append(file, lines)
	for _, i := range file {
		parsed := checkline(i)
		switch parsed {
		case 0:
			v.Header = strings.Join(i, "")
		case 1:
			var pos Position
			v.Subtitle.Content = append(v.Subtitle.Content, *CreateSubtitle(true, "", "", "", i, pos))
		case 2:
			if err != nil {
				return err
			}
			start, end, haspos := parsetimecode(i[1])
			var pos Position
			if haspos {
				pos = *CreatePosition(parsepos(i[1]))
			}
			v.Subtitle.Content = append(v.Subtitle.Content, *CreateSubtitle(false, i[0], start, end, i[2:], pos))
		}

	}
	defer f.Close()
	return nil
}

//ParseWebVtt takes a filename and returns a WebVtt structure and any errors
func ParseWebVtt(filename string) (*WebVtt, error) {
	v := &WebVtt{}
	err := LoadWebVtt(v, filename)
	if err != nil {
		return v, err
	}
	return v, nil
}