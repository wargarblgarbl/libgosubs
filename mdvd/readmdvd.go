package mdvd

import (
	//	"fmt"
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//helperfunctions

func uint(in string) (out int64, err error) {
	out, err = strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0, err
	}
	return
}

//CreateEvent Creates a microDVD event
func CreateEvent(start, end int64, tags []Tag, def bool, text string) *Event {
	return &Event{
		Start:     start,
		End:       end,
		Tags:      tags,
		IsDefault: def,
		Text:      text,
	}
}

//CreateDefault creates a Default event which does not have a timestamp
func CreateDefault(tags []Tag, text string) *Event {
	return &Event{
		IsDefault: true,
		Tags:      tags,
		Text:      text,
	}
}

//CreateTag creates a tag object
func CreateTag(tag string) *Tag {
	split := strings.Split(tag, ":")
	return &Tag{
		Type:  split[0],
		Value: split[1],
	}
}

//TagBash converts a set of []strings to a set of Tags
func TagBash(in []string) (tags []Tag) {
	for _, i := range in {
		tags = append(tags, *CreateTag(i))
	}
	return
}

//ParseLine parses an individual Microdvd line file
func ParseLine(in string) (start, end int64, text string, tags []string, err error) {
	rgx := regexp.MustCompile(`\{(.*?)\}`)
	rs := rgx.Split(in, -1)
	ts := rgx.FindAllStringSubmatch(in, -1)
	text = rs[len(rs)-1]
	start, err = uint(ts[0][1])
	end, err = uint(ts[1][1])
	if len(ts) > 2 {
		for _, y := range ts[2:] {
			tags = append(tags, y[1])

		}
	}
	return
}

//ParseDefault parses the default line
func ParseDefault(in string) (text string, tags []string, err error) {
	rgx := regexp.MustCompile(`\{(.*?)\}`)
	rs := rgx.Split(in, -1)
	text = rs[len(rs)-1]
	ts := rgx.FindAllStringSubmatch(in, -1)
	if len(ts) > 1 {
		for _, y := range ts[1:] {
			tags = append(tags, y[1])
		}
	}
	return
}

//LoadMdvd loads a file into an mdvd object
func LoadMdvd(v *Mdvd, filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "{DEFAULT}") {
			text, tags, err := ParseDefault(scanner.Text())
			if err != nil {
				return (err)
			}
			v.Body = append(v.Body, *CreateDefault(TagBash(tags), text))
		} else {
			start, end, text, tags, err := ParseLine(scanner.Text())
			if err != nil {
				return (err)
			}
			v.Body = append(v.Body, *CreateEvent(start, end, TagBash(tags), false, text))
		}
	}
	return nil
}

//ParseMdvd parses a file into a Mdvd object
func ParseMdvd(filename string) (*Mdvd, error) {
	v := &Mdvd{}
	err := LoadMdvd(v, filename)
	if err != nil {
		return v, err
	}
	return v, nil
}
