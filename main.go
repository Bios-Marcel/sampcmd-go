package main

import (
	"os"
	"strings"

	"github.com/Bios-Marcel/sampcmd-go/sampcmd"
)

func main() {
	sampcmd.LaunchSAMP(strings.Join(os.Args[1:], " "))
}
