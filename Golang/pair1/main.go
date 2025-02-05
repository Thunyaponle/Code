/*
package main

import (

	"fmt"
	"pair1/bank"

)

	func main() {
		account := bank.NewInterestAccount(120)
		fmt.Println("-10")
		err := account.Deposit(-10)

		if err != nil {
			fmt.Println("error : ", err)
		}

		err = account.Withdrawing
	}
*/
package main

import (
	"banking/bank"
	"fmt"
)

func main() {
	// สร้างบัญชีที่มีอัตราดอกเบี้ย
	account := bank.NewInterestAccount(120, 0.05)

	// ทดลองฝากเงิน
	fmt.Println("Depositing -10:")
	err := account.Deposit(-10)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// ทดลองถอนเงิน
	fmt.Println("Withdrawing 50:")
	err = account.Withdraw(50)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Withdraw successful!")
	}

	fmt.Println("Current Balance:", account.GetBalance())

	account.AddInterest()
	fmt.Println("After Adding Interest, Balance:", account.GetBalance())
}
