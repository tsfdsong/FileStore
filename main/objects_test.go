package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_CreateObjects(t *testing.T) {
	user := "0x62415a4b78e054b200244010ace64ed2bdc7c69a"
	bname := "www.example.com"

	obj := CreateObjects(user)
	obj.AddBucket(bname)

	//create dir
	path := "path1/"
	err := obj.CreateDir(bname, path)
	if err != nil {
		t.Errorf("CreateDir : %v %v", err, path)
		return
	}

	path = "path1/path2/"
	err = obj.CreateDir(bname, path)
	if err != nil {
		t.Errorf("CreateDir : %v %v", err, path)
		return
	}

	path = "path1/path3/"
	err = obj.CreateDir(bname, path)
	if err != nil {
		t.Errorf("CreateDir : %v %v", err, path)
		return
	}

	//create file
	filepath := "path4/2.jpg"
	file := &File{
		Name:     "2.jpg",
		Hash:     "656fc88239fed34577fca4084cf2add0",
		Owner:    "0x62415a4b78e054b200244010ace64ed2bdc7c69a",
		IsPublic: false,
		Time:     time.Now().Unix(),
		Size:     100,
	}
	err = obj.CreateFile(bname, filepath, file)
	if err != nil {
		t.Errorf("CreateFile : %v %v", err, filepath)
		return
	}

	//create file
	filepath = "path1/4.jpg"
	file = &File{
		Name:     "4.jpg",
		Hash:     "656fc88239fed34577fca4084cf2add1",
		Owner:    "0x62415a4b78e054b200244010ace64ed2bdc7c69a",
		IsPublic: false,
		Time:     time.Now().Unix(),
		Size:     120,
	}
	err = obj.CreateFile(bname, filepath, file)
	if err != nil {
		t.Errorf("CreateFile : %v %v", err, filepath)
		return
	}

	//create file
	filepath = "3.jpg"
	file = &File{
		Name:     "3.jpg",
		Hash:     "656fc88239fed34577fca4084cf2add2",
		Owner:    "0x62415a4b78e054b200244010ace64ed2bdc7c69a",
		IsPublic: false,
		Time:     time.Now().Unix(),
		Size:     140,
	}
	err = obj.CreateFile(bname, filepath, file)
	if err != nil {
		t.Errorf("CreateFile : %v %v", err, filepath)
		return
	}

	//create file
	filepath = "path1/path2/src/1.jpg"
	file = &File{
		Name:     "1.jpg",
		Hash:     "656fc88239fed34577fca4084cf2add3",
		Owner:    "0x62415a4b78e054b200244010ace64ed2bdc7c69a",
		IsPublic: false,
		Time:     time.Now().Unix(),
		Size:     160,
	}
	err = obj.CreateFile(bname, filepath, file)
	if err != nil {
		t.Errorf("CreateFile : %v %v", err, filepath)
		return
	}

	//marshal to json
	bytes, err := obj.MarshalToJSON()
	if err != nil {
		t.Errorf("MarshalToJSON : %v", err)
		return
	}

	hash, err := GetFileHash(bytes)
	if err != nil {
		t.Errorf("GetFileHash : %v", err)
		return
	}

	err = PutDHT(bytes, hash)
	if err != nil {
		t.Errorf("PutDHT : %v", err)
		return
	}
}

//
func Test_ListDir(t *testing.T) {
	hash := "f86437192dbc7256e1d309e107a646c9"
	bytes, err := GetDHT(hash)
	if err != nil {
		t.Errorf("GetDHT : %v %v", err, hash)
		return
	}

	user := "0x62415a4b78e054b200244010ace64ed2bdc7c69b"
	bname := "www.example.com"
	obj, err := UnmarshalToObjects(user, bytes)
	if err != nil {
		t.Errorf("UnmarshalToObjects : %v %v", err, hash)
		return
	}

	path := ""
	dir, files, err := obj.ListDir(bname, path)
	if err != nil {
		t.Errorf("ListDir : %v %v", err, path)
		return
	}

	fmt.Printf("ListDir dir: %v\n", dir)
	for _, v := range files {
		fmt.Printf("ListDir file: %v\n", v)
	}
}

