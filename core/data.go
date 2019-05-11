package core

type AdDataType struct {
    RequestList []BiddingRequest
    AdData map[int]*AdSlot
}

type AdSlot struct {
    BidPrice float32
    Available bool
    AdPlacementId string
}

var AdData AdDataType

// InitData will initialize the Data Required for Auction
func InitData(){

    /*
    Assuming AdSlot is to be unique across system having
    multiple bid of different ads on it.
     */

    AdData = AdDataType{
        AdData:map[int]*AdSlot{
            10:{10.5,true,"12345"},
            11:{11, true, "12346"},
            12:{12,false,"12347"},
        },
    }
    // adding each request in the list to have a unique id for each request.
    AdData.RequestList = make([]BiddingRequest,0)
}