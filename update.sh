#!/bin/bash

# 版本号
VERSION="1.0.5"
SERVER_IP="您的服务器IP"
SERVER_PATH="/www/wwwroot/redamancy/backend/updates"

# 构建项目
echo "正在构建项目..."
quasar build -m capacitor -T android

# 进入构建目录
cd dist/capacitor/android

# 创建更新包
echo "正在创建更新包..."
zip -r bundle-${VERSION}.zip .

# 计算 SHA256
echo "正在计算 SHA256..."
SHA256=$(sha256sum bundle-${VERSION}.zip | awk '{print $1}')

# 创建 manifest.json
echo "正在创建 manifest.json..."
cat > manifest.json << EOF
{
    "version": "${VERSION}",
    "bundleUrl": "https://redamancy.com.cn/bundles/bundle-${VERSION}.zip",
    "sha256": "${SHA256}"
}
EOF

# 上传文件
echo "正在上传文件..."
scp manifest.json root@${SERVER_IP}:${SERVER_PATH}/
scp bundle-${VERSION}.zip root@${SERVER_IP}:${SERVER_PATH}/bundles/

# 设置权限
echo "正在设置权限..."
ssh root@${SERVER_IP} "chown -R www:www ${SERVER_PATH} && chmod -R 755 ${SERVER_PATH}"

echo "更新完成！"
echo "版本: ${VERSION}"
echo "SHA256: ${SHA256}" 