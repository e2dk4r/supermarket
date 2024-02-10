package mock

type DbErrorService struct {
	IsDuplicateErrorFn      func(err error) bool
	isDuplicateErrorInvoked bool

	IsNotFoundErrorFn      func(err error) bool
	isNotFoundErrorInvoked bool
}

func (dbs *DbErrorService) IsDuplicateError(err error) bool {
	dbs.isDuplicateErrorInvoked = true
	return dbs.IsDuplicateErrorFn(err)
}

func (dbs *DbErrorService) IsDuplicateErrorInvoked() bool {
	return dbs.isDuplicateErrorInvoked
}

func (dbs *DbErrorService) IsNotFoundError(err error) bool {
	dbs.isNotFoundErrorInvoked = true
	return dbs.IsNotFoundErrorFn(err)
}

func (dbs *DbErrorService) IsNotFoundErrorInvoked() bool {
	return dbs.isNotFoundErrorInvoked
}
