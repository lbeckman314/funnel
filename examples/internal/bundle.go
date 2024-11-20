// Code generated by go-bindata. DO NOT EDIT.
// sources:
// examples/capture-stdout-stderr.json (995B)
// examples/google-storage.json (462B)
// examples/resource-request.json (299B)
// examples/input-content.json (736B)
// examples/hello-world.json (187B)
// examples/safe-tag.json (233B)
// examples/log-streaming.json (304B)
// examples/s3-sleep.json (518B)
// examples/md5sum.json (737B)
// examples/nextflow.json (551B)
// examples/s3.json (524B)
// examples/malicious-tag.json (283B)
// examples/full-hello.json (577B)

package examples

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _examplesCaptureStdoutStderrJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x3f\x0f\xda\x30\x10\xc5\xf7\x7c\x8a\x27\xcf\x94\x4c\x5d\x98\xdb\x4a\x48\xdd\x18\x2b\x06\x2b\x3e\x88\xa5\xf8\x8f\xec\x73\x05\xa2\x7c\xf7\xea\x9c\xa4\x01\x54\x10\x2c\x11\xba\x7b\xf7\x78\xbf\x3b\x5f\x1a\x40\x79\xed\x48\x6d\xa0\x76\x6c\x42\x61\x68\x6f\xb0\x63\x43\x29\xa1\xd3\x91\x4b\x22\xb5\x12\x99\xa1\xdc\x25\x1b\xd9\x06\x2f\xea\x6f\xe4\x82\xcf\x9c\x34\x53\x9e\x84\xd6\x1f\x91\x17\x93\x3c\x9a\xc8\xcf\x12\x87\xa0\x8d\xf4\xb9\x27\x07\x9d\x11\x0a\xc7\xc2\x38\xd8\x81\xf2\x7a\xfc\x03\xeb\x63\xe1\xac\x36\xf8\xd5\x00\xc0\xa5\x7e\x6f\xf2\xd5\x7e\x95\xd6\xf2\x43\x9e\xad\x74\xc1\x01\x5d\x70\x4e\x7b\xb3\x46\xcb\x2e\xb6\x75\x08\xae\x64\x06\x9d\x6c\x66\x04\x2f\x19\xd0\x87\xcc\xc8\xe7\xcc\xe4\xd6\x8b\x67\x49\x83\x78\x49\xaa\x4d\xdb\x2e\x06\x8b\x82\xcf\xb1\x86\xf9\xb1\xfd\xf9\x7d\xa9\x46\xcd\xbd\x54\xa7\x09\x55\xeb\xd7\x06\xd8\x57\xb2\x11\xf6\x05\xda\xb8\xb5\xa7\x6c\xd3\x65\xc2\x61\x86\x83\x9d\x77\x4e\xa6\x2e\xb8\x0b\xd1\x92\x11\xfc\x1a\x61\xba\xc2\xa7\xac\x8f\x31\xde\x80\x15\xfd\x48\xbb\x7a\x8a\x46\x29\xbd\x42\x93\x47\xf2\x11\x5a\x1d\xf8\x1c\xed\x2e\xc6\x1b\x68\xa2\x7f\x38\x24\x9d\xa8\x2b\x1c\xd2\x7f\x4e\x69\x9d\x3e\x56\x47\x3d\x44\xeb\x69\xf1\x9c\xb8\x64\x42\xe5\x5e\xad\xa0\xbe\x74\xf2\x75\xe6\x6b\x2e\x6e\x7e\xa5\xf8\x03\x26\x42\x6b\xe8\xf7\x9c\x75\xff\xcf\x62\xba\xcb\xed\xce\x9f\xf5\xee\x43\x37\xd7\xe6\x6f\x00\x00\x00\xff\xff\x94\x21\xf6\x91\xe3\x03\x00\x00")

func examplesCaptureStdoutStderrJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesCaptureStdoutStderrJson,
		"examples/capture-stdout-stderr.json",
	)
}

func examplesCaptureStdoutStderrJson() (*asset, error) {
	bytes, err := examplesCaptureStdoutStderrJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/capture-stdout-stderr.json", size: 995, mode: os.FileMode(0640), modTime: time.Unix(1706586712, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc1, 0x51, 0xcf, 0xff, 0x9e, 0x30, 0x6e, 0xd6, 0x85, 0x28, 0xb5, 0x98, 0x44, 0xa1, 0x6b, 0x5c, 0x28, 0xb4, 0xba, 0x3e, 0x40, 0x82, 0x3, 0x58, 0x1b, 0x89, 0xca, 0xf7, 0xc0, 0x29, 0x58, 0x96}}
	return a, nil
}