func Test_GetFileInfo(t *testing.T) {
	hash := "f86437192dbc7256e1d309e107a646c9"
	bytes, err := GetDHT(hash)
	if err != nil {
		t.Errorf("GetDHT : %v %v", err, hash)
		return
	}

	user := "0x62415a4b78e054b200244010ace64ed2bdc7c69b"
	bname := "www.example.com"
	obj, err := UnmarshalToObjects(user, bytes)
	if err != nil {
		t.Errorf("UnmarshalToObjects : %v %v", err, hash)
		return
	}

	filepath := "3.jpg"
	file, err := obj.GetFileInfo(bname, filepath)
	if err != nil {
		t.Errorf("GetFileInfo : %v %v", err, hash)
		return
	}

	fmt.Printf("%v= %v", filepath, *file)
}

func Test_ShareFile(t *testing.T) {
	hash := "f86437192dbc7256e1d309e107a646c9"
	bytes, err := GetDHT(hash)
	if err != nil {
		t.Errorf("GetDHT : %v %v", err, hash)
		return
	}

	user := "0x62415a4b78e054b200244010ace64ed2bdc7c69b"
	bname := "www.example.com"
	obj, err := UnmarshalToObjects(user, bytes)
	if err != nil {
		t.Errorf("UnmarshalToObjects : %v %v", err, hash)
		return
	}

	filepath := "3.jpg"
	obj.ShareFile(bname, filepath)

	//marshal to json
	jsonBytes, err := obj.MarshalToJSON()
	if err != nil {
		t.Errorf("MarshalToJSON : %v", err)
		return
	}

	err = PutDHT(jsonBytes, "test")
	if err != nil {
		t.Errorf("PutDHT : %v", err)
		return
	}
}

//
func Test_DeleteFile(t *testing.T) {
	hash := "f86437192dbc7256e1d309e107a646c9"
	bytes, err := GetDHT(hash)
	if err != nil {
		t.Errorf("GetDHT : %v %v", err, hash)
		return
	}

	user := "0x62415a4b78e054b200244010ace64ed2bdc7c69b"
	bname := "www.example.com"
	obj, err := UnmarshalToObjects(user, bytes)
	if err != nil {
		t.Errorf("UnmarshalToObjects : %v %v", err, hash)
		return
	}

	filepath := "3.jpg"
	err = obj.DeleteFile(bname, filepath)
	if err != nil {
		t.Errorf("DeleteFile : %v %v", err, hash)
		return
	}

	//marshal to json
	jsonBytes, err := obj.MarshalToJSON()
	if err != nil {
		t.Errorf("MarshalToJSON : %v", err)
		return
	}

	err = PutDHT(jsonBytes, "test_deletefile")
	if err != nil {
		t.Errorf("PutDHT : %v", err)
		return
	}
}

//
func Test_DeleteDir(t *testing.T) {
	hash := "f86437192dbc7256e1d309e107a646c9"
	bytes, err := GetDHT(hash)
	if err != nil {
		t.Errorf("GetDHT : %v %v", err, hash)
		return
	}

	user := "0x62415a4b78e054b200244010ace64ed2bdc7c69b"
	bname := "www.example.com"
	obj, err := UnmarshalToObjects(user, bytes)
	if err != nil {
		t.Errorf("UnmarshalToObjects : %v %v", err, hash)
		return
	}

	filepath := "path1/"
	err = obj.DeleteDir(bname, filepath)
	if err != nil {
		t.Errorf("DeleteFile : %v %v", err, hash)
		return
	}

	//marshal to json
	jsonBytes, err := obj.MarshalToJSON()
	if err != nil {
		t.Errorf("MarshalToJSON : %v", err)
		return
	}

	err = PutDHT(jsonBytes, "test_deletedir")
	if err != nil {
		t.Errorf("PutDHT : %v", err)
		return
	}
}
