package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HanderFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):] // /list/fib.txt
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
		return err
	}
	writer.Write(all)
	return nil
}