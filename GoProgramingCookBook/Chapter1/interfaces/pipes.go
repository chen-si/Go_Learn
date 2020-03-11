package interfaces

import (
	"io"
	"os"
)

//PipeExamples help give more example of using io interfaces
func PipeExamples() error {
	//the pipe reader and pipe writer implent io.Reader and io.Writer
	r, w := io.Pipe()
	//this needs to be run in a separate go routine as it will block waiting for
	//the reader close at the end to cleanup
	go func() {
		//for now we'll write soming basic, this could also be used to encode json
		//base64 encode ,etc...
		w.Write([]byte("testing..."))
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
