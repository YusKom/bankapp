package card

import (
	"bank/pkg/bank/types"
	// "fmt"
)

func Withdraw(card *types.Card, amount types.Money) {
	const limitWithdraw = 20_000_00
	if amount < 0 {
		return
	}

	if amount > limitWithdraw {
		return
	}

	if !(*card).Active {
		return
	}

	if card.Balance < amount {
		return
	}

	card.Balance -= amount
}


func Deposit(card *types.Card, amount types.Money) {
	const amountLimit = 50_000_00
	
	if !(*card).Active {
		return
	}

	if amount > amountLimit {
		return
	}

	
	if amount< 0 {
		
		return
	}

	card.Balance += amount
		return


}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {


	const bonusLimit = 5_000_00

	// bonus := card.MinBalance * types.Money(percent)  * types.Money(daysInMonth) / types.Money(daysInYear)

	bonus := (card.MinBalance * types.Money(percent) / types.Money(100) * types.Money(daysInMonth)) / types.Money(daysInYear)

	if !(*card).Active {
		return
	}

	if card.Balance < 0 {
		return
	}

	if bonus > bonusLimit {
		return 
	}

	card.Balance += bonus 
	return
}


// func Deposit(card *types.Card, amount types.Money) {
// 	const amountLimit = 50_000
	
// 	if card.Active == true && card.Balance < 0 && amount <= amountLimit {
// 		return card
// 		card.Balance += amount
// 	}
	
// 	return card
// }


func IssueCard(currency types.Currency, color string, name string) types.Card {
	myCard := types.Card {
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,
	}

	return myCard
}


// Total вычисляет общую сумму средств на всех картах.
// Отрицательные суммы и суммы на заблокированных картах игнорируются.
func Total(cards []types.Card) types.Money {
	
	sum := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}

		if card.Balance <= 0 {
			continue
		}
		sum += card.Balance
	}
	return sum

}