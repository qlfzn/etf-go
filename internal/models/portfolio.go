package models

type FinalPortfolio struct {
	TotalInvestment      float64            `json:"total_investment"`
	WeightedExpenseRatio float64            `json:"weighted_expense_ratio"`
	ExpectedDividend     float64            `json:"expected_dividend_income"`
	SectorExposure       map[string]float64 `json:"sector_exposure"`
	TopHoldings          []HoldingSummary   `json:"top_holdings"`
}

type HoldingSummary struct {
	Symbol string  `json:"symbol"`
	Weight float64 `json:"weight"`
}
