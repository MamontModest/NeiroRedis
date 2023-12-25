package cache

type Cache struct {
	Storage map[string]interface{}
}

type ICache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Delete(key string) bool
	Update(key string, value interface{})
	Len() int
}

// Get get value and flag, if value not exist return nil, false
func (c Cache) Get(key string) (interface{}, bool) {
	if value, flag := c.Storage[key]; flag {
		return value, flag
	}
	return nil, false
}

// Set value to Storage
func (c Cache) Set(key string, value interface{}) {
	c.Storage[key] = value
}

// Delete value from Storage if value not in storage return false
func (c Cache) Delete(key string) bool {
	if _, ok := c.Storage[key]; ok {
		delete(c.Storage, key)
		return true
	}
	return false
}

// Update value to Storage
func (c Cache) Update(key string, value interface{}) {
	c.Storage[key] = value
}

func (c Cache) Len() int {
	return len(c.Storage)
}
