package webvtt

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//CreateSubtitle creates a VTT formatted subtitle
func CreateSubtitle(note bool, cue string, start string, end string, text []string, pos Position, haspos bool) *Subtitle {
	return &Subtitle{
		Note:     note,
		Cue:      cue,
		Start:    start,
		End:      end,
		Position: pos,
		Line:     text,
		Haspos:   haspos,
	}
}

//CreatePosition creates a VTT subtiotle position
func CreatePosition(vertical string, line int, posstring string, pos int, lpercent bool, align string, size int) *Position {
	return &Position{
		Vertical:    vertical,
		Line:        line,
		Posstring:   posstring,
		Position:    pos,
		Linepercent: lpercent,
		Align:       align,
		Size:        size,
	}
}

//CreateStyle creates a style from a header string and an interface containing
//the variables
func CreateStyle(header string, variables map[string]interface{}) *Style {
	return &Style{
		Header: header,
		Value:  variables,
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
	split := strings.Split(tc, " ")
	start = split[0]
	end = split[2]
	if len(split) > 3 {
		pos = true
	}
	return
}

//a function to parse the positioning part of a timecode
func parsepos(tc string) (vertical string, line int, posstring string, pos int, lpercent bool, align string, size int) {
	split := strings.Split(tc, " ")
	line = -2
	pos = -2
	size = -2
	for _, a := range split {
		b := strings.Split(a, ":")
		for z, y := range b {
			if strings.Contains(y, ",") {
				t := strings.Split(y, ",")
				b[z] = t[0]
				posstring = t[1]
			}
		}
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
		} else if strings.Contains(i, "STYLE") {
			return 3
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
			haspos := false
			v.Subtitle.Content = append(v.Subtitle.Content, *CreateSubtitle(true, "", "", "", i, pos, haspos))
		case 2:
			if err != nil {
				return err
			}
			var start string
			var end string
			var haspos bool
			var hascue bool
			if strings.Contains(i[0], "-->") {
				start, end, haspos = parsetimecode(i[0])
			} else {
				start, end, haspos = parsetimecode(i[1])
				hascue = true
			}
			var pos Position
			if haspos && hascue {
				pos = *CreatePosition(parsepos(i[1]))
			} else {
				pos = *CreatePosition(parsepos(i[0]))
			}
			if strings.Contains(i[0], "-->") {
				v.Subtitle.Content = append(v.Subtitle.Content, *CreateSubtitle(false, "", start, end, i[1:], pos, haspos))
			} else {
				v.Subtitle.Content = append(v.Subtitle.Content, *CreateSubtitle(false, i[0], start, end, i[2:], pos, haspos))
			}
		case 3:
			var s Style
			m := make(map[string]interface{})
			s.Value = m
			header := strings.Replace(i[1], "{", "", -1)
			for _, h := range i[2:(len(i) - 1)] {
				j := strings.Split(h, ":")
				s.Value[j[0]] = strings.Replace(j[1], ";", "", -1)
			}
			v.Styles = append(v.Styles, *CreateStyle(header, s.Value))
			//Do the thing with the mapping our values
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
