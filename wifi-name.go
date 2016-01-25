package wifiname

import (
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
)

const osxCmd = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"

func WifiName() string {
	platform := runtime.GOOS
	if platform == "darwin" {
		return forOSX()
	} else if platform == "win32" {
		// TODO for Windows
		return ""
	} else {
		// TODO for Linux
		return ""
	}
}

func forOSX() string {
	cmd := exec.Command(osxCmd, "-I")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

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
