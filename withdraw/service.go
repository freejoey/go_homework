package withdraw

import (
	"context"
	"fmt"
	"homework/db"
	"log"
	"time"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	UnimplementedWithdrawServer
}

type Withdraw struct {
	wid int64

	uid int64

	amount float64

	createdAt int64

	updatedAt int64
}

type Account struct {
	aid int64

	uid int64

	properties float64

	status string

	createdAt int64

	updatedAt int64
}

func (s *server) Withdraw(_ context.Context, in *WithdrawRequest) (*WithdrawResponse, error) {
	var uid, amount = in.GetUid(), in.GetAmount()
	log.Printf("Received uid: %d, amount: %f\n", uid, amount)
	res, err := db.DB.Exec("update acount set properties = properties + ? where uid = ? ", -amount, uid)
	if err != nil {
		fmt.Println("exec sql failed, ", err)
		panic(err)
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
		panic(err)
	}
	fmt.Println("update succeed:", row)

	res, err = db.DB.Exec("insert into withdraw(uid,amount,createdAt,updatedAt) values (?,?,?)", uid, amount, time.Now(), time.Now())
	if err != nil {
		fmt.Println("exec failed, ", err)
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		panic(err)
	}

	fmt.Println("insert succeed:", id)
	return &WithdrawResponse{}, nil
}
