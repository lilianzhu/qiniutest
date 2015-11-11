 if [ "$QBOXROOT" = "" ]; then
   QBOXROOT=$(cd ../; pwd)
   export QBOXROOT
 fi

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
"username":"liqian@qiniu.com",
"password":"happylq1989@736",
"AK":"oJ8ikIGwZmPkKqpFSF23YuY9dluxWhyZFcJ6XhZD",
"SK":"afDM6KaoIq2-6PqoU9GFDT6AOkepzVEfWtzsyDob",
"bucket_username":"general_storage_002@test.qiniu.io",
"bucket_password":"QA@7niu.com",
"bucketaccess_key":"HahrREkvH_bsbDTL8oAhDobiMQXcBUDXfH4RVGBv",
"bucketsecret_key":"RSQnjQGjl4RIiCCZ98hd0_VZMxpVcuWSZFMJGV88"
}'
