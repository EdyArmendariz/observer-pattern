package main

// Concrete Observer: StockObserver
type StockObserver struct {
	name string
}

func (s *StockObserver) Update(t string) {
	// do something
	println("StockObserver:", s.name, "has been updated,", "received subject string:", t)
}
