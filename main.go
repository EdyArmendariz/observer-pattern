// https://go.dev/doc/code

// go install .
// go build
// .\main.exe

package main

func main() {

	// Create a new stockMonitor object
	stockMonitor := &StockMonitor{
		ticker: "AAPL",
		price:  0.0,
	}

	observerA := &StockObserver{
		name: "Observer A",
	}
	observerB := &StockObserver{
		name: "Observer B",
	}

	// Attach our Observers to the stockMonitor
	stockMonitor.Attach(observerA)
	stockMonitor.Attach(observerB)

	// Start the stockMonitor
	stockMonitor.Notify()

	// Change the price of the stockMonitor
	stockMonitor.SetPrice(500)

	// Detach an Observer from the stockMonitor
	stockMonitor.Detach(observerA)

	// Change the price of the stockMonitor
	stockMonitor.SetPrice(528)
}
