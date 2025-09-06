package models

type ETFProfile struct {
	NetAssets float64 `json:"net_assets"`
	NetExpenseRatio float64 `json:"net_expense_ratio"`
	PortfolioTurnover float64 `json:"portfolio_turnover"`
	DividendYield float64 `json:"dividend_yield"`
	Sectors []Sectors `json:"sectors"`
	Holdings []Holdings `json:"holdings"`
}

type Sectors struct {
	Name string `json:"sector"`
	Weight float64 `json:"weight"`
}

type Holdings struct {
	Symbol string `json:"symbol"`
	Description string `json:"description"`
	Weight float64 `json:"weight"`
}