var _examplesGoogleStorageJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xcd\x4a\x03\x31\x14\x85\xf7\xf3\x14\x87\xac\x14\x6a\x67\x5c\x54\xb0\x5b\x05\x37\x82\x62\xed\x4a\xba\xb8\x4d\x32\x63\x30\x7f\x4c\x6e\xb0\xb4\xf4\xdd\x25\x29\x75\xf0\x67\x33\x0c\xe4\x7c\xe7\xdc\xef\xd0\x00\xc2\x93\xd3\x62\x09\xf1\x10\xc2\x60\x35\xee\x6c\xc8\x0a\x2b\x0e\x23\x0d\x1a\x7a\x47\x2e\x5a\x2d\x66\x25\xa9\x74\x92\xa3\x89\x6c\x82\x2f\xc0\x2b\xa5\x0f\x18\x1f\x33\x27\x90\x57\x08\x99\xeb\xbf\x24\x8f\xad\xc6\xbf\x7d\xeb\x97\xc7\x34\x3f\xb5\xe9\x9d\x96\x99\xc3\x98\xc4\x12\x6f\x0d\x00\x1c\xea\x17\x10\xc6\xd1\x50\x6f\xca\xdb\xec\x39\xd7\x7c\x7d\x90\xc1\x39\xf2\xaa\x10\xc2\xa9\x45\xca\x4e\xcc\x20\x5a\x76\xb1\xed\x8d\xd5\xf3\xbd\x89\x62\x53\xc3\xc7\x06\xd8\xd4\x9d\xd3\x85\x7f\x47\xce\xde\x7b\x13\x0b\x3b\x8d\xfc\xd2\xbc\x0f\x9f\xde\x06\x52\x20\xc4\xbc\xb5\x46\xa2\xc4\xd1\x8f\xc1\x9d\x1d\xcf\x76\x17\xeb\x15\x9e\x89\xb5\x67\x3c\xf5\xbd\x91\x1a\x8a\x98\x2e\xa7\xea\x3c\xda\x52\x39\xa4\x65\xdb\xe6\x14\x39\x5c\x45\x32\x63\x4b\x31\x5a\x23\xa9\x2c\xa6\xb6\x5b\xdc\x76\x5d\x77\x7d\x53\x65\xbe\xc9\x48\xfc\x5e\xd0\x9f\xae\x93\x6a\x73\x6c\xbe\x02\x00\x00\xff\xff\x40\x00\x1c\x30\xce\x01\x00\x00")

func examplesGoogleStorageJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesGoogleStorageJson,
		"examples/google-storage.json",
	)
}

func examplesGoogleStorageJson() (*asset, error) {
	bytes, err := examplesGoogleStorageJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/google-storage.json", size: 462, mode: os.FileMode(0640), modTime: time.Unix(1706586712, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3b, 0x28, 0x52, 0xb5, 0x1b, 0x70, 0x7b, 0xa1, 0x77, 0xb8, 0x17, 0x96, 0xc6, 0x3, 0xd8, 0xa0, 0x46, 0x5d, 0xe7, 0x44, 0x62, 0x78, 0xce, 0xec, 0x46, 0x20, 0x25, 0x74, 0xf3, 0xb6, 0xaf, 0x42}}
	return a, nil
}

var _examplesResourceRequestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8f\x39\x6b\x03\x31\x10\x85\x7b\xfd\x8a\xc7\xd4\x22\xec\x82\xdd\x6c\x97\x03\x5c\x05\x82\x21\x55\x30\x61\xa2\x1d\x8c\xb0\x75\x44\xa3\x85\x40\xd8\xff\x6e\xa4\xb5\x4b\xbd\xe3\x7b\x9a\x7f\x03\x50\xe4\x20\x34\x81\x8e\xa2\x69\x29\x4e\x50\xe4\x77\x11\xad\x64\x9b\x3b\x8b\xba\xe2\x73\xf5\x29\xb6\xd0\x9b\x84\x14\xb5\x16\xae\xa2\x60\x54\xd6\xcb\x23\xef\xe3\x19\x3b\xbc\x7e\x7c\xc2\xa5\x22\x6a\x31\xee\x71\x78\xc1\xf1\xf9\xdd\x82\xe3\x8c\x71\x18\xda\x7b\xf6\x7a\x81\x66\x76\xf2\xb4\x2d\x94\xfb\xae\xd2\x84\xf6\x21\x80\x5c\x5e\xbe\x3b\x84\x26\xec\xec\xa6\x15\x0e\x87\x1f\x9a\x30\xee\xef\x42\x03\x6d\xca\x30\x18\x60\xed\x30\xf9\x13\xb7\xd4\x54\x5a\xf3\xab\xe7\x36\x24\x40\x3e\xf0\xb9\x1f\xca\xd7\xec\xa3\x90\x7d\x18\x2e\x85\xc0\x71\x6e\x0d\xd2\xab\x48\x26\x0b\x1a\xe9\xd4\xfd\xd5\x00\x27\xb3\x9a\x5b\x00\x00\x00\xff\xff\xc0\xca\xf1\x83\x2b\x01\x00\x00")

func examplesResourceRequestJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesResourceRequestJson,
		"examples/resource-request.json",
	)
}

func examplesResourceRequestJson() (*asset, error) {
	bytes, err := examplesResourceRequestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/resource-request.json", size: 299, mode: os.FileMode(0640), modTime: time.Unix(1706586712, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa5, 0x4a, 0x7f, 0x6d, 0xd4, 0x38, 0x22, 0xf1, 0x24, 0xb5, 0xc6, 0x61, 0xce, 0x9c, 0x40, 0xc9, 0x6, 0x44, 0xc8, 0x46, 0xd0, 0xd5, 0x88, 0x95, 0x76, 0x48, 0x8b, 0xa1, 0x33, 0x39, 0xbd, 0xc1}}
	return a, nil
}

