package collector

import "encoding/json"

// ClusterStatsResponse is a representation of a Elasticsearch Cluster Stats
type ClusterStatsResponse struct {
	ClusterName string `json:"cluster_name"`
	Timestamp   int64  `json:"timestamp"`
	Status      string `json:"status"`
	Indices     ClusterStatsIndicesResponse `json:"indices"`
	Nodes       ClusterStatsNodeResponse `json:"nodes"`
}

type ClusterStatsIndicesResponse struct {
	Count int64 `json:"count"`
	Shards int64 `json:"shards"`
	
  
}
