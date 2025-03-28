package main

import (
	"errors"
	"sync"
)

// Cache Define an interface
type Cache interface {
	ArtifactCache
	LocalArtifactCache
}
type ArtifactCache interface {
	MissingBlob()
	PutArtifact(artifactId int64, artifactInfo string) (err error)
	PutBlob()
	DeleteBlob(blobIDs []string) (err error)
}

type LocalArtifactCache interface {
	GetArtifact()
	GetBlob(blobId string) (string, error)
	Close()
	Clear()
}

/*
🧠 What Does It Do?
Type Assertion: The var _ Cache = &MemoryCache{} statement is assigning a pointer to an MemoryCache instance (&MemoryCache{}) to a variable of type Cache.
Compile-Time Check: If MemoryCache does not implement all the methods defined in the Cache interface (which includes both ArtifactCache and LocalArtifactCache), the code will fail to compile.
Zero Value: It does not create any object at runtime—this is purely a type-checking mechanism.
*/
var _ Cache = &MemoryCache{}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{}
}

type MemoryCache struct {
	artifact sync.Map
	blobs    sync.Map
}

func (m *MemoryCache) MissingBlob() {
	panic("implement me")
}

func (m *MemoryCache) PutArtifact(artifactId int64, artifactInfo string) error {
	m.artifact.Store(artifactId, artifactInfo)
	return nil
}

func (m *MemoryCache) PutBlob() {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryCache) DeleteBlob(blobIDs []string) error {
	for _, blob := range blobIDs {
		m.blobs.Delete(blob)
	}
	return nil
}

func (m *MemoryCache) GetArtifact() {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryCache) GetBlob(blobId string) (string, error) {
	info, ok := m.blobs.Load(blobId)
	if !ok {
		return "", errors.New("no blob found")
	}
	blob, ok := info.(string)
	if !ok {
		return "", errors.New("no blob found")
	}
	return blob, nil
}

func (m *MemoryCache) Close() {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryCache) Clear() {
	//TODO implement me
	panic("implement me")
}
