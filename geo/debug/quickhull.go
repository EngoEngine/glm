package main

import (
	"fmt"
	"github.com/luxengine/glm"
	"github.com/luxengine/glm/geo"
	"math/rand"
	"os"
)

const (
	numtest = 50
	size    = 500
)

func main() {
	rand.Seed(999)

	ftxt, err := os.OpenFile("hull.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for n := 1; n <= numtest; n++ {

		var points []glm.Vec2
		for i := 0; i < 10; i++ {
			points = append(points, glm.Vec2{float32(int(rand.Float32() * size)), float32(int(rand.Float32() * size))})
		}

		hull := geo.Quickhull2(points)

		ftxt.WriteString(fmt.Sprintf(
			`		{
			points: %#v,
			hull: %#v,
		},
`, points, hull))

	}
	ftxt.Close()

}
