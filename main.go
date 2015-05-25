package main

import (
    "runtime"
    "encoding/json"
    "github.com/garyburd/redigo/redis"
    "github.com/gographics/imagick/imagick"
)

var workers = make(chan int, MAXWORKER)

func processImage(messageData map[string]interface{} ) {
    // max worker process at the same time
    workers <- 1
    defer func(){ 
      if err:=recover();err!=nil{
        logger.Error(err)
      }
      logger.Info("processImage finished")
       <- workers
    }()

    logger.Info("I am in processImage now")
   
    originalImg := messageData["originalImg"].(string)
    destinalImg := messageData["destinalImg"].(string)
    resizeParam := messageData["resize"].(map[string]interface{})

    mw := imagick.NewMagickWand()
    defer mw.Destroy()
    
    err := mw.ReadImage(originalImg)
    if err != nil {
      panic(err)
    }

    if resizeParam != nil{
      hWidth := uint ( resizeParam["width"].(float64) )
      hHeight := uint( resizeParam["height"].(float64) )
      // Resize the image using the Lanczos filter
      // The blur factor is a float, where > 1 is blurry, < 1 is sharp
      err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
      if err != nil {
        panic(err)
      }
    }

    // Set the compression quality to 95 (high quality = low compression)
    err = mw.SetImageCompressionQuality(95)
    if err != nil {
      panic(err)
    }

    err = mw.WriteImage(destinalImg)
    if err != nil {
      panic(err)
    }
}

func main() {
    logger.Info(REDIS_CONFIG.Host)
    logger.Info(RESOURCES["image_process_channel"])
    runtime.GOMAXPROCS(runtime.NumCPU())
    imagick.Initialize()
    defer imagick.Terminate()

    c, err := redis.Dial("tcp", REDIS_CONFIG.Host)
    if err != nil {
        // logger.Error("redis.Dial error %v", err)
        panic(err)
        return
    }

    psc := redis.PubSubConn{c}
    psc.PSubscribe(RESOURCES["image_process_channel"])
    for {
        switch v := psc.Receive().(type) {
        case redis.PMessage:
            var messageData = map[string]interface{}{}

            json.Unmarshal(v.Data, &messageData)
            go processImage(messageData)
            //spawn a goroutine to process
            // logger.Info(messageData["nsp"].(string), messageData["room"].(string), content)

        case error:
            return
        }
    }

}