package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

//GetUserAndBucketName ...
func GetUserAndBucketName(r *http.Request) (string, string, error) {
	userid := r.Header.Get("Authorization")
	bucketname := r.Header.Get("BucketName")

	return userid, bucketname, nil
}

//initStoreContract init user file store contract
var initStoreContract = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Init store contract panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("create store contract : %v", err), "", w)
		}
	}()

	vars := mux.Vars(r)
	userid := vars["userid"]

	Infof("InitStoreContract user=%v", userid)

	type ETHKey struct {
		PrivateKey string `json:"privatekey"`
	}

	decoder := json.NewDecoder(r.Body)
	var key ETHKey
	err := decoder.Decode(&key)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Parse request body : %v", err), "", w)
		return
	}

	if ok := ValidateAccountPrivateKey(userid, key.PrivateKey); !ok {
		HandleError(http.StatusInternalServerError, fmt.Errorf("init store contract user and private key not matach"), "", w)
		return
	}

	contractAddr, err := CheckUserContract(userid, key.PrivateKey)
	if err != nil {
		HandleError(http.StatusInternalServerError, err, "", w)
		return
	}

	//user store contract not exist,create it
	if contractAddr == "" {
		contractAddr, err = InitStoreContract(userid, key.PrivateKey)
		if err != nil {
			HandleError(http.StatusInternalServerError, err, "", w)
			return
		}
	}

	if contractAddr == "" {
		HandleError(http.StatusInternalServerError, fmt.Errorf(" Get user=%v store contract address empty", userid), "", w)
		return
	}

	//user store contract exist, setup state channel and authorize
	err = InitChannelState(userid, key.PrivateKey, contractAddr)
	if err != nil {
		HandleError(http.StatusInternalServerError, err, "", w)
		return
	}

	HandleError(http.StatusOK, nil, fmt.Sprintf("%v", contractAddr), w)
}

//Buckket API
var createBucketHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Create bucket panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("create bucket : %v", err), "", w)
		}
	}()

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get useid and bucket name, %v ", err), "", w)
		return
	}

	Infof("Create bucket user=%v bucket=%v", userid, bucketname)

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v create bucket=%v failed, %v", userid, bucketname, err), w)
		return
	}

	obj.AddBucket(bucketname)

	err = UpdateDHTAndContract(userid, obj)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v create bucket=%v failed, %v", userid, bucketname, err), w)
		return
	}

	HandleError(http.StatusOK, nil, fmt.Sprintf("user=%v create bucket=%v success", userid, bucketname), w)
}

var getUserBucketsHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Get user bucket panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("get user bucket : %v", err), "", w)
		}
	}()

	vars := mux.Vars(r)
	userid := vars["userid"]

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v get all bucket failed, %v", userid, err), w)
		return
	}

	Infof("Get user all buckets user=%v", userid)

	buckets := obj.GetBuckets()

	type BucketResponse struct {
		BucketName string `json:"bucketname"`
		Time       int64  `json:"time"`
	}

	var res []BucketResponse

	for _, v := range buckets {
		var temp BucketResponse
		temp.BucketName = v.Name
		temp.Time = v.Time

		res = append(res, temp)
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v get all bucket failed, %v", userid, err), "", w)
	} else {
		HandleError(http.StatusOK, nil, string(bytes), w)
	}
}

var getBucketInfoHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Get bucket info panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("get bucket info: %v", err), "", w)
		}
	}()

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get useid and bucket name, %v ", err), "", w)
		return
	}

	Infof("Get bucket info user=%v bucket=%v", userid, bucketname)

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v get bucket=%v info failed, %v", userid, bucketname, err), w)
		return
	}

	type BucketResult struct {
		Files []*File  `json:"files"`
		Dir   []string `json:"dir"`
	}
	var response BucketResult

	dir, files, err := obj.ListDir(bucketname, "")
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v get bucket=%v failed, %v", userid, bucketname, err), "", w)
		return
	}

	response.Files = append(response.Files, files...)
	response.Dir = append(response.Dir, dir...)

	bytes, err := json.Marshal(response)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v get bucket=%v info failed, %v", userid, bucketname, err), "", w)
	} else {
		HandleError(http.StatusOK, nil, string(bytes), w)
	}
}

var deleteBucketHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Delete bucket panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("delete bucket: %v", err), "", w)
		}
	}()

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get useid and bucket name, %v ", err), "", w)
		return
	}

	Infof("Delete bucket user=%v bucket=%v", userid, bucketname)

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v delete bucket=%v info failed, %v", userid, bucketname, err), w)
		return
	}

	obj.Delete(bucketname)

	err = UpdateDHTAndContract(userid, obj)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v delete bucket=%v failed, %v", userid, bucketname, err), w)
		return
	}

	HandleError(http.StatusOK, nil, fmt.Sprintf("user=%v delete bucket=%v success", userid, bucketname), w)
}

//File/Dir API
var putHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Put file/dir panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("put file/dir: %v", err), "", w)
		}
	}()

	r.Body = http.MaxBytesReader(w, r.Body, 200*1024*1024)

	vars := mux.Vars(r)
	file := vars["file"]

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get useid and bucket name, %v ", err), "", w)
		return
	}

	Infof("Create file/dir user=%v bucket=%v file=%v", userid, bucketname, file)

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v create file/dir=%v failed, %v", userid, file, err), w)
		return
	}

	if file != "" && !strings.HasSuffix(file, "/") {
		//read file content
		subStrs := strings.Split(file, "/")
		fileName := subStrs[len(subStrs)-1]

		inbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("put file=%v read body failed, user=%v %v", userid, file, err), "", w)
			return
		}

		fileSize := uint64(len(inbody))
		Infof("Create file/dir bodysize, user=%v bucket=%v file=%v size=%v", userid, bucketname, file, fileSize)

		hashStr, err := PutFile(inbody, fileName)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("put file=%v store system failed, user=%v %v", userid, file, err), "", w)
			return
		}

		filecontext := &File{
			Name:     file,
			Hash:     hashStr,
			Owner:    userid,
			IsPublic: false,
			Time:     time.Now().Unix(),
			Size:     fileSize,
			FileType: false,
		}
		err = obj.CreateFile(bucketname, file, filecontext)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v create file=%v failed, %v", userid, file, err), "", w)
			return
		}

		err = UpdateDHTAndContract(userid, obj)
		if err != nil {
			CheckStateChannel(fmt.Errorf("user=%v create file=%v failed, %v", userid, file, err), w)
			return
		}

		HandleError(http.StatusOK, nil, fmt.Sprintf("user=%v upload file=%v hash=%v success", userid, file, hashStr), w)
	} else {
		//dir create
		err := obj.CreateDir(bucketname, file)
		if err != nil {
			CheckStateChannel(fmt.Errorf("user=%v create dir failed, %v", userid, err), w)
			return
		}

		err = UpdateDHTAndContract(userid, obj)
		if err != nil {
			CheckStateChannel(fmt.Errorf("user=%v create dir=%v failed, %v", userid, file, err), w)
			return
		}

		HandleError(http.StatusOK, nil, fmt.Sprintf("user=%v create dir=%v success", userid, file), w)
	}
}

var getHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Get file/dir panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("get file/dir: %v", err), "", w)
		}
	}()

	vars := mux.Vars(r)
	file := vars["file"]

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get useid and bucket name, %v ", err), "", w)
		return
	}

	Infof("Get file user=%v bucket=%v file=%v", userid, bucketname, file)

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v get file/dir=%v failed, %v", userid, file, err), w)
		return
	}

	if file != "" && !strings.HasSuffix(file, "/") {
		//file
		fileObj, err := obj.GetFileInfo(bucketname, file)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v get file=%v failed, %v", userid, file, err), "", w)
			return
		}

		Infof("Download file user=%v bucket=%v hash=%v", userid, bucketname, fileObj.Hash)

		body, err := GetFileWithBody(fileObj.Hash)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("Download file user=%v bucket=%v shardhash=%v %v", userid, bucketname, fileObj.Hash, err), "", w)
			return
		}

		Infof("Download file success, user=%v bucket=%v hash=%v", userid, bucketname, fileObj.Hash)

		//read file content
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	} else {
		//dir
		HandleError(http.StatusBadRequest, nil, string("The fuction not support"), w)
	}
}

var deleteHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Delete file/dir panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("delete file/dir: %v", err), "", w)
		}
	}()

	vars := mux.Vars(r)
	file := vars["file"]

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get useid and bucket name, %v ", err), "", w)
		return
	}

	Infof("Delete file/dir user=%v bucket=%v file=%v", userid, bucketname, file)

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v delete file/dir=%v failed, %v", userid, file, err), w)
		return
	}

	if file != "" && !strings.HasSuffix(file, "/") {
		//file
		err = obj.DeleteFile(bucketname, file)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v delete file=%v failed, %v", userid, file, err), "", w)
			return
		}

		err = UpdateDHTAndContract(userid, obj)
		if err != nil {
			CheckStateChannel(fmt.Errorf("user=%v delete file=%v failed, %v", userid, file, err), w)
			return
		}

		HandleError(http.StatusOK, nil, fmt.Sprintf("user=%v delete file=%v success", userid, file), w)
	} else {
		obj.DeleteDir(bucketname, file)

		err = UpdateDHTAndContract(userid, obj)
		if err != nil {
			CheckStateChannel(fmt.Errorf("user=%v delete dir=%v failed, %v", userid, file, err), w)
			return
		}

		HandleError(http.StatusOK, nil, fmt.Sprintf("user=%v delete dir=%v success", userid, file), w)
	}
}

var getinfoHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Get info panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("get info: %v", err), "", w)
		}
	}()

	userid := r.Header.Get("Authorization")

	Infof("Get info user=%v", userid)

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("get info failed, object user=%v %v", userid, err), w)
		return
	}

	hash, err := GetFileIndex(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("get info failed, index user=%v %v", userid, err), w)
		return
	}

	Infof("Get info, get file index user=%v indexhash=%v", userid, hash)

	//get index content
	bytes, err := obj.MarshalToJSON()
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get info marshal onjects failed, user=%v %v ", userid, err), "", w)
		return
	}

	reply := &struct {
		IndexHash    string `json:"indexhash"`
		IndexContent []byte `json:"indexcontent"`
	}{
		IndexHash:    hash,
		IndexContent: bytes,
	}

	res, err := json.Marshal(reply)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get info marshal response failed, user=%v %v ", userid, err), "", w)
		return
	}

	Infof("Get info success, user=%v indexhash=%v", userid, hash)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

var getAttributeHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Get file/dir attribute panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("Get file/dir attribute: %v", err), "", w)
		}
	}()

	vars := mux.Vars(r)
	file := vars["file"]

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get file/dir attribute get useid and bucket name failed, %v ", err), "", w)
		return
	}

	Infof("Get file/dir attribute user=%v bucket=%v file=%v", userid, bucketname, file)

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v get file attribute file/dir=%v failed, %v", userid, file, err), w)
		return
	}

	if file != "" && !strings.HasSuffix(file, "/") {
		//file
		fileinfo, err := obj.GetFileInfo(bucketname, file)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v get file attribute file=%v failed, %v", userid, file, err), "", w)
			return
		}

		bytes, err := json.Marshal(fileinfo)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v get file attribute  dir=%v failed, %v", userid, file, err), "", w)
			return
		}

		HandleError(http.StatusOK, nil, string(bytes), w)
	} else {
		//folder
		curBucket, ok := obj.Buckets[bucketname]
		if !ok {
			HandleError(http.StatusInternalServerError, fmt.Errorf("get file attribute bucket not found, user=%v bucketname=%v %v", userid, bucketname, err), "", w)
			return
		}

		item, ok := curBucket.Folders[file]
		if !ok {
			HandleError(http.StatusInternalServerError, fmt.Errorf("get file attribute folders not found, user=%v bucketname=%v foldername=%v %v", userid, bucketname, file, err), "", w)
			return
		}

		resBytes, err := json.Marshal(item)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("get file attribute marshal failed, user=%v bucketname=%v foldername=%v %v", userid, bucketname, file, err), "", w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resBytes)
	}
}

var headAttributeHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("file/dir attribute panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("file/dir attribute: %v", err), "", w)
		}
	}()

	vars := mux.Vars(r)
	file := vars["file"]

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get useid and bucket name, %v ", err), "", w)
		return
	}

	Infof("Get file/dir attribute user=%v bucket=%v file=%v", userid, bucketname, file)

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v get file attribute file/dir=%v failed, %v", userid, file, err), w)
		return
	}

	if file != "" && !strings.HasSuffix(file, "/") {
		//file
		fileinfo, err := obj.GetFileInfo(bucketname, file)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v get file attribute file=%v failed, %v", userid, file, err), "", w)
			return
		}

		bytes, err := json.Marshal(fileinfo)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v get file attribute  dir=%v failed, %v", userid, file, err), "", w)
			return
		}

		HandleError(http.StatusOK, nil, string(bytes), w)
	} else {
		type DirResult struct {
			Files []*File  `json:"files"`
			Dir   []string `json:"dir"`
		}
		var response DirResult

		dir, files, err := obj.ListDir(bucketname, file)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v get dir attribute dir=%v failed, %v", userid, file, err), "", w)
			return
		}

		response.Files = append(response.Files, files...)
		response.Dir = append(response.Dir, dir...)

		bytes, err := json.Marshal(response)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v get dir attribute dir=%v failed, %v", userid, file, err), "", w)
			return
		}

		HandleError(http.StatusOK, nil, string(bytes), w)
	}
}

var shareFileHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Share file/dir panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("share file/dir: %v", err), "", w)
		}
	}()

	vars := mux.Vars(r)
	file := vars["file"]

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get useid and bucket name, %v ", err), "", w)
		return
	}

	Infof("Share file/dir user=%v bucket=%v file=%v", userid, bucketname, file)

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v share file=%v failed, %v", userid, file, err), w)
		return
	}

	if file != "" && !strings.HasSuffix(file, "/") {
		//file
		obj.ShareFile(bucketname, file)

		err = UpdateDHTAndContract(userid, obj)
		if err != nil {
			CheckStateChannel(fmt.Errorf("user=%v share file=%v failed, %v", userid, file, err), w)
			return
		}

		HandleError(http.StatusOK, nil, fmt.Sprintf("user=%v share file=%v success", userid, file), w)
	} else {
		//dir
		HandleError(http.StatusBadRequest, fmt.Errorf("user=%v share dir=%v not supported", userid, file), "", w)
	}
}

var moveHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Move file/dir panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("move file/dir: %v", err), "", w)
		}
	}()

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("move file/dir get useid and bucket name, %v ", err), "", w)
		return
	}

	type MoveItem struct {
		From     string `json:"from"`
		To       string `json:"to"`
		MoveType int    `json:"type"`
	}

	decoder := json.NewDecoder(r.Body)
	var key MoveItem
	err = decoder.Decode(&key)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Move file/dir Parse request body : %v", err), "", w)
		return
	}

	Infof("Move user=%v from=%v to=%v type=%v", userid, key.From, key.To, key.MoveType)

	obj, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("user=%v move file/dir failed, %v", userid, err), w)
		return
	}

	//type: 0: folder; 1 : file
	if key.MoveType == 1 {
		//file
		err = obj.MoveFile(bucketname, key.From, key.To)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v move file from=%v to=%v failed, %v", userid, key.From, key.To, err), "", w)
			return
		}

		err = UpdateDHTAndContract(userid, obj)
		if err != nil {
			CheckStateChannel(fmt.Errorf("user=%v move file from=%v to=%v update obj failed, %v", userid, key.From, key.To, err), w)
			return
		}

		HandleError(http.StatusOK, nil, fmt.Sprintf("user=%v move file from=%v to=%v success", userid, key.From, key.To), w)
	} else {
		err = obj.MoveDir(bucketname, key.From, key.To)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("user=%v move dir from=%v to=%v failed, %v", userid, key.From, key.To, err), "", w)
			return
		}

		err = UpdateDHTAndContract(userid, obj)
		if err != nil {
			CheckStateChannel(fmt.Errorf("user=%v move dir from=%v to=%v update obj failed, %v", userid, key.From, key.To, err), w)
			return
		}

		HandleError(http.StatusOK, nil, fmt.Sprintf("user=%v move dir from=%v to=%v success", userid, key.From, key.To), w)
	}
}

var copyHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Copy file/dir panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("move file/dir: %v", err), "", w)
		}
	}()

	type CopyItem struct {
		FromUser   string `json:"fromuser"`
		From       string `json:"from"`
		FromBucket string `json:"frombucket"`
		ToUser     string `json:"touser"`
		To         string `json:"to"`
		ToBucket   string `json:"tobucket"`
		MoveType   int    `json:"type"`
	}

	decoder := json.NewDecoder(r.Body)
	var key CopyItem
	err := decoder.Decode(&key)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Copy file/dir Parse request body : %v", err), "", w)
		return
	}

	Infof("Copy file/dir from=%v fromuser=%v frombucket=%v to=%v touser=%v tobucket=%v type=%v\n", key.From, key.FromUser, key.FromBucket, key.To, key.ToUser, key.ToBucket, key.MoveType)

	fromObj, err := GetUserObjects(key.FromUser)
	if err != nil {
		CheckStateChannel(fmt.Errorf("Copy from user=%v get user object failed, %v", key.FromUser, err), w)
		return
	}

	toObj, err := GetUserObjects(key.ToUser)
	if err != nil {
		CheckStateChannel(fmt.Errorf("Copy to user=%v get user object failed, %v", key.ToUser, err), w)
		return
	}

	toBucket, ok := toObj.Buckets[key.ToBucket]
	if !ok {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Copy file/dir to bucketname=%v not found", key.ToBucket), "", w)
		return
	}

	//type: 0: folder; 1 : file
	if key.MoveType == 1 {
		//file
		newBucket, err := fromObj.CopyFile(key.FromBucket, key.From, key.To, toBucket)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("Copy file %v from=%v %v to=%v failed, %v", key.FromUser, key.From, key.ToUser, key.To, err), "", w)
			return
		}

		toObj.Buckets[key.ToBucket] = newBucket

		err = UpdateDHTAndContract(key.ToUser, toObj)
		if err != nil {
			CheckStateChannel(fmt.Errorf("Copy file %v from=%v %v to=%v update obj failed, %v", key.FromUser, key.From, key.ToUser, key.To, err), w)
			return
		}

		HandleError(http.StatusOK, nil, fmt.Sprintf("Copy file %v from=%v %v to=%v success", key.FromUser, key.From, key.ToUser, key.To), w)
	} else {
		newBucket, err := fromObj.CopyDir(key.FromBucket, key.From, key.To, toBucket)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("Copy dir %v from=%v %v to=%v failed, %v", key.FromUser, key.From, key.ToUser, key.To, err), "", w)
			return
		}

		toObj.Buckets[key.ToBucket] = newBucket

		err = UpdateDHTAndContract(key.ToUser, toObj)
		if err != nil {
			CheckStateChannel(fmt.Errorf("Copy dir %v from=%v %v to=%v update obj failed, %v", key.FromUser, key.From, key.ToUser, key.To, err), w)
			return
		}

		HandleError(http.StatusOK, nil, fmt.Sprintf("Copy dir %v from=%v %v to=%v success", key.FromUser, key.From, key.ToUser, key.To), w)
	}
}

var renameHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Rename file/dir panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("Rename file/dir: %v", err), "", w)
		}
	}()

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Rename useid and bucket name, %v ", err), "", w)
		return
	}

	type CopyItem struct {
		OldName    string `json:"oldname"`
		NewName    string `json:"newname"`
		RenameType int    `json:"type"`
	}

	decoder := json.NewDecoder(r.Body)
	var key CopyItem
	err = decoder.Decode(&key)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Rename file/dir Parse request body : %v", err), "", w)
		return
	}

	Infof("Rename file/dir user=%v bucketname=%v oldname=%v newname=%v  renametype=%v\n", userid, bucketname, key.OldName, key.NewName, key.RenameType)

	currentObject, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("Rename from user=%v get user object failed, %v", userid, err), w)
		return
	}

	//type: 0: folder; 1 : file
	if key.RenameType == 1 {
		//file
		err = currentObject.RenameFile(bucketname, key.OldName, key.NewName)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("Rename file oldname=%v newname=%v failed, %v", key.OldName, key.NewName, err), "", w)
			return
		}

		err = UpdateDHTAndContract(userid, currentObject)
		if err != nil {
			CheckStateChannel(fmt.Errorf("Rename file oldname=%v newname=%v update obj failed, %v", key.OldName, key.NewName, err), w)
			return
		}

		HandleError(http.StatusOK, nil, fmt.Sprintf("Rename file user=%v bucketname=%v oldname=%v newname=%v success", userid, bucketname, key.OldName, key.NewName), w)
	} else {
		err = currentObject.RenameDir(bucketname, key.OldName, key.NewName)
		if err != nil {
			HandleError(http.StatusInternalServerError, fmt.Errorf("Rename dir oldname=%v newname=%v failed, %v", key.OldName, key.NewName, err), "", w)
			return
		}

		err = UpdateDHTAndContract(userid, currentObject)
		if err != nil {
			CheckStateChannel(fmt.Errorf("Rename dir oldname=%v newname=%v update obj failed, %v", key.OldName, key.NewName, err), w)
			return
		}

		HandleError(http.StatusOK, nil, fmt.Sprintf("Rename dir user=%v bucketname=%v oldname=%v newname=%v success", userid, bucketname, key.OldName, key.NewName), w)
	}
}

var getsizeHandle = func(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Panicf("Get size panic, %s", string(debug.Stack()))
			HandleError(http.StatusInternalServerError, fmt.Errorf("Get size: %v", err), "", w)
		}
	}()

	userid, bucketname, err := GetUserAndBucketName(r)
	if err != nil {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get size useid and bucket name, %v ", err), "", w)
		return
	}

	Infof("Get size user=%v bucket=%v.", userid, bucketname)

	currentObject, err := GetUserObjects(userid)
	if err != nil {
		CheckStateChannel(fmt.Errorf("Rename from user=%v get user object failed, %v", userid, err), w)
		return
	}

	curBucket, ok := currentObject.Buckets[bucketname]
	if !ok {
		HandleError(http.StatusInternalServerError, fmt.Errorf("Get size bucketname=%v not found", bucketname), "", w)
		return
	}

	var resSize float64
	for _, item := range curBucket.Files {
		temp := (float64)(item.Size / 1024.00)

		resSize = resSize + temp
	}

	Infof("Get size user=%v bucket=%v size=%.2fKB", userid, bucketname, resSize)

	HandleError(http.StatusOK, nil, fmt.Sprintf("%.2f", resSize), w)
}
