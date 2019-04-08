package types

// MotdExtra contains information about the motd decoration things
type MotdExtra struct {
	Bold  bool   `json:"bold"`
	Color string `json:"color"`
	Text  string `json:"text"`
}

// ServerStatusPlayers contains information about the min and max numbers of players
type ServerStatusPlayers struct {
	Max    int                 `json:"max"`
	Now    int                 `json:"now"`
	Sample []map[string]string `json:"sample,omitempty"`
}

// ServerStatusServer contains information about the server version.
// As it is a ping request, it is fairly basic information.
type ServerStatusServer struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

// ServerStatus contains all information available from a ping request.
// It also includes fields about the success of a request.
type ServerStatus struct {
	Status        string              `json:"status"`
	Online        bool                `json:"online"`
	Motd          string              `json:"motd"`
	MotdExtra     interface{}         `json:"motd_extra,omitempty"`
	MotdFormatted string              `json:"motd_formatted,omitempty"`
	Favicon       string              `json:"favicon,omitempty"`
	Error         string              `json:"error,omitempty"`
	Players       ServerStatusPlayers `json:"players"`
	Server        ServerStatusServer  `json:"server"`
}
