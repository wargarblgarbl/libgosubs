# libgosubs
Golang library to read and write subtitles in the following formats

- Advanced SubStation Alpha v4
- standard SRT
- TTML v1.0 - This is based on the spec provided by Netflix in their documentation

# notes

TTML is somewhat complex to implement in Go due to the way that Go handles XML namespaces. Until this issue is fixed, two different structs for reading and writing, as well as a lengthy conversion function will probably be necessary. See the test file for a sample (and probably poor) implementation. 

# todo

- Clean up the ASSv4 format, specifically do something about the way headers are handled
- Implement WebVtt - should be approximately the same as SRT. 
- Look into a better way of handling SRT, currently the parser has somewhat ugly handling of whitespace in a file. 


## libgosubs project garbage

--------
# Documentation
Available via Godoc

|Godoc | Format | 
| ------------- | ------------- |
|[![GoDoc](https://godoc.org/github.com/wargarblgarbl/libgosubs/ass?status.svg)](https://godoc.org/github.com/wargarblgarbl/libgosubs/ass) | ASS |
|[![GoDoc](https://godoc.org/github.com/wargarblgarbl/libgosubs/srt?status.svg)](https://godoc.org/github.com/wargarblgarbl/libgosubs/srt) | SRT |
|[![GoDoc](https://godoc.org/github.com/wargarblgarbl/libgosubs/ttml?status.svg)](https://godoc.org/github.com/wargarblgarbl/libgosubs/ttml) | TTML  |

--------
# Test Coverage

| Coverage | Format | 
| ------------- | ------------- |
|![cover.run go](https://cover.run/go/github.com/wargarblgarbl/libgosubs/ass.svg)| ASS| 
|![cover.run go](https://cover.run/go/github.com/wargarblgarbl/libgosubs/srt.svg)| SRT| 
|![cover.run go](https://cover.run/go/github.com/wargarblgarbl/libgosubs/ttml.svg)| TTML| 


