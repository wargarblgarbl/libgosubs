package webvtt
import (
	"bufio"
	//"errors"
	"os"
//	"strconv"
	"strings"
	"fmt"
//	"regexp"

)

func CreateSubtitle(id int, start string, end string, text []string, pos []Position) *Subtitle {
	return &Subtitle{
		Id:    id,
		Start: start,
		End:   end,
		Position: pos, 
		Line:  text,
	}
}

func CreatePosition(pos int, align string, size int) *Position {
	return &Position{
		Position: pos,
		Align: align,
		Size: size,
	}
}

func CreateWebVtt(*[]Subtitle, *WebVtt) *WebVtt {
	return &WebVtt{

	}
}

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
		if scanner.Text() == "" {
			lines = append(lines, "")
		}
		lines = append(lines, scanner.Text())
	}
	what := strings.Join(lines, "\n")
	why := strings.Split(what, "\n\n")
	for z, i := range why {
		fmt.Println(z, i)
	}
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
