package sampcmd

/*
#include <stddef.h>
int _LaunchSAMP(wchar_t *args);
*/
import "C"

import (
	"github.com/orofarne/gowchar"
)

//LaunchSAMP launches the SA-MP client passing along the given string as its commandline parameters.
func LaunchSAMP(args string) int {
	wcharString, _ := gowchar.StringToWcharT(args)
	castedShit := (*C.wchar_t)(wcharString)
	return int(C._LaunchSAMP(castedShit))
}
