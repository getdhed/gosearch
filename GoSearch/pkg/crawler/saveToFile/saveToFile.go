package fileWork

import (
	"1dz/GoSearch/pkg/crawler"
	"fmt"
	"io"
	"log"
	"os"
)

func CreateFile() *os.File {
	f, err := os.Create("documents.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}
func WriteDocuments(slc []crawler.Document, w io.Writer) error {
	for i := range slc {
		_, err := fmt.Fprintln(w, fmt.Sprint(slc[i].ID), slc[i].URL)
		if err != nil {
			return err
		}
	}
	return nil
}
func ReadFromFile(r io.Reader)([]byte,error){
var b[]byte
var buf = make([]byte,10)
for {
	n,err:=r.Read(buf)
	if n>0{
		b=append(b, buf[:n]...)
	}
	if err ==io.EOF{
		break
	}
	if err!=nil{
		return nil,err
	}
}
return b,nil
}
func ReadAll(r io.Reader)([]byte,error){
	b,err:=io.ReadAll(r)
	if err !=nil{
		return nil,err
	}
	return b,nil
}