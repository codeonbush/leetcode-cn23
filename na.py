import os
import re
import shutil

# 文件夹路径
parent_folder_path = '/Users/chester.chen/Pro/workflow/leetcode-cn23'

# 列出父文件夹中的所有子文件夹
folders = os.listdir(parent_folder_path)

# 创建一个字典来保存规范化的文件夹名称
folder_dict = {}

# 正则表达式匹配两种类型的文件夹名
pattern_brackets = re.compile(r'\[(\d+)\](.+)')
pattern_no_brackets = re.compile(r'(\d+)\.(.+)')

for folder in folders:
    brackets_match = pattern_brackets.match(folder)
    no_brackets_match = pattern_no_brackets.match(folder)

    if brackets_match:
        number = brackets_match.group(1)
        name = brackets_match.group(2)
        normalized_name = f'[{number}]{name}'
        folder_dict[number] = normalized_name  # 使用有括号的文件夹名
    elif no_brackets_match:
        number = no_brackets_match.group(1)
        name = no_brackets_match.group(2)
        normalized_name = f'[{number}]{name}'
        if number in folder_dict:
            # 如果已存在相同编号的有括号文件夹，直接删除无括号的文件夹
            no_brackets_folder_path = os.path.join(parent_folder_path, folder)
            shutil.rmtree(no_brackets_folder_path)
        else:
            # 如果没有相同编号的有括号文件夹，将无括号的记录下来
            folder_dict[number] = normalized_name

# 重命名剩余文件夹为统一格式
for folder in folders:
    brackets_match = pattern_brackets.match(folder)
    no_brackets_match = pattern_no_brackets.match(folder)
    
    if brackets_match or no_brackets_match:
        number = brackets_match.group(1) if brackets_match else no_brackets_match.group(1)
        old_folder_path = os.path.join(parent_folder_path, folder)
        new_folder_path = os.path.join(parent_folder_path, folder_dict[number])
        if old_folder_path != new_folder_path:
            if not os.path.exists(new_folder_path):
                os.rename(old_folder_path, new_folder_path)

print("文件夹名合并和规范化完成。")

