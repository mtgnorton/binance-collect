frontProjectPath="/Users/mtgnorton/Coding/vue/gf-admin-ui"
backendProjectPath="/Users/mtgnorton/Coding/go/src/github.com/mtgnorton/gf-admin"
packName="binance-collect"
#前端文件打包并且移动到后端目录中
cd $frontProjectPath || exit
npm run build:prod
rm -rf ${backendProjectPath}/public/static/*
rm -f ${backendProjectPath}/public/index.html
rm -f ${backendProjectPath}/public/favicon.ico

mv ${frontProjectPath}/dist/* ${backendProjectPath}/public/

# 打包后端代码
cd $backendProjectPath || exit
gf build main.go -y --name ${packName} --pack public,config -s linux -a amd64 -p ./bin #将静态资源文件和配置文件一起打包

scp -r $backendProjectPath/bin/linux_amd64/binance-collect  261:/www/wwwroot/binance-collect/

#将编译后的文件和配置文件上传到指定的服务器
#scp -r $backendProjectPath/config/casbin_model.conf 261:/www/wwwroot/binance-collect/config/
#scp -r $backendProjectPath/config/config-prod.toml 261:/www/wwwroot/binance-collect/config/
