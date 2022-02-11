package api

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "drive", drive}.register()
	route{http.MethodPost, "drive/files", driveFiles}.register()
	route{http.MethodPost, "drive/files/attached-notes", driveFilesAttachedNotes}.register()
	route{http.MethodPost, "drive/files/check-existence", driveFilesCheckExistence}.register()
	route{http.MethodPost, "drive/files/create", driveFilesCreate}.register()
	route{http.MethodPost, "drive/files/delete", driveFilesDelete}.register()
	route{http.MethodPost, "drive/files/find", driveFilesFind}.register()
	route{http.MethodPost, "drive/files/find-by-hash", driveFilesFindByHash}.register()
	route{http.MethodPost, "drive/files/show", driveFilesShow}.register()
	route{http.MethodPost, "drive/files/update", driveFilesUpdate}.register()
	route{http.MethodPost, "drive/files/upload-from-url", driveFilesUploadFromUrl}.register()
	route{http.MethodPost, "drive/folders", driveFolders}.register()
	route{http.MethodPost, "drive/folders/create", driveFoldersCreate}.register()
	route{http.MethodPost, "drive/folders/delete", driveFoldersDelete}.register()
	route{http.MethodPost, "drive/folders/find", driveFoldersFind}.register()
	route{http.MethodPost, "drive/folders/show", driveFoldersShow}.register()
	route{http.MethodPost, "drive/folders/update", driveFoldersUpdate}.register()
	route{http.MethodPost, "drive/stream", driveStream}.register()
}

// drive - drive
func drive(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFiles - drive/files
func driveFiles(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFilesAttachedNotes - drive/files/attached-notes
func driveFilesAttachedNotes(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFilesCheckExistence - drive/files/check-existence
func driveFilesCheckExistence(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFilesCreate - drive/files/create
func driveFilesCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFilesDelete - drive/files/delete
func driveFilesDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFilesFind - drive/files/find
func driveFilesFind(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFilesFindByHash - drive/files/find-by-hash
func driveFilesFindByHash(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFilesShow - drive/files/show
func driveFilesShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFilesUpdate - drive/files/update
func driveFilesUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFilesUploadFromUrl - drive/files/upload-from-url
func driveFilesUploadFromUrl(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFolders - drive/folders
func driveFolders(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFoldersCreate - drive/folders/create
func driveFoldersCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFoldersDelete - drive/folders/delete
func driveFoldersDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFoldersFind - drive/folders/find
func driveFoldersFind(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFoldersShow - drive/folders/show
func driveFoldersShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveFoldersUpdate - drive/folders/update
func driveFoldersUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// driveStream - drive/stream
func driveStream(ctx Context) {
	// TODO
	placeholder(ctx)
}
