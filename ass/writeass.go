package ass
import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func WriteAss(v *Ass, outpath string) {
	f, err := os.Create(outpath)
	if err != nil {
		fmt.Println(err)
	}
	var outout []string
	// Write header
	outout = append(outout, v.ScriptInfo.Header)
	outout = append(outout, "Title: "+v.ScriptInfo.Body.Title)
	outout = append(outout, "ScriptType: "+v.ScriptInfo.Body.ScriptType)
	outout = append(outout,  "WrapStyle: "+v.ScriptInfo.Body.WarpStyle )
	outout = append(outout,  "ScaledBorderAndShadow: "+v.ScriptInfo.Body.SBaShadow )
	outout = append(outout,  "YCbCr Matrix: "+v.ScriptInfo.Body.YCbCrMatrix )
	outout = append(outout,  "PlayResX: "+strconv.Itoa(v.ScriptInfo.Body.PlayResX))
	outout = append(outout,  "PlayResY: "+strconv.Itoa(v.ScriptInfo.Body.PlayResY))
	outout = append(outout, v.PGarbage.Header)
	outout = append(outout,  "Audio File: "+v.PGarbage.Body.AudioFile )
	outout = append(outout,  "Video File: "+v.PGarbage.Body.VideoFile )
	outout = append(outout,  "Video AR Mode: "+v.PGarbage.Body.VideoARMode)
	outout = append(outout,  "Video AR Value: "+strconv.FormatFloat(v.PGarbage.Body.VideoARValue, 'E', -1, 64))
	outout = append(outout,  "Video Zoom Percent: "+strconv.FormatFloat(v.PGarbage.Body.VideoZoomPercent, 'E', -1, 64) )
	outout = append(outout,  "Scroll Position: "+strconv.Itoa(v.PGarbage.Body.ScrollPosition) )
	outout = append(outout,  "Active Line: "+strconv.Itoa(v.PGarbage.Body.ActiveLine) )
	outout = append(outout,  "Video Position: "+strconv.Itoa(v.PGarbage.Body.VideoPos) )
	outout = append(outout, v.Styles.Header)
	outout = append(outout, v.Styles.Format)

	//Super ugly solution for merging entire object as deliniated by a string. There has to be a way to do this in a sane, programmatic
	// and idiomatic manner. 
	
	for _, z := range v.Styles.Body {
		var astyle string
		astyle = z.Format+","+z.Name+","+z.Fontname+","+strconv.Itoa(z.Fontsize)
		astyle = z.PrimaryColour+","+z.SecondaryColour
		astyle = z.OutlineColour+","+z.Backcolour+","+strconv.Itoa(z.Bold)+","+strconv.Itoa(z.Italic)
		astyle = strconv.Itoa(z.Underline)+","+strconv.Itoa(z.StrikeOut)+","+strconv.Itoa(z.ScaleX)+","+strconv.Itoa(z.ScaleY)
		astyle = strconv.Itoa(z.Spacing)+","+strconv.Itoa(z.Angle)+","+strconv.Itoa(z.BorderStyle)+","+strconv.Itoa(z.Outline)
		astyle = strconv.Itoa(z.Shadow)+","+strconv.Itoa(z.Alignment)+","+strconv.Itoa(z.MarginL)+","+strconv.Itoa(z.MarginR)
		astyle = strconv.Itoa(z.MarginV)+","+strconv.Itoa(z.Encoding)+"\n"
		outout = append(outout, astyle)
		astyle = ""
	}

	outout = append(outout, v.Events.Header)
	outout = append(outout, v.Events.Format)
	for _, e := range v.Events.Body {
		var anevent string
		anevent = e.Format+","+strconv.Itoa(e.Layer)+","+e.Start+","+e.End
		anevent = e.Style+","+e.Name+","+strconv.Itoa(e.MarginL)
		anevent = strconv.Itoa(e.MarginR)+","+strconv.Itoa(e.MarginV)
		anevent = e.Effect+","+e.Text
		outout = append(outout, anevent)
		anevent = ""
	}

	fmt.Fprintf(f, "%", strings.Join(outout, ""))



}


