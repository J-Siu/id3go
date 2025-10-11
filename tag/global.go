package tag

import (
	"github.com/wtolson/go-taglib"
)

const Version = "v1.4.0"

// Tags : helper list for tags supported
var Tags = []TagStruct{
	{"artist", "a", "Artist", "set artist", taglib.Artist},
	{"title", "t", "Title", "set title", taglib.Title},
	{"album", "A", "Album", "set album", taglib.Album},
	{"year", "y", "Year", "set year", taglib.Year},
	{"track", "T", "Track", "set track", taglib.Track},
	{"comment", "c", "Comments", "set comments", taglib.Comments},
}
