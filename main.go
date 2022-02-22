package main

import (
	"flag"
	"time"

	"github.com/bjarneo/superstonk/api"
	"github.com/bjarneo/superstonk/msg"
	"github.com/bjarneo/superstonk/utils"
	"github.com/pterm/pterm"
)

func main() {
	symbol := flag.String("symbol", "gme", "The stock symbol to watch")
	interval := flag.Int("interval", 5, "The refresh interval in seconds")
	// We are using float64 since the shares can be fractional
	shares := flag.Float64("shares", 0, "The quantity of stocks you own")

	flag.Parse()

	// Clear the terminal before we display our stonk
	utils.Clear()

	pterm.Print("\n\n")

	pterm.DefaultCenter.Println("[/r/Superstonk]")

	// Continuously update the area forever
	area, _ := pterm.DefaultArea.Start()
	for {
		stock := api.Quote(*symbol)

		price := msg.StockPrice(stock.Price(), stock.PriceState())

		utils.TerminalTitle(*symbol, stock.Price())

		statistics := msg.Statistics(stock, *shares)

		area.Update(price + statistics)

		time.Sleep(time.Second * time.Duration(*interval))
	}
}
