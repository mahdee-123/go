package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"unsafe"
)

// 1. একটা Book struct বানাও — Title (string), Author (string), Pages (int), Price (float64)। দুটো book তৈরি করো এবং সব fields print করো।

type Book struct {
	title string
	author string
	pages int
	price float64
}


// 2. উপরের Book struct এর জন্য একটা NewBook() constructor function লেখো যেটা সব field validate করে। Pages ≤ 0 বা Price < 0 হলে error return করবে। 

func NewBook(title string, author string, pages int , price float64) (newBook Book,err error)  {

	if pages <=0  && price < 0 {
		err = errors.New("book is invalid!")
		return
	}

	newBook = Book {
		title : title,
		author:  author,
		pages : pages,
		price : price ,
	} 

	return 
	
}

// 3. Book-এ দুটো method লেখো: Describe() string (value receiver) এবং ApplyDiscount(pct float64) (pointer receiver)। পার্থক্য বোঝো — কোনটা original change করে, কোনটা করে না।

func (b Book) Describe() string {
		b.author = "new name"
		return fmt.Sprintf("the book named %s written by %s which hase  %d pages and cost is %0.2f", b.title, b.author, b.pages, b.price)
}




func (b *Book) ApplyDiscount(pct float64) float64 {
	b.price = b.price - b.price*pct  
	return b.price 
}


// 5. একটা function লেখো updateAge(p Person, age int) যেটা Person এর age change করার চেষ্টা করে। তারপর দেখাও যে original Person unchanged — কারণ ব্যাখ্যা করো comment-এ।
type Person struct {
	Name string
	Age int
}
// func updateAge(p Person, age int) {
// 	p.Age = age
// }



// 6. উপরের problem টা fix করো। updateAge কে *Person নিতে দাও। এখন original change হবে। দুটো version পাশাপাশি রেখে compare করো।

func updateAge(p *Person, age int) {
	p.Age = age
}


// 7. একটা struct বানাও যেখানে bool, int64, bool আছে। আরেকটা বানাও int64, bool, bool দিয়ে। unsafe.Sizeof() দিয়ে দুটোর size print করো। কেন size আলাদা?

type struct1 struct {
	e1 bool 
	e2 int64
	e3 bool
}

type struct2 struct {
	e1 int64 
	e2 bool
	e3 bool
}


// struct embedding
// 8. Address struct বানাও (City, Country)। Person struct-এ Address embed করো। দেখাও যে p.City সরাসরি access করা যায় — p.Address.City না লিখেও।


type Address struct {
	City string 
	Country string 
}

type Student struct {
	Name string 
	Address
}

// 9. Animal struct বানাও Breathe() method সহ। Dog struct-এ Animal embed করো এবং Bark() method যোগ করো। দেখাও যে dog.Breathe() এবং dog.Bark() দুটোই কাজ করে।

type Animal struct {
	name string
}

func (a Animal) Breathe() {
	fmt.Println("breathing...")
} 
type Dog struct {
	Animal
}
func (d Dog) Bark() {
	fmt.Println("barking...")
} 
// 10. দুটো version লেখো — একটায় Engine struct কে Car-এ embed করো, আরেকটায় named field হিসেবে রাখো (Engine Engine)। দেখাও কোথায় syntax আলাদা হয়। 

type Engine struct {
	name string
}

type CarWithEmbedding struct {
	Engine
}

type CarWithNameFeild struct {
	Engine Engine
}
func (c Engine ) Start() {
	fmt.Println("car started...")
}
// 11. একটা User struct বানাও JSON tags সহ — id, name, email, created_at। encoding/json দিয়ে JSON marshal করো এবং দেখাও Go field names এর বদলে tag names আসছে।

// 12. User struct-এ optional field যোগ করো (Bio string) যেটা empty হলে JSON-এ আসবে না। আরেকটা field (Password) যেটা কখনো JSON-এ যাবে না।
type User struct {
	ID        		int    		`json:"id"`
	Name      		string 		`json:"name"`
	Email     		string 		`json:"email"`
	Password 		string 		`json:"-"`
	Bio 				string 		`json:"bio,omitempty"`
	CreatedAt 		time.Time 		`json:"created_at"`
}

/// 14. একটা Server struct বানাও (Host, Port, Timeout)। NewServer() function লেখো যেটা functional options pattern follow করে — caller যেটুকু দরকার সেটুকুই set করবে, বাকিটা default।


type Option func(*Server)
type Server struct {
	Host string
	Port int
	Timeout time.Duration
}

func withHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}



func main() {
	// 1.... 
	book1, err1:= NewBook("go programming", "arnob", 100, 100.0)
	book2 , err2 := NewBook("ts programming", "john", 200, 300.0)

  // 2.. 
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(book1)
	}

	if err2 != nil {
		fmt.Println(err2)
	} else {
			fmt.Println(book2) 
	}


	fmt.Println(book1)
	fmt.Println(book1.Describe())
	fmt.Println(book1)
	book1.ApplyDiscount(0.5)
	fmt.Println(book1)



	
	// 4.   Variable declare না করে সরাসরি একটা anonymous struct বানাও যেটায় config data রাখবে (host string, port int)। Real use case: quick one-off data grouping

	configData := struct{
		host string
		port int
	}{
		host: "localhost",
		port: 8080,
	}

	fmt.Println(configData)
	
	/// 5....+ 6....
	person1 := Person{Name: "Arnob", Age: 20}
	updateAge(&person1, 40)
	fmt.Println(person1)
	// doesn't change

	// 7...
	// 7...


	fmt.Println(unsafe.Sizeof(struct1{}))
	fmt.Println(unsafe.Sizeof(struct2{}))



	// 8. .....
	student1 := Student {
		Name : "mahdee hossasin arnob",
		Address: Address{
			City: "Sylhet",
			Country:  "Bangladesh",
		},
	}

	fmt.Println(student1.Address.City)
	fmt.Println(student1.City)

	// 9......
	dog1 := Dog{
		Animal: Animal{
			name: "tommy",
		},
	}

	dog1.Breathe()
	dog1.Bark() 



	// 10......
	car1 := CarWithEmbedding{
		Engine: Engine{
			name: "v8",
		},
	} 

	car2 := CarWithNameFeild{
		Engine: Engine{
			name: "v8",
		},
	}

	fmt.Println(car1.Engine.name)
	fmt.Println(car2.Engine.name)
	fmt.Println(car1.name)
	// fmt.Println(car2.name) // doesn't work

	car1.Start()
	car2.Engine.Start()


	// 11. 

	u := User{
		ID: 1,
		Name: "ar",
		Email: "ar",
		CreatedAt: time.Now(),
	} 

 
	jsonOfU , _ := json.MarshalIndent(u, "", "  ")

	fmt.Println(string(jsonOfU))

	
	var u1 User

	jsonData := `{"id":"10","name":24,"email":"12@gmail.com", "created_at":"2020-01-01T00:00:00Z"}`

	json.Unmarshal([]byte(jsonData), &u )
	fmt.Println(u1)
	


	
	//

}