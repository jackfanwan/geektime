package delete_slice

type Slice[T any] struct {
	SliceArray []T
}

func (slice *Slice[T]) DeleteByIndex(index int) bool {
	length := len(slice.SliceArray)
	if length == 0 || length <= index {
		return false
	}
	if index == 0 {
		if length == 1 {
			slice.SliceArray = []T{}
		} else {
			slice.SliceArray = slice.SliceArray[index+1:]
		}
	} else if index == length-1 {
		if length == 1 {
			slice.SliceArray = []T{}
		} else {
			slice.SliceArray = slice.SliceArray[0:index]
			slice.reduceCap()
		}
	} else {
		slice.SliceArray = append(slice.SliceArray[0:index], slice.SliceArray[index+1:]...)
		slice.reduceCap()
	}
	return true
}

func (slice *Slice[T]) reduceCap() {
	// 缩容机制为，容量 = 实际长度 * 1.5时，启动缩容，并把容量缩为1.25倍
	realLength := len(slice.SliceArray)
	virLength := cap(slice.SliceArray)
	if virLength >= int(float64(realLength)*1.5) {
		length := int(float64(realLength) * 1.25)
		array := make([]T, 0, length)
		array = append(array, slice.SliceArray...)
		slice.SliceArray = array
	}
}
