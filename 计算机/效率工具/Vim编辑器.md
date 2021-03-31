**1.Vim编辑器简介**

Vim是从 vi 发展出来的一个文本编辑器。代码补完、编译及错误跳转等方便编程的功能特别丰富，在程序员中被广泛使用。Vim编辑器避免了鼠标和上下左右键的使用，当你熟练使用快捷键之后就会发现Vim编辑器非常顺手。

**2.Vim编辑器模式**

Vim编辑器一共有五种模式

- **Normal**: 正常模式，查看文件，编辑文件
- **Insert**: 插入模式，编辑文件
- **Replace**: 删除模式，可以删除文本for replacing text
- **Visual:** 选择模式，选择文本
- **Command-line**: 命令行模式，运行命令

最常用的normal insert和command-line模式之间的转化关系如下图所示，使用vim命令进入一个文档之后，首先进入normal模式，也就是图中的命令模式，使用i进入插入模式，用esc退出，在命令模式使用: 进入命令行模式，回车结束命令。

![img](https://pic4.zhimg.com/v2-5fa1ca0e113bd81b5493333a45bf7ad3_b.png)

**3.Vim编辑器快捷操作**

Vim不同模式下有不同的操作

大部分时间我们会在Normal模式下操作，键盘的快捷键如下所示

![img](https://pic1.zhimg.com/v2-4e1c2d3216d397c0885a00a10bbdad48_b.png)

- Basic movement: hjkl (left, down, up, right)
- Words: w (next word), b (beginning of word), e (end of word)
- Lines: 0 (beginning of line), ^ (first non-blank character), $ (end of line)
- Screen: H (top of screen), M (middle of screen), L (bottom of screen)
- Scroll: Ctrl-u (up), Ctrl-d (down)
- File: gg (beginning of file), G (end of file)
- Line numbers: :{number}<CR> or {number}G (line {number})
- Misc: % (corresponding item)
- Find: f{character}, t{character}, F{character}, T{character} 
- find/to forward/backward {character} on the current line
- , / ; for navigating matches 
- Search: /{regex}, n / N for navigating matches

命令行模式的参数如下所示：

- :q quit (close window)
- :w save (“write”)
- :wq save and quit
- :e {name of file} open file for editing
- :ls show open buffers
- :help {topic} open help     
- :help :w opens help for the :w command
- :help w opens help for the w movement

选择模式：

- Visual: v
- Visual Line: V
- Visual Block: Ctrl-v

编辑模式：

- i enter Insert mode     
- o / O insert line below / above
- d{motion} delete {motion}     
- e.g. dw is delete word, d$ is delete to end of line, d0 is delete  to beginning of line
- c{motion} change {motion}     
- e.g. cw is change word
- like d{motion} followed by i 
- x delete character (equal do dl)
- s substitute character (equal to xi)
- Visual mode + manipulation     
- select text, d to delete it or c to change it
- u to undo, <C-r> to redo
- y to copy / “yank” (some other commands like d also copy)
- p to paste
- Lots more to learn: e.g. ~ flips the case of a character 