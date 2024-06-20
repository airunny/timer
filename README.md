# WikiGlobal Golang微服务项目模板 

## 安装 Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## 创建微服务项目 
```
# 基于模板创建项目 
kratos new netwon -r http://git55.fxeyeinterface.com/public-projects/timer.git
cd netwon
kratos run

通过上面三条命令即可创建并运行一个微服务项目
```
## 项目结构说明1
```
新创建的项目是作为一个独立的微服务，项目的proto文件定义在api文件下即可
├── Makefile    用于项目中的快捷指令 
├── README.md
├── api         定义proto文件
├── cmd         项目的入口
├── internal    项目的逻辑
└── third_party 三方proto依赖
```

## 项目结构说明2
#### 目前整个公司层名将整体proto作为一个单独仓库独立了出去，如果想把项目的proto也拉进整体的proto则需要做如下处理
```
make proto  // 获取整体proto的依赖，执行改命令之后会在项目目录下多出来一个wiki_protos文件夹，内部是公司所有的protos文件
rm -rf ./third_parth // 删除项目中的 third_party 文件；因为在wiki_protos 已经添加了三方的依赖；所以不再需要该文件
```

## 构建镜像 
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```
