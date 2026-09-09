package voice
import "fmt"
type Client struct{ addr string }
func New(a string) *Client{ return &Client{addr:a} }
func (c *Client) Send(id int, x,y,z float32) error { fmt.Printf("voice %d %.1f,%.1f,%.1f\n",id,x,y,z); return nil }
func (c *Client) Mute(id int, m bool) error { fmt.Printf("mute %d %v\n",id,m); return nil }
