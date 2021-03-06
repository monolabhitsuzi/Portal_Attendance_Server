/*
 * Portal Client API
 *
 * description
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"
	"os"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "github.com/monolabhitsuzi/Portal_Attendance_Server/go"
)

func main() {
	log.Printf("Server started")

	port := os.Getenv("PORT")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":" + port, router))
}
