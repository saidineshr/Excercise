package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)
func update(TotalAmount *int,Amount int,m *sync.Mutex)error{
	m.Lock()
	fmt.Println("Current Bank Balance:",*TotalAmount)
	*TotalAmount+=Amount
	if Amount<0{
		fmt.Println("withdraw:",Amount,"BankBalance:",*TotalAmount)
	}else{
		fmt.Println("Deposit:",Amount,"BankBalance:",*TotalAmount)
	}
	fmt.Println(".......")
	if *TotalAmount<0{
		*TotalAmount-=Amount
		fmt.Println("Final Balance:",*TotalAmount)
		return errors.New("you cannot withdraw as balance is less than 0")
	}else{
		defer m.Unlock()
		return nil
	}

}
func main() {
	BankBalance:=500

	var flag int
	var AmountToBeDeposited int
	var m sync.Mutex
	for {
		//deposit
		go func() {
			AmountToBeDeposited = int(rand.Int31n(10000))
			update(&BankBalance, AmountToBeDeposited, &m)

		}()
		//withdraw
		go func() {
			AmountToBeDeposited = -int(rand.Int31n(10000))
			temp := update(&BankBalance, AmountToBeDeposited, &m)

			if temp != nil {
				fmt.Println(temp)
				flag=1
				return
			}
		}()
		if flag==1 {
			break
		}
	}
}
