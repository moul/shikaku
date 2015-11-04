package main

import (
	"fmt"
	"log"
	"math/rand"
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
	}

	if _, err := flags.Parse(&opts); err != nil {
		log.Fatalf("Parsing error: %v", err)
	}

	shikakuMap := shikaku.NewShikakuMap(opts.Width, opts.Height, 0, 0)
	if err := shikakuMap.GenerateBlocks(opts.Blocks); err != nil {
		log.Fatalf("Failed to generate %d blocks: %v", opts.Blocks, err)
	}
	fmt.Println(shikakuMap.String())
	// fmt.Println(shikakuMap.Draw())
}
