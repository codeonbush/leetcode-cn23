import os

files = os.listdir('.')

for file in files:
    items = file.rsplit('.', 1)
    if len(items) != 2:
        continue
    name, extension = items
    if extension != 'go':
        continue
    os.mkdir(name)
    os.rename(file, f'{name}/solution_test.go')
