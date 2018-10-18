package ass

import (
	"bufio"
	//	"fmt"
	//	"errors"
	"os"
	"strconv"
	"strings"
)

//Helper functions
func floatit(in string) (out float64) {
	if in == "" {
		out = 0.0
	} else {
		outa, err := strconv.ParseFloat(in, 64)
		if err != nil {
			panic(err)
		}
		out = outa
	}
	return
}

func reversesplit(in []string) (out string) {
	out = strings.Join(in, ":")
	return
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

//Loadass loads the .ass file and parses out the various possible valid lines.
func Loadass(v *Ass, filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		splitter := strings.Split(scanner.Text(), ":")
		prefix := splitter[0]
		suffix := strings.Trim(reversesplit(splitter[1:]), " ")
		switch prefix {
		case "Title":
			v.ScriptInfo.Body.Title = suffix
		case "ScriptType":
			v.ScriptInfo.Body.ScriptType = suffix
		case "WrapStyle":
			v.ScriptInfo.Body.WarpStyle = suffix
		case "ScaledBorderAndShadow":
			v.ScriptInfo.Body.SBaShadow = suffix
		case "YCbCr Matrix":
			v.ScriptInfo.Body.YCbCrMatrix = suffix
		case "PlayResX":
			v.ScriptInfo.Body.PlayResX = intit(suffix)
		case "PlayResY":
			v.ScriptInfo.Body.PlayResY = intit(suffix)
		case "Audio File":
			v.PGarbage.Body.AudioFile = suffix
		case "Video File":
			v.PGarbage.Body.VideoFile = suffix
		case "Video AR Mode":
			v.PGarbage.Body.VideoARMode = suffix
		case "Video AR Value":
			v.PGarbage.Body.VideoARValue = floatit(suffix)
		case "Video Zoom Percent":
			v.PGarbage.Body.VideoZoomPercent = floatit(suffix)
		case "Scroll Position":
			v.PGarbage.Body.ScrollPosition = intit(suffix)
		case "Active Line":
			v.PGarbage.Body.ActiveLine = intit(suffix)
		case "Video Position":
			v.PGarbage.Body.VideoPos = intit(suffix)
		case "Style":
			v.Styles.Body = append(v.Styles.Body, *Parsestyle(suffix))
		case "Dialogue":
			v.Events.Body = append(v.Events.Body, *Parseevent(suffix, prefix))
		case "Comment":
			v.Events.Body = append(v.Events.Body, *Parseevent(suffix, prefix))
		}
	}
	return nil
}

//Parseevent creates the event from an event string split up on :.
//For example `Dialogue: 0,0:03:20.10,0:03:21.36,Default,,0,0,0,,` would parse to
//in = `0,0:03:20.10,0:03:21.36,Default,,0,0,0,`
//etype = Dialogue
func Parseevent(in string, etype string) *Event {
	split := strings.Split(in, ",")
	return &Event{
		Format:  etype,
		Layer:   intit(split[0]),
		Start:   split[1],
		End:     split[2],
		Style:   split[3],
		Name:    split[4],
		MarginL: intit(split[5]),
		MarginR: intit(split[6]),
		MarginV: intit(split[7]),
		Effect:  split[8],
		Text:    strings.Join(split[9:], ","),
	}
}

//Createevent creates event from a lit of elements
func Createevent(format string, layer int, start string, end string, style string, name string, marginl int, marginr int, marginv int, effect string, text string) *Event {
	return &Event{
		Format:  format,
		Layer:   layer,
		Start:   start,
		End:     end,
		Style:   style,
		Name:    name,
		MarginL: marginl,
		MarginR: marginr,
		MarginV: marginv,
		Effect:  effect,
		Text:    text,
	}
}

//Createstyle creates tyle from a list of elements.
func Createstyle(name string, fontname string, fontsize float64, pcolour string, scolour string, ocolour string, bcolour string, b int, i int, u int, so int, sx float64, sy float64, spacing float64, angle float64, bstyle int, outline float64, shadow float64, align int, marginl int, marginr int, marginv int, encoding int) *Style {
	return &Style{
		//Style is always static.
		Format:          "Style",
		Name:            name,
		Fontname:        fontname,
		Fontsize:        fontsize,
		PrimaryColour:   pcolour,
		SecondaryColour: scolour,
		OutlineColour:   ocolour,
		Backcolour:      bcolour,
		Bold:            b,
		Italic:          i,
		Underline:       u,
		StrikeOut:       so,
		ScaleX:          sx,
		ScaleY:          sy,
		Spacing:         spacing,
		Angle:           angle,
		BorderStyle:     bstyle,
		Outline:         outline,
		Shadow:          shadow,
		Alignment:       align,
		MarginL:         marginl,
		MarginR:         marginr,
		MarginV:         marginv,
		Encoding:        encoding,
	}
}

//Parsestyle - creates a style from a string. Takes a full .ass style line as an argument.
//Similar to Createevent, except Styles don't have multiple Formats, so we only take the format-less style string.
func Parsestyle(in string) *Style {
	split := strings.Split(in, ",")
	return &Style{
		//Style is always static.
		Format:          "Style",
		Name:            split[0],
		Fontname:        split[1],
		Fontsize:        floatit(split[2]),
		PrimaryColour:   split[3],
		SecondaryColour: split[4],
		OutlineColour:   split[5],
		Backcolour:      split[6],
		Bold:            intit(split[7]),
		Italic:          intit(split[8]),
		Underline:       intit(split[9]),
		StrikeOut:       intit(split[10]),
		ScaleX:          floatit(split[11]),
		ScaleY:          floatit(split[12]),
		Spacing:         floatit(split[13]),
		Angle:           floatit(split[14]),
		BorderStyle:     intit(split[15]),
		Outline:         floatit(split[16]),
		Shadow:          floatit(split[17]),
		Alignment:       intit(split[18]),
		MarginL:         intit(split[19]),
		MarginR:         intit(split[20]),
		MarginV:         intit(split[21]),
		Encoding:        intit(split[22]),
	}
}

//Setheaders sets default headers for the various fields.
//These are static and should not change in ass v4
func Setheaders(v *Ass) {
	v.ScriptInfo.Header = "[Script Info]"
	v.PGarbage.Header = "[Aegisub Project Garbage]"
	v.Styles.Header = "[V4+ Styles]"
	v.Events.Header = "[Events]"
	v.Styles.Format = "Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding"
	v.Events.Format = "Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text"
}

//ParseAss - Parses an .ass file to a structure
func ParseAss(filename string) (*Ass, error) {
	v := &Ass{}
	Setheaders(v)
	err := Loadass(v, filename)
	if err != nil {
		return v, err
	}
	return v, nil
}
