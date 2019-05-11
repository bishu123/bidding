package core

import (
    "bidding/util"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

// AuctionView which will handle all the bidding services
func AuctionView(context *gin.Context) {
    biddingRequestArr := []BiddingRequest{
        {"12345", 10}, {"12345", 11}, {"12345", 12},
    }
    // Channel will have the response of BiddingView.
    responseChannel := make(chan BiddingResponse, 30)
    for _, val := range biddingRequestArr {
        // Hitting Bidding for each slot with 10 concurrent request
        for i := 0; i < 10; i++ {
            go Bid(responseChannel, val)
        }
    }

    //waiting for the result of  go routine.
    for i:=0;i<30;i++ {
        select {
        case res := <-responseChannel:
            context.JSON(http.StatusAccepted, res)
        case <-time.After(5 * time.Second):
            context.JSON(http.StatusGatewayTimeout, nil)
        }
    }
}

//BiddingView handle a single bidding request
func BiddingView(context *gin.Context) {
    var biddingRequest BiddingRequest
    if !util.BindRequestFromJson(context, &biddingRequest) {
        return
    }
    response, err := Bidding(biddingRequest)
    if err != nil {
        context.JSON(http.StatusInternalServerError, CommonResponse{
            APIStatus: false,
        })
    } else {
        context.JSON(http.StatusOK, response)
    }
    return
}
