package ass

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//WriteAss takes an Ass object and the path to where to write the file
func WriteAss(v *Ass, outpath string) error {
	f, err := os.Create(outpath)
	if err != nil {
		return (err)
	}
	var outout []string
	// Write header
	outout = append(outout, v.ScriptInfo.Header)
	outout = append(outout, "Title: "+v.ScriptInfo.Body.Title)
	outout = append(outout, "ScriptType: "+v.ScriptInfo.Body.ScriptType)
	outout = append(outout, "WrapStyle: "+v.ScriptInfo.Body.WarpStyle)
	outout = append(outout, "ScaledBorderAndShadow: "+v.ScriptInfo.Body.SBaShadow)
	outout = append(outout, "YCbCr Matrix: "+v.ScriptInfo.Body.YCbCrMatrix)
	outout = append(outout, "PlayResX: "+strconv.Itoa(v.ScriptInfo.Body.PlayResX))
	outout = append(outout, "PlayResY: "+strconv.Itoa(v.ScriptInfo.Body.PlayResY))
	outout = append(outout, "\n"+v.PGarbage.Header)
	outout = append(outout, "Audio File: "+v.PGarbage.Body.AudioFile)
	outout = append(outout, "Video File: "+v.PGarbage.Body.VideoFile)
	outout = append(outout, "Video AR Mode: "+v.PGarbage.Body.VideoARMode)
	outout = append(outout, "Video AR Value: "+strconv.FormatFloat(v.PGarbage.Body.VideoARValue, 'f', -1, 64))
	outout = append(outout, "Video Zoom Percent: "+strconv.FormatFloat(v.PGarbage.Body.VideoZoomPercent, 'f', -1, 64))
	outout = append(outout, "Scroll Position: "+strconv.Itoa(v.PGarbage.Body.ScrollPosition))
	outout = append(outout, "Active Line: "+strconv.Itoa(v.PGarbage.Body.ActiveLine))
	outout = append(outout, "Video Position: "+strconv.Itoa(v.PGarbage.Body.VideoPos))
	outout = append(outout, "\n"+v.Styles.Header)
	outout = append(outout, v.Styles.Format)

	//Super ugly solution for merging entire object as deliniated by a string. There has to be a way to do this in a sane, programmatic
	// and idiomatic manner.

	for _, z := range v.Styles.Body {

		/*
			 "Format: Name, Fontname, Fontsize,
			PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold,
			Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle,
			Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding"
		*/
		var astyle []string
		//astyle = append(astyle, z.Format)
		astyle = append(astyle, z.Name)
		astyle = append(astyle, z.Fontname)
		astyle = append(astyle, strconv.Itoa(z.Fontsize))
		astyle = append(astyle, z.PrimaryColour)
		astyle = append(astyle, z.SecondaryColour)
		astyle = append(astyle, z.OutlineColour)
		astyle = append(astyle, z.Backcolour)
		astyle = append(astyle, strconv.Itoa(z.Bold))
		astyle = append(astyle, strconv.Itoa(z.Italic))
		astyle = append(astyle, strconv.Itoa(z.Underline))
		astyle = append(astyle, strconv.Itoa(z.StrikeOut))
		astyle = append(astyle, strconv.Itoa(z.ScaleX))
		astyle = append(astyle, strconv.Itoa(z.ScaleY))
		astyle = append(astyle, strconv.Itoa(z.Spacing))
		astyle = append(astyle, strconv.Itoa(z.Angle))
		astyle = append(astyle, strconv.Itoa(z.BorderStyle))
		astyle = append(astyle, strconv.Itoa(z.Outline))
		astyle = append(astyle, strconv.Itoa(z.Shadow))
		astyle = append(astyle, strconv.Itoa(z.Alignment))
		astyle = append(astyle, strconv.Itoa(z.MarginL))
		astyle = append(astyle, strconv.Itoa(z.MarginR))
		astyle = append(astyle, strconv.Itoa(z.MarginV))
		astyle = append(astyle, strconv.Itoa(z.Encoding))
		outout = append(outout, z.Format+": "+strings.Join(astyle, ",")+"\n")

	}

	outout = append(outout, v.Events.Header)
	outout = append(outout, v.Events.Format)
	for _, e := range v.Events.Body {
		var anevent []string
		/*
			Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
		*/
		anevent = append(anevent, strconv.Itoa(e.Layer))
		anevent = append(anevent, e.Start)
		anevent = append(anevent, e.End)
		anevent = append(anevent, e.Style)
		anevent = append(anevent, e.Name)
		anevent = append(anevent, strconv.Itoa(e.MarginL))
		anevent = append(anevent, strconv.Itoa(e.MarginR))
		anevent = append(anevent, strconv.Itoa(e.MarginV))
		anevent = append(anevent, e.Effect)
		anevent = append(anevent, e.Text)

		outout = append(outout, e.Format+": "+strings.Join(anevent, ","))
	}

	fmt.Fprint(f, "", strings.Join(outout, "\n")+"\n")
	return nil
}
