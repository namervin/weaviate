//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package lsmkv

import (
	"bufio"
	"bytes"
	"io"

	"github.com/pkg/errors"
)

type compactorMap struct {
	// c1 is always the older segment, so when there is a conflict c2 wins
	// (because of the replace strategy)
	c1 *segmentCursorCollection
	c2 *segmentCursorCollection

	// the level matching those of the cursors
	currentLevel        uint16
	secondaryIndexCount uint16

	w    io.WriteSeeker
	bufw *bufio.Writer
}

func newCompactorMapCollection(w io.WriteSeeker,
	c1, c2 *segmentCursorCollection, level, secondaryIndexCount uint16) *compactorMap {
	return &compactorMap{
		c1:                  c1,
		c2:                  c2,
		w:                   w,
		bufw:                bufio.NewWriterSize(w, 1e6),
		currentLevel:        level,
		secondaryIndexCount: secondaryIndexCount,
	}
}

func (c *compactorMap) do() error {
	if err := c.init(); err != nil {
		return errors.Wrap(err, "init")
	}

	kis, err := c.writeKeys()
	if err != nil {
		return errors.Wrap(err, "write keys")
	}

	if err := c.writeIndices(kis); err != nil {
		return errors.Wrap(err, "write index")
	}

	// flush buffered, so we can safely seek on underlying writer
	if err := c.bufw.Flush(); err != nil {
		return errors.Wrap(err, "flush buffered")
	}

	dataEnd := uint64(kis[len(kis)-1].valueEnd)

	if err := c.writeHeader(c.currentLevel+1, 0, c.secondaryIndexCount,
		dataEnd); err != nil {
		return errors.Wrap(err, "write header")
	}

	return nil
}

func (c *compactorMap) init() error {
	// write a dummy header, we don't know the contents of the actual header yet,
	// we will seek to the beginning and overwrite the actual header at the very
	// end

	if _, err := c.bufw.Write(make([]byte, SegmentHeaderSize)); err != nil {
		return errors.Wrap(err, "write empty header")
	}

	return nil
}

func (c *compactorMap) writeKeys() ([]keyIndex, error) {
	key1, value1, _ := c.c1.first()
	key2, value2, _ := c.c2.first()

	// the (dummy) header was already written, this is our initial offset
	offset := SegmentHeaderSize

	var kis []keyIndex

	for {
		if key1 == nil && key2 == nil {
			break
		}
		if bytes.Equal(key1, key2) {
			values := append(value1, value2...)
			valuesMerged, err := newMapDecoder().DoPartial(values)
			if err != nil {
				return nil, err
			}

			var mergedEncoded []value
			for _, merged := range valuesMerged {
				encoded, err := newMapEncoder().Do(merged)
				if err != nil {
					return nil, err
				}

				mergedEncoded = append(mergedEncoded, encoded...)
			}

			ki, err := c.writeIndividualNode(offset, key2, mergedEncoded)
			if err != nil {
				return nil, errors.Wrap(err, "write individual node (equal keys)")
			}

			offset = ki.valueEnd
			kis = append(kis, ki)

			// advance both!
			key1, value1, _ = c.c1.next()
			key2, value2, _ = c.c2.next()
			continue
		}

		if (key1 != nil && bytes.Compare(key1, key2) == -1) || key2 == nil {
			// key 1 is smaller
			ki, err := c.writeIndividualNode(offset, key1, value1)
			if err != nil {
				return nil, errors.Wrap(err, "write individual node (key1 smaller)")
			}

			offset = ki.valueEnd
			kis = append(kis, ki)
			key1, value1, _ = c.c1.next()
		} else {
			// key 2 is smaller
			ki, err := c.writeIndividualNode(offset, key2, value2)
			if err != nil {
				return nil, errors.Wrap(err, "write individual node (key2 smaller)")
			}

			offset = ki.valueEnd
			kis = append(kis, ki)

			key2, value2, _ = c.c2.next()
		}
	}

	return kis, nil
}

func (c *compactorMap) writeIndividualNode(offset int, key []byte,
	values []value) (keyIndex, error) {
	return (&segmentCollectionNode{
		values:     values,
		primaryKey: key,
		offset:     offset,
	}).KeyIndexAndWriteTo(c.bufw)
}

func (c *compactorMap) writeIndices(keys []keyIndex) error {
	indices := &segmentIndices{
		keys:                keys,
		secondaryIndexCount: c.secondaryIndexCount,
	}

	_, err := indices.WriteTo(c.bufw)
	return err
}

// writeHeader assumes that everything has been written to the underlying
// writer and it is now safe to seek to the beginning and override the initial
// header
func (c *compactorMap) writeHeader(level, version, secondaryIndices uint16,
	startOfIndex uint64) error {
	if _, err := c.w.Seek(0, io.SeekStart); err != nil {
		return errors.Wrap(err, "seek to beginning to write header")
	}

	h := &segmentHeader{
		level:            level,
		version:          version,
		secondaryIndices: secondaryIndices,
		strategy:         SegmentStrategyMapCollection,
		indexStart:       startOfIndex,
	}

	if _, err := h.WriteTo(c.w); err != nil {
		return err
	}

	return nil
}