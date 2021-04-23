package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jboursiquot/go-proverbs"
	"github.com/mitchellh/go-wordwrap"
	"github.com/zzwx/interval"
)

func main() {
	min := flag.Uint("min", 2, "min seconds")
	max := flag.Uint("max", 5, "max seconds")
	width := flag.Uint("width", 20, "text width")
	link := flag.Bool("link", true, "include source web link")
	flag.Parse()
	*min = interval.ClampUint(0, math.MaxUint32-1, *min)
	*max = interval.ClampUint(*min+1, math.MaxUint32, *max)
	*width = interval.ClampUint(1, math.MaxUint32, *width)
	rand.Seed(time.Now().UnixNano())
	var w io.Writer
	if rand.Float64() < 0.5 {
		w = os.Stdout
	} else {
		w = os.Stderr
	}
	howLong := int(*min) + rand.Intn(int(*max-*min))
	proverb := proverbs.Random()
	wr := strings.Split(wordwrap.WrapString(proverb.Saying, *width), "\n")
	for li, line := range wr {
		if li > 0 {
			fmt.Fprintf(w, "\n")
		}
		ln := utf8.RuneCountInString(line)
		for i := 0; i < ln; i++ {
			time.Sleep(time.Duration(howLong*1000/(len(wr)*ln)) * time.Millisecond)
			fmt.Fprintf(w, "\r%v", strings.Repeat("_", i+1))
		}
		fmt.Fprintf(w, "\r%v", line)
	}
	if *link {
		fmt.Fprintf(w, "\n%v", proverb.Link)
	}
}
