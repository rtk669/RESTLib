package lib

import "sync"

type Library struct {
	libMap map[string]*Book
	mtx    sync.Mutex
}

func NewLibrary() *Library {
	return &Library{
		libMap: make(map[string]*Book),
		mtx:    sync.Mutex{},
	}
}

func (l *Library) AddBook(title, author string, pageAmount int) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.libMap[title]; ok {
		return Book{}, ErrBookAlreadyExist
	}

	book, err := NewBook(title, author, pageAmount)
	if err != nil {
		return Book{}, err
	}

	l.libMap[title] = book

	return *book, nil
}

func (l *Library) CompleteBook(title string) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.libMap[title]; !ok {
		return Book{}, ErrBookNotFound
	}
	if l.libMap[title].Completed {
		return Book{}, ErrBookAreadyCompleted
	}

	l.libMap[title].CompleteBook()

	book := l.libMap[title]

	return *book, nil
}

func (l *Library) GetBook(title string) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.libMap[title]; !ok {
		return Book{}, ErrBookNotFound
	}

	book := l.libMap[title]

	return *book, nil
}

func (l *Library) GetBooks(
	authorExist, completedExist bool,
	author string, completed bool,
) map[string]Book {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	copyLibMap := make(map[string]Book)

	// switch {
	// case authorExist && completedExist:
	// 	{
	// 		copyLibMapFiltered := make(map[string]Book)

	// 		for k, v := range copyLibMap {
	// 			if copyLibMap[k].Author == author && copyLibMap[k].Completed == completed {
	// 				copyLibMapFiltered[k] = v
	// 			}
	// 		}
	// 		return copyLibMapFiltered
	// 	}

	// case authorExist:
	// 	{
	// 		copyLibMapFiltered := make(map[string]Book)

	// 		for k, v := range copyLibMap {
	// 			if copyLibMap[k].Author == author {
	// 				copyLibMapFiltered[k] = v
	// 			}
	// 		}
	// 		return copyLibMapFiltered
	// 	}

	// case completedExist:
	// 	{
	// 		copyLibMapFiltered := make(map[string]Book)

	// 		for k, v := range copyLibMap {
	// 			if copyLibMap[k].Completed == completed {
	// 				copyLibMapFiltered[k] = v
	// 			}
	// 		}
	// 		return copyLibMapFiltered
	// 	}
	// default:
	// 	return copyLibMap
	// }

	for k, v := range l.libMap {
		if authorExist && v.Author != author {
			continue
		}

		if completedExist && v.Completed != completed {
			continue
		}

		copyLibMap[k] = *v
	}

	return copyLibMap

}

func (l *Library) DeleteBook(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.libMap[title]; !ok {
		return ErrBookNotFound
	}

	delete(l.libMap, title)
	return nil
}
