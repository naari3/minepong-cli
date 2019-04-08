package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Syfaro/mcapi/types"
	"github.com/syfaro/minepong"
)

// ServerStatus contains all information available from a ping request.
// It also includes fields about the success of a request.
type ServerStatus struct {
	Status        string                    `json:"status"`
	Online        bool                      `json:"online"`
	Motd          string                    `json:"motd"`
	MotdExtra     interface{}               `json:"motd_extra,omitempty"`
	MotdFormatted string                    `json:"motd_formatted,omitempty"`
	Favicon       string                    `json:"favicon,omitempty"`
	Error         string                    `json:"error"`
	Players       types.ServerStatusPlayers `json:"players"`
	Server        types.ServerStatusServer  `json:"server"`
}

// Execute ping to minecraft server
func Execute(hostname string, pretty bool) {
	status := getServerStatus(hostname)

	var (
		outputJSON []byte
		err        error
	)
	if pretty {
		outputJSON, err = json.MarshalIndent(&status, "", "  ")
	} else {
		outputJSON, err = json.Marshal(&status)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println(string(outputJSON))
}

func getServerStatus(hostname string) *ServerStatus {
	var status = &ServerStatus{}

	pong, err := minepong.Ping(hostname)
	if err != nil {
		status.Online = false
		status.Status = "error"
		status.Error = err.Error()
	} else {
		status.Status = "success"
		status.Online = true
		status.Favicon = pong.FavIcon
		status.Players.Max = pong.Players.Max
		status.Players.Now = pong.Players.Online
		status.Server.Name = pong.Version.Name
		status.Server.Protocol = pong.Version.Protocol
		status.Error = ""

		status.Motd, status.MotdExtra, status.MotdFormatted = getMotd(pong.Description)
	}
	return status
}

func getMotd(description interface{}) (string, interface{}, string) {
	var (
		motd          string
		motdExtra     interface{}
		motdFormatted string
	)

	switch desc := description.(type) {
	case string:
		motd = desc
		motdExtra = nil
		motdFormatted = ""
	case map[string]interface{}:
		if val, ok := desc["extra"]; ok {
			texts := val.([]interface{})

			b := bytes.Buffer{}
			f := bytes.Buffer{}

			f.WriteString("<span>")

			for id, text := range texts {
				m := text.(map[string]interface{})
				extra := types.MotdExtra{}

				for k, v := range m {
					if k == "text" {
						b.WriteString(v.(string))
						extra.Text = v.(string)
					} else if k == "color" {
						extra.Color = v.(string)
					} else if k == "bold" {
						extra.Bold = v.(bool)
					}
				}

				f.WriteString("<span")

				if extra.Color != "" || extra.Bold {
					f.WriteString(" style='")

					if extra.Color != "" {
						f.WriteString("color: ")
						f.WriteString(extra.Color)
						f.WriteString("; ")
					}

					if extra.Bold {
						f.WriteString(" font-weight: bold; ")
					}

					f.WriteString("'")
				}

				f.WriteString(">")
				f.WriteString(extra.Text)
				f.WriteString("</span>")

				if id != len(texts)-1 {
					f.WriteString(" ")
				}
			}

			f.WriteString("</span>")

			motd = b.String()
			motdExtra = val
			motdFormatted = strings.Replace(f.String(), "\n", "<br>", -1)
		} else if val, ok := desc["text"]; ok {
			motd = val.(string)
			motdExtra = types.MotdExtra{}
			motdFormatted = ""
		}
	default:
		motd = ""
		motdExtra = types.MotdExtra{}
		motdFormatted = ""
	}

	return motd, motdExtra, motdFormatted
}
