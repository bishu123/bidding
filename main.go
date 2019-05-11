package main

import (
    "bidding/core"
    "bidding/util"
    "github.com/gin-gonic/gin"
)

func main(){
    ginEngine := gin.Default()
    core.InitData()
    util.InitHttpClient(10)
    InitRouterPattern(ginEngine)
    ginEngine.Run("localhost:9091")
}

func InitRouterPattern(ginEngine *gin.Engine)  {
    //Core Apis
    ginEngine.GET("/auction", core.AuctionView)
    ginEngine.POST("/bidding", core.BiddingView)

}

