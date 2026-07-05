package permissions

type PermissionCode string

const (
	// User management
	PermissionCodeUsersRead       PermissionCode = "users.read"
	PermissionCodeUsersCreate     PermissionCode = "users.create"
	PermissionCodeUsersUpdate     PermissionCode = "users.update"
	PermissionCodeUsersDelete     PermissionCode = "users.delete"
	PermissionCodeUsersActivate   PermissionCode = "users.activate"
	PermissionCodeUsersDeactivate PermissionCode = "users.deactivate"

	// Permission management
	PermissionCodePermissionsRead   PermissionCode = "permissions.read"
	PermissionCodePermissionsAssign PermissionCode = "permissions.assign"
	PermissionCodePermissionsRevoke PermissionCode = "permissions.revoke"

	// Role management
	PermissionCodeRolesRead   PermissionCode = "roles.read"
	PermissionCodeRolesCreate PermissionCode = "roles.create"
	PermissionCodeRolesUpdate PermissionCode = "roles.update"
	PermissionCodeRolesDelete PermissionCode = "roles.delete"
	PermissionCodeRolesAssign PermissionCode = "roles.assign"

	// Workspace management
	PermissionCodeWorkspacesRead          PermissionCode = "workspaces.read"
	PermissionCodeWorkspacesCreate        PermissionCode = "workspaces.create"
	PermissionCodeWorkspacesUpdate        PermissionCode = "workspaces.update"
	PermissionCodeWorkspacesDelete        PermissionCode = "workspaces.delete"
	PermissionCodeWorkspacesManageMembers PermissionCode = "workspaces.manage_members"

	// Project management
	PermissionCodeProjectsRead           PermissionCode = "projects.read"
	PermissionCodeProjectsCreate         PermissionCode = "projects.create"
	PermissionCodeProjectsUpdate         PermissionCode = "projects.update"
	PermissionCodeProjectsDelete         PermissionCode = "projects.delete"
	PermissionCodeProjectsManageMembers  PermissionCode = "projects.manage_members"
	PermissionCodeProjectsManageSettings PermissionCode = "projects.manage_settings"

	// Ticket management
	PermissionCodeTicketsRead           PermissionCode = "tickets.read"
	PermissionCodeTicketsCreate         PermissionCode = "tickets.create"
	PermissionCodeTicketsUpdate         PermissionCode = "tickets.update"
	PermissionCodeTicketsDelete         PermissionCode = "tickets.delete"
	PermissionCodeTicketsAssign         PermissionCode = "tickets.assign"
	PermissionCodeTicketsChangeStatus   PermissionCode = "tickets.change_status"
	PermissionCodeTicketsChangePriority PermissionCode = "tickets.change_priority"
	PermissionCodeTicketsChangeDueDate  PermissionCode = "tickets.change_due_date"

	// Ticket comments
	PermissionCodeTicketCommentsRead   PermissionCode = "ticket_comments.read"
	PermissionCodeTicketCommentsCreate PermissionCode = "ticket_comments.create"
	PermissionCodeTicketCommentsUpdate PermissionCode = "ticket_comments.update"
	PermissionCodeTicketCommentsDelete PermissionCode = "ticket_comments.delete"

	// Ticket attachments
	PermissionCodeTicketAttachmentsRead   PermissionCode = "ticket_attachments.read"
	PermissionCodeTicketAttachmentsUpload PermissionCode = "ticket_attachments.upload"
	PermissionCodeTicketAttachmentsDelete PermissionCode = "ticket_attachments.delete"

	// Kanban board
	PermissionCodeKanbanRead       PermissionCode = "kanban.read"
	PermissionCodeKanbanUpdate     PermissionCode = "kanban.update"
	PermissionCodeKanbanMoveTicket PermissionCode = "kanban.move_ticket"

	// Workflow
	PermissionCodeWorkflowsRead   PermissionCode = "workflows.read"
	PermissionCodeWorkflowsCreate PermissionCode = "workflows.create"
	PermissionCodeWorkflowsUpdate PermissionCode = "workflows.update"
	PermissionCodeWorkflowsDelete PermissionCode = "workflows.delete"

	// Dashboard / reports
	PermissionCodeDashboardRead PermissionCode = "dashboard.read"
	PermissionCodeReportsRead   PermissionCode = "reports.read"

	// Audit logs
	PermissionCodeAuditLogsRead PermissionCode = "audit_logs.read"

	// System settings
	PermissionCodeSystemSettingsRead   PermissionCode = "system_settings.read"
	PermissionCodeSystemSettingsUpdate PermissionCode = "system_settings.update"
)
