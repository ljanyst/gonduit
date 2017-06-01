package gonduit

import (
	"github.com/danieldanciu/gonduit/requests"
	"github.com/danieldanciu/gonduit/responses"
)

// MacroCreateMeme performs a call to macro.creatememe.
func (c *Conn) MacroCreateMeme(
	req requests.MacroCreateMemeRequest,
) (*responses.MacroCreateMemeResponse, error) {
	var res responses.MacroCreateMemeResponse

	if err := c.Call("macro.creatememe", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
