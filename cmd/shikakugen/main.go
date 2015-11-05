package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/moul/shikaku"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	var opts struct {
		Width  int `short:"W" long:"width" description:"Width of the grid" required:"true"`
		Height int `short:"H" long:"height" description:"Height of the grid" required:"true"`
		Blocks int `short:"B" long:"blocks" description:"Blocks in the grid" required:"true"`

		DrawMap         bool  `short:"m" long:"draw-map" description:"Draw the map in ascii-art"`
		DrawSolution    bool  `short:"s" long:"draw-solution" description:"Draw the solution in ascii-art"`
		NoMachineOutput bool  `short:"q" long:"no-machine-output" description:"No machine output"`
		Srand           int64 `long:"srand" description:"Random seed"`
	}

	if _, err := flags.Parse(&opts); err != nil {
		log.Fatalf("Parsing error: %v", err)
	}

	if opts.Srand > 0 {
		rand.Seed(opts.Srand)
	}

	shikakuMap := shikaku.NewShikakuMap(opts.Width, opts.Height, 0, 0)
	if err := shikakuMap.GenerateBlocks(opts.Blocks); err != nil {
		log.Fatalf("Failed to generate %d blocks: %v", opts.Blocks, err)
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
	fmt.Println(strings.Join(outputs, "\n\n"))
}
