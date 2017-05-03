package readsrt
import (
	"os"
	"fmt"
	"strconv"
	"bufio"
	"strings"
)

type SubRip struct {
	subtitle struct {
		content []Subtitle
	}
}

type Subtitle struct {
	id int
	start string
	end string
	line []string
}


func LoadSrt(v *SubRip, filepath string) {
	f, err := os.Open(filepath)
	if err!= nil {
		fmt.Println("Cannot read file", filepath)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	z := &Subtitle{}
	for scanner.Scan() {
		line := scanner.Text()
		//Ufeff problem fix
		if strings.HasPrefix(line, "\ufeff") {
			line = strings.Replace(line, "\ufeff", "", -1)
		}
		//The most jackass SRT parser ever. Appends to an object. 
		i, err := strconv.Atoi(line)
		if err == nil && z.id == 0 {
			z.id = int(i)
		} else if strings.Contains(line, "-->")  {
			split := strings.Split(line, "-->")
			z.start = split[0]
			z.end = split[1]
		} else if line != "" && z.start != "" && z.end != "" && z.id != 0 {
			z.line = append(z.line, line)
		} else if line == ""  {			
			v.subtitle.content = append(v.subtitle.content, *z)
			z = &Subtitle{}
		}
	}
	v.subtitle.content = append(v.subtitle.content, *z)
	z = &Subtitle{}
	defer f.Close()	
}

func ParseSrt(filename string) *SubRip {
	v := &SubRip{}
	LoadSrt(v, filename)
	return v
}