var _examplesInputContentJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x51\x3f\x4f\xeb\x30\x10\xdf\xf3\x29\x7e\xca\xd2\xa5\x6a\xa6\xb7\x74\x7e\xef\x89\x4a\x6c\x0c\x0c\x50\x21\x13\x5f\x5a\x4b\xb6\x2f\xb2\x2f\x82\xaa\xea\x77\x47\xbe\x04\x52\xaa\xaa\xb0\x64\xf0\xfd\xfe\xe7\x58\x01\x75\x34\x81\xea\x35\xea\x4d\xec\x07\x41\xe7\x3c\xa1\xe5\x28\x14\x05\x26\x5a\xf0\x20\x9f\xef\xf5\xb2\xe0\x2d\xe5\x36\xb9\x5e\x1c\xc7\x42\xfb\x4b\x81\x63\x96\x64\x84\x32\x86\xec\xe2\x0e\xb2\x27\x2c\x26\x8d\x05\x3a\x47\xde\xa2\xe3\x04\x57\x1c\x32\x84\xd1\x26\x32\x42\x30\xa3\x1d\x47\xa5\xec\x39\x0b\xf2\x21\x0b\x85\xd1\x69\xc4\xd7\x6b\x3c\x55\x00\x70\xd4\xef\x59\xe2\xd6\xc8\xa8\xa9\x70\x3d\x5d\x84\x1b\x3b\x09\x23\xd8\x3f\x79\x08\x2b\x34\x12\xfa\xc6\x45\xbc\x39\xef\xf1\x4a\x53\x10\x7b\x25\xc2\x6a\x16\x95\x43\xaf\x7e\xff\x37\xf7\xff\xe6\xd7\xde\xc8\xbe\xbc\x4e\x92\xf3\x61\x6a\x5e\x6e\x77\xe4\x3d\xe3\x91\x93\xb7\xcf\xb1\x56\xc0\xa9\x02\xb6\x5a\x6f\x5c\xf6\x87\x7e\x59\x2c\xdf\x28\xf8\xa0\x67\x70\x07\x1d\x23\xa3\x35\xbd\x0c\x89\x74\x65\x4d\x26\x94\xe5\x45\x31\xb7\x2a\x0e\xc9\x17\xb9\xf2\x3b\xd6\x4d\xa3\xc4\xd6\x28\xef\xdb\xbc\xbf\x58\xa2\xa4\xbd\x28\x4a\xef\xd4\x0e\xc2\xe9\x4a\x55\x17\xcc\x4e\x15\x8d\xef\x5d\xa4\xf3\x11\x43\x30\xd1\x16\x46\xd9\xa1\x5e\xce\x43\x6f\xbf\x40\xd3\x38\xd7\xad\xab\x53\xf5\x11\x00\x00\xff\xff\x9b\xe3\x40\x40\xe0\x02\x00\x00")

func examplesInputContentJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesInputContentJson,
		"examples/input-content.json",
	)
}

func examplesInputContentJson() (*asset, error) {
	bytes, err := examplesInputContentJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/input-content.json", size: 736, mode: os.FileMode(0640), modTime: time.Unix(1706586712, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x49, 0x70, 0x12, 0xa9, 0x99, 0xea, 0x42, 0x7d, 0xd2, 0x79, 0x3c, 0xcc, 0xad, 0xf8, 0x22, 0xd7, 0xb8, 0xd8, 0xbb, 0xb6, 0x4, 0xbb, 0xa4, 0xf0, 0x5a, 0xa6, 0x1a, 0xfd, 0xa9, 0x28, 0xba, 0x20}}
	return a, nil
}

var _examplesHelloWorldJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x31\xce\xc2\x30\x0c\x46\xf7\x9c\xe2\x93\xe7\xea\x3f\x40\xe7\x7f\xe0\x0e\xa8\x83\x49\x2d\x1a\x91\xc4\x55\x6c\x04\x52\xd5\xbb\xa3\x84\x85\xc5\xcb\x7b\xcf\xf6\x11\x00\xaa\x5c\x84\x66\xd0\x45\x72\x56\xbc\xb4\xe5\x95\xa6\x0e\x56\xb1\xd8\xd2\xee\x49\x6b\xe7\xff\x52\xb4\x9a\x37\x76\x31\xf8\x26\x28\x6a\x8e\x1b\x5b\x8a\x90\xb8\x29\x9c\xed\xf1\xf7\x4d\xe5\x2d\xf1\xe9\xda\x8c\x66\x5c\x03\x00\x1c\x63\x02\x94\x0a\xdf\xc7\x3d\xce\x7b\xaa\x32\xfc\x01\xa2\x96\xc2\x75\xed\x05\xf5\x7d\x34\x81\xb6\x9f\x9f\x96\x21\x9e\x01\x58\xc2\x19\x3e\x01\x00\x00\xff\xff\x1b\x6b\x00\xfc\xbb\x00\x00\x00")

func examplesHelloWorldJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesHelloWorldJson,
		"examples/hello-world.json",
	)
}

func examplesHelloWorldJson() (*asset, error) {
	bytes, err := examplesHelloWorldJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/hello-world.json", size: 187, mode: os.FileMode(0640), modTime: time.Unix(1706586712, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5f, 0x7f, 0x2d, 0x92, 0x82, 0x32, 0xc2, 0x3a, 0xc8, 0x1d, 0x60, 0x4d, 0xdc, 0x56, 0xfc, 0xf8, 0x75, 0x64, 0x9d, 0xa3, 0x9e, 0x45, 0x16, 0x97, 0xca, 0xc4, 0x1b, 0x81, 0xc3, 0x8c, 0xb6, 0x7}}
	return a, nil
}

