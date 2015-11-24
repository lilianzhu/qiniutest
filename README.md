#qiniutest


### 配置编译环境

1. 安装go环境
2. make get
3. source env.sh
4. make test

### 创建＊_suit_test.go和*_test.go
1. cd src/qiniutest.com/ && mkdir <filename>
2. bootstrap --nodot && bootstrap --nodot <filename>

### 更新identifiers
ginkgo nodot

### 执行测试用例
1. cd qiniutest/ && source env.sh
2. cd src/qiniutest.com/biz/acc
3. ginkgo -v

### 正则匹配执行
- ginkgo --focus=<REGEXP> 只执行匹配REGEXP的case
- ginkgo --skip=<REGEXP> 不执行匹配REGEXP的case

### 并行执行
－ ginkgo -p 自动检测cpu数量
－ ginkgo -nodes=N 设置并行运行的node数量
－ 加 -stream 可查看每个node的处理情况

