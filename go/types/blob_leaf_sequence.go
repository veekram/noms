// Copyright 2016 The Noms Authors. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package types

type blobLeafSequence struct {
	data []byte
	vr   ValueReader
}

func newBlobLeafSequence(vr ValueReader, data []byte) indexedSequence {
	return blobLeafSequence{data, vr}
}

// indexedSequence interface
func (bl blobLeafSequence) getOffset(idx int) uint64 {
	return uint64(idx)
}

func (bl blobLeafSequence) equalsAt(idx int, other interface{}) bool {
	return bl.data[idx] == other
}

// sequence interface
func (bl blobLeafSequence) getItem(idx int) sequenceItem {
	return bl.data[idx]
}

func (bl blobLeafSequence) seqLen() int {
	return len(bl.data)
}

func (bl blobLeafSequence) numLeaves() uint64 {
	return uint64(len(bl.data))
}

func (bl blobLeafSequence) valueReader() ValueReader {
	return bl.vr
}

func (bl blobLeafSequence) Chunks() []Ref {
	return []Ref{}
}

func (bl blobLeafSequence) Type() *Type {
	return BlobType
}