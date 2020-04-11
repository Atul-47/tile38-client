package main

import (
	"log"

	"github.com/k0kubun/pp"

	t38c "github.com/lostpeer/tile38-client"
	"github.com/lostpeer/tile38-client/geofence"
)

func main() {
	geo, err := geofence.New(geofence.NewRadixFencer("localhost:9851"), true)
	if err != nil {
		log.Fatal(err)
	}

	req := geofence.Within("people", t38c.AreaCircle(0, 0, 10000)).
		WithOptions(t38c.Where("speed", 0, 60)).
		Actions(geofence.Enter, geofence.Exit).
		Format(t38c.OutputBounds)

	ch, err := geo.Fence(req)
	if err != nil {
		log.Fatal(err)
	}
	for event := range ch {
		pp.Println(event)
	}
	// tile38, err := t38c.New(t38c.NewRadixPool("localhost:9851", 1), t38c.Debug)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := tile38.IntersectsHashes("fleet", t38c.AreaCircle(0, 0, 999999), 5)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
