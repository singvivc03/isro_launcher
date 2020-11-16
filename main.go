package main

import (
	"github/singvivc03/isro_launcher/src/launcher"
)

func main() {
	launchSite1 := launcher.New("LS1")
	launchSite2 := launcher.New("LS2")

	startsFrom, totalNumberOfSatellite, groupOf := 1, 500, 4
	for startsFrom < totalNumberOfSatellite {
		launchSatellite(launchSite1, &startsFrom, groupOf)
		launchSatellite(launchSite2, &startsFrom, groupOf)
	}
}

func launchSatellite(launcher launcher.Launcher, satelliteSequenceStartsFrom *int, groupOf int) {
	if *satelliteSequenceStartsFrom < 500 {
		signal := launcher.Launch(*satelliteSequenceStartsFrom, groupOf)
		*satelliteSequenceStartsFrom = <-signal
	}
}
