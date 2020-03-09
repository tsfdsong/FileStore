package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

//File ...
type File struct {
	Name     string `json:"name"`
	Hash     string `json:"hash"`
	Owner    string `json:"owner"`
	IsPublic bool   `json:"ispublic"`
	Time     int64  `json:"time"`
	Size     uint64 `json:"size"`
	FileType bool   `json:"filetype"`
}

//Folder ...
type Folder struct {
	Name string `json:"name"`
	Time int64  `json:"time"`
}

//Bucket ...
type Bucket struct {
	Name    string             `json:"name"`
	Files   map[string]*File   `json:"files"`   //allpath => File
	Folders map[string]*Folder `json:"folders"` //allpath => File
	Index   []string           `json:"index"`
	Time    int64              `json:"time"`
}

//Objects ...
type Objects struct {
	User    string             `json:"user"` //eth address
	Buckets map[string]*Bucket `json:"buckets"`
}

//CreateObjects ...
func CreateObjects(user string) *Objects {
	return &Objects{
		User:    user,
		Buckets: make(map[string]*Bucket),
	}
}

//RemoveDuplicateItem ...
func RemoveDuplicateItem(strs []string) []string {
	result := make([]string, 0, len(strs))
	temp := map[string]struct{}{}

	for _, item := range strs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

//AddItem ...
func AddItem(src []string, in string) []string {
	result := src

	isFound := false
	for _, item := range src {
		if item == in {
			isFound = true
			break
		}
	}

	if !isFound {
		result = append(result, in)
	}

	return result
}

//RemoveItem ...
func RemoveItem(src []string, in string) []string {
	res := src
	for i, item := range src {
		if item == in {
			res = append(src[:i], src[i+1:]...)
			break
		}
	}

	return res
}

//AddBucket ...
func (o *Objects) AddBucket(bucketname string) {
	if _, ok := o.Buckets[bucketname]; !ok {
		o.Buckets[bucketname] = &Bucket{
			Name:    bucketname,
			Files:   make(map[string]*File),
			Folders: make(map[string]*Folder),
			Index:   make([]string, 0),
			Time:    time.Now().Unix(),
		}
	}
}

//GetBuckets ...
func (o *Objects) GetBuckets() []*Bucket {
	var result []*Bucket
	for _, v := range o.Buckets {
		result = append(result, v)
	}
	return result
}

//GetBucketInfo ...
func (o *Objects) GetBucketInfo(name string) *Bucket {
	if item, ok := o.Buckets[name]; ok {
		return item
	}

	return nil
}

//Delete ...
func (o *Objects) Delete(name string) {
	delete(o.Buckets, name)
}

//CreateDir ...
func (o *Objects) CreateDir(bucketname, path string) error {
	if path == "" {
		return fmt.Errorf("CreateDir dir name %v is empty", path)
	}

	if !strings.HasSuffix(path, "/") {
		//path must end with "/"
		return fmt.Errorf("CreateDir input %v must end with '/'", path)
	}

	item, ok := o.Buckets[bucketname]
	if !ok {
		return fmt.Errorf("CreateDir bucketname=%v not exist", bucketname)
	}

	item.Index = AddItem(item.Index, path)

	tempFolder := &Folder{
		Name: path,
		Time: time.Now().Unix(),
	}
	item.Folders[path] = tempFolder
	return nil
}

//ListDir ...
func (o *Objects) ListDir(bucketname, path string) ([]string, []*File, error) {
	if path != "" && !strings.HasSuffix(path, "/") {
		//path must end with "/"
		return nil, nil, fmt.Errorf("ListDir input %v must end with '/'", path)
	}

	dir := make([]string, 0)
	files := make([]*File, 0)
	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return nil, nil, fmt.Errorf("ListDir bucketname=%v not found", path)
	}

	for _, v := range curBucket.Index {
		if !strings.HasPrefix(v, path) {
			continue
		}

		substr := strings.TrimPrefix(v, path)

		res := strings.Split(substr, "/")

		if len(res) == 1 && !strings.HasSuffix(substr, "/") {
			//file
			if file, ok := curBucket.Files[v]; ok {
				temp := file
				idx := strings.Split(temp.Name, "/")
				temp.Name = idx[len(idx)-1]
				files = append(files, temp)
			}
		} else if len(res) == 0 {
			dir = AddItem(dir, "")
		} else {
			//sub folder
			dir = AddItem(dir, res[0])
		}
	}

	return dir, files, nil
}

