 if [ "$QBOXROOT" = "" ]; then
   QBOXROOT=$(cd ../; pwd)
   export QBOXROOT
 fi

export GOPATH=$GOPATH:$QBOXROOT/qiniutest:$QBOXROOT/qiniutest/src:$QBOXROOT/qiniutest/bin
export PATH=$PATH:$QBOXROOT/qiniutest:$QBOXROOT/qiniutest/bin

export TEST_ENV=product
export TEST_ZONE=z0
export QiniuTestEnv_product='{
"acchost":"https://acc.qbox.me",
"rshost":"http://rs.qiniu.com",
"rsfhost":"http://rsf.qbox.me",
"uchost":"http://uc.qbox.me",
"apihost":"http://api.qiniu.com",
"dochost":"http://docsync.qiniuapi.com",
"ufophost":"https://api.qiniu.com/v1/ufop",
"username":"general_storage_001@test.qiniu.io",
"password":"QA@7niu.com",
"access_key":"ByzN-V0c4TLdp_G04E_f27UQqapQTRVaAWC27Y2E",
"secret_key":"XRr4ikIQRze2530-EiApT_8xhV-VXqmhIcFCx1Nz",
"bucket_username":"general_storage_002@test.qiniu.io",
"bucket_password":"QA@7niu.com",
"bucketaccess_key":"HahrREkvH_bsbDTL8oAhDobiMQXcBUDXfH4RVGBv",
"bucketsecret_key":"RSQnjQGjl4RIiCCZ98hd0_VZMxpVcuWSZFMJGV88",
"quota_username":"general_storage_004@test.qiniu.io",
"quota_password":"QA@7niu.com",
"quota_access_key":"L7dkNKDbhyO7zrjs8IVMF6nb7x4Ya8gpdKpXLSvr",
"quota_secret_key":"-ni2s6h9T2vfqAbGsuROkF0OQoHwXnmSl1JTaDh2",
"aksk_username":"general_storage_003@test.qiniu.io",
"aksk_password":"QA@7niu.com",
"multizone_ak":"5WBnj2i9tMhtBXJYMd81l0kjRmnbfzUMOHFk0cFA",
"multizone_sk":"1NV14VwS84SWEZauU37C-yZFxZz_6kBTznnDsxzP"
}'
export QiniuTestEnv_product_z0='{
"iohost":"http://iovip.qbox.me",
"uphost":"http://up.qiniu.com"
}'