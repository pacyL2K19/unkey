package api

type nodeConfig struct {
	Platform  string `json:"platform,omitempty" description:"The platform this agent is running on"`
	Image     string `json:"image,omitempty" description:"The image this agent is running"`
	HttpPort  int    `json:"httpPort" default:"7070" description:"Port to listen on"`
	Schema    string `json:"$schema,omitempty" description:"Make jsonschema happy"`
	Region    string `json:"region,omitempty" description:"The region this agent is running in"`
	Heartbeat *struct {
		URL      string `json:"url" minLength:"1" description:"URL to send heartbeat to"`
		Interval int    `json:"interval" min:"1" description:"Interval in seconds to send heartbeat"`
	} `json:"heartbeat,omitempty" description:"Send heartbeat to a URL"`

	Cluster *struct {
		NodeID        string `json:"nodeId,omitempty" description:"A unique node id"`
		AdvertiseAddr string `json:"advertiseAddr,omitempty" description:"The address to advertise to other nodes"`
		RpcPort       int    `json:"rpcPort" default:"7071" description:"The port used for RPC"`
		GossipPort    int    `json:"gossipPort" default:"7072" description:"The port used for gossip"`
		Discovery     *struct {
			Static *struct {
				Addrs []string `json:"addrs" minLength:"1" description:"List of node addresses"`
			} `json:"static,omitempty" description:"Static cluster discovery configuration"`
			AwsCloudmap *struct {
				ServiceName string `json:"serviceName" minLength:"1" description:"Cloudmap service name"`
				Region      string `json:"region" minLength:"1" description:"Cloudmap region"`
			} `json:"awsCloudmap,omitempty" description:"Cloudmap cluster discovery configuration"`
		} `json:"discovery,omitempty" description:"Cluster discovery configuration, only one supported: static, cloudmap"`
	} `json:"cluster,omitempty" description:"Cluster configuration"`

	Logs *struct {
		Color bool `json:"color" description:"Display color in logs"`
	} `json:"logs,omitempty"`
	RedisUrl   string `json:"redisUrl"`
	Clickhouse *struct {
		Url string `json:"url" minLength:"1"`
	} `json:"clickhouse,omitempty"`

	Database struct {
		// DSN of the primary database for reads and writes.
		Primary string `json:"primary"`

		// An optional read replica DSN.
		ReadonlyReplica string `json:"readonlyReplica,omitempty"`
	} `json:"database"`
}