//DeleteDir ...
func (o *Objects) DeleteDir(bucketname, path string) error {
	if path == "" {
		return fmt.Errorf("DeleteDir dir name %v is empty", path)
	}

	if !strings.HasSuffix(path, "/") {
		//path must end with "/"
		return fmt.Errorf("DeleteDir input %v must end with '/'", path)
	}

	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return fmt.Errorf("DeleteDir buckname=%v not found", bucketname)
	}
	index := curBucket.Index[:0]

	for _, v := range curBucket.Index {
		if !strings.HasPrefix(v, path) {
			index = append(index, v)
			continue
		}

		if !strings.HasSuffix(v, "/") {
			//file
			delete(curBucket.Files, v)
		}
	}

	curBucket.Index = index

	delete(curBucket.Folders, path)

	return nil
}

//CreateFile ...
// filepath : path+filename
func (o *Objects) CreateFile(bucketname, filepath string, file *File) error {
	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return fmt.Errorf("CreateFile bucketname=%v not found", bucketname)
	}

	curBucket.Index = AddItem(curBucket.Index, filepath)
	curBucket.Files[filepath] = file

	return nil
}

//GetFileInfo ...
func (o *Objects) GetFileInfo(bucketname, filepath string) (*File, error) {
	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return nil, fmt.Errorf("GetFileInfo bucketname=%v not found", bucketname)
	}

	if file, ok := curBucket.Files[filepath]; ok {
		return file, nil
	}

	return nil, fmt.Errorf("GetFileInfo bucketname=%v file=%v info get failed", bucketname, filepath)
}

//ShareFile ...
func (o *Objects) ShareFile(bucketname, filepath string) {
	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return
	}

	if file, ok := curBucket.Files[filepath]; ok {
		file.IsPublic = true
	}

	return
}

//DeleteFile ...
func (o *Objects) DeleteFile(bucketname, filepath string) error {
	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return fmt.Errorf("DeleteFile bucketname=%v not found", bucketname)
	}

	for i, v := range curBucket.Index {
		if v == filepath {
			curBucket.Index = append(curBucket.Index[:i], curBucket.Index[i+1:]...)
			break
		}
	}

	if _, ok := curBucket.Files[filepath]; ok {
		delete(curBucket.Files, filepath)
		return nil
	}

	return fmt.Errorf("DeleteFile buxcketname=%v filepath=%v cannot get file", bucketname, filepath)
}

//MarshalToJSON ...
func (o *Objects) MarshalToJSON() ([]byte, error) {
	bytes, err := json.Marshal(o)
	return bytes, err
}

//String ...
func (o *Objects) String() string {
	bytes, err := json.Marshal(o)
	if err != nil {
		return ""
	}

	return string(bytes)
}

//UnmarshalToObjects ...
func UnmarshalToObjects(user string, bytes []byte) (*Objects, error) {
	result := CreateObjects(user)

	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//MoveFile ...
func (o *Objects) MoveFile(bucketname, frompath, topath string) error {
	if frompath == topath {
		return fmt.Errorf("Move same file")
	}

	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return fmt.Errorf("Move bucketname=%v not found", bucketname)
	}

	for i, v := range curBucket.Index {
		if v == frompath {
			if item, ok := curBucket.Files[frompath]; ok {
				curBucket.Index = append(curBucket.Index[:i], curBucket.Index[i+1:]...)

				delete(curBucket.Files, frompath)

				subStrs := strings.Split(frompath, "/")
				fileName := subStrs[len(subStrs)-1]
				newPath := topath + fileName

				curBucket.Files[newPath] = item
				curBucket.Index = AddItem(curBucket.Index, newPath)

				if strings.HasSuffix(newPath, "/") {
					tempFolder := &Folder{
						Name: newPath,
						Time: time.Now().Unix(),
					}

					curBucket.Folders[newPath] = tempFolder
				}

				return nil
			}
		}
	}

	return fmt.Errorf("Move from=%v to=%v failed", frompath, topath)
}

//MoveDir ...
func (o *Objects) MoveDir(bucketname, frompath, topath string) error {
	if frompath == topath {
		return fmt.Errorf("Move same dir")
	}

	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return fmt.Errorf("Move bucketname=%v not found", bucketname)
	}

	isSuccess := false
	for i := len(curBucket.Index) - 1; i >= 0; i-- {
		v := curBucket.Index[i]
		if strings.HasPrefix(v, frompath) {
			isSuccess = true

			subStrs := strings.Split(frompath, "/")
			dirName := subStrs[len(subStrs)-2]

			tempPath := topath + dirName + "/"
			newpath := strings.Replace(v, frompath, tempPath, 1)

			Debugf("Move dir function frompath=%v topath=%v newpath=%v \n", frompath, topath, newpath)

			curBucket.Index = AddItem(curBucket.Index, newpath)

			if strings.HasSuffix(newpath, "/") {
				tempFolder := &Folder{
					Name: newpath,
					Time: time.Now().Unix(),
				}

				curBucket.Folders[newpath] = tempFolder
			}

			if item, ok := curBucket.Files[v]; ok {
				delete(curBucket.Files, v)
				curBucket.Files[newpath] = item
			}

			curBucket.Index = append(curBucket.Index[:i], curBucket.Index[i+1:]...)
		}
	}

	if isSuccess {
		return nil
	}

	return fmt.Errorf("Move dir from=%v to=%v failed", frompath, topath)
}

