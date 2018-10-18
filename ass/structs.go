package ass

//Ass is the main struct for an .ass subtitle file.
type Ass struct {
	//Script info portion
	ScriptInfo struct {
		Header string
		Body   Scriptinfo
	}
	//Aegisub project garbage portion
	PGarbage struct {
		Header string
		Body   Projectgarbage
	}
	//Styles portion
	Styles struct {
		Header string
		Format string
		Body   []Style
	}
	//Events which are usually subtitles or commented out lines.
	Events struct {
		Header string
		Format string
		Body   []Event
	}
}

//Scriptinfo outlines the script information for the format
type Scriptinfo struct {
	Title       string
	ScriptType  string
	WarpStyle   string
	SBaShadow   string
	YCbCrMatrix string
	PlayResX    int
	PlayResY    int
}

//Projectgarbage - aegisub project garbage section
//Generally useless, until it suddenly isn't.
//Players will not care about this, this section is only here for compatibility.
type Projectgarbage struct {
	AudioFile        string
	VideoFile        string
	VideoARMode      string
	VideoARValue     float64
	VideoZoomPercent float64
	ScrollPosition   int
	ActiveLine       int
	VideoPos         int
}

//Style is a struct for the ass styles
//A script can have multiple styles.
type Style struct {
	Format          string
	Name            string
	Fontname        string
	Fontsize        float64
	PrimaryColour   string
	SecondaryColour string
	OutlineColour   string
	Backcolour      string
	Bold            int
	Italic          int
	Underline       int
	StrikeOut       int
	ScaleX          float64
	ScaleY          float64
	Spacing         float64
	Angle           float64
	BorderStyle     int
	Outline         float64
	Shadow          float64
	Alignment       int
	MarginL         int
	MarginR         int
	MarginV         int
	Encoding        int
}

//Event contains all of the required variables for an event.
//Start and End are required, everything else is optional
//A script will definitely have multiple events.
type Event struct {
	Format  string
	Layer   int
	Start   string
	End     string
	Style   string
	Name    string
	MarginL int
	MarginR int
	MarginV int
	Effect  string
	Text    string
}
