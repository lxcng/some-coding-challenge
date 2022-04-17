package api

type CreateAccountReq struct {
	FullName string
	Balance  float64
}

type CreateAccountResp struct {
	ID    int64
	Error string `json:",omitempty"`
}

type GetBalanceResp struct {
	Balance float64
	Error   string `json:",omitempty"`
}

type DepositResp struct {
	Ok    bool
	Error string `json:",omitempty"`
}

type WithdrawResp struct {
	Ok    bool
	Error string `json:",omitempty"`
}
