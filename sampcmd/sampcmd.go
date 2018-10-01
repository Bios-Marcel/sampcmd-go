package sampcmd

/*
#include <stddef.h>
int _LaunchSAMP(wchar_t *args);
int _LaunchSAMPSetWorkingDir(wchar_t *working_dir, wchar_t *args);
*/
import "C"

import (
	"github.com/orofarne/gowchar"
)

//LaunchSAMP launches the SA-MP client passing along the given string as its commandline parameters.
//This has to be executed from within the GTA SA folder.
func LaunchSAMP(args string) int {
	argsAsWchar, _ := gowchar.StringToWcharT(args)
	argsAsWcharCPointer := (*C.wchar_t)(argsAsWchar)

	return int(C._LaunchSAMP(argsAsWcharCPointer))
}

//LaunchSAMPSetWorkingDir launches the SA-MP client passing along the given string as its commandline parameters.
//You must pass the GTA SA folder as the working directory here.
func LaunchSAMPSetWorkingDir(workingDir, args string) int {
	argsAsWchar, _ := gowchar.StringToWcharT(args)
	argsAsWcharCPointer := (*C.wchar_t)(argsAsWchar)

	workingDirAsWchar, _ := gowchar.StringToWcharT(workingDir)
	workingDirAsWcharCPointer := (*C.wchar_t)(workingDirAsWchar)

	return int(C._LaunchSAMPSetWorkingDir(workingDirAsWcharCPointer, argsAsWcharCPointer))
}
