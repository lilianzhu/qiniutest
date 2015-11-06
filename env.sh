# if [ "$QBOXROOT" = "" ]; then
#   QBOXROOT=$(cd ../; pwd)
#   export QBOXROOT
# fi

export GOPATH=$GOPATH:$QBOXROOT/qiniutest:$QBOXROOT/qiniutest/src:$QBOXROOT/qiniutest/bin
export PATH=$PATH:$QBOXROOT/qiniutest:$QBOXROOT/qiniutest/bin

export QiniuTestEnv=product
export QiniuTestEnv_product='{
"acchost":"acc.qbox.me",
"rshost":"rs.qiniu.com",
"rsfhost":"rsf.qbox.me",
"uchost":"uc.qbox.me",
"apihost":"api.qiniu.com",
"dochost":"docsync.qiniuapi.com",
"ufophost":"https://api.qiniu.com/v1/ufop",
"username":"",
"password":"",
"AK":"",
"SK":"",
"bucket_username":"",
"bucket_password":"",
"bucketaccess_key":"",
"bucketsecret_key":""
}'
