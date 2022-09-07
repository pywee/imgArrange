# imgArrange
文件整理，一键整理所有文件

直接将配置文件 config.json 和 imgArrange.exe 文件复制到要整理文件的文件夹中，双击运行即可。

通过修改配置文件 config.json 来决定要整理哪些文件，指定某些文件创建到某个文件夹中。
配置文件示例：

```json
{
    "dirName": "相片",
    "suffix": ["jpg", "jpeg", "png", "gif"],
    "ymd": 0
}
```

表示将当前文件夹中带有 "jpg", "jpeg", "png", "gif" 后缀的文件移动到名为【相片】的文件夹中，如果该文件夹不存在，则自动创建。