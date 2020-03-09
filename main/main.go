package main

import (
	"context"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
)

func main() {
	LoadEnv()

	errorLogger = initLogger(MyEnv["LOGFILE"], MyEnv["LOGLEVEL"])

	kjson, err := ioutil.ReadFile(MyEnv["KEYSTORE"])
	if err != nil {
		Errorf("Get node private key read file=%v failed: %v\n", MyEnv["KEYSTORE"], err)
		return
	}

	key, err := keystore.DecryptKey(kjson, MyEnv["KEYSTOREPASS"])
	if err != nil {
		Errorf("Get node private key get decrypt key failed: %v\n", err)
		return
	}

	NodeKeyPrivate = key.PrivateKey
	NodeKeyStore = hex.EncodeToString(crypto.FromECDSA(NodeKeyPrivate))

	r := mux.NewRouter()

	//handler
	r.HandleFunc("/contract/{userid}", initStoreContract).Methods("PUT")

	//Bucket
	r.HandleFunc("/bucket", createBucketHandle).Methods("PUT")
	r.HandleFunc("/bucket", getBucketInfoHandle).Methods("GET")
	r.HandleFunc("/bucket", deleteBucketHandle).Methods("DELETE")
	r.HandleFunc("/bucket/{userid}", getUserBucketsHandle).Methods("GET")

	//File/Dir
	r.HandleFunc("/storage/{file:.*}", putHandle).Methods("PUT")
	r.HandleFunc("/storage/{file:.*}", getHandle).Methods("GET")
	r.HandleFunc("/storage/{file:.*}", deleteHandle).Methods("DELETE")

	//Attribute
	r.HandleFunc("/attribute/{file:.*}", headAttributeHandle).Methods("GET")

	//share file
	r.HandleFunc("/file/{file:.*}", shareFileHandle).Methods("POST")

	//move
	r.HandleFunc("/opcode/move", moveHandle).Methods("POST")
	r.HandleFunc("/opcode/copy", copyHandle).Methods("POST")
	r.HandleFunc("/opcode/rename", renameHandle).Methods("POST")
	r.HandleFunc("/opcode/getsize", getsizeHandle).Methods("GET")

	//Big file upload and download
	r.HandleFunc("/uploads/upload/{file:.*}", uploadHandle).Methods("PUT")
	r.HandleFunc("/uploads/complete/{file:.*}", completeMultipartHandle).Methods("POST")
	r.HandleFunc("/download/{hash:.*}", getMultipartHandle).Methods("GET")

	r.HandleFunc("/common/authorize", authorizeHandle).Methods("GET")
	r.HandleFunc("/getinfo", getinfoHandle).Methods("GET")
	r.HandleFunc("/getattribute/{file:.*}", getAttributeHandle).Methods("GET")

	http.Handle("/", r) // enable the router

	srv := &http.Server{
		Handler:      r,
		Addr:         ":23000",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			Errorf("http server failed: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	srv.Shutdown(ctx)

	Infof("server shutting down\n")
	os.Exit(0)

}
