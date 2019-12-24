 #1 环境变量
 export PKG_CONFIG_PATH=/usr/lib/pkgconfig
 export LD_LIBRARY_PATH=/usr/local/lib64/oracle/12.2/client64
 
 #2 编辑文件 oci8.pc
 
 prefix=/usr/local
 includedir=${prefix}/include/oracle/12.2/client64
 libdir=${prefix}/lib64/oracle/12.2/client64
 
 Name: oci8
 Description: Oracle Instant Client
 Version: 12.2
 Cflags: -I${includedir}
 Libs: -L${libdir} -lclntsh


文件说明
/usr/local 此路径为 以下官方下载依赖包的存放路径

#3 依赖项
官网下载 
instantclient-basiclite-linux.x64-12.2.0.1.0.zip
instantclient-sdk-linux.x64-12.2.0.1.0.zip

https://www.oracle.com/database/technologies/instant-client/linux-x86-64-downloads.html

#终端执行以下命令
unzip -ojd /usr/lib64  instantclient-basiclite-linux.x64-12.2.0.1.0.zip
ln -snf /usr/lib64/libclntsh.so.12.1 /usr/lib64/libclntsh.so
unzip -ojd /usr/include  instantclient-sdk-linux.x64-12.2.0.1.0.zip

#4编译错误
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1

#执行以下命令
export CGO_ENABLED=0


#run 选项


-L /usr/local/lib64/oracle/12.2/client64 -lclntsh -I /usr/local/lib64/oracle/12.2/client64
