package withdraw

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
