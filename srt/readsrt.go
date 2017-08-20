package srt

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

//LoadSRT loads the provided file into the given object.
//It fixes the \ufeff problem that some parsers have.
func LoadSrt(v *SubRip, filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return (err)
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
		if err == nil && z.Id == 0 {
			z.Id = int(i)
			//A bit more jackassery, because if z.Start and z.End are set, then welp.
		} else if strings.Contains(line, "-->") && z.Start == "" && z.End == "" {
			split := strings.Split(line, "-->")

			//Oh hey, we occasionally can pick up extra whitespace, let's strip it
			z.Start = strings.Replace(split[0], " ", "", -1)
			z.End = strings.Replace(split[1], " ", "", -1)
		} else if line != "" && z.Start != "" && z.End != "" && z.Id != 0 {
			z.Line = append(z.Line, line)
		} else if line == "" {
			//Clear object on newline
			if z.Start != "" && z.End != "" && z.Line != nil {
				v.Subtitle.Content = append(v.Subtitle.Content, *z)
			}
			z = &Subtitle{}
		} else {
			//At some point, we need to start actually returning errors.
			//Wouldn't that be nice?
			return errors.New("Error parsing .srt. Stray newline?")
		}

	}
	//Since the last subtitle often won't have a newline, append everything and clear object one last time
	v.Subtitle.Content = append(v.Subtitle.Content, *z)
	z = &Subtitle{}
	defer f.Close()
	return nil
}

//ParseSrt is the loader for srt files. Takes the path of the file being opened as the argument.
func ParseSrt(filename string) *SubRip {
	v := &SubRip{}
	err := LoadSrt(v, filename)
	if err != nil {
		panic(err)
	}
	return v
}
