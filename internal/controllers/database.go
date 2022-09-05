package controllers

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var db IDatabase

const LIKES_TABLE string = "account:%s:likes"

type database struct {
	client *redis.Client
}

type IDatabase interface {
	SelectAmountLikesForAccount(string) (int, error)
	IsAccountLikes(string, string) (bool, error)
	AddLikeToAccount(string, string) error
}

func NewDatabase(addr string, pass string) database {
	return database{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pass,
			DB:       0,
		})}
}

func InitDatabase(dbHost string, dbPort int, pass string) {
	db = controllersFactory.NewIDatabase(fmt.Sprintf("%s:%d", dbHost, dbPort), pass)
}

func (d database) SelectAmountLikesForAccount(targetAcc string) (int, error) {
	res := d.client.SCard(ctx, fmt.Sprintf(LIKES_TABLE, targetAcc))
	return int(res.Val()), res.Err()
}

func (d database) IsAccountLikes(currentAcc string, targetAcc string) (bool, error) {
	res := d.client.SIsMember(ctx, fmt.Sprintf(LIKES_TABLE, targetAcc), currentAcc)
	return res.Val(), res.Err()
}

func (d database) AddLikeToAccount(targetAcc string, currentAcc string) error {
	res := d.client.SAdd(ctx, fmt.Sprintf(LIKES_TABLE, targetAcc), currentAcc, 0)
	return res.Err()
}
