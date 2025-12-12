package schemas

import base "github.com/MatheusMikio/Base"

type Professional struct {
	base.BaseUser
	Appointments []Scheduling `json:"appointments"`
}
