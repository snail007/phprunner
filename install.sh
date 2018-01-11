#!/bin/bash
wget  https://github.com/snail007/phprunner/releases/download/v1.0/dist.tar.xz
tar jxf dist.tar.xz
mv dist /usr/local/phprunner
ln -s /usr/local/phprunner/phprunner /usr/bin/
echo "install done"
echo "usage:"
phprunner
