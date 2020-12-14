post /baas/admin/createApplication  创建一个application
get /baas/admin//getApplication/{id} 查看一个application
get /baas/admin/getApplications  获取所有application 清单
put /baas/admin/update/application/{id} 修改一个application
delete /baas/admin/deleteApplication/{id} 删除
get /baas/admin/application/{application_id}/usages 查看一个application 的使用
get /baas/admin/application/{application id}/users  查看一个application中所有用户


post /upload 上传一个文件到ipfs
post /chunk-upload 上传一个文件的一部分到 ipfs
post /chunks-completed 上传文件结束  

以上api必须和一个具有 ipfs 服务端相连接的 application 配合使用。 所以首先创建app with ipfs 并且创建ipfs服务，配置信息在baas.toml中体现

get /{application_name}/{date}/{file_name} 获取文件


example：
//创建一个用ipfs作为storage的application
curl "http://localhost:4200/applications" -H 'Content-Type: application/json' -X POST  -d '{
  "name": "moac", "description": "ipfs storage and other service", "storageEngine": "ipfs","storageAccessKey": "", "storageSecretKey": "", "storageRegion": "chn-sh", "deliveryUrl": "https://baas.moacchain.net:14200"}' 

upload file to server 
   curl "http://baas.moacchain.net:14200/upload"  -H 'Content-Type: multipart/form-data'  -X POST  -F 'application_id= d6313e0b-6fa9-4570-a7bf-8cce667ab988'  -F 'file_name=girl.jpg'   -F 'file=@girl.jpg'

access jpg url is 
    http://127.0.0.1:4200/moac/xxxxxxxxxxxxx/girl.jpg   x需替换为时间, moac 为application的 名称




  如果文件很大，可通过分段先传输，再合并的chunk方式，所以分为两步
   - chunk upload
    比如一个文件分成了3个部分，每个部分10M
    curl "http://localhost:4200/chunk-upload"  -H 'Content-Type: multipart/form-data'  -X POST   -F 'upload_id=任意字串(建议使用 uuids,如果是程序的话)'  -F 'chunk_number=1'  -F 'total_chunks=5'  -F 'total_file_size=20000000' -F  file_name=test.jpg'  -F 'data=@part1.bin' 
    其中解释一下，
    upload_id 为任意字串为了区分上载内容，属于一个文件的，用同一个id
    chunk_number 表示第几个部分，1表示，第一块，2表示第二块，依次
    total_chunks 表示这个文件总共几块
    total_file_size 表示该文件共有多少字节
    file_name 表示该文件名称
    data 表示os中的文件名

    然后上传文件的其余的部分，可同时上传，完成后进行completed

   - Completed process: merge chunk
   curl "http://localhost:4200/chunks-completed" -H 'Content-Type: application/json' -X POST  -d '{
  "uploadId": "8889998888", "fileName": "combine.mp4", "applicationId": "d6313e0b-6fa9-4570-a7bf-8cce667ab988"
}' 
   解释：
   uploadId  为上传chunk时创建的id，要保持一致
   filename是合并后文件名
   applicationsid  是用户应用的id，必须


   返回：
   {"data":{"id":"049873b3-9500-4885-8da5-f0ed317ca94e","fileName":"combine.mp4","fileBlobID":"moac/1595299171294756724/combine","size":60533146,"mimeType":"video","mimeValue":"video/mp4","extension":"mp4","url":"http://127.0.0.1:4200/moac/1595299171294756724/combine.mp4","width":0,"height":0,"createdAt":"2020-07-21T10:39:32.063102626+08:00","updatedAt":"2020-07-21T10:39:32.063102856+08:00","hash":"QmWUN4sXaVyex55avHxfUPKVZRJuV2eANmbPjEevSXTufD"}}

   -  使用 
        可以直接回放 mpv http://127.0.0.1:4200/moac/1595299171294756724/combine.mp4

