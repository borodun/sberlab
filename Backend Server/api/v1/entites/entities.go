package entites

type Details struct {
	Error   string `json:"error"`
	Details string `json:"details"`
}

type EntityArray struct {
	EntityInfos []EntityInfo `json:"entities"`
}

type EntityInfo struct {
	Endpoint   string `json:"endpoint"`
	Link       string `json:"link"`
	Type       string `json:"type"`
	TypeInJSON string `json:"type_in_json"`
	Details    string `json:"details"`
}

type AnswerEntities struct {
	Error          string   `json:"error"`
	AnswerEntities []Entity `json:"entities_array"`
}

type Entities struct {
	Entities []Entity `json:"entities"`
}

type Entity struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

type ECSs struct {
	Entities []Entity `json:"servers"`
}

type VPCs struct {
	Entities []Entity `json:"vpcs"`
}

type Volumes struct {
	Entities []Entity `json:"volumes"`
}

type Subnets struct {
	Entities []Entity `json:"subnets"`
}

type SGs struct {
	Entities []Entity `json:"security_groups"`
}

type NATs struct {
	Entities []Entity `json:"nat_gateways"`
}

type SNATEntity struct {
	Name   string `json:"floating_ip_address"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type SNATs struct {
	Entities []SNATEntity `json:"snat_rules"`
}

type DNATEntity struct {
	Internal int    `json:"internal_service_port"`
	External int    `json:"external_service_port"`
	ID       string `json:"id"`
	Status   string `json:"status"`
}

type DNATs struct {
	Entities []DNATEntity `json:"dnat_rules"`
}

type EIPs struct {
	Entities []Entity `json:"publicips"`
}

type ELBEntity struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"operating_status"`
}

type ELBs struct {
	Entities []ELBEntity `json:"loadbalancers"`
}

type ListenerEntity struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Port int    `json:"protocol_port"`
}

type Listeners struct {
	Entities []ListenerEntity `json:"listeners"`
}

type PoolEntity struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	Protocol string `json:"protocol"`
}

type Pools struct {
	Entities []PoolEntity `json:"pools"`
}

type BackendEntity struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"operating_status"`
}

type BackendServers struct {
	Entities []BackendEntity `json:"members"`
}
