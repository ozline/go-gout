package gout

import (
	"math"
	"net/http"

	"github.com/ozline/go-gout/binding"
	"github.com/ozline/go-gout/render"
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

const abortIndex int8 = math.MaxInt8 >> 1 //最大int/2

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

// Status sets the HTTP response code.
func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Render(code, render.JSON{Data: obj})
}

// bodyAllowedForStatus is a copy of http.bodyAllowedForStatus non-exported function.
func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == http.StatusNoContent:
		return false
	case status == http.StatusNotModified:
		return false
	}
	return true
}

// Render writes the response headers and calls render.Render to render data.
func (c *Context) Render(code int, r render.Render) {
	c.Status(code)

	if !bodyAllowedForStatus(code) {
		r.WriteContentType(c.Writer)
		c.Writer.WriteHeaderNow()
		return
	}

	if err := r.Render(c.Writer); err != nil {
		panic(err)
	}
}
