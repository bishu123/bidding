package util

import (
    "net/http"
    "time"
)

type HttpClientType struct {
    timeout time.Duration
}

var HttpClient HttpClientType

func InitHttpClient(timeout int){
    HttpClient = HttpClientType{
        timeout: time.Duration(timeout) *time.Second,
    }
}

// Create HttpClient
func (this *HttpClientType)GetHttpClient()*http.Client  {
    return  &http.Client{
        Timeout: HttpClient.timeout,
    }
}