package domain_user

type userEntity struct {
	id string
}

func (e *userEntity) ID() string {
	return e.id
}
