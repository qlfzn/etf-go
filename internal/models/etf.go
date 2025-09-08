package models

type ETFProfile struct {
	AmountInvested float64 `json:"amount_invested"`
	NetAssets string `json:"net_assets"`
	NetExpenseRatio string `json:"net_expense_ratio"`
	PortfolioTurnover string `json:"portfolio_turnover"`
	DividendYield string `json:"dividend_yield"`
	Sectors []Sectors `json:"sectors"`
	Holdings []Holdings `json:"holdings"`
}

type Sectors struct {
	Name string `json:"sector"`
	Weight string `json:"weight"`
}

type Holdings struct {
	Symbol string `json:"symbol"`
	Description string `json:"description"`
	Weight string `json:"weight"`
}