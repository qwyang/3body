# 3body 
    a project to place my codes and scripts.  

## rsa key
    ssh-keygen -t rsa -C '419139939@qq.com'

## basic config for vim
    vim ~/.vimrc:  
```
        filetype on
        filetype plugin on
        filetype indent on
        "使能鼠标
        "set mouse=a
        "设置编码
        set encoding=utf-8
        syntax enable
        "显示行号
        set nu
        "用空格替换tab，并且设置tab为4个空格
        set expandtab
        set ts=4
        "set autoindent
        set sw=4
        "set smartindent
        "高亮显示
        set cursorline
        hi cursorline guibg=#00ff00
        hi CursorColumn guibg=#00ff00
```
## basic config for git
```
    git config --global user.email "yangqunwei@huawei.com"
    git config --global user.name "qunwei"
    git config --global core.editor vim
    git config --global push.default simple
    git clone https://github.com/qwyang/pysh.git
    git remote add origin git@github.com:qwyang/3body.git
```
## useful help for git
    git config -l --global
    git ls-files
    git ls-files -d
    git ls-files -d | xargs git checkout
    git config core.eol crlf
    git config core.autocrlf true
    git archive --format tar.gz -o abc.tar.gz HEAD

## other information
    IDE for shell and py: pycharm
    Online Editor for markdown: http://mahua.jser.me/
    markdown introduction: http://blog.csdn.net/skykingf/article/details/45536231
