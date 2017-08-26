package webvtt
import (
	"bufio"
	//"errors"
	"os"
	"strconv"
	"strings"
	"fmt"
//	"regexp"

)

func CreateSubtitle(note bool, cue string, start string, end string, text []string, pos Position) *Subtitle {
	return &Subtitle{
		Note: note,
		Cue:    cue,
		Start: start,
		End:   end,
		Position: pos, 
		Line:  text,
	}
}

func CreatePosition(vertical string, line int, pos int, lpercent bool,  align string, size int) *Position {
	return &Position{
		Vertical: vertical,
		Line: line,
		Position: pos,
		Linepercent: lpercent,
		Align: align,
		Size: size,
	}
}

func CreateWebVtt(*[]Subtitle, *WebVtt) *WebVtt {
	return &WebVtt{

	}
}
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



func parsetimecode(tc string)(start string, end string, pos bool){
	pos = false
	split := strings.Split(tc, " ")
	start = split[0]
	end = split[2]
	if len(split) > 3 {
		pos = true
	}
	return
}

func parsepos(tc string)(vertical string, line int, pos int, lpercent bool, align string, size int) {
	split := strings.Split(tc, " ")
	fmt.Println(split)
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


func checkline(line []string)(out int) {
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

//func CreateSubtitle(note bool, id int, start string, end string, text []string, pos []Position) *Subtitle {
//func parsepos(tc st(vertical string, line int, lpercent bool, pos int, align string, size int) {
//func CreatePosition(vertical string, line int, pos int, lpercent bool,   align string, size int) *Position {

func LoadWebVtt(v *WebVtt, filepath string) error {
	f, err := os.Open(filepath)
	if err != nil{
		fmt.Println("whoops")
		return err
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines) 
  	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	join := strings.Join(lines, "\n")
	split := strings.Split(join, "\n\n")
	
	for _, i := range split {
		b := strings.Split(i, "\n")
		parsed := checkline(b)

		switch parsed {
		case 0:
			v.Header = strings.Join(b, "")
		case 1:
			var pos Position

			v.Subtitle.Content = append(v.Subtitle.Content, *CreateSubtitle(true, "", "", "", b, pos))
		case 2:
			id :=  b[0]
			if err != nil {
				return err
			}
			start, end, haspos := parsetimecode(b[1])
			var pos Position
			if haspos {
				pos = *CreatePosition(parsepos(b[1]))
			}
			v.Subtitle.Content = append(v.Subtitle.Content, *CreateSubtitle(false, id, start, end, b[2:], pos))
		}
		
		fmt.Println(parsed)
	}
	fmt.Println(v)
	return nil
}

func ParseWebVtt(filename string)(*WebVtt, error) {
	v := &WebVtt{}
	err := LoadWebVtt(v, filename)
	if err != nil {
		return v, err
	}
	return v, nil
}
