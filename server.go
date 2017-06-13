package main

import (
  "github.com/julienschmidt/httprouter"
  "net/http"
  "log"
)

func main() {
  router := httprouter.New()
  router.GET("/product/:productId", product)

  log.Fatal(http.ListenAndServe(":8080", router))
}

func product(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  productId := ps.ByName("productId")

  product, err := getProduct(productId)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  w.Write(product)
}