var _examplesSafeTagJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8d\x41\x8b\x83\x40\x0c\x46\xef\xf3\x2b\x3e\x72\x76\xf1\xee\x7d\x61\xef\xdb\x5b\xf1\x90\x6a\x1c\x0b\x8e\x23\x9a\x80\xa5\xf8\xdf\x4b\xa6\xa5\xed\x2d\xe1\xbd\x97\xdc\x03\x40\x33\x27\xa1\x06\xf4\xcf\x83\xe0\xc4\x11\xbf\x3b\xa7\x65\x12\xaa\x9c\x2a\xc7\x8d\x1a\xb8\x09\xd0\xc6\x83\xfc\x28\x47\xf7\x7d\xa6\x00\x1c\xc5\x93\x5d\x3a\xd3\xbc\xba\x7c\x2e\xf2\x33\x01\xe8\x9a\x38\x96\x0f\x76\xb1\x59\xad\xdc\x2d\xa0\xcb\x29\xf1\xdc\x7b\x41\xd2\x8d\x99\x2a\xd0\x9f\xf0\xa4\xe3\x8d\xda\xb7\xb5\x69\x9f\x4d\xbd\xaf\x7b\x56\xae\xb3\xe9\x62\x4a\xdf\x5c\xd6\xf5\xc3\x5f\x7b\xc1\x47\x00\xda\x70\x84\x47\x00\x00\x00\xff\xff\x12\x1d\xbf\x86\xe9\x00\x00\x00")

func examplesSafeTagJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesSafeTagJson,
		"examples/safe-tag.json",
	)
}

func examplesSafeTagJson() (*asset, error) {
	bytes, err := examplesSafeTagJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/safe-tag.json", size: 233, mode: os.FileMode(0644), modTime: time.Unix(1727826126, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb, 0x31, 0x20, 0x2e, 0x11, 0x9a, 0x4a, 0xd1, 0x11, 0xca, 0x83, 0x8e, 0x9a, 0x8, 0xf3, 0x28, 0xa1, 0xd3, 0xeb, 0x6b, 0x12, 0x9b, 0x92, 0xd, 0x72, 0xce, 0xb0, 0x8a, 0x26, 0x65, 0x9e, 0xf}}
	return a, nil
}

var _examplesLogStreamingJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8f\xb1\x6e\xf3\x30\x10\x83\x77\x3f\x05\xa1\x39\x7f\x82\x7f\x4d\xd6\xa2\x53\x87\xee\x45\x86\x8b\x44\xd8\x02\xa4\x93\xa1\x3b\xb7\x0d\x82\xbc\x7b\x61\xb5\x0b\x87\xfb\x78\x04\xf9\x98\x80\xa0\x52\x19\xce\x08\x6f\x6d\x86\x79\xa7\xd4\xac\x73\x38\xec\x28\xd1\x62\xcf\xab\xe7\xa6\xbb\xe3\x85\xb5\xa9\x79\x17\xa7\xc1\x17\x71\x98\xa7\xb6\xf9\x89\xbd\xa3\xb4\xd9\x20\x9d\x7f\x19\x4c\xb8\xdd\xe1\x0d\xbe\x10\xaf\x9b\x2a\x0b\x92\xd8\x72\x6b\xd2\x93\x1d\xf1\xde\xb3\xba\x0d\x9a\xc4\x79\xf2\x5c\x09\x7e\xb2\xdf\x61\x8c\x4d\xd3\x01\x9b\x7a\x2e\x88\xa2\x91\x85\xe9\xf8\x5b\x89\xdf\x8c\x9b\xb7\x6e\xe1\x8c\x8f\x09\x00\x1e\x43\x81\x90\xab\xcc\x63\x89\x94\x35\x2b\x87\x7f\x80\xd8\x6a\x15\x4d\xfb\x47\xb0\x25\x1c\x10\xfe\xc5\x5d\xbf\x96\x5c\x08\xef\x1b\x2f\x48\x6d\xf4\xb8\xc0\x0a\xb9\xe2\xff\x7e\x51\x86\xeb\xc8\x78\x4e\xc0\x75\x7a\x4e\x3f\x01\x00\x00\xff\xff\x88\xae\x93\xb2\x30\x01\x00\x00")

func examplesLogStreamingJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesLogStreamingJson,
		"examples/log-streaming.json",
	)
}

func examplesLogStreamingJson() (*asset, error) {
	bytes, err := examplesLogStreamingJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/log-streaming.json", size: 304, mode: os.FileMode(0640), modTime: time.Unix(1706586712, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8d, 0xe2, 0xc2, 0x8b, 0x68, 0xe0, 0xaf, 0x64, 0x8c, 0x5, 0xf0, 0xbf, 0x9a, 0xef, 0x1c, 0x4c, 0x87, 0xa1, 0xba, 0xf1, 0xbe, 0xae, 0xb8, 0xfc, 0x65, 0x86, 0xc9, 0xfe, 0x58, 0xe0, 0x46, 0xea}}
	return a, nil
}

