package wifiname

import (
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

const osxCmd = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
const osxArgs = "-I"
const linuxCmd = "iwgetid"
const linuxArgs = "--raw"

func WifiName() string {
	platform := runtime.GOOS
	if platform == "darwin" {
		return forOSX()
	} else if platform == "win32" {
		// TODO for Windows
		return ""
	} else {
		// TODO for Linux
		return forLinux()
	}
}

func forLinux() string {
	cmd := exec.Command(linuxCmd, linuxArgs)
	stdout, err := cmd.StdoutPipe()
	panicIf(err)

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	var str string

	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += (string(b) + "\n")
	}

	name := strings.Replace(str, "\n", "", -1)
	return name
}

func forOSX() string {

	cmd := exec.Command(osxCmd, osxArgs)

	stdout, err := cmd.StdoutPipe()
	panicIf(err)

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	var str string

	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += (string(b) + "\n")
	}

	r := regexp.MustCompile(`s*SSID: (.+)s*`)

	name := r.FindAllStringSubmatch(str, -1)

	if len(name) <= 1 {
		return "Could not get SSID"
	} else {
		return name[1][1]
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
