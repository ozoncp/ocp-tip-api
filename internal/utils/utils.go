package utils

// SplitByBatches разделяет  слайс на батчи - исходный слайс конвертировать в слайс
// слайсов - чанки одинкового размера (кроме последнего)
func SplitByBatches(sourceSlice []int, batchSize int) [][]int {

	sourceLen := len(sourceSlice)
	fullBatchesCount := sourceLen / batchSize
	remainderBatchSize := sourceLen % batchSize
	splittedSlice := make([][]int, (sourceLen+batchSize-1)/batchSize)

	for i := 0; i < fullBatchesCount; i++ {
		splittedSlice[i] = sourceSlice[i*batchSize : (i+1)*batchSize]
	}

	if remainderBatchSize > 0 {
		splittedSlice[fullBatchesCount] = sourceSlice[sourceLen-remainderBatchSize:]
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

	mapFilter := make(map[int]bool)
	for _, value := range excludedValues {
		mapFilter[value] = true
	}

	for _, value := range sourceSlice {
		if !mapFilter[value] {
			filteredSlice = append(filteredSlice, value)
		}
	}

	return filteredSlice
}