var _examplesS3SleepJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\x4d\x8f\xd3\x30\x10\x86\xef\xf9\x15\x23\x9f\x40\xda\xc4\x2b\x2a\x71\xc8\x35\xad\xb8\x54\x1c\x58\xb8\x80\xf6\x30\x71\xa6\xa9\xc1\x5f\xf2\x8c\x69\x50\xd5\xff\x8e\xe2\x02\x01\xf6\x62\x59\x7a\xdf\x79\xf4\xcc\x5c\x1b\x00\x15\xd0\x93\xea\x41\x3d\xed\xe0\x49\x62\xc6\x99\x80\x16\xf4\xc9\x91\x7a\x58\xf3\x89\xd8\x64\x9b\xc4\xc6\xb0\xd6\x3e\x22\x7f\x03\x1b\x52\x11\x06\x0c\x13\xc4\x22\xf5\x6f\x30\xc0\x48\x30\xb8\x58\xa6\x3f\xa0\x4f\x1f\x8e\xdc\xdd\x31\xb4\x90\x29\x12\x33\xab\x1e\xbe\x34\x00\x00\xd7\xfa\x02\x28\xeb\x71\xae\x0a\x65\x2c\x41\x4a\xed\xd7\xc0\x44\xef\x31\x4c\xeb\x84\x62\x47\x94\xd4\x03\xa8\xdd\xdb\xc7\x47\xf5\x5c\x2b\xb7\x06\xe0\xb9\xd2\xef\x42\x2f\xd1\xbf\x97\xab\xf9\x06\xfe\x6f\xa7\x7d\xbc\x04\x17\x71\x02\x84\x54\x46\x67\x0d\x9c\xac\x23\x38\xe5\xe8\xe1\xaf\xb3\xbc\x1a\x3e\xc3\x70\x38\x1e\x97\x77\x87\xf7\x07\xd8\x5b\x36\xf1\x3b\x65\x18\x28\x70\x61\xd8\xa3\x60\x0f\x67\x91\xc4\xbd\xd6\x99\x66\xcb\x92\x7f\x74\x31\x51\x98\x50\xb0\xc3\x0b\x6b\xc1\x59\x8f\x36\xda\x70\x8a\xd9\xa3\x58\xc3\xfa\xf5\x66\x55\xb2\x5b\x6d\x78\xd7\x6b\x6d\xc8\xb9\x65\xa6\x40\xad\xa9\xf8\xf6\x2e\xd6\x16\x6e\x2f\xc4\xd2\xbe\xa9\x8d\x5f\xa1\xce\xe4\x08\x99\xba\xaf\x1c\xc3\xc6\x4b\x28\xe7\x15\xa8\xc5\xa7\x7f\x2b\xdb\xf5\x9a\x5b\xf3\x33\x00\x00\xff\xff\xfc\xad\xcd\x36\x06\x02\x00\x00")

func examplesS3SleepJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesS3SleepJson,
		"examples/s3-sleep.json",
	)
}

func examplesS3SleepJson() (*asset, error) {
	bytes, err := examplesS3SleepJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/s3-sleep.json", size: 518, mode: os.FileMode(0644), modTime: time.Unix(1731526968, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x95, 0x15, 0xc0, 0xd, 0xc2, 0x5, 0x9b, 0x35, 0xda, 0xb5, 0x99, 0xc3, 0x58, 0x1e, 0x46, 0xa3, 0x86, 0x77, 0x20, 0x1c, 0xe5, 0x1f, 0xa8, 0x5f, 0xa9, 0x31, 0x14, 0x7d, 0xfe, 0xe4, 0x3, 0x46}}
	return a, nil
}

var _examplesMd5sumJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xcf\x6a\xf3\x30\x10\xc4\xef\x7e\x8a\x41\xe7\x10\x9f\xbe\x4b\xce\x5f\x0b\x81\xde\x7a\x2c\x21\x08\x6b\x93\x08\xac\x3f\x78\x57\xe0\x10\xf2\xee\x45\x6b\xa7\x86\x24\xb4\xb4\x17\x63\x6b\x77\x66\xf7\x37\xf2\xa5\x01\x4c\xb4\x81\xcc\x06\x26\xb8\x7f\x5c\x02\x68\xb4\x21\xf7\x64\x56\xb5\xe6\x88\xbb\xc1\x67\xf1\x29\xd6\x96\xff\x14\x52\x64\x19\xac\x10\xc3\xc7\x5c\x04\x36\x3a\xa4\x22\xf5\xf5\xe0\x7b\x62\x14\xf6\xf1\x08\x0b\xf6\xd5\x06\xb3\x6b\x97\x42\xb0\xd1\xad\x27\x5b\x95\xb2\xd9\xe0\xa3\x01\x80\x8b\x3e\x1f\x57\xd1\x36\x55\x68\xf5\x6e\x99\xad\xce\x97\x34\x8f\x58\xa3\x95\x90\xdb\xe9\x63\x3f\x2d\x17\x0a\x0b\x68\xf4\x2c\x48\x11\x72\x22\x9c\x12\x0b\xf8\xcc\x42\x61\xbd\x38\x97\xa1\xaf\x8e\x15\x60\xd3\xb6\x0f\x3e\x4b\xa3\x9c\xb3\x2e\xf8\xba\x7d\x7b\x59\x4e\xb3\x95\x53\x3d\x55\xa1\x8f\x46\xcf\xaf\x0d\xb0\x53\xda\x29\x9e\x9f\x71\x59\x5c\xfa\x86\xf7\x5d\xcb\x48\x87\x5b\xa6\x9e\xd1\xd9\x2c\x65\x20\xa7\xd7\xd0\xa5\xec\xc9\xd5\x44\x6e\x04\xfb\xf9\x66\xfe\x08\x3f\xa9\x7f\x45\x5f\x01\xee\xf0\x69\xa4\xae\x48\x1a\x9e\x04\xe0\x83\x3d\xaa\xa3\xed\xb3\x8f\xb4\x78\xce\xbf\x4b\x55\xcc\xe9\x98\xd5\x92\xef\xee\xab\x6f\x8e\xec\xf9\xf4\xe6\xda\x7c\x06\x00\x00\xff\xff\xb9\xf5\xf2\xe7\xe1\x02\x00\x00")

