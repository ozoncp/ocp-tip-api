package utils

// SplitByBatches разделяет  слайс на батчи - исходный слайс конвертировать в слайс
// слайсов - чанки одинкового размера (кроме последнего)
func SplitByBatches(sourceSlice []int, batchSize int) [][]int {
	var splittedSlice [][]int

	sourceLen := len(sourceSlice)
	fullBatchesCount := sourceLen / batchSize
	remainderBatchSize := sourceLen % batchSize

	if remainderBatchSize > 0 {
		splittedSlice = make([][]int, fullBatchesCount+1)
		splittedSlice[fullBatchesCount] = sourceSlice[sourceLen-remainderBatchSize:]
	} else {
		splittedSlice = make([][]int, fullBatchesCount)
	}

	for i := 0; i < fullBatchesCount; i++ {
		splittedSlice[i] = sourceSlice[i*batchSize : (i+1)*batchSize]
	}

	return splittedSlice
}

// ReverseMap выполняет конвертацию отображения (“ключ-значение“) в отображение (“значение-ключ“)
func ReverseMap(sourceMap map[string]int) map[int]string {
	reversedMap := make(map[int]string)
	for key, value := range sourceMap {
		reversedMap[value] = key
	}
	return reversedMap
}

// FilterSlice фильтрует входной слайс по критерию вхождения элемента в другой слайс
func FilterSlice(sourceSlice []int, excludedValues []int) []int {
	var filteredSlice []int

	mapFilter := make(map[int]int)
	for _, value := range excludedValues {
		mapFilter[value] = 1
	}

	for _, value := range sourceSlice {
		if _, found := mapFilter[value]; !found {
			filteredSlice = append(filteredSlice, value)
		}
	}

	return filteredSlice
}
