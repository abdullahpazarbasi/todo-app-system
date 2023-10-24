package domain_user_port

type Factory interface {
	CreateUserEntity(id string) (UserEntity, error)
}
