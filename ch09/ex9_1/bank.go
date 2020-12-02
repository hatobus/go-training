package bank

var deposits = make(chan int)
var balances = make(chan int)

type withDraw struct {
	amount int
	res    chan bool
}

var drawch = make(chan withDraw)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	r := make(chan bool)
	drawch <- withDraw{amount, r}
	return <-r
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case wd := <-drawch:
			if balance < wd.amount {
				wd.res <- false
			} else {
				balance -= wd.amount
				wd.res <- true
			}
		}
	}
}

func init() {
	go teller()
}
