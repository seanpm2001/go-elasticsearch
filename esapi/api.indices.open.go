// Code generated from specification version 5.6.16 (052c67e4ebe): DO NOT EDIT

package esapi

import (
	"context"
	"strconv"
	"strings"
	"time"
)

func newIndicesOpenFunc(t Transport) IndicesOpen {
	return func(index []string, o ...func(*IndicesOpenRequest)) (*Response, error) {
		var r = IndicesOpenRequest{Index: index}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// IndicesOpen opens an index.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-open-close.html.
//
type IndicesOpen func(index []string, o ...func(*IndicesOpenRequest)) (*Response, error)

// IndicesOpenRequest configures the Indices Open API request.
//
type IndicesOpenRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   string
	IgnoreUnavailable *bool
	MasterTimeout     time.Duration
	Timeout           time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IndicesOpenRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len(strings.Join(r.Index, ",")) + 1 + len("_open"))
	path.WriteString("/")
	path.WriteString(strings.Join(r.Index, ","))
	path.WriteString("/")
	path.WriteString("_open")

	params = make(map[string]string)

	if r.AllowNoIndices != nil {
		params["allow_no_indices"] = strconv.FormatBool(*r.AllowNoIndices)
	}

	if r.ExpandWildcards != "" {
		params["expand_wildcards"] = r.ExpandWildcards
	}

	if r.IgnoreUnavailable != nil {
		params["ignore_unavailable"] = strconv.FormatBool(*r.IgnoreUnavailable)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = time.Duration(r.MasterTimeout * time.Millisecond).String()
	}

	if r.Timeout != 0 {
		params["timeout"] = time.Duration(r.Timeout * time.Millisecond).String()
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f IndicesOpen) WithContext(v context.Context) func(*IndicesOpenRequest) {
	return func(r *IndicesOpenRequest) {
		r.ctx = v
	}
}

// WithAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (this includes `_all` string or when no indices have been specified).
//
func (f IndicesOpen) WithAllowNoIndices(v bool) func(*IndicesOpenRequest) {
	return func(r *IndicesOpenRequest) {
		r.AllowNoIndices = &v
	}
}

// WithExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both..
//
func (f IndicesOpen) WithExpandWildcards(v string) func(*IndicesOpenRequest) {
	return func(r *IndicesOpenRequest) {
		r.ExpandWildcards = v
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
//
func (f IndicesOpen) WithIgnoreUnavailable(v bool) func(*IndicesOpenRequest) {
	return func(r *IndicesOpenRequest) {
		r.IgnoreUnavailable = &v
	}
}

// WithMasterTimeout - specify timeout for connection to master.
//
func (f IndicesOpen) WithMasterTimeout(v time.Duration) func(*IndicesOpenRequest) {
	return func(r *IndicesOpenRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
//
func (f IndicesOpen) WithTimeout(v time.Duration) func(*IndicesOpenRequest) {
	return func(r *IndicesOpenRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IndicesOpen) WithPretty() func(*IndicesOpenRequest) {
	return func(r *IndicesOpenRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IndicesOpen) WithHuman() func(*IndicesOpenRequest) {
	return func(r *IndicesOpenRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IndicesOpen) WithErrorTrace() func(*IndicesOpenRequest) {
	return func(r *IndicesOpenRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IndicesOpen) WithFilterPath(v ...string) func(*IndicesOpenRequest) {
	return func(r *IndicesOpenRequest) {
		r.FilterPath = v
	}
}
