./smart migrate -a true -g true
 ./smart migrate -c config/settings.yml

生成迁移文件
go run main.go migrate -g true -c config/settings.yml

修改完迁移文件后，执行下面命令开始变更
go run main.go migrate -c config/settings.yml

go run main.go migrate -h # 帮助
go run main.go migrate -g true -a true  -c config/settings.dev.yml # 生成 go-admin系统预置 迁移文件
go run main.go migrate -g true -c config/settings.dev.yml  # 生成 自定义功能 迁移文件 自己开发新功能用这个功能
go run main.go migrate -c config/settings.dev.yml # 执行迁移命令 迁移 未迁移过的 文件


生成迁移文件
go run main.go migrate -g true -c config/settings.yml

修改完迁移文件后，执行下面命令开始变更
go run main.go migrate -c config/settings.yml

生成迁移文件--系统相关
go run main.go migrate -g true -a true  -c config/settings.yml
