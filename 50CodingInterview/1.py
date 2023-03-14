def Median(array1, array2):
    len1 = len(array1)
    len2 = len(array2)
    return (array1[len1//2] + array2[len2//2]) / 2

array1 = [1,3,5]
array2 = [2,4,6]
print(Median(array1, array2))