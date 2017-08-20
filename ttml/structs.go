package ttml

/*
Not the same struct used in gass2ttml. Reasons?
Golang apparently doesn't like prefixed xml namespaces, and xml.Unmarshal() ignores them.
If you have them in here, it will ignore them completely.
This is going to be hell down the line when we have to encode things, because xml.Marshal() does work
with them. ARGH.
*/

//Tt read struct - outlines the file on read
type Tt struct {
	Xmlns        string `xml:"xmlns,attr"`
	XmlnsTtp     string `xml:"ttp,attr"`
	XmlnsTts     string `xml:"tts,attr"`
	XmlnsTtm     string `xml:"ttm,attr"`
	XmlnsXML     string `xml:"xml,attr"`
	TtpTimeBase  string `xml:"timeBase,attr"`
	TtpFrameRate string `xml:"frameRate,attr"`
	XMLLang      string `xml:"lang,attr"`
	Head         struct {
		Metadata struct {
			TtmTitle string `xml:"title"`
		} `xml:"metadata"`
		Styling struct {
			Style []Style `xml:"style"`
		} `xml:"styling"`
		Layout struct {
			Region []Region `xml:"region"`
		} `xml:"layout"`
	} `xml:"head"`
	Body struct {
		Region string `xml:"region,attr"`
		Style  string `xml:"style,attr"`
		Div    struct {
			P []Subtitle `xml:"p"`
		} `xml:"div"`
	} `xml:"body"`
}
//Region struct
type Region struct {
	XMLID           string `xml:"id,attr"`
	TtsDisplayAlign string `xml:"displayAlign,attr"`
	TtsExtent       string `xml:"extend,attr"`
	TtsOrigin       string `xml:"origin,attr"`
}

//Style struct
type Style struct {
	XMLID         string `xml:"id,attr"`
	TtsTextAlign  string `xml:"textAlign,attr"`
	TtsFontFamily string `xml:"fontFamily,attr"`
	TtsFontSize   string `xml:"fontSize,attr"`
}

//Subtitle - struct for subtitles
type Subtitle struct {
	Id     string `xml:"id,attr"`
	Begin  string `xml:"begin,attr"`
	End    string `xml:"end,attr"`
	Style  string `xml:"style,attr,omitempty"`
	Region string `xml:"region,attr,omitempty"`
	Text   string `xml:",innerxml"`
}

//WTt is the Write TTML struct.
//Primary reason for using a seperate struct is Go's strange handling of complex XML parameters
type WTt struct {
	Xmlns        string `xml:"xmlns,attr"`
	XmlnsTtp     string `xml:"xmlns:ttp,attr"`
	XmlnsTts     string `xml:"xmlns:tts,attr"`
	XmlnsTtm     string `xml:"xmlns:ttm,attr"`
	XmlnsXML     string `xml:"xmlns:xml,attr"`
	TtpTimeBase  string `xml:"ttp:timeBase,attr"`
	TtpFrameRate string `xml:"ttp:frameRate,attr"`
	XMLLang      string `xml:"xml:lang,attr"`
	Head         struct {
		Metadata struct {
			TtmTitle string `xml:"ttm:title"`
		} `xml:"metadata"`
		Styling struct {
			Style []Style `xml:"style"`
		} `xml:"styling"`
		Layout struct {
			Region []Region `xml:"region"`
		} `xml:"layout"`
	} `xml:"head"`
	Body struct {
		Region string `xml:"region,attr"`
		Style  string `xml:"style,attr"`
		Div    struct {
			P []Subtitle `xml:"p"`
		} `xml:"div"`
	} `xml:"body"`
}
