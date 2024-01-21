package main

import (
	"context"
	"fmt"
	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
	"os"
	"runtime"
)

func main() {
	file, _ := os.Open("akademgorodok.pbf")
	defer file.Close()

	scanner := osmpbf.New(context.Background(), file, runtime.GOMAXPROCS(-1))
	scanner.SkipRelations = true

	nodes, ways := osm.Nodes{}, osm.Ways{}
	objectsTotal := 0
	fmt.Println("loading pbf...")

	for scanner.Scan() {
		switch o := scanner.Object().(type) {
		case *osm.Node:
			nodes = append(nodes, o)
		case *osm.Way:
			ways = append(ways, o)
		}
		objectsTotal++

		if objectsTotal%100000 == 0 {
			println(".")
		}
	}
	fmt.Printf("done!\nloaded %v objects\n", objectsTotal)
}

// pbf has a total of 102916477 objects
