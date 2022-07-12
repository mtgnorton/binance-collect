package main



gf build main.go -y --name  Artifacts_${PIPELINE_ID}  --pack public -s linux -a amd64 -p ./bin #将静态资源文件和配置文件一起打包
