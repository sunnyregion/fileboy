package main

var exampleFileGirl string = `# 主配置
core:
    # 配置版本号
    version: 1

# 监控配置
monitor:
    # 要监听的目录
    # test1       监听当前目录下 test1 目录
    # test1/test2 监听当前目录下 test1/test2 目录
    # test1,*     监听当前目录下 test1 目录及其所有子目录（递归）
    # .,*         监听当前目录及其所有子目录（递归）
    includeDirs:
        - .,*

    # 不监听的目录
    # .idea   忽略.idea目录及其所有子目录的监听
    exceptDirs:
        - .idea
        - .git
        - .vscode

    # 监听文件的格式，此类文件更改会执行 commend 中的命令
    # .go   后缀为 .go 的文件更改，会执行 commend 中的命令
    # .*    所有的文件更改都会执行 commend 中的命令
    types:
        - .go

# 命令
command:
    # 监听的文件有更改会执行的命令
    # 可以有多条命令，会依次执行
    # 如有多条命令，每条命令都会等待上一条命令执行完毕后才会执行
    # 如遇交互式命令，允许外部获取输入
    exec:
        - go version
        - go env
`

var firstRunHelp = `第一次运行 fileboy ?
你可能需要先执行 fileboy init 生成配置。
更多信息查看帮助:
`

var helpStr = `fileboy [option]
Usage of fileboy:
    无参数 
        读取当前目录下的 filegirl.yaml 配置，开始监听并工作
    init 
        初始化 fileboy, 在当前目录生成 filegirl.yaml 配置文件
    exec 
        尝试运行定义的 commend 命令
`

var englishSay = []string{
	`      Have you, the darkness is no darkness.`,
	`    Why do the good girls always love bad boys?`,
	`              If love is not madness.`,
	`         This world is so lonely without you.`,
	`         You lie. Silence in front of me.`,
	`    I need him like I need the air to breathe.`,
	`  Happiness is when the desolated soul meets love.`,
	`   What I can lose, but do not want to lose you.`,
	`     The same words, both miss, is also missed.`,
	`  Each bathed in the love of the people is a poet.`,
}

var logo = `
 _______ _____ _       _______ ______   _____  _     _ 
(_______|_____) |     (_______|____  \ / ___ \| |   | |
 _____     _  | |      _____   ____)  ) |   | | |___| |
|  ___)   | | | |     |  ___) |  __  (| |   | |\_____/ 
| |      _| |_| |_____| |_____| |__)  ) |___| |  ___   
|_|     (_____)_______)_______)______/ \_____/  (___)
`
