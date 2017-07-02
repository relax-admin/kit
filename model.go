package kit

import "fmt"

type Result struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
	Error   Error       `json:"error"`
}

type Error struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Details interface{} `json:"details,omitempty"`
}
type StatusMessage struct {
	StatusCode int
	Result     Result
}

type ArrayResult struct {
	TotalCount int         `json:"totalCount"`
	Items      interface{} `json:"items"`
}

type (
	APIParam struct {
		Q              string
		SkipCount      int
		MaxResultCount int
		Fields         string
		Sort           string
		SortAsc        string
		SortDesc       string
	}
)

func (e Error) Error() string {
	return fmt.Sprintf("[%d]%v-%v", e.Code, e.Message, e.Details)
}
