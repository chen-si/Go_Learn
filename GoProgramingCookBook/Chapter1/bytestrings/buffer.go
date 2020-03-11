package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

//Buffer demonstrates some tricks for initializing bytes Buffers
//These buffers implements an io.Reader interface
func Buffer(rawString string) *bytes.Buffer {
	//we will start with a string encodeed into raw bytes
	rawBytes := []byte(rawString)
	//there is a number of way to create a buffer from
	//the raw bytes or from the original string
	var b = new(bytes.Buffer)
	b.Write(rawBytes)
	//alternatively
	b = bytes.NewBuffer(rawBytes)
	//and avoiding the initial byte altogether
	b = bytes.NewBufferString(rawString)

	return b
}

//ToString is an example of taking an io.Reader an consuming
//it all ,then return a string
func ToString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
