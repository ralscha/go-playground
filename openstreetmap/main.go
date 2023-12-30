package main

import (
	"fmt"
	"github.com/qedus/osmpbf"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

func main() {
	f, err := os.Open("E:\\_download\\switzerland-latest.osm.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d := osmpbf.NewDecoder(f)

	// use more memory from the start, it is faster
	d.SetBufferSize(osmpbf.MaxBlobSize)

	// start decoding with several goroutines, it is faster
	err = d.Start(runtime.GOMAXPROCS(-1))
	if err != nil {
		log.Fatal(err)
	}

	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			isPlace := false
			var placeId int64
			tags := make(map[string]string)

			switch obj := v.(type) {
			case *osmpbf.Node:
				tags = obj.Tags
				placeId = obj.ID
				isPlace = true
			case *osmpbf.Way:
				tags = obj.Tags
				placeId = obj.ID
				isPlace = true
			}

			if isPlace {
				//if strings.ToLower(tags["amenity"]) == "cafe" {
				// if name contains starbucks
				if strings.Contains(strings.ToLower(tags["name"]), "starbucks") {
					fmt.Println(placeId, tags["amenity"])
					fmt.Println(tags["name"], tags["cuisine"], tags["addr:city"], tags["addr:postcode"], tags["addr:street"], tags["addr:housenumber"])
				}
				//}
			}
		}
	}

}
