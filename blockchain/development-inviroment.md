# 2018/12/05:install ethereum development invironment.
## install virtualbox on windows
    360安装管理器
## install gola
## install jetbrains goland
## dowload ubuntu16.04 ISO
    ubuntu-gnome-16.04-desktop-amd64.iso
## install ubuntu virtual machine
    memory:2G,hard disk:50G 
## install virtualbox enhancement package for vm
    实现host和vm之间复制粘贴功能
    虚拟机启动后，设备->拖放->双向
## vim编辑器设置
    set ts=4
    set expandtab    
## 输入法设置
    设置->区域和语言->设置输入源语言添加汉语
    设置->键盘->设置快捷键CTRL+SPACE切换语言
## 安装git工具，github新建blcokchain仓库
	git init
	git add README.md
	git commit -m "first commit"
	git remote add origin git@github.com:qwyang/blockchain.git
	git push -u origin master
## 虚拟机安装geth客户端
	sudo apt-get install software-properties-common
	sudo add-apt-repository -y ppa:ethereum/ethereum
	sudo apt-get update
	sudo apt-get install ethereum
## 安装nodejs
	curl -sL https://deb.nodesource.com/setup_10.x | sudo -E bash -
	sudo apt-get install -y nodejs
	dpkg -L nodejs | grep /usr/bin
	sudo npm install n -g
    sudo npm list -g --depth=0
    dpkg -S /bin/ping
## 虚拟机安装remix-ide编译器(遇到权限问题，npm的bug，需要选择版本）
	git clone https://github.com/ethereum/remix-ide.git
	cd remix-ide
	npm install
	npm start
### ***按照上述方式安装remix-ide失败的解决办法***  
    sudo apt-get install docker.io
    sudo docker pull 4c0n/remix-ide 
    docker run -itd -p8080:8080 -p65520:65520 -v /home/qwyang/solidity_files/:/app 4c0n/remix-ide
## 安装chrome浏览器(firefox打不开remix-ide)
    download corresponding deb
    sudo dpkg -i google-chrome-stable_current_amd64.deb
    sudo apt-get -f -y install 
    sudo dpkg -i google-chrome-stable_current_amd64.deb
