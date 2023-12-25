package cache

type Cache struct {
	storage map[string]interface{}
}

type ICache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Delete(key string) bool
}

// Get get value and flag, if value not exist return nil, false
func (c Cache) Get(key string) (interface{}, bool) {
	if value, flag := c.storage[key]; flag {
		return value, flag
	}
	return nil, false
}

// Set value to Storage
func (c Cache) Set(key string, value interface{}) {
	c.storage[key] = value
}

// Delete value from Storage if value not in storage return false
func (c Cache) Delete(key string) bool {
	if _, ok := c.storage[key]; ok {
		delete(c.storage, key)
		return true
	}
	return false
}
