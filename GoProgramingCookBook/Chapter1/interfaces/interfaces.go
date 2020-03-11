package interfaces

import (
	"fmt"
	"io"
	"os"
)

//Copy copies data from in to out first directly,
//then use a buffer. It also writes to stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	//We write to out but also stdout
	w := io.MultiWriter(out, os.Stdout)
	// a standerd copy, this can be dangerous if there's lot of data in in
	if _, err := io.Copy(w, in); err != nil {
		return err
	}
	//use this to write to then twice
	in.Seek(0, 0)

	//buffered write using 64 byte chunks
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	// lets print a new line
	fmt.Println()
	return nil
}

