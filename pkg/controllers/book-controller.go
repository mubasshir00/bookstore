package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"bookstore/pkg/utlis"
	"bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter,r *http.Request){
	newBooks:=models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookNyId(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("Error While parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}