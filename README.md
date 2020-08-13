<img src="baas.png" height="300" />

# BAAS block as a service  

### curl -XPOST http://127.0.0.1:14200/applications -H "Content-Type=application/json" -d '{
> "name":"智慧农业",
> "description":"区块链系统BAAS服务",
> "storageEngine":"bolt",
> "storageAccessKey":"mmaskdjfasdf",
> "storageSecretKey":"masdfsafdsafdsaf",
> "storageRegion":"XiongAN",
> "deliveryUrl":"http://127.0.0.1:4200/api"
> }'
{"data":{"id":"954348da-948c-4bb8-86c2-d9a53ef395ff","name":"智慧农业","description":"区块链系统BAAS服务","storageAccessKey":"mmaskdjfasdf","storageBucket":"","storageEndpoint":"","storageRegion":"XiongAN","storageEngine":"bolt","cdnUrl":"http://127.0.0.1:4200/api","createdAt":"2020-08-13T14:54:38.776070197+08:00","updatedAt":"2020-08-13T14:54:38.776070317+08:00"}}


### curl -XPOST http://127.0.0.1:14200 -d '{
> "application_id":"954348da-948c-4bb8-86c2-d9a53ef395ff",
> "userid":"jihua.gao@gmail.com",
> "password":"qweasdzxc",
> "name":"西方大熊"
> }'


{"status":0,"data":{"id":"d212ca92-310a-47ad-a014-7767d8d5a8a2","email":"jihua.gao@gmail.com","name":"","application_id":"954348da-948c-4bb8-86c2-d9a53ef395ff","description":"","address":"0x360af35eac2e66420905a2e2f451301b912ca663","createdAt":"2020-08-13T14:57:11.068651697+08:00","updatedAt":"2020-08-13T14:57:11.068652169+08:00"}}


### curl -XPOST http://127.0.0.1:14200/api/baas/getKey -d '{
> "userid":"jihua.gao@gmail.com",
> "password":"qweasdzxc",
> "application_id":"954348da-948c-4bb8-86c2-d9a53ef395ff"
> }'
{"data":{"key":"c72a8cf72de549944f1df31bb3f1f19d6dd65fac88e76b3e79a0c877fe12e5ef"},"status":1}


### curl -XPOST http://127.0.0.1:14200/api/baas/writeMsg -d '{
"userId":"jihua.gao@gmail.com",
"toUserId":"jihua.gao@gmail.com",
"message":"大马哈上线了",
"password":"qweasdzxc"
}'
{"data":{"txhash":"0xe4a5f20b6645aafc8bcf016976cd175681239e9638504cd58b72481878dd7915"},"error":"","status":1}



### curl -XPOST http://127.0.0.1:14200/api/baas/readMsg -d '{
"hash":"0xe4a5f20b6645aafc8bcf016976cd175681239e9638504cd58b72481878dd7915"
}'
{"data":{"timestamp":"2020-08-13 15:09:59 +0800 CST","txhash":"大马哈上线了"},"error":"","status":1}
