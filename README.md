<img src="gisvs-gopher.png" height="300" />

# Gisvs 

Gisvs  is an HTTP service for file(image) uploading, processing, serving, and storage service integrate with IPFS . 
Gicvs supports chunked, resumable, and concurrent uploads. it uses Libvips behind the scenes making it memory efficient.

gisvs use go bolt as  databases for storing file meta-data and analytics. Currently,

## Features
 - upload files easily via our Web API.
 - upload large files by chunking them and uploading them to our Web API.
 - IPFS as backend storage service.
 - Image resize, convert, corp, blind watermask, and verification.


## Install
Libvips must be installed on your machine. 

### Ubuntu/Debian
```bash
sudo apt install libvips libvips-dev libvips-tools
```
### MacOS
```bash
brew install vips
```
For other systems check out instructions [here](https://github.com/libvips/libvips/wiki#building-and-installing).

Installing Gisvs server.
```bash
go get -u github.com/agrefox/gisvs/...
```
This will install the `gisvs` command in your `$GOPATH/bin` folder.

## Usage
```bash
gisvsd -config=/path/to/config.toml
```
If no config is passed gisvs will look for a `gisvs.toml` file in the current directory.

## Large File Uploads and IPFS 
When dealing with large files, it is best to split the file into small chunks and upload each chunk separately. gisvs easily handles chunked uploads storing each chunk and then re-building the whole file. Once the whole file is re-built gisvs uploads the file to the application's storage engine.
To view an example upload response check out the [Web API](https://gisvs-api-docs.threeaccents.com/#req_649a25397026402b82397975292fbc4f)

Other benefits of chunking up files are the ability to resume uploads and uploading multiple chunks concurrently. gisvs handles both scenarios for you with ease.
## Image  Transformations 
Gisvs supports image file transformations via URL query params. Currently, the supported operations are:
 - Resize (width, height) `?width=100&height=100`
 - Smart Crop `?crop=true`
 - Flip `?flip=true`
 - Flop `?flop=true`
 - Zoom `?zoom=2`
 - Black and White `?bw=true`
 - Quality(JPEG), Compression(PNG) `?quality=100` `?compression=10`
 - Format conversion `format is based on the file extension. To transform a png to webp, just use the .webp extension.`
 - WaterMark with fft
 - verified WaterMark


## Config
Gisvs's is configured via a toml file. Here are toml [config examples]

Configuration options include:
 - ***db_engine:string***(default: bolt) The main database for  `bolt`.  Storage engine is default IPFS, Maybe will support more in there future.
 - **http**
    - ***port:int***(default: 6600) the port to run gisvs on.
    - ***https:boolean***(default: false) configures server to accept https requests.
    - ***ssl_cert_path:string*** path to ssl certificate. Only required if `https` is set to true.
    - ***ssl_key_path:string*** path to ssl key. Only required if `https` is set to true.
 - **security**
    - ***auth_token:string*** token for authenticating requests
    - ***aes_key:string*** key for use with AES-256 encryption. This is used to encrypt storage secrets.
 - **upload**
    - ***chunk_upload_dir:string***(default: ./data/chunks) directory for storing chunks while an upload is happening. Once an upload is completed, the chunks are deleted.
    - ***full_file_dir:string***(default: ./data/files) full_files are temp files used while building chunks or downloading files from the storage engine. These temp files are removed once the request is completed.
    - ***max_chunk_size:int64***(default: 10MB) max size of a file chunk in bytes.
    - ***max_file_size_upload:int64***(default: 50MB) max size of a file for a regular upload in bytes.
    - ***max_transform_file_size:int64***(default: 50MB) max size of a file that can be transformed in bytes.
 - **bolt(only used if `db_engine` is set to bolt**
    - ***dir:string***(default: ./data/gisvs/gisvs.db) directory for bolt db file.
 