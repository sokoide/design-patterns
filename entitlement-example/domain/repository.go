package domain

import "errors"

// EntitlementRepository is the interface that the UseCase depends on.
// It abstracts the data retrieval mechanism (Cache, AD, DB, etc.).
type EntitlementRepository interface {
	// IsMember checks if the user belongs to the group.
	IsMember(groupName string, userName string) (bool, error)
}

var (
	ErrDataSourceDown = errors.New("data source unavailable")
)
