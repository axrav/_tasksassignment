package db

import (
	"sync"
)

// inserting book into db
func CreateBookinDb(b Book, mu *sync.Mutex, wg *sync.WaitGroup) (Book, error) {
	mu.Lock()         // locking the mutex to make sure that only one goroutine can access the db at a time
	defer mu.Unlock() // unlocking the mutex after the goroutine is done
	defer wg.Done()   // decrementing the waitgroup counter by 1 to make sure that the goroutine is done before the next one starts
	err := Pgres.Create(&b).Error
	if err != nil {
		return Book{}, err
	}

	return b, nil
}

// inserting movie into db
func CreateMovieinDb(m Movie, mu *sync.Mutex, wg *sync.WaitGroup) (Movie, error) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	err := Pgres.Create(&m).Error
	if err != nil {
		return Movie{}, err
	}
	return m, nil
}

// inserting music into db
func CreateMusicinDb(m Music, mu *sync.Mutex, wg *sync.WaitGroup) (Music, error) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	err := Pgres.Create(&m).Error
	if err != nil {
		return Music{}, err
	}
	return m, nil
}

// getting all books from db
func GetBooksfromDb(mu *sync.Mutex, wg *sync.WaitGroup) ([]Book, error) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var books []Book
	err := Pgres.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

// getting all movies from db
func GetMoviesfromDb(mu *sync.Mutex, wg *sync.WaitGroup) ([]Movie, error) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var movies []Movie
	err := Pgres.Find(&movies).Error
	if err != nil {
		return nil, err
	}
	return movies, nil
}

// getting all music from db
func GetAllMusicfromDb(mu *sync.Mutex, wg *sync.WaitGroup) ([]Music, error) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var music []Music
	err := Pgres.Find(&music).Error
	if err != nil {
		return nil, err
	}
	return music, nil
}

// getting book from db
func GetBookfromDb(id string, mu *sync.Mutex, wg *sync.WaitGroup) (Book, error) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var book Book
	err := Pgres.First(&book, id).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

// getting movie from db
func GetMoviefromDb(id string, mu *sync.Mutex, wg *sync.WaitGroup) (Movie, error) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var movie Movie
	err := Pgres.First(&movie, id).Error
	if err != nil {
		return movie, err
	}
	return movie, nil
}

// getting music from db
func GetMusicfromDb(id string, mu *sync.Mutex, wg *sync.WaitGroup) (Music, error) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var music Music
	err := Pgres.First(&music, id).Error
	if err != nil {
		return music, err
	}
	return music, nil
}

// deleting book in db
func DeleteBookfromDb(id string, mu *sync.Mutex, wg *sync.WaitGroup) error {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var book Book
	err := Pgres.First(&book, id).Error
	if err != nil {
		return err
	}
	err = Pgres.Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}

// deleting movie in db
func DeleteMoviefromDb(id string, mu *sync.Mutex, wg *sync.WaitGroup) error {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var movie Movie
	err := Pgres.First(&movie, id).Error
	if err != nil {
		return err
	}
	err = Pgres.Delete(&movie).Error
	if err != nil {
		return err
	}
	return nil
}

// deleting music in db
func DeleteMusicfromDb(id string, mu *sync.Mutex, wg *sync.WaitGroup) error {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var music Music
	err := Pgres.First(&music, id).Error
	if err != nil {
		return err
	}
	err = Pgres.Delete(&music).Error
	if err != nil {
		return err
	}
	return nil
}

// updating book in db
func UpdateBookinDb(id string, b Book, mu *sync.Mutex, wg *sync.WaitGroup) (Book, error) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var book Book
	err := Pgres.First(&book, id).Error
	if err != nil {
		return Book{}, err
	}
	err = Pgres.Model(&book).Updates(b).Error
	if err != nil {
		return Book{}, err
	}
	return b, nil
}

// updating movie in db
func UpdateMovieinDb(id string, m Movie, mu *sync.Mutex, wg *sync.WaitGroup) (Movie, error) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var movie Movie
	err := Pgres.First(&movie, id).Error
	if err != nil {
		return Movie{}, err
	}
	err = Pgres.Model(&movie).Updates(m).Error
	if err != nil {
		return Movie{}, err
	}
	return m, nil
}

// updating music in db
func UpdateMusicinDb(id string, m Music, mu *sync.Mutex, wg *sync.WaitGroup) (Music, error) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	var music Music
	err := Pgres.First(&music, id).Error
	if err != nil {
		return Music{}, err
	}
	err = Pgres.Model(&music).Updates(m).Error
	if err != nil {
		return Music{}, err
	}
	return m, nil
}
