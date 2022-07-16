package gout

import (
	"net/http"

	"github.com/ozline/go-gout/binding"
)

// Content-Type MIME of the most common data formats.
const (
	MIMEJSON = binding.MIMEJSON
	// MIMEHTML              = binding.MIMEHTML
	// MIMEXML               = binding.MIMEXML
	// MIMEXML2              = binding.MIMEXML2
	MIMEPlain = binding.MIMEPlain

// MIMEPOSTForm          = binding.MIMEPOSTForm
// MIMEMultipartPOSTForm = binding.MIMEMultipartPOSTForm
// MIMEYAML              = binding.MIMEYAML
// MIMETOML              = binding.MIMETOML
)

type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	Params   Params
	handlers HandlersChain
	index    int8
	fullPath string

	engine       *Engine
	params       *Params
	skippedNodes *[]skippedNode

	// mu sync.RWMutex
	Keys map[string]interface{}

	Errors errorMsgs

	Accepted []string

	// queryCache url.Values

	// formCache url.Values

	sameSite http.SameSite
}

func (c *Context) reset() {
	// c.Writer = &c.writermem
	// c.Params = c.Params[:0]
	// c.handlers = nil
	c.index = -1

	c.fullPath = ""
	c.Keys = nil
	// c.Errors = c.Errors[:0]
	c.Accepted = nil
	// c.queryCache = nil
	// c.formCache = nil
	c.sameSite = 0
	// *c.params = (*c.params)[:0]
	// *c.skippedNodes = (*c.skippedNodes)[:0]
}

func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}
