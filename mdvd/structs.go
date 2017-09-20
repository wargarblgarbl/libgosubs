package mdvd

//Mdvd is the main mdvd subtitle object
type Mdvd struct {
	Body []Event
}

//Event - these compose the body and are the main subtitle events
type Event struct {
	Start     int64
	End       int64
	Tags      []Tag
	IsDefault bool
	Text      string
}

//Tag - tags are formatting tags that contain styling information
type Tag struct {
	Type  string
	Value string
}
