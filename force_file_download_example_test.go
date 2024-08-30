package utils_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/loickreitmann/utils"
)

type forceFileDownloadExample struct{}

func ExampleUtils_ForceFileDownload() {
	ffde := &forceFileDownloadExample{}

	mux := ffde.routes()
	fmt.Println("Starting Server on localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func (ffde *forceFileDownloadExample) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./testdata/force_file_download/"))))
	mux.HandleFunc("/download", ffde.download)
	return mux
}

func (ffde *forceFileDownloadExample) download(resw http.ResponseWriter, req *http.Request) {
	var u utils.Utils

	u.ForceFileDownload(resw, req, "./testdata/force_file_download", "download_test.jpg", "Be like the Mandalorian, this is the way.jpg")
}
