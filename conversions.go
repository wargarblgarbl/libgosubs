package libgosubs

import (
	"strings"
	"strconv"
)

//Timecode Helper functions
//0:00:05.58 -> 00:02:20,476
//0:00:05.58 <- 00:02:20,476


//AssToSrtTimecode returns an srt timecode value
func AssToSrtTimecode(input string) (srt string) {
	if len(input) == 10 {
		input = "0" + input + "0"
	} else if len(input) >= 11 {
		input = input + "0"
	}
	srt = strings.Replace(input, ".", ",", -1)
	return
}



//SrtToAssTimecode returns an ass timecode value. 
func SrtToAssTimecode(input string)(ass string){
	//A horrible converter that manages to round integers up and down.

	a := strings.Split(input, ",")
	//The most garbage rounding function
	var output []string
	for i, r := range a[1] {
		//Return the string of the rune
		val := string(r)
		ival, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		//The timing percision on SRT is three digits, so the i val is 2.
		if i == 2 && ival >= 5 {
			//Join the first two runes
			c, err := strconv.Atoi(strings.Join(output, ""))
			if err != nil {
				panic(err)
			}
			//Urgh, we're clearing the function here'
			output = nil 
			output = append(output, strconv.Itoa(c+1))
		} else if i < 2 {
			output = append(output, val)
		}
	
	}
	a[1] = strings.Join(output, "")
	return strings.Join(a, ".")
}

//TtmlToSrtTimecode returns an srt timecode value. 
func TtmlToSrtTimecode(input string)(srt string){
	srt = strings.Replace(input, ".", ",", -1)
	return
}

//TtmlToSrtTimecode returns a ttml timecode value. 
func SrtToTtmlTimecode(input string)(ttml string) {
	ttml = strings.Replace(input, ",", ".", -1)
	return
}

//AssToTtmlTimecode returns a ttml timecode value
func AssToTtmlTimecode(input string)(ttml string) {
	ttml = SrtToTtmlTimecode(AssToSrtTimecode(input))
	return
}

//TtmlToAssTimecode returns an ass timecode value
func TtmlToAssTimecode(input string)(ass string) {
	ass = SrtToAssTimecode(TtmlToSrtTimecode(input))
	return
}