func examplesMd5sumJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesMd5sumJson,
		"examples/md5sum.json",
	)
}

func examplesMd5sumJson() (*asset, error) {
	bytes, err := examplesMd5sumJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/md5sum.json", size: 737, mode: os.FileMode(0644), modTime: time.Unix(1722035929, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x27, 0x60, 0x3a, 0x1, 0xdf, 0xa8, 0x51, 0x9e, 0xcd, 0xc7, 0x97, 0xce, 0x6d, 0xcf, 0x8a, 0x93, 0x72, 0xb2, 0xaf, 0x16, 0xb8, 0xdd, 0x25, 0xfd, 0xc2, 0x5, 0x98, 0xa, 0x14, 0x13, 0x2, 0xc7}}
	return a, nil
}

var _examplesNextflowJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\x31\x6f\xe3\x30\x0c\x85\x77\xff\x0a\x82\x73\x6c\xef\x39\xdc\x94\xdc\x70\xc3\x5d\x83\x20\x4b\x51\x64\x50\x65\x39\x56\x62\x91\xaa\x44\xa5\x31\x8a\xfc\xf7\x42\x4a\x0b\x17\x45\x87\x76\xd1\x40\xbe\xf7\xf1\x3d\xbd\x54\x00\x48\xca\x19\x5c\x02\xfe\x37\x17\xe9\x47\x7e\x86\x95\x22\x15\x26\xd8\x99\x28\x11\x17\x59\xd2\x99\xa8\x83\xf5\x62\x99\xb2\x72\x9b\x08\x64\x30\xf0\xd9\x21\xd9\x01\xc2\x70\x36\xc1\xf6\xd3\xbc\x8f\xc9\x7b\x0e\x72\x83\x89\x3a\x44\x5c\x42\xbe\x0d\x80\x9b\xc0\x47\xa3\xe5\x63\x00\xac\x00\xae\x45\xca\x49\x7c\x92\xac\x7e\x28\xea\x9b\x07\x00\x53\x18\xb3\xa3\x15\xe7\x0b\xb4\x0c\xbd\x92\xa1\x4c\x1b\x7a\x23\xb5\xf3\x52\x26\x5f\x5a\xae\xff\x6e\xff\xac\x76\x77\xdb\x7b\x2c\x9b\xeb\xe2\x27\xe0\xc0\x2c\xdf\x60\x56\x00\xfb\x92\xdf\x5c\x8c\x4e\xc2\xe1\x8b\x06\xd6\xa9\x43\x31\x3f\x25\x35\x35\x96\x5b\x1e\x62\xaa\x35\x3b\x5f\x3f\x5a\x6e\xa9\xaf\x75\xf9\xd4\xf9\x9a\x66\xe7\x14\x75\x19\x85\x71\xc0\x05\x60\xad\xf3\xab\x3b\x28\xc1\x7e\xc1\x7b\x6d\xa8\xd7\xc7\xee\xd4\x8c\x8a\x0e\xcd\x26\xb0\x36\x31\x36\xa3\x4a\xa4\x87\x7f\x46\x0f\x8a\x6c\x74\xbf\xcf\x3d\x87\x13\x84\x44\x30\xdf\x6a\x9d\xb2\xd4\x50\x8f\xfb\xb9\x47\x75\xad\xaa\xd7\x00\x00\x00\xff\xff\xdf\x3a\x6a\xcf\x27\x02\x00\x00")

func examplesNextflowJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesNextflowJson,
		"examples/nextflow.json",
	)
}

func examplesNextflowJson() (*asset, error) {
	bytes, err := examplesNextflowJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/nextflow.json", size: 551, mode: os.FileMode(0644), modTime: time.Unix(1728514108, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5, 0xe2, 0xb9, 0x78, 0xd2, 0xd3, 0x17, 0x6e, 0xe7, 0xe5, 0x4a, 0xf, 0xd3, 0x78, 0x15, 0x87, 0x85, 0xe6, 0x68, 0x53, 0xbd, 0x7e, 0xb9, 0x1b, 0x2f, 0xa8, 0xd9, 0xb3, 0x73, 0xd0, 0x9f, 0x49}}
	return a, nil
}

