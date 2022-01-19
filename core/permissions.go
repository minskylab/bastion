package core

import hasura "github.com/minskylab/ent-hasura"

const memberRoleName = "member"

func MemberPermissions(
	sel *hasura.PermissionSelect,
	ins *hasura.PermissionInsert,
	upd *hasura.PermissionUpdate,
	del *hasura.PermissionDelete,
) hasura.PermissionsRoleAnnotation {
	return hasura.PermissionsRoleAnnotation{
		Role:             memberRoleName,
		SelectPermission: sel,
		InsertPermission: ins,
		UpdatePermission: upd,
		DeletePermission: del,
	}
}
