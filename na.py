import os
import re

def rename_directory(old_name):
    """Remove brackets and replace them with an underscore."""
    # 使用正则表达式匹配并替换
    new_name = re.sub(r'^\[(\d+)\]', r'\1_', old_name)
    return new_name

def rename_directories(root_dir):
    """Rename all directories in the root directory that match the pattern."""
    for dirpath, dirnames, filenames in os.walk(root_dir):
        for dirname in dirnames:
            if re.match(r'^\[\d+\]', dirname):
                new_name = rename_directory(dirname)
                old_path = os.path.join(dirpath, dirname)
                new_path = os.path.join(dirpath, new_name)
                os.rename(old_path, new_path)
                print(f'Renamed: {old_path} -> {new_path}')

if __name__ == "__main__":
    root_directory = "."  # 设置为要处理的根目录路径
    rename_directories(root_directory)
