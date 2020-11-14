package main
import(
	"net"
	"fmt"
	"io"
	"bytes"
	)

func main(){
	conn, _:= net.Dial("tcp","golang.org:80")
	defer conn.Close()


	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	fmt.Println(" belge: ", buf.String())
	return
}
