package cli

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/james4k/rcon"
)

// Execute counting number of online players
func Execute(hostname string, password string) {
	resp, err := execCmd(hostname, password, "list")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	onlineRgx := regexp.MustCompile("There are (\\d) of")
	groups := onlineRgx.FindSubmatch([]byte(resp))

	onlinePlayerCount, err := strconv.Atoi(string(groups[1]))
	if err != nil {
		fmt.Fprintln(os.Stderr, wrapError("some error", err))
		os.Exit(1)
	}
	fmt.Println(onlinePlayerCount)
}

func execCmd(hostname string, password string, cmd string) (string, error) {
	remoteConsole, err := rcon.Dial(hostname, password)
	if err != nil {
		return "", wrapError("Failed to connect to RCON server", err)
	}
	defer remoteConsole.Close()

	reqID, err := remoteConsole.Write(cmd)
	if err != nil {
		return "", wrapError("Failed to send command", err)
	}
	resp, respReqID, err := remoteConsole.Read()
	if respReqID != reqID {
		fmt.Println("Weird. This response is for another request.")
	}
	return resp, nil
}

func wrapError(msg string, err error) error {
	return errors.New(msg + ": " + err.Error())
}
