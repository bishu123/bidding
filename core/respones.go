package core

type CommonResponse struct {
    APIStatus bool `json:"api_status"`
}

type BiddingResponse struct {
    CommonResponse
    AdId     int     `json:"ad_id"`
    BidPrice float32 `json:"bid_price"`
}
