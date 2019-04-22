package models

import "github.com/qhenkart/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
