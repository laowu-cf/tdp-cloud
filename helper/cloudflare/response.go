package cloudflare

type Response struct {
	Success    bool           `json:"success"`
	Errors     []ResponseInfo `json:"errors"`
	Messages   []ResponseInfo `json:"messages"`
	Result     any            `json:"result"`
	ResultInfo ResultInfo
}

type ResponseInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResultInfo struct {
	Page       int               `json:"page" url:"page,omitempty"`
	PerPage    int               `json:"per_page" url:"per_page,omitempty"`
	TotalPages int               `json:"total_pages" url:"-"`
	Count      int               `json:"count" url:"-"`
	Total      int               `json:"total_count" url:"-"`
	Cursor     string            `json:"cursor" url:"cursor,omitempty"`
	Cursors    ResultInfoCursors `json:"cursors" url:"cursors,omitempty"`
}

type ResultInfoCursors struct {
	Before string `json:"before" url:"before,omitempty"`
	After  string `json:"after" url:"after,omitempty"`
}