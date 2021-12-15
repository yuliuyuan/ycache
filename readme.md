# 调用
get
curl localhost:8080/set -d 'key=yuliuyuan&value=guoyingying'
curl localhost:8080/get/yuliuyuan 
curl localhost:8080/del/yuliuyuan -X DELETE
curl localhost:8080/getstat