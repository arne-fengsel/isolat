package core

type Mottak struct {
	innKo chan IsolatFange
}

func NyttMottak() *Mottak {
	m := &Mottak{innKo: make(chan IsolatFange)}
	go m.HandterMottaksKo()
	return m
}

func (m *Mottak) Motta(f IsolatFange) {
	m.innKo <- f
}

func (m *Mottak) HandterMottaksKo() {
	for f := range m.innKo {
		isolat := &Isolat{fange: f.FangeTilIsolat, IsoleringsTid: f.IsoleringsTid}
		go isolat.StartSoning()
	}

	close(m.innKo)
}
