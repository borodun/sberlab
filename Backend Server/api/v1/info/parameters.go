package info

import (
	"backend/api/v1/ecss"
	"backend/api/v1/vpcs"
)

type Info struct {
	Servers ecss.Servers `json:"ecss"`
	VPCs    vpcs.VPCs    `json:"vpcs"`
}