//CopyFile ...
func (o *Objects) CopyFile(bucketname, frompath, topath string, buk *Bucket) (*Bucket, error) {
	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return nil, fmt.Errorf("Copy file bucketname=%v not found", bucketname)
	}

	if item, ok := curBucket.Files[frompath]; ok {
		subStrs := strings.Split(frompath, "/")
		fileName := subStrs[len(subStrs)-1]
		newPath := topath + fileName

		Infof("\nCopy from=%v to=%v newpath=%v\n", frompath, topath, newPath)

		if _, ok := buk.Files[newPath]; ok {
			return nil, fmt.Errorf("Copy file=%v exist", newPath)
		}
		buk.Index = AddItem(buk.Index, newPath)

		if strings.HasSuffix(newPath, "/") {
			tempFolder := &Folder{
				Name: newPath,
				Time: time.Now().Unix(),
			}

			buk.Folders[newPath] = tempFolder
		}

		buk.Files[newPath] = item

		return buk, nil
	}

	return nil, fmt.Errorf("Copy file=%v not found", frompath)
}

//CopyDir ...
func (o *Objects) CopyDir(bucketname, frompath, topath string, buk *Bucket) (*Bucket, error) {
	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return nil, fmt.Errorf("Copy dir bucketname=%v not found", bucketname)
	}

	isSuccess := false
	for _, v := range curBucket.Index {
		if strings.HasPrefix(v, frompath) {
			isSuccess = true

			subStrs := strings.Split(frompath, "/")
			dirName := subStrs[len(subStrs)-2]

			tempPath := topath + dirName + "/"
			newpath := strings.Replace(v, frompath, tempPath, -1)

			Debugf("Copy dir function frompath=%v topath=%v newpath=%v raw=%v\n", frompath, topath, newpath, v)

			if _, ok := buk.Files[newpath]; ok {
				return nil, fmt.Errorf("Copy dir=%v exist", newpath)
			}

			if item, ok := curBucket.Files[v]; ok {
				buk.Files[newpath] = item
			}

			buk.Index = AddItem(buk.Index, newpath)

			if strings.HasSuffix(newpath, "/") {
				tempFolder := &Folder{
					Name: newpath,
					Time: time.Now().Unix(),
				}

				buk.Folders[newpath] = tempFolder
			}
		}
	}

	if isSuccess {
		return buk, nil
	}

	return nil, fmt.Errorf("Copy dir=%v failed", frompath)
}

//RenameFile ...
func (o *Objects) RenameFile(bucketname, old, new string) error {
	if old == new {
		return fmt.Errorf("Rename file same name")
	}

	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return fmt.Errorf("Rename file bucketname=%v not found", bucketname)
	}

	if item, ok := curBucket.Files[old]; ok {
		newitem := item
		newitem.Name = new
		curBucket.Files[new] = newitem

		Debugf("Rename: item old=%v new=%v", item.Name, newitem.Name)

		delete(curBucket.Files, old)

		curBucket.Index = RemoveItem(curBucket.Index, old)
		curBucket.Index = AddItem(curBucket.Index, new)

		return nil
	}

	return fmt.Errorf("Rename file=%v not found", old)
}

//RenameDir ...
func (o *Objects) RenameDir(bucketname, old, new string) error {
	if old == new {
		return fmt.Errorf("Rename dir same name")
	}

	if !strings.HasSuffix(old, "/") || !strings.HasSuffix(new, "/") {
		return fmt.Errorf("Rename dir path must hash '/' suffix")
	}

	curBucket, ok := o.Buckets[bucketname]
	if !ok {
		return fmt.Errorf("Rename file bucketname=%v not found", bucketname)
	}

	isSuccess := false

	var newIndex []string
	for _, v := range curBucket.Index {
		if strings.HasPrefix(v, old) {
			isSuccess = true
			newPath := strings.Replace(v, old, new, 1)

			Infof("Rename dir raw=%v oldname=%v newname=%v newpath=%v \n", v, old, new, newPath)

			if item, ok := curBucket.Files[v]; ok {
				curBucket.Files[newPath] = item
				curBucket.Files[newPath].Name = newPath
				delete(curBucket.Files, v)
			}

			newIndex = append(newIndex, newPath)
		} else {
			newIndex = append(newIndex, v)
		}
	}

	curBucket.Index = newIndex

	if isSuccess {
		return nil
	}

	return fmt.Errorf("Rename dir oldname=%v newname=%v failed", old, new)
}
