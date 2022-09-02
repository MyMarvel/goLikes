package controllers

type account struct {
	Name string
}

type IAccount interface {
	IncrementLikeCounter(string) (int, error)
	IsAccountLikes(string) (bool, error)
	GetLikesCount(string) (int, error)
}

func NewAccount(name string) account {
	return account{name}
}

func (acc account) IncrementLikeCounter(targetAccount string) (int, error) {
	err := db.AddLikeToAccount(targetAccount, acc.Name)
	likesCount, _ := db.SelectAmountLikesForAccount(targetAccount)
	return likesCount, err
}

func (acc account) IsAccountLikes(targetAccount string) (bool, error) {
	return db.IsAccountLikes(acc.Name, targetAccount)
}

func (acc account) GetLikesCount(targetAccount string) (int, error) {
	return db.SelectAmountLikesForAccount(targetAccount)
}
