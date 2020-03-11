package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

//WorkWithBuffer will make use of the buffer created by Buffer function
func WorkWithBuffer() error {
	rawString := "It's easy to encode unicode into a byte array"
	b := Buffer(rawString)

	//We can quickly convert a buffer back into bytes with b.Bytes()
	//or a string with b.String()
	fmt.Println(b.String())
	s, err := ToString(b)
	if err != nil{
		return err
	}
	fmt.Println(s)

	//we can also take our bytes and creates a bytes reader
	//these reader implements io.Reader io.ReaderAt,io.WriteTo,io.seeker,
	//io.ByteScanner and io.RuneScanner interface
	reader := bytes.NewReader([]byte(rawString))
	//we can also plug it into a scanner that allows buffered reading and tookenzation
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	//iterate over all of the scan events
	for scanner.Scan(){
		fmt.Print(scanner.Text())
	}
	return nil
}
