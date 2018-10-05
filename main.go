package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Bios-Marcel/sampcmd-go/sampcmd"
)

func main() {
	launchError := sampcmd.LaunchSAMPDetectGTADirectory(strings.Join(os.Args[1:], " "))
	if launchError != 0 {
		fmt.Println("Couldn't start GTA SA via automatic installation detection. Returned", launchError)
		launchError = sampcmd.LaunchSAMP(strings.Join(os.Args[1:], " "))
		if launchError != 0 {
			fmt.Println("Couldn't start GTA SA using the current directory. Returned", launchError)
		}
	}
}
