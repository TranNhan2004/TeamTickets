CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO permissions (id, code, description)
VALUES
    -- User management
    (gen_random_uuid(), 'users.read', 'View users'),
    (gen_random_uuid(), 'users.create', 'Create users'),
    (gen_random_uuid(), 'users.update', 'Update users'),
    (gen_random_uuid(), 'users.delete', 'Delete users'),
    (gen_random_uuid(), 'users.activate', 'Activate users'),
    (gen_random_uuid(), 'users.deactivate', 'Deactivate users'),

    -- Permission management
    (gen_random_uuid(), 'permissions.read', 'View permissions'),
    (gen_random_uuid(), 'permissions.assign', 'Assign permissions'),
    (gen_random_uuid(), 'permissions.revoke', 'Revoke permissions'),

    -- Role management
    (gen_random_uuid(), 'roles.read', 'View roles'),
    (gen_random_uuid(), 'roles.create', 'Create roles'),
    (gen_random_uuid(), 'roles.update', 'Update roles'),
    (gen_random_uuid(), 'roles.delete', 'Delete roles'),
    (gen_random_uuid(), 'roles.assign', 'Assign roles to users'),

    -- Workspace management
    (gen_random_uuid(), 'workspaces.read', 'View workspaces'),
    (gen_random_uuid(), 'workspaces.create', 'Create workspaces'),
    (gen_random_uuid(), 'workspaces.update', 'Update workspaces'),
    (gen_random_uuid(), 'workspaces.delete', 'Delete workspaces'),
    (gen_random_uuid(), 'workspaces.manage_members', 'Manage workspace members'),

    -- Project management
    (gen_random_uuid(), 'projects.read', 'View projects'),
    (gen_random_uuid(), 'projects.create', 'Create projects'),
    (gen_random_uuid(), 'projects.update', 'Update projects'),
    (gen_random_uuid(), 'projects.delete', 'Delete projects'),
    (gen_random_uuid(), 'projects.manage_members', 'Manage project members'),
    (gen_random_uuid(), 'projects.manage_settings', 'Manage project settings'),

    -- Ticket management
    (gen_random_uuid(), 'tickets.read', 'View tickets'),
    (gen_random_uuid(), 'tickets.create', 'Create tickets'),
    (gen_random_uuid(), 'tickets.update', 'Update tickets'),
    (gen_random_uuid(), 'tickets.delete', 'Delete tickets'),
    (gen_random_uuid(), 'tickets.assign', 'Assign tickets'),
    (gen_random_uuid(), 'tickets.change_status', 'Change ticket status'),
    (gen_random_uuid(), 'tickets.change_priority', 'Change ticket priority'),
    (gen_random_uuid(), 'tickets.change_due_date', 'Change ticket due date'),

    -- Ticket comments
    (gen_random_uuid(), 'ticket_comments.read', 'View ticket comments'),
    (gen_random_uuid(), 'ticket_comments.create', 'Create ticket comments'),
    (gen_random_uuid(), 'ticket_comments.update', 'Update ticket comments'),
    (gen_random_uuid(), 'ticket_comments.delete', 'Delete ticket comments'),

    -- Ticket attachments
    (gen_random_uuid(), 'ticket_attachments.read', 'View ticket attachments'),
    (gen_random_uuid(), 'ticket_attachments.upload', 'Upload ticket attachments'),
    (gen_random_uuid(), 'ticket_attachments.delete', 'Delete ticket attachments'),

    -- Kanban board
    (gen_random_uuid(), 'kanban.read', 'View kanban board'),
    (gen_random_uuid(), 'kanban.update', 'Update kanban board'),
    (gen_random_uuid(), 'kanban.move_ticket', 'Move tickets on kanban board'),

    -- Workflow
    (gen_random_uuid(), 'workflows.read', 'View workflows'),
    (gen_random_uuid(), 'workflows.create', 'Create workflows'),
    (gen_random_uuid(), 'workflows.update', 'Update workflows'),
    (gen_random_uuid(), 'workflows.delete', 'Delete workflows'),

    -- Dashboard / reports
    (gen_random_uuid(), 'dashboard.read', 'View dashboard'),
    (gen_random_uuid(), 'reports.read', 'View reports'),

    -- Audit logs
    (gen_random_uuid(), 'audit_logs.read', 'View audit logs'),

    -- System settings
    (gen_random_uuid(), 'system_settings.read', 'View system settings'),
    (gen_random_uuid(), 'system_settings.update', 'Update system settings')
ON CONFLICT (code) DO NOTHING;