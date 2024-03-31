package main

import (
	"net/http"
	"sync"
	"time"
)

type downloadStatsItem struct {
	Bytes int64
	Count int64
}

var (
	// Use a map to track downloads, with each file's path as the key
	downloadStats             = make(map[string]downloadStatsItem)
	totalBytesSentForListings int64
	statsMutex                sync.Mutex
)

// Adjust countingWriter to immediately print download progress for each file
type countingWriter struct {
	http.ResponseWriter
	bytesWritten int64  // Track the number of bytes written
	path         string // The path of the file being served
}

// FileInfo is a struct to hold detailed information about files.
type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
}
