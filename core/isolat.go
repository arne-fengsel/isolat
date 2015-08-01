package core

import (
	"fmt"
	"time"
)

type Isolat struct {
	fange         Fange
	IsoleringsTid int
}

func (i *Isolat) StartSoning() {
	time.Sleep(time.Duration(i.IsoleringsTid) * time.Second)
	fmt.Println("Fange " + i.fange.String() + " løslates.")
}
