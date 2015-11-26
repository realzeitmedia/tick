package tick_test

import (
	"fmt"
	"time"

	"github.com/realzeitmedia/tick"
)

func Example() {
	tick := tick.NewTicker(10 * time.Second)

	for i := 0; i < 6; i++ {
		t := <-tick.C
		fmt.Printf("tick %d: %s", i, t)
	}
}
