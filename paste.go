package gonduit

import (
	"github.com/etcinit/gonduit/requests"
	"github.com/etcinit/gonduit/responses"
)

// PasteCreate calls the paste.create endpoint.
func (c *Conn) PasteCreate(req *requests.PasteCreateRequest) (responses.PasteCreateResponse, error) {
	req.Session = c.Session

	var res responses.PasteCreateResponse
	if err := c.Call("paste.create", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// PasteQuery calls the paste.query endpoint.
func (c *Conn) PasteQuery(req *requests.PasteQueryRequest) (responses.PasteQueryResponse, error) {
	req.Session = c.Session

	var res responses.PasteQueryResponse

	if err := c.Call("paste.query", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}