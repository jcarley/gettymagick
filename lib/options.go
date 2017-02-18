//go:generate codecgen -o options_generated.go options.go

package lib

type Options struct {
	Compression int    `codec:"compression,omitempty"`
	Destination string `codec:"destination,omitempty"`
	Height      int    `codec:"height,omitempty"`
	Quality     int    `codec:"quality,omitempty"`
	Source      string `codec:"source,omitempty"`
	Width       int    `codec:"width,omitempty"`
}
