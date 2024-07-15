package authority

import user "soms/repository/user"

func AuthorityFilterWithRole(authority []string, uuid string) bool { // authority list["Admin","Student","Master","Researcher","Others"]
	// Get User Role from DB
	userAuth, err := user.Repository.GetRoleByUUID(uuid)
	if err != nil {
		return false
	}

	// if userAuth in authority return true
	for _, auth := range authority {
		if userAuth.Role == auth {
			return true
		}
	}
	// else return false
	return true
}
