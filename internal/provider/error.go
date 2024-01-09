package provider

import (
	"github.com/paloaltonetworks/scm-go/api"
)

func IsObjectNotFound(e error) bool {
	return e == api.ObjectNotFoundError
}
