package ass
import (
	"os"
	"fmt"
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
	outout = append(outout,  "PlayResX: "+v.ScriptInfo.Body.PlayResX)
	outout = append(outout,  "PlayResY: "+v.ScriptInfo.Body.PlayResY)
	outout = append(outout, v.PGarbage.Header)
	outout = append(outout,  "Audio File: "+v.PGarbage.Body.AudioFile )
	outout = append(outout,  "Video File: "+v.PGarbage.Body.VideoFile )
	outout = append(outout,  "Video AR Mode: "+v.PGarbage.Body.VideoARMode )
	outout = append(outout,  "Video AR Value: "+v.PGarbage.Body.VideoARValue)
	outout = append(outout,  "Video Zoom Percent: "+v.PGarbage.Body.VideoZoomPercent )
	outout = append(outout,  "Scroll Position: "+v.PGarbage.Body.ScrollPosition )
	outout = append(outout,  "Active Line: "+v.PGarbage.Body.ActiveLine )
	outout = append(outout,  "Video Position: "+v.PGarbage.Body.VideoPos )
	outout = append(outout, v.Styles.Header)
	outout = append(outout, v.Styles.Format)
	var astyle string
	for _, z := range v.Styles.Body {
		for _, zz := range z {
			astyle = astyle+zz+","
		}
		outout = append(outout, astyle)
		astyle = ""
	}
	outout = append(outout, v.Events.Header)
	outout = append(outout, v.Events.Format)
	var anevent string
	for _, e := range v.Events.Body {
		for _, ee := range e {
			anevent = anevent+ee+","
		}
		outout = append(outout, anevent)
		anevent = ""
	}
	fmt.Fprintf(f, "%", string(outout))



}


