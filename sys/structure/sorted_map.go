package structure

type Entry struct {
	Key   string
	Value any
}

type SortedMap struct {
	Entries []Entry
}

func (sm *SortedMap) Len() int {
	return len(sm.Entries)
}

func (sm *SortedMap) Less(i, j int) bool {
	return sm.Entries[i].Key < sm.Entries[j].Key
}

func (sm *SortedMap) Swap(i, j int) {
	sm.Entries[i], sm.Entries[j] = sm.Entries[j], sm.Entries[i]
}

func (sm *SortedMap) Append(key string, value any) {
	entry := Entry{Key: key, Value: value}
	sm.Entries = append(sm.Entries, entry)
}

func (sm *SortedMap) Set(key string, value any) (any, bool) {
	for idx, entry := range sm.Entries {
		if entry.Key == key {
			sm.Entries[idx] = sm.Entries[idx]
			return entry.Value, true
		}
	}
	return "", false
}

func (sm *SortedMap) Get(key string) (any, bool) {
	for _, entry := range sm.Entries {
		if entry.Key == key {
			return entry.Value, true
		}
	}
	return "", false
}

func (sm *SortedMap) Remove(key string) {
	index := -1
	for i, entry := range sm.Entries {
		if entry.Key == key {
			index = i
			break
		}
	}
	if index != -1 {
		sm.Entries = append(sm.Entries[:index], sm.Entries[index+1:]...)
	}
}
