package sampcmd

/*
#include <stddef.h>
int _LaunchSAMP(wchar_t *args);
int _LaunchSAMPSetWorkingDir(wchar_t *working_dir, wchar_t *args);
*/
import "C"

import (
	"path/filepath"

	"github.com/orofarne/gowchar"
	"golang.org/x/sys/windows/registry"
)

//RegistryError error used for problems during registry calls.
const RegistryError = 12345

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

//LaunchSAMPDetectGTADirectory launches the SA-MP client passing along the given string as its commandline parameters.
//This automatically detects the location where GTA SA is installed.
func LaunchSAMPDetectGTADirectory(args string) int {
	argsAsWchar, _ := gowchar.StringToWcharT(args)
	argsAsWcharCPointer := (*C.wchar_t)(argsAsWchar)

	key, keyError := registry.OpenKey(registry.CURRENT_USER, "Software\\SAMP", registry.READ)
	if keyError != nil {
		return RegistryError
	}

	path, _, valueError := key.GetStringValue("gta_sa_exe")
	if valueError != nil {
		return RegistryError
	}

	pathAsWchar, _ := gowchar.StringToWcharT(filepath.Dir(path))
	pathAsWcharCPointer := (*C.wchar_t)(pathAsWchar)

	return int(C._LaunchSAMPSetWorkingDir(pathAsWcharCPointer, argsAsWcharCPointer))
}
