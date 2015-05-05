// Generated by vfsgen; do not edit.

package vfsgen_test

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type assetsFS map[string]interface{}

var AssetsFS = func() http.FileSystem {
	assetsFS := assetsFS{
		"/": &dir{
			name:    "/",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/folderA": &dir{
			name:    "folderA",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/folderA/file1.txt": &compressedFile{
			name:              "file1.txt",
			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0a\x2e\x29\x4d\x4b\xd3\x03\x04\x00\x00\xff\xff\x27\xbb\x40\xc8\x06\x00\x00\x00"),
			uncompressedSize:  6,
			modTime:           mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/folderA/file2.txt": &compressedFile{
			name:              "file2.txt",
			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0a\x2e\x29\x4d\x4b\xd3\x03\x04\x00\x00\xff\xff\x27\xbb\x40\xc8\x06\x00\x00\x00"),
			uncompressedSize:  6,
			modTime:           mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/folderB": &dir{
			name:    "folderB",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/folderB/folderC": &dir{
			name:    "folderC",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/folderB/folderC/file3.txt": &compressedFile{
			name:              "file3.txt",
			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0a\x2e\x29\x4d\x4b\xd3\x03\x04\x00\x00\xff\xff\x27\xbb\x40\xc8\x06\x00\x00\x00"),
			uncompressedSize:  6,
			modTime:           mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/not-worth-compressing-file.txt": &compressedFile{
			name:              "not-worth-compressing-file.txt",
			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xf2\x2c\x29\x56\xc8\xcb\x2f\xca\x4d\xcc\x51\x48\xce\xcf\x2b\x49\xcd\x03\xf2\x13\x8b\x52\x15\x32\x52\x8b\x52\xf5\x00\x01\x00\x00\xff\xff\xdc\xc7\xff\x13\x1d\x00\x00\x00"),
			uncompressedSize:  29,
			modTime:           mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/sample-file.txt": &compressedFile{
			name:              "sample-file.txt",
			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0a\xc9\xc8\x2c\x56\x48\xcb\xcc\x49\x55\x48\xce\xcf\x2d\x28\x4a\x2d\x2e\x4e\x2d\x56\x28\x4f\xcd\xc9\xd1\x53\x70\xca\x49\x1c\xd4\x20\x43\x11\x10\x00\x00\xff\xff\xe7\x47\x81\x3a\xbd\x00\x00\x00"),
			uncompressedSize:  189,
			modTime:           mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
	}

	assetsFS["/"].(*dir).entries = []os.FileInfo{
		assetsFS["/folderA"].(os.FileInfo),
		assetsFS["/folderB"].(os.FileInfo),
		assetsFS["/not-worth-compressing-file.txt"].(os.FileInfo),
		assetsFS["/sample-file.txt"].(os.FileInfo),
	}
	assetsFS["/folderA"].(*dir).entries = []os.FileInfo{
		assetsFS["/folderA/file1.txt"].(os.FileInfo),
		assetsFS["/folderA/file2.txt"].(os.FileInfo),
	}
	assetsFS["/folderB"].(*dir).entries = []os.FileInfo{
		assetsFS["/folderB/folderC"].(os.FileInfo),
	}
	assetsFS["/folderB/folderC"].(*dir).entries = []os.FileInfo{
		assetsFS["/folderB/folderC/file3.txt"].(os.FileInfo),
	}

	return assetsFS
}()

func (fs assetsFS) Open(path string) (http.File, error) {
	f, ok := fs[path]
	if !ok {
		return nil, os.ErrNotExist
	}

	switch f := f.(type) {
	case *compressedFile:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &compressedFileInstance{
			compressedFile: f,
			gr:             gr,
		}, nil
	case *dir:
		return &dirInstance{
			dir: f,
		}, nil
	default:
		return f.(http.File), nil
	}
}

func mustUnmarshalTextTime(text string) time.Time {
	var t time.Time
	err := t.UnmarshalText([]byte(text))
	if err != nil {
		panic(err)
	}
	return t
}

// compressedFile is ...
type compressedFile struct {
	name              string
	compressedContent []byte
	uncompressedSize  int64
	modTime           time.Time
}

func (f *compressedFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *compressedFile) Stat() (os.FileInfo, error) { return f, nil }

func (f *compressedFile) GzipBytes() []byte {
	log.Println("using GzipBytes for", f.name)
	return f.compressedContent
}

func (f *compressedFile) Name() string       { return f.name }
func (f *compressedFile) Size() int64        { return f.uncompressedSize }
func (f *compressedFile) Mode() os.FileMode  { return 0444 }
func (f *compressedFile) ModTime() time.Time { return f.modTime }
func (f *compressedFile) IsDir() bool        { return false }
func (f *compressedFile) Sys() interface{}   { return nil }

type compressedFileInstance struct {
	*compressedFile
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *compressedFileInstance) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedFile.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.ReadFull(f.gr, make([]byte, f.seekPos-f.grPos))
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos += int64(n)
	return n, err
}
func (f *compressedFileInstance) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case os.SEEK_SET:
		f.seekPos = 0 + offset
	case os.SEEK_CUR:
		f.seekPos += offset
	case os.SEEK_END:
		f.seekPos = f.compressedFile.uncompressedSize + offset
	}
	return f.seekPos, nil
}
func (f *compressedFileInstance) Close() error {
	return f.gr.Close()
}

// dir is ...
type dir struct {
	name    string
	entries []os.FileInfo // Not nil.
	modTime time.Time
}

func (d *dir) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *dir) Close() error               { return nil }
func (d *dir) Stat() (os.FileInfo, error) { return d, nil }

func (d *dir) Name() string       { return d.name }
func (d *dir) Size() int64        { return 0 }
func (d *dir) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *dir) ModTime() time.Time { return d.modTime }
func (d *dir) IsDir() bool        { return true }
func (d *dir) Sys() interface{}   { return nil }

// dirInstance is an opened dir instance.
type dirInstance struct {
	*dir
	pending []os.FileInfo
}

func (d *dirInstance) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == os.SEEK_SET {
		d.pending = nil
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.dir.name)
}

func (d *dirInstance) Readdir(count int) ([]os.FileInfo, error) {
	if d.pending == nil {
		d.pending = d.dir.entries
	}

	if len(d.pending) == 0 && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.pending) {
		count = len(d.pending)
	}
	e := d.pending[:count]
	d.pending = d.pending[count:]
	return e, nil
}
