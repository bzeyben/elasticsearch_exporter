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
	Docs 
}

type ClusterStatsIndicesShardsResponse struct {
	Total int64 `json:"total"`
	Primaries int64 `json:"primaries"`
	Replication float64 `json:"replication"`
	Index map[string]ClusterStatsIndicesShardsIndexResponse `json:"index"`
}

type ClusterStatsIndicesShardsIndexResponse struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
}

type ClusterStatsNodeResponse struct {
	Count ClusterStatsNodeCountResponse `json:"count"`
	Versions []string `json:"versions"`
}

type ClusterStatsNodeCountResponse struct {
	Total int64 `json:"total"`
	Data int64 `json:"data"`
	CoordinatingOnly int64 `json:"coordinating_only"`
	Master int64 `json:"master"`
	Ingest int64 `json:"ingest"`
}
