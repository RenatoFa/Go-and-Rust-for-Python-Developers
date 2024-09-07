count, total = 0, 0

for n in range(1, 101):
    if n % 3 == 0 or n % 5 == 0:
        count += 1
        total += n

print(total/count)
