package sapinapp

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
	"github.com/moul/shikaku"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// extract query from url
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Fprintf(w, "URL error: %v:\n", err)
		return
	}

	// parse query
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		fmt.Fprintf(w, "URL query error: %v:\n", err)
		return
	}

	// check if no arguments
	if len(m) == 0 {
		http.Redirect(w, r, "?width=8&height=8&blocks=10&draw-map=1&draw-solution=1&no-machine-output=0", http.StatusFound)
	}

	// unmarshal arguments
	var opts struct {
		Width  int `schema:"width"`
		Height int `schema:"height"`
		Blocks int `schema:"blocks"`

		DrawMap         bool  `schema:"draw-map"`
		DrawSolution    bool  `schema:"draw-solution"`
		NoMachineOutput bool  `schema:"no-machine-output"`
		Srand           int64 `schema:"srand"`
	}
	decoder := schema.NewDecoder()
	err = decoder.Decode(&opts, m)
	if err != nil {
		fmt.Fprintf(w, "Parameters error: %v:\n", err)
		return
	}

	// check arguments
	if opts.Width > 20 {
		fmt.Fprintf(w, "Max width is: 20\n")
		return
	}
	if opts.Width < 4 {
		fmt.Fprintf(w, "Min width is: 4\n")
		return
	}
	if opts.Height > 20 {
		fmt.Fprintf(w, "Max height is: 20\n")
		return
	}
	if opts.Height < 4 {
		fmt.Fprintf(w, "Min height is: 4\n")
		return
	}

	if opts.Blocks < 1 {
		fmt.Fprintf(w, "Min value for blocks is 0\n")
		return
	}
	if opts.Blocks > 70 {
		fmt.Fprintf(w, "Max value for blocks is 70\n")
		return
	}
	if opts.Blocks > opts.Width*opts.Height/3+1 {
		fmt.Fprintf(w, "Max value for blocks is height*value/3 = %d\n", opts.Width*opts.Height/3)
		return
	}

	// draw
	if opts.Srand > 0 {
		rand.Seed(opts.Srand)
	}

	shikakuMap := shikaku.NewShikakuMap(opts.Width, opts.Height, 0, 0)
	if err := shikakuMap.GenerateBlocks(opts.Blocks); err != nil {
		fmt.Fprintf(w, "Failed to generate %d blocks: %v", opts.Blocks, err)
		return
	}

	outputs := []string{}

	if !opts.NoMachineOutput {
		outputs = append(outputs, shikakuMap.String())
	}
	if opts.DrawMap {
		outputs = append(outputs, shikakuMap.DrawMap())
	}
	if opts.DrawSolution {
		outputs = append(outputs, shikakuMap.DrawSolution())
	}
	fmt.Fprintln(w, strings.Join(outputs, "\n\n"))

}
