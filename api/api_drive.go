package openapi

import (
	"github.com/gin-gonic/gin"
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
func drive(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFiles - drive/files
func driveFiles(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFilesAttachedNotes - drive/files/attached-notes
func driveFilesAttachedNotes(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFilesCheckExistence - drive/files/check-existence
func driveFilesCheckExistence(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFilesCreate - drive/files/create
func driveFilesCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFilesDelete - drive/files/delete
func driveFilesDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFilesFind - drive/files/find
func driveFilesFind(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFilesFindByHash - drive/files/find-by-hash
func driveFilesFindByHash(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFilesShow - drive/files/show
func driveFilesShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFilesUpdate - drive/files/update
func driveFilesUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFilesUploadFromUrl - drive/files/upload-from-url
func driveFilesUploadFromUrl(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFolders - drive/folders
func driveFolders(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFoldersCreate - drive/folders/create
func driveFoldersCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFoldersDelete - drive/folders/delete
func driveFoldersDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFoldersFind - drive/folders/find
func driveFoldersFind(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFoldersShow - drive/folders/show
func driveFoldersShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveFoldersUpdate - drive/folders/update
func driveFoldersUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// driveStream - drive/stream
func driveStream(c *gin.Context) {
	// TODO
	placeholder(c)
}
