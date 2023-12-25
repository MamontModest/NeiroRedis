package heap

type Heap struct {
	//Сами объекты
	Objects []*Object
	//Индексы объектов
	IndexObject map[string]int
}

// Object объекты, которые мы будем хранить в куче
type Object struct {
	//Ключ
	Key string
	//Время, после которого объект должен само удалиться
	Exp int
}

type IHeap interface {
	Push(o *Object)
	Pop() (*Object, bool)
	GetLastItem() (*Object, bool)
	ChangeObject(new *Object)
}

// Push добавляет объект в кучу
func (h *Heap) Push(o *Object) {
	if len(h.Objects) == 0 {
		h.IndexObject[o.Key] = 0
		h.Objects = append(h.Objects, o)
		return
	}
	h.Objects = append(h.Objects, o)
	indexNewObject := len(h.Objects) - 1
	h.IndexObject[o.Key] = indexNewObject
	h.swapUp(indexNewObject)
	return
}

// Pop удаляет объект из кучи
func (h *Heap) Pop() (*Object, bool) {
	if len(h.Objects) == 0 {
		return nil, false
	}
	//swap 0 element with lsat
	h.Objects[0], h.Objects[len(h.Objects)-1] = h.Objects[len(h.Objects)-1], h.Objects[0]

	//delete last element
	lastObject := h.Objects[len(h.Objects)-1]
	h.Objects = h.Objects[:len(h.Objects)-1]
	delete(h.IndexObject, lastObject.Key)

	h.swapDown(0)

	return lastObject, true
}

// GetLastItem Получаем последний элемент
func (h *Heap) GetLastItem() (*Object, bool) {
	if len(h.Objects) == 0 {
		return nil, false
	}
	return h.Objects[0], true
}

func (h *Heap) ChangeObject(new *Object) {
	if index, ok := h.IndexObject[new.Key]; ok {
		h.Objects[index].Exp = new.Exp
		h.swapUp(index)
		h.swapDown(h.IndexObject[new.Key])
	}
}

func (h *Heap) swapUp(indexDown int) {
	for indexDown != 0 && h.Objects[(indexDown-1)/2].Exp > h.Objects[indexDown].Exp {
		indexUp := (indexDown - 1) / 2
		h.Objects[indexUp], h.Objects[indexDown] = h.Objects[indexDown], h.Objects[indexUp]
		//swap indexes
		h.IndexObject[h.Objects[indexUp].Key], h.IndexObject[h.Objects[indexDown].Key] = indexUp, indexDown
		indexDown = indexUp
	}
}

func (h *Heap) swapDown(indexUp int) {
	for indexUp < (len(h.IndexObject)-1)/2 {
		index1, index2 := indexUp*2+1, indexUp*2+2
		if index2 < len(h.Objects) {
			if h.Objects[index1].Exp < h.Objects[index2].Exp {
				if h.Objects[index1].Exp < h.Objects[indexUp].Exp {
					h.Objects[indexUp], h.Objects[index1] = h.Objects[index1], h.Objects[indexUp]
					//swap indexes
					h.IndexObject[h.Objects[indexUp].Key], h.IndexObject[h.Objects[index1].Key] = indexUp, index1
					indexUp = index1
				} else {
					break
				}
			} else {
				if h.Objects[index2].Exp < h.Objects[indexUp].Exp {
					h.Objects[indexUp], h.Objects[index2] = h.Objects[index2], h.Objects[indexUp]
					//swap indexes
					h.IndexObject[h.Objects[indexUp].Key], h.IndexObject[h.Objects[index2].Key] = indexUp, index2
					indexUp = index2
				} else {
					break
				}
			}
		} else if index1 < len(h.Objects) {
			if h.Objects[index1].Exp < h.Objects[indexUp].Exp {
				h.Objects[indexUp], h.Objects[index1] = h.Objects[index1], h.Objects[indexUp]
				//swap indexes
				h.IndexObject[h.Objects[indexUp].Key], h.IndexObject[h.Objects[index1].Key] = indexUp, index1
				indexUp = index1
			} else {
				break
			}
		}
	}
}
