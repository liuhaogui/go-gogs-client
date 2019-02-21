// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"fmt"
)

// User represents a API user.
type User struct {
	ID        int64  `json:"id"`
	UserName  string `json:"username"` // LEGACY [Gogs 1.0]: remove field(s) for backward compatibility
	Login     string `json:"login"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
}

type UserList struct {
	Data []User `json:"data""`
	Ok   bool   `json:"ok"`
}

func (c *Client) GetUserInfo(user string) (*User, error) {
	u := new(User)
	err := c.getParsedResponse("GET", fmt.Sprintf("/users/%s", user), nil, nil, u)
	return u, err
}

func (c *Client) SearchUser(user string) (*UserList,error) {
	u := new(UserList)
	err := c.getParsedResponse("GET", fmt.Sprintf("/users/search?q=%s", user), nil, nil, u)
	return u, err
}
