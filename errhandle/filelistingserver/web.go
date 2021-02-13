package main

import (
	"log"
	"net/http"
	"os"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if re := recover(); re != nil {
				log.Panicf("Panic: %v", re)
				http.Error(w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(w, r)
		if err != nil {
			log.Printf("Error occurred handling request :%s", err.Error())
		}

		//user error
		if userError, ok := err.(userError); ok {
			http.Error(w, userError.Message(), http.StatusBadRequest)
		}

		//system error
		code := http.StatusOK
		switch {
		case os.IsNotExist(err):
			code = http.StatusNotFound
		case os.IsPermission(err):
			code = http.StatusForbidden
		default:
			code = http.StatusInternalServerError
		}
		http.Error(w, http.StatusText(code), code)
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	//http.HandleFunc("/",errWrapper(HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
