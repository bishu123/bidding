package core

import (
    "bidding/util"
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

func Bid(channel chan BiddingResponse ,request BiddingRequest){
    url := "http://localhost:9091/bidding"
    var response *BiddingResponse
    reqByte,_ := ProcessRequest(request)
    req, _ := http.NewRequest("POST", url, bytes.NewReader(reqByte))
    req.Header.Add("Content-Type","application/json")
    client := util.HttpClient.GetHttpClient()
    resp, _ := client.Do(req)
    responseBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return
    }
    if err := GetUsableResponse(responseBytes, &response); err != nil {
        return
    }
    channel <- *response
    return
}

func Bidding(request BiddingRequest) (resp BiddingResponse, err error)  {
   if bidData, ok := AdData.AdData[request.AdSlot]; ok && bidData.Available {
       /*
           Fetching Data from data go for the requested slot.
           Also , The bidding price will go up by .5 price after every request.
           and updating the slot information with latest bid
        */
       resp.BidPrice = bidData.BidPrice
       AdData.RequestList = append(AdData.RequestList, request)
       resp.AdId = len(AdData.RequestList)
       resp.APIStatus = true
       bidData.BidPrice += .5
       bidData.Available = true
       bidData.AdPlacementId = request.AdPlacementId
   }
   return
}


func ProcessRequest(requestObj interface{}) ([]byte, error) {
    request, err := json.Marshal(requestObj)
    return request, err
}

func GetUsableResponse(responseBytes []byte, responseObj interface{}) error {
    return json.Unmarshal(responseBytes, &responseObj)
}

