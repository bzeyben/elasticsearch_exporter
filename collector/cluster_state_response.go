package collector

import "encoding/json"

// NodeStatsResponse is a representation of a Elasticsearch Node Stats
type ClusterStateResponse struct {
    ClusterName string `json:"cluster_name"`
    Version     int64 `json:"version"`
	  Nodes       map[string]ClusterStateNodeResponse `json:"nodes"`
    Metadata    ClusterStateMetadataResponse `json:"metadata"`
}

type ClusterStateMetadataResponse struct {
  ClusterUUID string `json:"cluster_uuid"`
  
}
