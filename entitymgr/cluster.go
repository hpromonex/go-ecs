package entitymgr

//ClusterSize determents the maximum Clustersize
const ClusterSize = 8

//Cluster represents a cluster of entities
type Cluster struct {
	entityComponents Bitmatrix
	entities         []EntityID
}
