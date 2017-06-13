package main

import (
  "fmt"
  "gopkg.in/couchbaselabs/gocb.v1"
  "encoding/json"
)

var bucket *gocb.Bucket

func getProduct(productId string) ([]byte, error) {
  cluster, err := gocb.Connect("couchbase://db")
  if err != nil {
    fmt.Println("ERRROR CONNECTING TO CLUSTER:", err)
  }
  // Open Bucket
  bucket, err = cluster.OpenBucket("product", "")
  if err != nil {
    fmt.Println("ERRROR OPENING BUCKET:", err)
  }

  // New query, a really generic one with high selectivity
  myQuery := gocb.NewN1qlQuery("SELECT * FROM `product` " +
    "WHERE productId = '" + productId + "'")
  rows, err := bucket.ExecuteN1qlQuery(myQuery, nil)
  if err != nil {
    fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
  }

  // Interface for handling streaming return values
  var row interface{}

  // Stream the first result only into the interface
  err = rows.One(&row)
  if err != nil {
    fmt.Println("ERROR ITERATING QUERY RESULTS:", err)
  }

  // Marshal single result in interface
  jsonOut, err := json.Marshal(row)
  if err != nil {
    fmt.Println("ERROR PROCESSING STREAMING OUTPUT:", err)
  }

  return jsonOut, nil
}