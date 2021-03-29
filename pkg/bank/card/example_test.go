package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: true,
		},
		{
			Balance: 2_000_00,
			Active: true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active: true,
		},
	}))

	// Output
	// 100000
	// 300000
	// 0
	// 0
}




// func ExampleAddBonus_positive() {
// 	card := types.Card{MinBalance: 10_000_00, Balance: 20_000_00, Active: true}
// 	AddBonus(&card, 3, 30, 365)
// 	fmt.Println(card.Balance)
// 	// Output: 2002465
// }

// func ExampleAddBonus_inactive() {
// 	card := types.Card{Balance: 0, Active: false}
// 	AddBonus(&card, 3, 30, 365)
// 	fmt.Println(card.Balance)
// 	// Output: 0
// }


// func ExampleAddBonus_negative() {
// 	card := types.Card{MinBalance: -10_000_00, Balance: 10_000_00, Active: true}
// 	AddBonus(&card, 3, 30, 365)
// 	fmt.Println(card.Balance)
// 	// Output: 997535
// }


// func ExampleAddBonus_limit() {
// 	card := types.Card{Balance: 100_000, Active: true}
// 	AddBonus(&card, 3, 30, 365)
// 	fmt.Println(card.Balance)
// 	// Output: 100000
// }



// func ExampleDeposit_positive() {
// 	card := types.Card{Balance: 20_000, Active: true}
// 	Deposit(&card, 10_000)
// 	fmt.Println(card.Balance)
// 	// Output: 30000
// }

// func ExampleDeposit_inactive() {
// 	card := types.Card{Balance: 0, Active: false}
// 	Deposit(&card, 10_000)
// 	fmt.Println(card.Balance)
// 	// Output: 0
// }


// func ExampleDeposit_limit() {
// 	card := types.Card{Balance: 40_000_00, Active: true}
// 	Deposit(&card, 60_000_00)
// 	fmt.Println(card.Balance)
// 	// Output: 4000000
// }


// func ExampleWithdraw() {
// 	card := types.Card{Balance: 20_000_00, Active: true}
// 	Withdraw(&card, 10_000_00)
// 	fmt.Println(card.Balance)
// 	// Output: 1000000
// }





// func ExampleCard() {
// 	card := types.Card{Balance: 20_000_00, Active: true}
// 	clone := card
// 	card.Balance -= 10_000_00
// 	fmt.Println(clone.Balance)
// 	// Output: 1000000
// }

// func ExampleWithdraw_positive() {
// 	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
// 	fmt.Println(result.Balance)
// 	// Output: 1000000
// }

// func ExampleWithdraw_noMoney() {
// 	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
// 	fmt.Println(result.Balance)
// 	// Output: 0
// }

// func ExampleWithdraw_inactive() {
// 	result := Withdraw(types.Card{Balance: 0, Active: false}, 10_000_00)
// 	fmt.Println(result.Balance)
// 	// Output: 0
// }

// func ExampleWithdraw_limit() {
// 	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 40_000_00)
// 	fmt.Println(result.Balance)
// 	// Output: 2000000
// }