var _examplesS3Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xcf\x4e\xf3\x30\x10\xc4\xef\x79\x8a\x95\x4f\xdf\x27\x35\xb1\x44\xc5\x25\xd7\xb4\xe2\x52\x71\xa0\x70\x01\xf5\xb0\x71\xb6\xa9\xc1\x7f\x22\xaf\x4d\x83\xaa\xbe\x3b\x8a\x0b\x2d\x85\x4b\x14\x69\xc6\x33\xbf\x9d\x43\x01\x20\x1c\x5a\x12\x35\x88\xf5\x1c\xd6\xd1\x07\xec\x09\x68\x44\x3b\x18\x12\xb3\x49\xef\x88\x55\xd0\x43\xd4\xde\x4d\xb6\x47\xe4\x37\xd0\x6e\x48\x91\x01\x5d\x07\x3e\xc5\xfc\xaf\xd0\x41\x4b\xd0\x18\x9f\xba\x73\xd0\xd3\xc3\x8a\xab\x53\x0c\x8d\xa4\x52\xf4\x81\x45\x0d\x2f\x05\x00\xc0\x21\x7f\x01\x84\xb6\xd8\x67\x84\xd4\x26\x17\x53\xf6\x67\x41\x79\x6b\xd1\x75\xd3\x0b\x61\xbb\x5b\x4e\x56\xcc\x40\xc8\x40\x86\x90\xa9\x7a\x65\xef\xc4\x26\x9b\x8f\x05\xc0\x26\xf7\x9c\xd0\xfe\x96\x7c\x9f\x99\xf5\x4b\xc5\xaf\xeb\x16\x7e\xef\x8c\xc7\x0e\x10\x86\xd4\x1a\xad\x60\xab\x0d\xc1\x36\x78\x0b\x3f\x06\xfa\xd7\x3c\x43\xb3\x5c\xad\xc6\xbb\xe5\xfd\x12\x16\x9a\x95\x7f\xa7\x00\x0d\x39\x4e\x0c\x0b\x8c\x58\xc3\x2e\xc6\x81\x6b\x29\x03\xf5\x9a\x63\xf8\xa8\xfc\x40\xae\xc3\x88\x15\xee\x59\x46\xec\x65\xab\xbd\x76\x5b\x1f\x2c\x46\xad\x58\xfe\xbf\x50\xa5\x60\x26\x1a\x9e\xd7\x52\x2a\x32\x66\xec\xc9\x51\xa9\x72\x7c\x79\x02\x2b\x13\x97\x7b\xe2\x58\xde\x64\xc7\x97\x78\xbd\xcd\x39\x6f\xc0\xb8\x9b\x02\xaf\xe5\xcb\x72\xc5\xb1\xf8\x0c\x00\x00\xff\xff\xc0\x06\x69\xef\x0c\x02\x00\x00")

func examplesS3JsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesS3Json,
		"examples/s3.json",
	)
}

func examplesS3Json() (*asset, error) {
	bytes, err := examplesS3JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/s3.json", size: 524, mode: os.FileMode(0644), modTime: time.Unix(1732064231, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa5, 0x12, 0xab, 0xa6, 0x34, 0xd0, 0xb, 0xd6, 0x56, 0x79, 0x95, 0xa, 0x80, 0x42, 0x6e, 0x70, 0x53, 0xb8, 0xcb, 0x92, 0x4, 0xb6, 0x62, 0xe0, 0xb, 0x1c, 0xfb, 0xb1, 0xd5, 0x78, 0xa4, 0x37}}
	return a, nil
}

var _examplesMaliciousTagJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x31\x4f\xc3\x30\x10\x46\x77\xff\x8a\x8f\x5b\xba\x14\x75\x2f\x33\x12\x0b\x1b\x1b\xea\x70\x38\x47\x12\x61\xe7\xa2\xe4\x4e\x0a\x42\xf9\xef\xe8\x4c\x81\x6e\xb6\xbe\xf7\x9e\xfd\x95\x00\x9a\xb8\x0a\x9d\x41\xcf\x5c\xc6\x3c\xaa\xaf\x78\xe1\x1e\x8f\x1b\xd7\xb9\x08\x1d\x03\x31\xee\x57\x3a\x23\x70\x80\x64\x9b\x8b\x8e\x76\x6f\xdc\x87\x37\xe9\x52\xb9\x3c\x40\xf2\xa0\x38\x0c\x9c\x3f\xa4\xbb\x3b\x34\x11\xa0\x95\xdf\xe5\x97\x8c\x33\x25\x60\x6f\x51\xd9\x24\xbb\xe9\x12\xe5\xd7\x06\xff\xf4\x01\x1a\x2b\xf7\xed\x4f\xfe\xe6\x93\xf9\xb5\x05\x50\xd6\x5a\x79\xea\xc2\xa0\x78\x8f\x8e\xa0\x27\xe1\x62\xc3\x27\x5d\xfe\xa8\xd5\x3a\x75\x0b\xff\xd4\xb1\xf1\x49\xdd\x66\x37\xba\xdd\x65\x59\xfe\xf7\xeb\xbd\xcd\x7b\x02\x2e\x69\x4f\xdf\x01\x00\x00\xff\xff\x13\x80\xdc\x46\x1b\x01\x00\x00")

func examplesMaliciousTagJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesMaliciousTagJson,
		"examples/malicious-tag.json",
	)
}

func examplesMaliciousTagJson() (*asset, error) {
	bytes, err := examplesMaliciousTagJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/malicious-tag.json", size: 283, mode: os.FileMode(0644), modTime: time.Unix(1727826052, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xce, 0xa3, 0x34, 0x4d, 0x66, 0x52, 0xc5, 0x96, 0x67, 0xd7, 0xea, 0xff, 0x5, 0x1c, 0x99, 0x98, 0xcb, 0xc3, 0xc5, 0xb5, 0xa8, 0x52, 0x91, 0x18, 0x5b, 0x89, 0xfb, 0xc8, 0xa5, 0x29, 0x67, 0x51}}
	return a, nil
}

