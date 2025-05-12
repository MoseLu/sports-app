#!/bin/bash

# 创建密钥目录
mkdir -p ../config/keys

# 生成私钥
openssl genrsa -out ../config/keys/private.pem 2048

# 从私钥生成公钥
openssl rsa -in ../config/keys/private.pem -pubout -out ../config/keys/public.pem

# 设置适当的文件权限
chmod 600 ../config/keys/private.pem
chmod 644 ../config/keys/public.pem

echo "RSA 密钥对已生成" 