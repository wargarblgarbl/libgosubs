package readass 
import (
	"os"
	"fmt"
	"strings"
	"bufio"
	"strconv"
)

type Ass struct { 
	ScriptInfo struct {
		Header string
		Body Scriptinfo
	}
	PGarbage struct {
		Header string
		Body Projectgarbage
	}

	Styles struct {
		Header string
		Format string
		Body []Style
	}
	Events struct {
		Header string
		Format string
		Body []Event
	}
}

type Scriptinfo struct { 
	Title string
	ScriptType string
	WarpStyle string
	SBaShadow string
	YCbCrMatrix string
	PlayResX int
	PlayResY int
}

type Projectgarbage struct { 
	AudioFile string
	VideoFile string
	VideoARMode string
	VideoARValue float64
	VideoZoomPercent float64
	ScrollPosition int
	ActiveLine int
	VideoPos int
}

type Style struct {
	Format string
	Name string
	Fontname string
	Fontsize int
	PrimaryColour string
	SecondaryColour string
	OUtlineColour string
	Backcolour string
	Bold int
	Italic int
	Underline int
	StrikeOut int
	ScaleX int
	ScaleY int
	Spacing int
	Angle int
	BorderStyle int
	Outline int
	Shadow int
	Alignment int
	MarginL int
	MarginR int
	MarginV int
	Encoding int
}

type Event struct { 
	Format string
	Layer int
	Start string
	End string
	Style string
	Name string
	MarginL int
	MarginR int
	MarginV int
	Effect string
	Text string
}

func floatit(in string)(out float64) {
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


func reversesplit(in []string)(out string) {
	out = strings.Join(in, ":")
	return
}

func intit(in string)(out int) {
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

func Loadass(v *Ass, filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Cannot read file")
		os.Exit(1)
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
		case "Video Zoom Percent" :
			v.PGarbage.Body.VideoZoomPercent = floatit(suffix)
		case "Scroll Position" :
			v.PGarbage.Body.ScrollPosition = intit(suffix)
		case "Active Line" :
			v.PGarbage.Body.ActiveLine = intit(suffix)
		case "Video Position" :
			v.PGarbage.Body.VideoPos = intit(suffix)
		case "Style" :
			v.Styles.Body = append(v.Styles.Body, *Createstyle(suffix))
		case "Dialogue":
			v.Events.Body = append(v.Events.Body, *Createevent(suffix, prefix))
		case "Comment":
			v.Events.Body = append(v.Events.Body, *Createevent(suffix, prefix))
		}
	}
}


func Createevent(in string, etype string) *Event{
	split := strings.Split(in, ",")
	return &Event {
		Format: etype,
		Layer: intit(split[0]),
		Start: split[1],
		End: split[2],
		Style: split[3],
		Name: split[4],
		MarginL: intit(split[5]),
		MarginR: intit(split[6]),
		MarginV: intit(split[7]),
		Effect: split[8],
		Text: split[9],
	}
}

func Createstyle(in string) *Style{
	split := strings.Split(in, ",")
	return &Style {
		Format: "Style: ",
		Name: split[0],
		Fontname: split[1],
		Fontsize: intit(split[2]),
		PrimaryColour: split[3],
		SecondaryColour: split[4],
		OUtlineColour: split[5],
		Backcolour: split[6],
		Bold: intit(split[7]),
		Italic: intit(split[8]),
		Underline: intit(split[9]),
		StrikeOut: intit(split[10]),
		ScaleX: intit(split[11]),
		ScaleY: intit(split[12]),
		Spacing: intit(split[13]),
		Angle: intit(split[14]),
		BorderStyle: intit(split[15]),
		Outline: intit(split[16]),
		Shadow: intit(split[17]),
		Alignment: intit(split[18]),
		MarginL: intit(split[19]),
		MarginR: intit(split[20]),
		MarginV: intit(split[21]),
		Encoding: intit(split[22]),
	}
}

func Setheaders(v *Ass) {
	v.ScriptInfo.Header = "[Script Info]"
	v.PGarbage.Header = "[Aegisub Project Garbage]"
	v.Styles.Header = "[V4+ Styles]"
	v.Events.Header = "[Events]"
	v.Styles.Format = "Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding"
	v.Events.Format = "Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text"
		
}

func ParseAss(filename string) *Ass{
	v := &Ass{}
	Setheaders(v)
	Loadass(v, filename)
	return v
}
