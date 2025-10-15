/*
The MIT License

Copyright (c) 2025 John Siu

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package tag

import (
	"github.com/J-Siu/go-helper/v2/basestruct"
	"github.com/wtolson/go-taglib"
)

// TagFile : struct holding filename and the taglib.File handle
type TagFile struct {
	basestruct.Base
	Path       string       `json:"Path"`       // filename
	TaglibFile *taglib.File `json:"TaglibFile"` // taglib file handle
}

func (t *TagFile) New(path string) *TagFile {
	t.Path = path
	t.TaglibFile, t.Err = taglib.Read(path)
	return t
}

// Save : save taglib FILE
func (t *TagFile) Save() *TagFile {
	t.TaglibFile.Save()
	return t
}

// Close : close taglib FILE
func (t *TagFile) Close() *TagFile {
	t.TaglibFile.Close()
	return t
}

// Set : get tag
func (t *TagFile) Set(tag *TagStruct, val string) *TagFile {
	t.TaglibFile.SetTag(tag.Tn, val)
	return t
}

// Get : get tag
func (t *TagFile) Get(tag *TagStruct) string {
	return t.TaglibFile.Tag(tag.Tn)
}
