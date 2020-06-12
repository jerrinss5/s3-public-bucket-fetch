env CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -a -o s3FileFetchLinux
env CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -a -o s3FileFetchMac

# to run the program
# ./s3FileFetchLinux -r <region-name> -b <s3-bucket> -k <key-name-for-the-bucket> -f <downloaded-file-name>