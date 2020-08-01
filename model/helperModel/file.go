package helpermodel

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request, fieldName string, uid string) (string, error) {
	r.ParseMultipartForm(32 << 20)
	r.FormValue("data")
	file, _, err := r.FormFile("picture")
	if err != nil {
		fmt.Println("case1", err)
		return "", nil
	}
	defer file.Close()
	path := "./files/" + fieldName + "/" + uid + ".jpg"
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("case2", err)
		return "", nil
	}
	defer f.Close()
	io.Copy(f, file)
	return path, nil
}

func RemoveFile(picture string) error {
	err := os.Remove(picture)
	if err != nil {
		return err
	}
	return nil
}
