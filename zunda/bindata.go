// Code generated by go-bindata.
// sources:
// ariaki.txt
// zunda.txt
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ariakiTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x96\x51\xae\x2b\x21\x08\x86\xdf\x5d\x45\x93\x09\xc6\x18\x62\xb7\xe0\xfe\x57\x75\x53\x04\x05\xc4\xce\x99\x5e\x9e\xec\x14\x3f\x7f\x10\xd4\xd7\x13\x83\x6e\x2c\xf0\x48\x4f\x70\x97\x66\x40\xef\xf0\x9f\x3c\xe8\x1d\xb5\xce\x72\xe2\x55\x89\xe0\x7d\xc7\xeb\x95\x46\x05\xa0\x85\x2e\xc4\xfb\x90\x10\x91\x56\xcd\xe4\x4e\xf0\x6b\x77\x9f\xa9\xc3\xc3\x92\x69\x38\x15\xa5\xe0\x4b\xc6\xf9\x2b\x79\x05\xc9\x1b\xbc\x4c\x38\x60\x3d\x6a\xf1\x66\x73\x54\xbb\x5b\xb8\xc6\x3c\x92\xc1\x62\x3e\x88\xe6\x04\xf1\x10\x9d\xdc\x58\x62\x9a\x41\x20\x3b\x6d\x11\x6e\xa3\xd3\x07\xe6\x55\xb5\xd2\xe6\x33\x4b\xe4\xef\xbc\xc8\x67\xfa\x42\xb0\x3b\x6d\xfa\xee\x35\x73\xe0\xc9\x9c\x40\x5f\xe3\x9a\x52\x7f\x1d\x79\xb0\xd5\xc8\x9a\x84\xc2\x11\x95\x61\x15\x26\x2b\x6e\x6b\x12\x25\xc2\xce\xa7\x62\xda\x2b\x26\x19\x5a\x13\xea\x28\xb3\xd1\x31\xed\x15\x00\xe1\x70\xc4\x24\x8d\xe3\x61\xd6\x7b\x00\x2e\x00\x00\xc8\x82\xaa\x7b\xc8\x69\xa9\x7f\x9b\x28\xe3\xe3\x63\x2e\x93\xa3\x1c\x2c\xde\x97\x0e\xff\x6a\xdb\x1e\x27\xc1\x05\x62\x7e\xe6\xfd\xaa\xee\x14\x6f\x7c\x15\xfc\xc1\xfc\x19\x31\x78\xbe\x52\x1f\xe1\x4a\xc8\xb3\x27\x4f\x8b\xab\xab\xd8\x33\x7b\x6f\xa6\xc5\x9b\x40\xbc\xb8\x1e\x1a\x55\xd7\x6a\xf8\x55\x91\x03\x59\x7b\x18\x55\x1a\x38\x63\x97\x66\x8c\x96\x6a\xf3\x73\x8e\x1c\x3d\x0f\xe4\x0a\xf2\xdb\xcc\xed\x5c\x4d\x64\xe3\xde\x3a\xdf\xe7\xe1\x41\xe6\xe3\xf4\xdf\x41\x7f\x43\xbc\xf8\x02\x4d\xe1\xae\xdf\x1b\xcf\xb1\xb9\x42\x75\xff\x3e\xc3\x49\x02\xa0\xf7\x77\x5b\x91\x48\xbf\x95\xf8\x30\x3b\x9b\xa4\x5a\xf7\x1b\xed\x43\x52\xfa\x6f\xdf\x1b\x62\x63\xef\x1a\x8f\x42\xde\x6b\x9d\x91\x80\xe7\xdd\xe1\x47\x16\x3d\x11\x90\xd1\x52\xba\xd5\xf3\x5e\x5b\x82\x0f\x56\x57\xa0\xb3\x0a\x4b\x57\xf9\x0b\xd2\x13\x18\xa8\xae\xa4\x0d\x91\x17\x08\x07\xd6\xcc\xfb\xc0\xf1\xee\x72\xb8\xa4\xd2\xcb\x43\x52\x34\x78\x88\x2c\x60\x5d\xdf\x7a\x6e\x51\x22\x75\x62\x68\x90\x4d\x3b\xa4\x29\x26\xdb\xe0\x4c\x64\x64\x99\x97\xb3\x31\xb8\x63\x21\x85\x0f\xa5\x35\xab\x38\xad\x6d\xfe\xfc\xc8\xf2\xff\xce\xfe\xed\xbd\x22\x19\x00\x20\xae\x06\xc4\xf1\x72\x45\x6a\xd0\x99\x0e\xfa\xeb\x9a\x3d\x66\x78\xff\x02\x00\x00\xff\xff\x19\x41\xad\xf3\x30\x0c\x00\x00")

func ariakiTxtBytes() ([]byte, error) {
	return bindataRead(
		_ariakiTxt,
		"ariaki.txt",
	)
}

