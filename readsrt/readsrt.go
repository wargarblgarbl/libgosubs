package readsrt
import (
	"os"
	"fmt"
	"strconv"
	"bufio"
	"strings"
)

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


//LoadSRT loads the provided file into the given object.
//It fixes the \ufeff problem that some parsers have. 
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
			//Clear object on newline
			v.subtitle.content = append(v.subtitle.content, *z)
			z = &Subtitle{}
		} else {
			fmt.Println("Error parsing .srt. Stray newline?")
		}
			
	}
	//Since the last subtitle often won't have a newline, append everything and clear object one last time
	v.subtitle.content = append(v.subtitle.content, *z)
	z = &Subtitle{}
	defer f.Close()	
}

//Generic loader for srt files. Takes the path of the file being opened as the argument. 
func ParseSrt(filename string) *SubRip {
	v := &SubRip{}
	LoadSrt(v, filename)
	return v
}

