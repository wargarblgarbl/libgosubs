package srt

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parsetimecode(tc string) (start string, end string) {
	split := strings.Split(tc, " ")
	start = split[0]
	end = split[2]
	return
}

func checkline(line []string) (isline bool) {
	for _, i := range line {
		if strings.Contains(i, "-->") {
			isline = true
		}
	}
	return
}

//CreateSubtitle creates a subtitle object.
func CreateSubtitle(id int, start string, end string, text []string) *Subtitle {
	return &Subtitle{
		Id:    id,
		Start: start,
		End:   end,
		Line:  text,
	}
}

//LoadSrt loads the provided file into the given object.
//It fixes the \ufeff problem that some parsers have.
func LoadSrt(v *SubRip, filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	var file [][]string
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "\ufeff") {
			line = strings.Replace(line, "\ufeff", "", -1)
			line = strings.Replace(line, "\xef\xbb\xbf", "", -1)

		}
		if strings.HasPrefix(line, "\xef\xbb\xbf") {
			line = strings.Replace(line, "\ufeff", "", -1)
			line = strings.Replace(line, "\xef\xbb\xbf", "", -1)
		}
		if line != "" && line != "\ufeff" {
			lines = append(lines, line)
		} else {
			file = append(file, lines)
			lines = nil
		}
	}
	file = append(file, lines)
	for _, i := range file {
		parsed := checkline(i)
		if parsed {
			if err != nil {
				return err
			}
			id, err := strconv.Atoi(i[0])
			if err != nil {
				return err
			}
			start, end := parsetimecode(i[1])
			v.Subtitle.Content = append(v.Subtitle.Content, *CreateSubtitle(id, start, end, i[2:]))
		}
	}

	return nil
}

//ParseSrt is the loader for srt files. Takes the path of the file being opened as the argument.
func ParseSrt(filename string) (*SubRip, error) {
	v := &SubRip{}
	err := LoadSrt(v, filename)
	if err != nil {
		return v, err
	}
	return v, nil
}
