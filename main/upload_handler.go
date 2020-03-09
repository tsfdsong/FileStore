package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

//InitMultipartResult ...
type InitMultipartResult struct {
	Bucket   string `json:"bucket"`
	Key      string `json:"key"`
	UploadID string `json:"uploadid"`
}

//CompleteMultipartUpload ...
type CompleteMultipartUpload struct {
	Shards []string `json:"shards"`
}

//upload API
var authorizeHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Get token panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("Initiate multipart upload : %v", err), "", w)
		}
	}()

	userid := r.Header.Get("Authorization")

	tokenString, err := GenerateToken(userid)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("get token : %v", err), "", w)
		return
	}

	Infof("Get token: user=%v token=%v", userid, tokenString)

	HandleError(http.StatusOK, nil, fmt.Sprintf("%v", tokenString), w)
}

var uploadHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Upload multipart panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("Upload multipart: %v", err), "", w)
		}
	}()

	start := time.Now().UnixNano() / 1000000

	userID := r.Header.Get("Authorization")
	bucketName := r.Header.Get("BucketName")

	vars := mux.Vars(r)
	file := vars["file"]

	Infof("BigFile upload file user=%v bucket=%v file=%v", userID, bucketName, file)

	subStrs := strings.Split(file, "/")
	fileName := subStrs[len(subStrs)-1]

	Infof("Upload file filename=%v\n", fileName)

	respbody, err := PutFileWithBody(r.Body, fileName)
	if err != nil {
		end := time.Now().UnixNano()/1000000 - start
		Errorf("BigFile upload put with body file=%v time=%v err=%v", file, end, err)
		HandleError(http.StatusInternalServerError, fmt.Errorf("Upload file failed:  user=%v bucket=%v %v", userID, bucketName, err), "", w)
		return
	}

	end := time.Now().UnixNano()/1000000 - start
	Infof("BigFile upload success file=%v time=%v", file, end)

	Infof("Upload file response data=%v\n", string(respbody))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respbody))
}

var completeMultipartHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Complete shards panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("Complete shards: %v", err), "", w)
		}
	}()

	userID := r.Header.Get("Authorization")
	bucketName := r.Header.Get("BucketName")
	size := r.Header.Get("Length")
	filesize, _ := strconv.Atoi(size)

	// tokenContent := r.Header.Get("X-Dynamic-Token")
	// if !ValidateToken(tokenContent, []byte(MyEnv["TOKENKEY"])) {
	// 	HandleError(http.StatusUnauthorized, fmt.Errorf("Complete shards permission denied"), "", w)
	// 	return
	// }

	vars := mux.Vars(r)
	file := vars["file"]

	var completeList CompleteMultipartUpload

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&completeList)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Parse request body : %v", err), "", w)
		return
	}

	bytes, err := json.Marshal(completeList)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("CompleteMultipartUpload marshal : %v", err), "", w)
		return
	}
	s := string(bytes)

	Infof("CompleteMultipartUpload user=%v bucket=%v", userID, bucketName)
	Infof("CompleteMultipartUpload shards=%v", s)

	para := strings.NewReader(s)
	req, err := http.NewRequest("PATCH", MyEnv["IPFSGATEWAY"]+"file", para)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("CompleteMultipartUpload new request : %v", err), "", w)
		return
	}

	req.Header.Add("X-Dynamic-Token", MyEnv["NODETOKEN"])

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("CompleteMultipartUpload do request : %v", err), "", w)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		HandleError(http.StatusInternalServerError, fmt.Errorf("CompleteMultipartUpload status code : %v", response.StatusCode), "", w)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Upload file user=%v bucket=%v failed, %v", userID, bucketName, err), "", w)
		return
	}

	hash := string(body)
	Infof("CompleteMultipartUpload shards file hash=%v size=%v", hash, filesize)

	//update user objects
	obj, err := GetUserObjects(userID)
	if err != nil {
		CheckStateChannel(fmt.Errorf("CompleteMultipartUpload user=%v get user object failed, %v", userID, err), w)
		return
	}

	filecontext := &File{
		Name:     file,
		Hash:     hash,
		Owner:    userID,
		IsPublic: false,
		Time:     time.Now().Unix(),
		Size:     uint64(filesize),
		FileType: true,
	}
	err = obj.CreateFile(bucketName, file, filecontext)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("CompleteMultipartUpload user=%v create file=%v failed, %v", userID, file, err), "", w)
		return
	}

	err = UpdateDHTAndContract(userID, obj)
	if err != nil {
		CheckStateChannel(fmt.Errorf("CompleteMultipartUpload user=%v file=%v update object and  index file failed, %v", userID, file, err), w)
		return
	}

	HandleError(http.StatusOK, nil, fmt.Sprintf("CompleteMultipartUpload user=%v bucket=%v success", userID, bucketName), w)
}

var getMultipartHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Get multipart panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("Get multipart failed: %v", err), "", w)
		}
	}()

	userID := r.Header.Get("Authorization")
	bucketName := r.Header.Get("BucketName")

	vars := mux.Vars(r)
	hash := vars["hash"]

	Infof("Download shard file user=%v bucket=%v shardhash=%v", userID, bucketName, hash)

	body, err := GetFileWithBody(hash)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Download shard file user=%v bucket=%v shardhash=%v %v", userID, bucketName, hash, err), "", w)
		return
	}

	Infof("Download shard file success, user=%v bucket=%v shardhash=%v", userID, bucketName, hash)

	//read file content
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
