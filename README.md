# redp
[![Build Status](https://travis-ci.org/hawkingrei/redp.svg?branch=master)](https://travis-ci.org/hawkingrei/redp)
[![Go Report Card](https://goreportcard.com/badge/github.com/hawkingrei/redp)](https://goreportcard.com/report/github.com/hawkingrei/redp)


# RUN
```shell
make
./build/redp-cli -config example.toml
```

#  How to play

## query and create account
	
Signature:<Username>:MD5(Username)

 ```shell
 curl http://127.0.0.1:9000/api/user --header "Signature:wwz:e235ac07af7a969a52bec0985f6a9f85" -v
 
{"Uid":1,"Username":"wwz","Memory":0}%
 ```
 
 
## create hongbao

```shell
curl -X POST http://127.0.0.1:9000/api/hongbao --header "Signature:wwz:e235ac07af7a969a52bec0985f6a9f85" --header "money:10" --header "num:10" -v

{"Hbid":1,"Username":"wwz","Money":10,"Num":10,"Password":"QwSOkYoF","Closed":0,"CreateTime":"2017-12-01T12:22:44.285775111+08:00"}%
```

## grab hongbao 
```shell
curl http://127.0.0.1:9000/api/hongbao/1 --header "Signature:wz:d0965c07d1a00fcc85d28b8a241ae35a" --header "Password:QwSOkYoF" -v

{"Gothbid":1,"Hbid":1,"Username":"wz","Money":0.87}%
```

## List grabed hongbao
```shell
curl http://127.0.0.1:9000/api/hongbao --header "Signature:wz:d0965c07d1a00fcc85d28b8a241ae35a" -v

[{"Gothbid":21,"Hbid":3,"Username":"wz","Money":2.15},{"Gothbid":22,"Hbid":3,"Username":"wz","Money":0.87},{"Gothbid":23,"Hbid":3,"Username":"wz","Money":0.87}]%
```



