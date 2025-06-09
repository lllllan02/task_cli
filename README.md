# Task CLI

这是一个简单的命令行任务管理工具，用于跟踪和管理您的待办事项。本项目实现了 [roadmap.sh 的任务追踪项目](https://roadmap.sh/projects/task-tracker) 指定的功能。

## 功能特点

- 添加、更新和删除任务
- 标记任务状态（进行中、已完成）
- 列出所有任务
- 按状态筛选任务
- 使用 JSON 文件持久化存储任务数据

## 使用方法

### 编译

```bash
go build -o task_cli
```

### 命令说明

1. 添加任务：
```bash
./task_cli add "买菜"
```

2. 更新任务：
```bash
./task_cli update 1 "买菜做饭"
```

3. 删除任务：
```bash
./task_cli delete 1
```

4. 标记任务状态：
```bash
./task_cli mark-in-progress 1  # 标记为进行中
./task_cli mark-done 1         # 标记为已完成
```

5. 列出任务：
```bash
./task_cli list           # 列出所有任务
./task_cli list todo      # 列出待办任务
./task_cli list done      # 列出已完成任务
./task_cli list in-progress  # 列出进行中任务
```

## 数据存储

任务数据以 JSON 格式存储在 `tasks.json` 文件中。每个任务包含以下属性：

- id: 任务唯一标识符
- name: 任务描述
- status: 任务状态（todo/progress/done）
- createdAt: 创建时间
- updatedAt: 更新时间
- deletedAt: 删除时间（如果已删除）

## 技术说明

- 使用 Go 语言开发
- 不依赖任何外部库
- 使用原生文件系统操作
- 支持中英文混合输出对齐 