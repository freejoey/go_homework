package account

import (
	"context"
	"fmt"
	"homework/db"
	"log"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	UnimplementedAccountServer
}

type Account struct {
	aid int64

	uid int64

	properties float64

	status string

	createdAt int64

	updatedAt int64
}

func GetAccountCount() int {
	var count = make([]int, 1)
	err := db.DB.Select(&count, "select count(*) from account")
	if err != nil {
		fmt.Println("exec sql failed, ", err)
		panic(err)
	}
	return count[0]
}

func GetAccountInfo(id int) *Account {
	var account []Account
	err := db.DB.Select(&account, "select * from account where id=?", id)
	if err != nil {
		fmt.Println("exec sql failed, ", err)
		panic(err)
	}
	return &account[0]
}

func (s *server) Info(ctx context.Context, in *AccountInfoRequest) (*AccountInfoResponse, error) {
	var aid = in.GetAid()
	log.Printf("Received: %d", aid)

	var account []Account
	err := db.DB.Select(&account, "select * from t where id=?", aid)
	if err != nil {
		fmt.Println("exec sql failed, ", err)
		panic(err)
	}
	return &AccountInfoResponse{}, nil
}
