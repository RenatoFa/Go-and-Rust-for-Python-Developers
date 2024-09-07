from operator import itemgetter


poem = '''those who do not feel this love pulling
        them like a river those who do not drink dawn
        like a cup of sring water or take in sunset like
        supper those who do not want to change let them sleep
        '''


counts = {}
count = 0

for word in poem:
    if word.isspace():
        continue
    counts[word] = counts.get(word, 0) + 1
    count += 1

for word, frenquecy in sorted(counts.items(), key=itemgetter(1), reverse=True):
    frenquecy = frenquecy/count * 100
    print(f'{word}: {frenquecy}')
