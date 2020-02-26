package entitymgr

import "math"

//EntityID typedef
type EntityID uint32

var enitybits = uint32(math.Log2(ClusterSize) + 1)

//ClusterID returns the ClusterID part of EntityID
func (id EntityID) ClusterID() uint32 {
	return uint32(id >> enitybits)
}

//ClusterElement returns the ID within the containing Cluster
func (id EntityID) ClusterElement() uint32 {
	return uint32(id & ((1 << enitybits) - 1))
}
