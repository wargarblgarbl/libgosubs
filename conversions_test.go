package libgosubs
import (
	"testing"
	"fmt"
)

func TestTimecodeAtS(t *testing.T) {
	wat := AssToSrtTimecode("0:00:05.58")
	if wat != "00:00:05,580" {
		t.Errorf("Timecode not converted correctly between ASS and SRT")
	}
	fmt.Println(wat)
}

func TestTimecodeStA(t *testing.T) {
	wat := SrtToAssTimecode("00:02:20,474")
	if wat != "00:02:20.47" {
		t.Errorf("Timecode not correctly rounded down")
	}
	wat2 := SrtToAssTimecode("00:02:20,475")
	if wat2 != "00:02:20.48" {
		t.Errorf("Timecode not correctly rounded up")
	}
	
}

func TestTimcodeTtS(t *testing.T) {
	wat := TtmlToSrtTimecode("00:00:05.580")
	if wat != "00:00:05,580" {
		t.Errorf("Somehow, we failed at replacing a . with a ,")
	}
}

func TestTimcodeStT(t *testing.T) {
	wat := SrtToTtmlTimecode("00:00:05,580")
	if wat != "00:00:05.580" {
		t.Errorf("Somehow, we failed at replacing a . with a ,")
	}
}
