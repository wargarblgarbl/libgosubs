# libgosubs
Golang library to read and write subtitles in the following formats

- Advanced SubStation Alpha v4
- SRT
- TTML v1.0 - This is based on the spec provided by Netflix in their documentation
- WebVTT experimental support
- MicroDVD experimental support
# notes

TTML is somewhat complex to implement in Go due to the way that Go handles XML namespaces. Until this issue is fixed, two different structs for reading and writing, as well as a lengthy conversion function will probably be necessary. See the test file for a sample (and probably poor) implementation.

# todo

- Clean up the ASSv4 format, specifically do something about the way headers are handled

# updates
- Experimental MicroDVD format support added

## libgosubs project garbage

### Documentation
Available via Godoc

|Godoc | Format |
| ------------- | ------------- |
|[![GoDoc](https://godoc.org/github.com/wargarblgarbl/libgosubs/ass?status.svg)](https://godoc.org/github.com/wargarblgarbl/libgosubs/ass) | ASS |
|[![GoDoc](https://godoc.org/github.com/wargarblgarbl/libgosubs/srt?status.svg)](https://godoc.org/github.com/wargarblgarbl/libgosubs/srt) | SRT |
|[![GoDoc](https://godoc.org/github.com/wargarblgarbl/libgosubs/ttml?status.svg)](https://godoc.org/github.com/wargarblgarbl/libgosubs/ttml) | TTML  |
|[![GoDoc](https://godoc.org/github.com/wargarblgarbl/libgosubs/wvtt?status.svg)](https://godoc.org/github.com/wargarblgarbl/libgosubs/wvtt) | WVTT  |
|[![GoDoc](https://godoc.org/github.com/wargarblgarbl/libgosubs/mdvd?status.svg)](https://godoc.org/github.com/wargarblgarbl/libgosubs/mdvd) | MicroDVD |

### Test Coverage

| Coverage | Format |
| ------------- | ------------- |
|![gocover.io go](http://gocover.io/_badge/github.com/wargarblgarbl/libgosubs/ass)| ASS|
|![gocover.io go](http://gocover.io/_badge/github.com/wargarblgarbl/libgosubs/srt)| SRT|
|![gocover.io go](http://gocover.io/_badge/github.com/wargarblgarbl/libgosubs/ttml)| TTML|
|![gocover.io go](http://gocover.io/_badge/github.com/wargarblgarbl/libgosubs/wvtt)| WVTT|
|![gocover.io go](http://gocover.io/_badge/github.com/wargarblgarbl/libgosubs/mdvd)| MDVD| 

### Other
 [![Go Report Card](https://goreportcard.com/badge/github.com/wargarblgarbl/libgosubs)](https://goreportcard.com/report/github.com/wargarblgarbl/libgosubs)
