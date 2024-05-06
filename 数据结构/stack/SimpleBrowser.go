package stack

import "fmt"

type Browser struct {
	forwardStack Stack
	backStack    Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewArrayStack(),
		backStack:    NewLinkedListStack(),
	}
}

func (this *Browser) PushBack(add string) {
	this.backStack.Push(add)
}

func (this *Browser) CanBack() bool {
	if this.backStack.IsEmpty() {
		return false

	}
	return true
}

func (this *Browser) Back() {
	if this.backStack.IsEmpty() {
		return
	}
	top := this.backStack.Pop()
	this.forwardStack.Push(top)
	fmt.Printf("back to %+v\n", top)
}

func (this *Browser) CanForward() bool {
	if this.forwardStack.IsEmpty() {
		return false
	}
	return true
}

func (this *Browser) Forward() {
	if this.forwardStack.IsEmpty() {
		return
	}
	top := this.forwardStack.Pop()
	this.backStack.Push(top)
	fmt.Printf("forward to %+v\n", top)
}

func (this *Browser) Open(url string) {
	fmt.Printf("open new url %+v\n", url)
	this.forwardStack.Flush()
	this.backStack.Push(url)
}
