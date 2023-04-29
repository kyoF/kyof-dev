環境構築方法

Dockerイメージの作成
```
# apache-vimディレクトリ内で実行
$ docker build -t apache-vim-image .
```

Dockerコンテナの起動
```
# apache-vimディレクトリ内で実行
$ docker run -dit --name apache-vim -p 80:80 -v $PWD/src:/usr/local/apache2/htdocs -v $PWD/apache/httpd.conf:/usr/local/apache2/conf/httpd.conf apache-vim-image
```
