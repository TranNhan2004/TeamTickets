package workspaces

import "time"

type Workspace struct {
	ID        string
	Name      string
	OwnerID   string
	DeletedAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type WorkspaceRole struct {
	ID   string
	Name string
}

type WorkspaceRolePermission struct {
	ID              string
	WorkspaceRoleID string
	PermissionID    string
}

type WorkspaceMember struct {
	ID              string
	WorkspaceID     string
	UserID          string
	WorkspaceRoleID string
	DeletedAt       *time.Time
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}
