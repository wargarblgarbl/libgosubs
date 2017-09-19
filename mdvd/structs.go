package mdvd

type Mdvd struct {
	Body []Event
}

type Event struct {
	Start     int64
	End       int64
	Tags      []Tag
	IsDefault bool
	Text      string
}

type Tag struct {
	Type  string
	Value string
}
