# phprunner
linux下独立程序执行php文件,一个二进制文件,不依赖系统其他文件,可以解释执行php文件.目前只支持linux 64位系统.
# 一键安装
`curl -L https://raw.githubusercontent.com/snail007/phprunner/master/install.sh | bash`  
# 使用
phprunner filename.php  

# 注意事项

如果php文件本身有错误,比如语法错误,函数不存在等等,执行的时候是没有提示的,所以要确保脚本是调试没问题的再执行.  

# 源码使用
首先下载php源码，然后解压，比如解压路径：/home/php-7.0.8  
然后执行下面语句：  
export C_INCLUDE_PATH=/home/php-7.0.8/:/home/php-7.0.8/main/:/home/php-7.0.8/Zend/:/home/php-7.0.8/TSRM/  
然后go get才不会出错  
`go get gogithub.com/deuill/go-php`  
