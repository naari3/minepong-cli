# minepong-cli

A CLI for get information about Minecraft server.

Just wrapped https://github.com/Syfaro/minepong for CLI

## Usage

```
minepong-cli is a CLI for get information about Minecraft server metadata.

Usage:
  minepong-cli [flags]

Examples:

minepong-cli --host mc1 --port 25575
minepong-cli --port 25575
MC_PORT=25575 minepong-cli


Flags:
  -h, --help          help for minepong-cli
      --host string   server's hostname (default "localhost")
      --port int      Server's port (default 25565)
  -p, --pretty        Use pretty printing
```

## Example

```json
// $ minepong-cli --host localhost --port 25565 -p
{
  "status": "success",
  "online": true,
  "motd": "A Vanilla Minecraft Server powered by Docker",
  "players": {
    "max": 20,
    "now": 0
  },
  "server": {
    "name": "1.13.2",
    "protocol": 404
  }
}
```
