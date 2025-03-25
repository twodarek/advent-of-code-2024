INPUT_FILE="/Users/twodarek/personal/advent-of-code-2024/01/01.2/input.txt"

def readInput():
    listA = []
    mapB = {}

    f = open(INPUT_FILE, "r")

    for line in f:
        tmp = line.split()
        listA.append(int(tmp[0]))
        intTmpOne = int(tmp[1])
        if intTmpOne in mapB:
            mapB[intTmpOne] = mapB[intTmpOne] + 1
        else:
            mapB[intTmpOne] = 1

    f.close()
    return listA, mapB


def calcSimilarity(listA, mapB):
    resultant = 0
    for item in listA:
        if item in mapB:
            resultant += item*mapB[item]
    return resultant


def main():
    listA, mapB = readInput()
    similarity = calcSimilarity(listA, mapB)
    print("Similarity score: {0}".format(similarity))


if __name__ == "__main__":
    main()