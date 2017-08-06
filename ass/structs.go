package ass

type Ass struct {
	//Script info portion
	ScriptInfo struct {
			   Header string
			   Body Scriptinfo
		   }
	//Aegisub project garbage portion
	PGarbage struct {
			   Header string
			   Body Projectgarbage
		   }
	//Styles portion
	Styles struct {
			   Header string
			   Format string
			   Body []Style
		   }
	//Events which are usually subtitles or commented out lines.
	Events struct {
			   Header string
			   Format string
			   Body []Event
		   }
}

//Info for the script.
type Scriptinfo struct {
	Title string
	ScriptType string
	WarpStyle string
	SBaShadow string
	YCbCrMatrix string
	PlayResX int
	PlayResY int
}

//Aegisub Project Garbage
//Generally useless, untill it suddenly isn't.
//Players will not care about this, this section is only here for compatability.
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

//A script can have multiple styles.
type Style struct {
	Format string
	Name string
	Fontname string
	Fontsize int
	PrimaryColour string
	SecondaryColour string
	OutlineColour string
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

//A script will definitely have multiple events.
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
