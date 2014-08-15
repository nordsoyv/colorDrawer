package main

import (
	//	"github.com/dradtke/go-allegro/allegro"
	//	"fmt"
	"github.com/nordsoyv/colorDrawer/colorCube"
	"github.com/nordsoyv/colorDrawer/config"
	"github.com/nordsoyv/colorDrawer/strategy"
)

func main() {
	/*	var (
			display    *allegro.Display
			eventQueue *allegro.EventQueue
			running    bool = true
			err        error
		)
	*/
	configuration := config.Read("config.json")
	cube := colorCube.New(uint8(configuration.ColorCubeBitSize))
	strat := strategy.NearestNeighbor(configuration)

	/*	if err := allegro.Install(); err != nil {
			panic(err)
		}
		defer allegro.Uninstall()

		// Create a 640x480 window and give it a title.
		allegro.SetNewDisplayFlags(allegro.WINDOWED)
		if display, err = allegro.CreateDisplay(512, 512); err == nil {
			defer display.Destroy()
			display.SetWindowTitle("colorDrawer")
		} else {
			panic(err)
		}

		// Create an event queue. All of the event sources we care about should
		// register themselves to this queue.
		if eventQueue, err = allegro.CreateEventQueue(); err == nil {
			defer eventQueue.Destroy()
		} else {
			panic(err)
		}

		// Calculate the timeout value based on the desired FPS.
		timeout := float64(1) / float64(FPS)

		// Register event sources.
		eventQueue.Register(display)

		// Set the screen to black.
		allegro.ClearToColor(allegro.MapRGB(0, 0, 0))
		allegro.FlipDisplay()

		imageUpdates = make(chan strategy.ImageUpdate)

	*/
	strat.GenerateImage(cube)
	//surface.ToPng(configuration.OutputFilename)
	/*

		// Main loop.
		var event allegro.Event
		for running {
			if e, found := eventQueue.WaitForEventUntil(allegro.NewTimeout(timeout), &event); found {
				switch e.(type) {
				case allegro.DisplayCloseEvent:
					running = false
					break

					// Handle other events here.
				}
			}
		}
	*/
}
