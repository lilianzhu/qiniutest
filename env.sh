 if [ "$QBOXROOT" = "" ]; then
   QBOXROOT=$(cd ../; pwd)
   export QBOXROOT
 fi

export GOPATH=$GOPATH:$QBOXROOT/qiniutest:$QBOXROOT/qiniutest/src:$QBOXROOT/qiniutest/bin
export PATH=$PATH:$QBOXROOT/qiniutest:$QBOXROOT/qiniutest/bin

export TEST_ENV=product
export TEST_ZONE=z0