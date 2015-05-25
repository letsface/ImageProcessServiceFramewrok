redis = require("redis")
client = redis.createClient()

parameter = 
  originalImg: "01.jpg"

for i in [1..200]
  parameter.destinalImg = "#{i}.jpg"
  parameter.resize = {width: i, height: i}
  client.publish("image_process_message#", JSON.stringify(parameter), (err, res) ->
    console.log err, res)
