package core

import (
	"fmt"
	"time"
)

type Isolat struct {
	fange    string
	sekunder int
}

func (i *Isolat) StartSoning() {
	time.Sleep((5 * time.Second))
	fmt.Println("Fange løslates: " + i.fange)
}
