package controllers

type Factory struct {
}

type IFactory interface {
	NewIAccount(string) IAccount
	NewIDatabase(string, string) IDatabase
}

var controllersFactory IFactory

func InitControllersFactory(initFactory IFactory) {
	controllersFactory = initFactory
}

func NewFactory() Factory {
	return Factory{}
}

func (f Factory) NewIAccount(name string) IAccount {
	return NewAccount(name)
}

func (f Factory) NewIDatabase(addr string, pass string) IDatabase {
	return NewDatabase(addr, pass)
}
