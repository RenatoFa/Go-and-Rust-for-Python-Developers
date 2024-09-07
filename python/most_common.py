poem = '''those who do not feel this love pulling
        them like a river those who do not drink dawn
        like a cup of sring water or take in sunset like
        supper those who do not want to change let them sleep
        '''


frequency = {}

for word in poem.split():
    frequency[word] = frequency.get(word, 0) + 1


most_common = max(frequency, key=frequency.get)

print(frequency)
print(most_common)
