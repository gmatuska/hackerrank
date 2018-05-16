import sys


def count_words(filename, threshold):
    results = dict()
    with open(filename, 'r') as f:
        for line in f:
            for word in line.split():
                results[word] = results.setdefault(word, 0) + 1
    count = 0
    for word, wc in sorted(results.items(), key=lambda x: x[1], reverse=True):
        if count < threshold:
            print('{} - {}'.format(wc, word))
            count = count + 1
        else:
            break

count_words(sys.argv[1], sys.argv[2])
