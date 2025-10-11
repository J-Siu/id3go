package tag

import "github.com/wtolson/go-taglib"

// TagStruct : hold attributes of a tag
type TagStruct struct {
	Ln string         `json:"Ln"` // flag long name
	Sn string         `json:"Sn"` // flag short name
	Dn string         `json:"Dn"` // tag display name for output
	Ms string         `json:"Ms"` // flag message
	Tn taglib.TagName `json:"Tn"` // taglib tag name constant
}
