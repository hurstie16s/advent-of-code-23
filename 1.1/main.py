numbers = [
    ("one","1"),
    ("two","2"),
    ("three","3"),
    ("four","4"),
    ("five","5"),
    ("six","6"),
    ("seven","7"),
    ("eight","8"),
    ("nine","9")
]
def main():
    with open("input.txt", "r") as input:
        inputData = input.read().split()
    
    for i in range(len(inputData)):
        for (string,number) in numbers:
            inputData[i] = inputData[i].replace(string, number)

    inputData = list(map(removeLetters, inputData))
    #print(inputData)
    inputData = list(map(lambda n: int(n[0]+n[len(n)-1]), inputData))
    print(inputData)
    
    print("Result: " + str(sum(inputData)))


def removeLetters(entry):
    chars = list(entry)
    newChars = [c for c in chars if c.isnumeric()]
    return ''.join(newChars)

if __name__ == "__main__":
    main()