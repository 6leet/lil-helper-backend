package helpermodel

import (
	"fmt"
	"io"
	"log"
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
	prePath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(prePath)
	// path := prePath + "/files/" + fieldName + "/" + uid + ".jpg"
	path := "/files/" + fieldName + "/" + uid + ".jpg"
	f, err := os.OpenFile(prePath+path, os.O_WRONLY|os.O_CREATE, 0666)
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
