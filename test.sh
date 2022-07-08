frontProjectPath="/Users/mtgnorton/Coding/vue/gf-admin-ui"
backendProjectPath="/Users/mtgnorton/Coding/go/src/github.com/mtgnorton/gf-admin"
packName="binance-collect"

# 打包后端代码
cd $backendProjectPath || exit
gf build main.go -y --name ${packName} --pack public -s linux -a amd64 -p ./bin #将静态资源文件和配置文件一起打包

# 将编译后的文件和配置文件上传到指定的服务器
scp -r $backendProjectPath/bin/linux_amd64/binance-collect 261:/www/wwwroot/binance-collect/
