# 编译成arm，可以在android上运行

CGO_ENABLED=0 GOARCH=arm GOOS=linux go build -o park main.go
<https://www.cnblogs.com/bigben0123/p/8398082.html>

## 编译 Win 桌面应用

set GOARCH=amd64&&set GOOS=windows&&go build -o server.exe -ldflags="-s -w -H=windowsgui" main.go

## 编译 Linux 二进制应用

## 注意：由于sqllite需要使用cgo，所以不支持交叉编译

set CGO_ENABLED=1&&set GOARCH=amd64&&set GOOS=linux&&go build -o server main.go
gomobile bind -ldflags "-s -w" -v -target=android/arm .
gomobile bind -v -target=android/arm .
go build -buildmode=c-shared -ldflags "-s -w" -o park.dll main.go
strip 工具来减少Android 的 so的大小

## 自动生成路由文件

bee generate routers

## 生成自动化文档

bee run -gendoc=true -downdoc=true

## 编译android包

cordova build android --release

## 签名apk

jarsigner -verbose -keystore testadmin.keystore -signedjar testadmin.apk testadmin_unsigned.apk testadmin
jarsigner -verbose -keystore testadmin.keystore -signedjar mark.apk app-release-unsigned.apk testadmin

cp testadmin.keystore platforms/android/app/build/outputs/apk/release/
cd platforms/android/app/build/outputs/apk/release/
jarsigner -verbose -storepass "123456" -keystore testadmin.keystore -signedjar mark-${BUILD_NUMBER}.apk app-release-unsigned.apk testadmin

## md语法参考

<https://blog.csdn.net/afei__/article/details/80717153>

## 代码统计

 find . "(" -name "*.js" -or -name "*.go" ")" -print | xargs wc -l


pipeline{
    agent none
    stages{
        stage("first"){
            agent{ label 'golang1.22.5' }
            steps {
                git credentialsId: '41064023-cd9e-4264-8402-ac8d4931e755', url: 'http://git.matrix-net.tech/paper/park.git'
                echo 'hello world'
                sh '''
                export GOPROXY=https://goproxy.io
                export PATH=$PATH:/usr/local/go/bin/
                export GO111MODULE=on
                go build -o park
                '''
                stash(name:'test',includes: 'Dockerfile,park,conf/,docker-compose-paper.yaml')
            }
        }
        stage('Package Docker Image') {
            agent{ label 'docker' }
            steps {
                script {
                    // 构建 Docker 镜像
                    unstash("test")
                    
                    sh '''
                    echo "Packaging Docker image..."
                    docker build -t registry.cn-zhangjiakou.aliyuncs.com/yuanyuexiang/park:V${BUILD_NUMBER} .
                    '''
                    withDockerRegistry(credentialsId: 'd05e283a-29be-47fb-81e6-e536652fca63', url: 'https://registry.cn-zhangjiakou.aliyuncs.com') {
                        sh '''
                        echo "Pushing Docker image..."
                        docker push registry.cn-zhangjiakou.aliyuncs.com/yuanyuexiang/park:V${BUILD_NUMBER}
                        '''
                    } 
                    
                    withDockerServer([credentialsId: '59aee53f-fcd7-4f48-a53e-e2e32191c251',uri: "tcp://47.92.146.162:2376"]) {
                         sh "ls -a"
                         //sh "docker-compose -f docker-compose-paper.yaml up -d"
                         sh "docker stop service-park"
                         sh "docker rm service-park"
                         sh "docker-compose -f docker-compose-paper.yaml up service-park -d"
                         sh "docker-compose ls"
                    }
                }
            }
        }
    }
}