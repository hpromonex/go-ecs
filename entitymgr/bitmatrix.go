package entitymgr

import (
	"fmt"
	"strings"

	"github.com/hpromonex/go-ecs/component"
)

const bytesPerEntity = component.Count/8 + 1

// Bitmatrix represents a bitmatrix
type Bitmatrix [ClusterSize * bytesPerEntity]byte

// Set sets the bit at coords
func (bm *Bitmatrix) Set(id, bit component.ID, val bool) {

	if bit >= component.Count {
		return
	}

	byteoff := bit / 8
	bitoff := bit % 8

	if val {
		bm[id*bytesPerEntity+byteoff] |= 1 << bitoff
	} else {
		bm[id*bytesPerEntity+byteoff] &^= 1 << bitoff
	}

}

// Get gets the bit at coords
func (bm *Bitmatrix) Get(id, bit component.ID) bool {
	byteoff := bit / 8
	bitoff := bit % 8
	return bm[id*bytesPerEntity+byteoff]&(1<<bitoff) > 0
}

func (bm Bitmatrix) String() string {
	var sb strings.Builder
	sb.WriteString("Bitmatrix:\n")
	sb.WriteString(fmt.Sprintf("Bytes per Entity: %d", bytesPerEntity))

	bpe := int(bytesPerEntity)

	for i := 0; i < len(bm)/bpe; i++ {
		sb.WriteString(fmt.Sprintf("\n% 3d : ", i))
		for b := 0; b < bpe; b++ {
			sb.WriteString(fmt.Sprintf("%08b ", bm[i*bpe+b]))
		}
	}

	return sb.String()
}

// Query component Ids
func (bm *Bitmatrix) Query(components ...component.ID) []uint32 {

	query := idsToBytes(components)
	bpe := int(bytesPerEntity)

	var ids []uint32
	for i := 0; i < len(bm)/bpe; i++ {
		entityBytes := bm[i*bpe : i*bpe+bpe]

		var byteEntity uint8
		var byteComparison uint8
		var byteResult uint8

		mark := true
		for bidx := 0; bidx < bpe; bidx++ {
			byteEntity = entityBytes[bidx]
			byteComparison = query[bidx]
			byteResult = byteEntity & byteComparison
			if byteResult != byteComparison {
				mark = false
				break
			}
		}

		if mark {
			ids = append(ids, uint32(i))
		}
	}

	return ids
}

// idsToBytes turns an array of componentIDs into a bitset representation
func idsToBytes(components []component.ID) [bytesPerEntity]uint8 {
	var bytes [bytesPerEntity]byte

	for _, c := range components {

		byteoff := c / 8
		bitoff := c % 8

		bytes[byteoff] |= 1 << bitoff

	}

	return bytes
}
