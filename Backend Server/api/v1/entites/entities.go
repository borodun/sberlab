package entites

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

type Entities struct {
	Error    string   `json:"error"`
	Entities []Entity `json:"entities"`
}

type Entity struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

type ECSs struct {
	Servers []Entity `json:"servers"`
}

type VPCs struct {
	VPCs []Entity `json:"vpcs"`
}

type NATs struct {
	NATs []Entity `json:"nat_gateways"`
}