func ariakiTxt() (*asset, error) {
	bytes, err := ariakiTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ariaki.txt", size: 3120, mode: os.FileMode(436), modTime: time.Unix(1543992578, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _zundaTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x41\x8a\xb4\xbe\x13\xdd\x7f\xa7\x10\x42\x22\x48\xf8\xf7\x01\x44\xcc\x36\x9b\xac\xb3\xf2\x0a\xb9\x81\x7b\x97\x9e\x41\x88\x88\xb8\x10\x1a\x37\xee\x5e\x5d\xec\x4f\xa5\xa7\xbf\xb6\xdb\xee\xe1\x9b\xf9\x65\x60\x84\xd6\x7a\x55\xa9\x7a\xf5\xaa\xb2\xdf\x1d\x5c\x3f\xbc\xf8\xf3\x73\xa4\xb2\x90\x18\xa4\x2c\xb0\xb7\x45\xf3\x5f\x21\x75\xa0\x18\xbc\xf7\x18\x31\x61\xa2\x88\xd1\xf8\x2a\x2f\x4c\xfd\x11\x55\x97\xb5\xa7\x0e\x23\x86\xa2\x6c\x5b\x5a\xca\x4b\x21\xe5\xe9\xf3\x2c\x04\x5a\x8c\xd7\xd8\xa4\xc5\xc0\x7f\x21\x60\xc0\x24\x1c\x36\x8c\x6d\x25\xf5\x09\x1c\x57\xac\xd4\x63\x6c\x6b\xe3\xeb\x80\x11\x6b\x71\x51\xd8\x28\x0a\x8c\x45\x9e\xd7\x34\x63\x7f\xf6\x51\xd3\x2c\x31\xd4\xd4\x53\xc7\x17\x60\xeb\x09\x13\x46\xea\x55\xc0\xe4\x04\x26\xa1\x32\x75\xb4\xba\x39\x1a\x0d\x75\xd8\xa8\x6f\x43\x8d\x81\xe3\xc1\x86\x09\x83\xa1\x88\x41\x61\x55\xd4\x2b\x55\x61\xb7\x34\x07\x4f\xcb\xc3\x9a\x16\x97\x07\x87\x31\x19\x8c\x58\xb1\x5a\x4b\x1d\x75\xd4\x5b\xea\xb1\xd2\x8c\xcd\x89\x22\xab\xeb\xa3\xbb\x10\x0a\x0c\x52\x51\xa4\x0e\x93\x10\x42\x38\x8a\x58\xb1\x39\x8c\x42\x38\x21\x38\xe4\xdb\x63\x63\x24\x9a\x19\x97\xbf\x35\x8c\x48\x51\x08\x6c\xd8\x1c\xf5\x14\x45\xc0\x40\x33\xa7\x3e\xaf\x1c\xb6\x36\x39\xed\x3d\x5f\x3b\xcf\xc2\x9d\x63\x37\xa7\x34\xbb\x82\x66\x0c\xd4\x71\x6c\x23\xf5\xd8\xb1\x89\xe4\xdf\x51\x8f\xad\xaa\xa8\xd7\x2a\xc8\xfc\x71\xbd\xa6\xa2\xd9\xb3\x43\x3e\xd8\xa4\x94\x34\x5b\x8a\x14\x31\x61\x2d\x38\x80\x5a\x34\x99\x7c\x78\x69\xa8\xa3\x58\xb3\x7b\xf6\x64\xb0\x27\x63\x97\x4a\x30\xb3\xcf\x40\x9d\x13\x14\x9b\xc2\xca\x4c\xab\x63\xe9\x72\x85\xeb\x3d\xa2\xf4\x3f\xdd\x58\xd2\x9c\xe5\x5c\x99\xe0\x64\x7b\xe4\x5d\x45\x7d\xce\xfc\xe0\x84\x31\x83\x06\x6c\x02\x13\x75\xc2\x58\xea\xd8\xe1\xce\x99\xf3\x9c\x02\x5d\x3c\x71\x04\xd7\xa2\xc2\xb5\x48\x17\xe7\xd0\x84\x33\x18\xf8\x3a\x1d\xc5\x2a\xbc\x90\xbb\x60\xb0\x8d\x81\xd3\x83\x2b\x10\x99\x17\x1d\x06\xea\xb9\x44\xa3\xa1\x39\x25\x40\xdb\x13\xe3\x73\x4b\xbd\x0b\xde\xa5\x12\x73\x25\x7b\xae\xa8\x54\xf6\xd9\x09\xae\x4c\x1f\x29\x71\xa5\xa5\xb0\x0a\x2b\x27\xca\x71\x99\x16\x5a\xb0\x5b\x59\xe8\x16\xbb\x55\xf5\x27\xcd\x68\xa8\xb3\x99\x62\xea\xaf\x4c\xac\x83\x06\xdc\x3c\x34\x42\x51\xff\x51\x70\x3e\x1f\x6d\x70\x6d\xa9\x7b\xb5\xfc\xea\x4d\x57\x99\x1f\x23\x32\x28\xd7\xe6\x82\xed\x05\x35\x81\x7a\x8a\x34\xbf\x09\x54\x7b\x8c\x41\x59\xff\xf7\x95\xc6\x1a\x3c\xd6\x60\x30\xb6\xc7\x0f\xad\xbf\x04\x79\xc6\x4d\x94\x1d\xce\x91\x0c\x15\xf3\x33\x35\x6c\xbc\x7c\xfd\x8a\x41\xd5\x49\x34\x36\xea\x99\x47\xdb\x41\x97\x1a\x2f\x1b\x73\x86\x67\x66\x84\x57\x74\x36\x65\xad\x0a\xd5\x4d\x4d\xee\x64\x97\x1a\x63\xea\xee\x0b\xcd\xc2\x3d\x9a\x0d\x57\x23\x99\x09\x6f\xa2\xef\xdc\x2b\x7a\x8b\x49\xb2\x98\x39\x69\xcc\x5f\x76\x87\x44\x35\x83\x5d\x1e\xc4\x1a\x57\x13\x68\xa1\xee\x0c\xeb\xc5\x39\x25\x17\x6c\x81\x79\xcd\xea\x30\x1e\x05\x39\xb7\x96\xe6\x9b\x46\x7a\xfd\x04\x8e\xf1\x4d\x46\x68\xe1\x96\x79\x45\xcf\xb2\xd6\xda\xbb\x7c\x36\x14\x2f\x1e\x43\x75\xfe\x28\xcb\x34\xcd\x86\x16\xac\x67\xe0\x2c\xf5\x50\xff\xce\xe8\x71\xca\xaf\xd8\x5f\x67\xa1\x4a\x23\x61\xd6\x6f\x60\xb5\x13\xef\xe2\x4d\x28\xb9\xfd\x1f\xb7\xa6\x2d\x59\x3a\xac\xb5\xcf\xf6\x99\x35\x2c\xc5\x4b\xf6\x06\x35\x6b\x52\xe3\x2f\xdf\x87\x7b\x3a\xb8\xb2\x22\x8d\xee\x35\xfe\xbb\x28\xe5\x98\x04\xeb\x4f\xfe\x13\x50\xeb\x6f\x63\xaa\x7c\x7d\xf1\xd8\x0e\x0c\x8f\x24\x6f\xf4\xbf\x42\x96\x2c\x67\xbc\x27\x48\x8a\x2d\x56\x55\x04\x55\x1e\xba\xf7\xb0\x76\x94\x0e\x5b\x10\x14\xdb\x7f\x40\xbd\x50\xc7\xb2\x6c\xf8\xf6\x52\x60\x10\x18\x25\xf7\xba\x3a\xc3\x26\xed\x1c\x30\xc9\x20\xa8\xab\xbe\x51\xb8\x34\x5a\x78\xd9\x10\x14\xab\xfb\xfd\x1a\xe3\x78\xe8\x4e\x82\xf7\x0a\xfb\x06\x3c\x4b\x9d\xb6\x62\x72\x14\x03\xcd\xd4\x3f\x29\xb0\x6e\x2e\x85\x37\x3c\x9d\xb9\x59\x30\x19\xee\xdf\xa3\xa9\xe4\xf6\xa7\x18\x2e\x56\x7d\x84\x4f\x27\xc7\xae\xa4\xe4\xd9\xcd\x83\x88\x97\x17\x2e\x04\xaf\x5e\x58\x15\xf6\xb2\x2c\x5b\x5d\xb2\x0b\x17\x9e\x3d\xb4\x3c\x55\x34\x93\xb3\x6c\x1f\x6f\x7e\xba\x41\xd6\xd4\xf1\xc8\xf5\x56\x49\x66\x15\x46\x8a\xe2\xa1\x87\xd9\x6f\x61\xb3\x94\x77\xad\x73\x5a\xca\x92\x96\x77\xac\xfa\x0d\xe4\x77\xe7\xcf\x9f\xff\x07\x00\x00\xff\xff\x15\x7b\xa6\xce\x8e\x0b\x00\x00")

func zundaTxtBytes() ([]byte, error) {
	return bindataRead(
		_zundaTxt,
		"zunda.txt",
	)
}

func zundaTxt() (*asset, error) {
	bytes, err := zundaTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "zunda.txt", size: 2958, mode: os.FileMode(436), modTime: time.Unix(1543993576, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"ariaki.txt": ariakiTxt,
	"zunda.txt":  zundaTxt,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"ariaki.txt": &bintree{ariakiTxt, map[string]*bintree{}},
	"zunda.txt":  &bintree{zundaTxt, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
