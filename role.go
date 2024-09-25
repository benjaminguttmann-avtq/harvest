package harvest

import (
	"fmt"
	"time"
)

type RolesResponse struct {
	PagedResponse
	Roles []*Role `json:"roles"`
}

type RoleStub struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UserIDs   int64     `json:"user_ids"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *API) GetRole(roleID int64, args Arguments) (role *Role, err error) {
	role = &Role{}
	path := fmt.Sprintf("/roles/%v", roleID)
	err = a.Get(path, args, &role)
	return role, err
}

func (a *API) GetRoles(args Arguments) (roles []*Role, err error) {
	roles = make([]*Role, 0)
	rolesResponse := RolesResponse{}
	err = a.GetPaginated("/roles", args, &rolesResponse, func() {
		for _, u := range rolesResponse.Roles {
			roles = append(roles, u)
		}
		rolesResponse.Roles = make([]*Role, 0)
	})
	return roles, err
}
