/*
The MIT License

Copyright (c) 2020 John Siu

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package helper

import (
	"fmt"

	"github.com/wtolson/go-taglib"
)

// TagStruct : hold attributes of a tag
type TagStruct struct {
	Ln string         // flag long name
	Sn string         // flag short name
	Dn string         // tag display name for output
	Ms string         // flag message
	tn taglib.TagName // taglib tag name constant
}

// Tags : helper list for tags supported
var Tags = []TagStruct{
	{"artist", "a", "Artist", "set artist", taglib.Artist},
	{"title", "t", "Title", "set title", taglib.Title},
	{"album", "A", "Album", "set album", taglib.Album},
	{"year", "y", "Year", "set year", taglib.Year},
	{"track", "T", "Track", "set track", taglib.Track},
	{"comment", "c", "Comments", "set comments", taglib.Comments},
}

// File : struct holding filename and the taglib.File handle
type File struct {
	name string       // filename
	tfh   *taglib.File // taglib file handle
}

// Open : open a taglib FILE
func Open(path string) *File {
	tfh, e := taglib.Read(path)
	// If taglib fail to open file, exit
	if e != nil {
		fmt.Println(path, ":", e)
		return nil
	}
	var fh = &File{path, tfh}
	return fh
}

// Save : save taglib FILE
func (fh *File) Save() {
	fh.tfh.Save()
}

// Close : close taglib FILE
func (fh *File) Close() {
	fh.tfh.Close()
}

// Set : get tag
func (fh *File) Set(tag *TagStruct, val string) {
	fh.tfh.SetTag(tag.tn, val)
}

// Get : get tag
func (fh *File) Get(tag *TagStruct) string {
	return fh.tfh.Tag(tag.tn)
}
