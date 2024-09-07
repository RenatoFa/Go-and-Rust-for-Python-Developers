message = 'Hi :)'

width = 8
padding = (width - len(message)) // 2

print(' ' * padding, end='')
print(message)
print('-' * width)
