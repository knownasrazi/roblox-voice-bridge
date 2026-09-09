# roblox-voice-bridge

**Voice, bridged.**

> Voice bridge for Roblox - proximity audio via Discord.

## Hand-crafted Go

\\go
func handlePosition(id int, x, y, z float32) error {
    if x < -3000 || x > 3000 { return ErrOutOfBounds }
    return voice.Send(fmt.Sprintf("player:%d %.1f,%.1f,%.1f", id, x, y, z))
}
\\\

## Run

\\\ash
git clone https://github.com/knownasrazi/roblox-voice-bridge.git
cd roblox-voice-bridge
go run main.go
\\\

## License

MIT