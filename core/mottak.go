package core

type Mottak struct {
	innKo chan string
}

func NyttMottak() *Mottak {
	ko := make(chan string)
	m := &Mottak{innKo: ko}
	go m.kjorIsolat()
	return m
}

func (m *Mottak) Motta(f string) {
	m.innKo <- f
}

func (m *Mottak) kjorIsolat() {
	for f := range m.innKo {
		isolat := &Isolat{fange: f, sekunder: 5}
		go isolat.StartSoning()
	}

	close(m.innKo)
}