var _examplesFullHelloJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\x4f\x6b\xdc\x30\x10\xc5\xef\xfe\x14\x83\xce\x71\x44\x7b\x4c\xa1\x97\x66\x43\x7a\xa9\xc1\x6c\x28\xa5\x84\x32\xd6\x8e\x57\xa2\x92\xc6\xe8\xcf\x66\x4b\xd9\xef\x5e\x34\x5e\x97\xc0\xfa\x20\xec\xa7\xe7\xdf\x3c\xe9\xfd\xed\x00\x54\xc4\x40\xea\x01\xd4\x53\xf5\x1e\x9e\xc9\x7b\x86\x37\x4e\xfe\xa0\xee\xda\xee\x81\xb2\x49\x6e\x29\x8e\x63\x33\x3d\x52\xe0\x98\x4b\xc2\x42\x19\x8a\x25\x08\x9c\x0b\x4c\x98\x9d\x01\x32\x96\xa1\x60\xfe\x0d\x6f\xae\x58\x40\xef\xc5\x41\xe7\x92\x30\xdf\xaf\xb8\x82\xc7\xac\x1e\xa0\x0d\x5e\xbf\xfa\xe1\xdb\xae\x81\xf7\xdf\x07\x71\x5c\xd5\xfd\xf3\xb8\x13\xfd\x69\x78\x19\x55\x07\x70\x91\xdf\xb9\x96\xa5\x96\x46\xf8\x79\x45\xd4\xe4\x9b\x4d\xbf\x64\x4a\x59\x4f\xd5\x58\x8c\x48\x3a\x27\xa3\x8f\xae\xd8\x3a\xdd\x1b\x0e\x9a\x6d\xae\xbd\xe1\xb0\xf4\x93\x63\x3d\xd7\x18\xc9\xeb\x5f\xd9\x24\x2c\xc6\xea\xb9\x7a\xdf\xaf\x68\xbd\x85\x58\xb0\x58\x01\x6f\x23\xb7\x70\x7f\x16\xb9\xad\xc7\xaf\xe3\xee\xcb\x7e\x18\x7f\x48\xb8\x57\x49\x77\x62\x5f\x03\x49\x3a\xa5\x4f\xec\x3f\xa8\x3b\x90\x97\x8f\x6a\x35\xd0\x99\x4c\x2d\x9c\xc4\x22\xbc\xf5\x14\x00\xca\x05\x3c\x0a\x18\xfd\xe2\x22\x5d\xc7\x01\x28\xc3\x21\x60\x3c\x08\x34\xdb\x46\xec\x4d\x5b\xe5\xb6\x67\x66\xf8\x0c\x5b\x46\x3d\x33\x7f\x5a\x6b\x98\x30\xbd\xdf\x98\x30\xad\x11\x04\x49\xf1\xf4\xbf\x03\x11\x6e\x3a\x10\xf5\xa6\x83\xf6\x5c\xba\x6d\x7d\xed\x2e\xdd\xbf\x00\x00\x00\xff\xff\x12\xe1\x02\xb9\x41\x02\x00\x00")

func examplesFullHelloJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesFullHelloJson,
		"examples/full-hello.json",
	)
}

func examplesFullHelloJson() (*asset, error) {
	bytes, err := examplesFullHelloJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/full-hello.json", size: 577, mode: os.FileMode(0640), modTime: time.Unix(1706586712, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa8, 0x44, 0x5c, 0x19, 0x7f, 0x64, 0x5f, 0xb2, 0x73, 0x48, 0x9f, 0x84, 0xc6, 0x2f, 0x14, 0x90, 0x37, 0xa0, 0x36, 0x7e, 0x8e, 0xff, 0x99, 0x31, 0xe7, 0xc4, 0xb9, 0xe, 0x7e, 0xdb, 0x3, 0x16}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"examples/capture-stdout-stderr.json": examplesCaptureStdoutStderrJson,
	"examples/google-storage.json":        examplesGoogleStorageJson,
	"examples/resource-request.json":      examplesResourceRequestJson,
	"examples/input-content.json":         examplesInputContentJson,
	"examples/hello-world.json":           examplesHelloWorldJson,
	"examples/safe-tag.json":              examplesSafeTagJson,
	"examples/log-streaming.json":         examplesLogStreamingJson,
	"examples/s3-sleep.json":              examplesS3SleepJson,
	"examples/md5sum.json":                examplesMd5sumJson,
	"examples/nextflow.json":              examplesNextflowJson,
	"examples/s3.json":                    examplesS3Json,
	"examples/malicious-tag.json":         examplesMaliciousTagJson,
	"examples/full-hello.json":            examplesFullHelloJson,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"examples": {nil, map[string]*bintree{
		"capture-stdout-stderr.json": {examplesCaptureStdoutStderrJson, map[string]*bintree{}},
		"full-hello.json":            {examplesFullHelloJson, map[string]*bintree{}},
		"google-storage.json":        {examplesGoogleStorageJson, map[string]*bintree{}},
		"hello-world.json":           {examplesHelloWorldJson, map[string]*bintree{}},
		"input-content.json":         {examplesInputContentJson, map[string]*bintree{}},
		"log-streaming.json":         {examplesLogStreamingJson, map[string]*bintree{}},
		"malicious-tag.json":         {examplesMaliciousTagJson, map[string]*bintree{}},
		"md5sum.json":                {examplesMd5sumJson, map[string]*bintree{}},
		"nextflow.json":              {examplesNextflowJson, map[string]*bintree{}},
		"resource-request.json":      {examplesResourceRequestJson, map[string]*bintree{}},
		"s3-sleep.json":              {examplesS3SleepJson, map[string]*bintree{}},
		"s3.json":                    {examplesS3Json, map[string]*bintree{}},
		"safe-tag.json":              {examplesSafeTagJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
