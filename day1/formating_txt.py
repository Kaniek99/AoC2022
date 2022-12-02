arr = []
with open('./input.txt', 'r') as f:
    data = f.read()
    for line in f:
        print(line)
        arr.append(line.strip())

print(arr)
# print(data)
