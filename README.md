# s3 file fetch
This program can be used to check if you s3 buckets are set to public and if you can fetch the file with any authenticated AWS user

# to run the program
./s3FileFetchLinux -r `region-name` -b `s3-bucket` -k `key-name-for-the-bucket` -f `downloaded-file-name`

# pre req
aws config with access key and id configured