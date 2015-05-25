# ImageProcess
重构密集计算，使用Go

1. install imagick
homebrew install imagemagick --universal    to build fat binaries.

2. install go, prepare go dev env.

3. install go modules(third party modules in import)

4. go build *.go

5. run ./main , run test.js 

---------------------------------------------------------------------------------
1. main goroutine listens to redis message
2. Once a message comming, start a goroutine to process image
3. In config.go MAXWORKER to control the maxium workers that process at the same time
4. Current, only wirte resize action. 
5. If there is a scalable issue, static language should be adopted. 
