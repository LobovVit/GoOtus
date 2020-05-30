package lesson06

import (
	"fmt"
	"testing"
)

func TestDLL(t *testing.T) {

	myList := List{}
	fmt.Println(myList.len)
	//myList.PushBack(1)
	myList.PushBack(2)
	fmt.Println(myList.len)
	myList.PushBack(3)
	myList.PushFirst(1)
	myList.PushBack(4)
	myList.PushBack(5)
	fmt.Println(myList.len)

	tmp := myList.First()
	tmp = tmp.next
	tmp = tmp.next
	fmt.Print(tmp.data)
	fmt.Print("___")
	fmt.Print(tmp.prev)
	fmt.Print("___")
	fmt.Println(tmp.next)

	tmp = tmp.next
	fmt.Print(tmp.data)
	fmt.Print("___")
	fmt.Print(tmp.prev)
	fmt.Print("___")
	fmt.Println(tmp.next)

	tmp = tmp.prev
	fmt.Print(tmp.data)
	fmt.Print("___")
	fmt.Print(tmp.prev)
	fmt.Print("___")
	fmt.Println(tmp.next)

	fmt.Print("________")
	fmt.Print(myList.len)
	fmt.Print("________")
	fmt.Print(*tmp)
	myList.Remove(*tmp)
	fmt.Print("________")
	fmt.Println(myList.len)

	tmp2 := myList.First()
	tmp2 = tmp2.next
	tmp2 = tmp2.next
	fmt.Print(tmp2.data)
	fmt.Print("___")
	fmt.Print(tmp2.prev)
	fmt.Print("___")
	fmt.Println(tmp2.next)

	tmp2 = tmp2.next
	fmt.Print(tmp2.data)
	fmt.Print("___")
	fmt.Print(tmp2.prev)
	fmt.Print("___")
	fmt.Println(tmp2.next)

	tmp2 = tmp2.prev
	fmt.Print(tmp2.data)
	fmt.Print("___")
	fmt.Print(tmp2.prev)
	fmt.Print("___")
	fmt.Println(tmp2.next)

}
