package main

type URLStore struct {
	urls map[string]string
}

func (s *URLStore) Get(key string) string {
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) bool {
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count()) // generate the short URL
		if ok := s.Set(key, url); ok {
			return key
		}
	}
}
