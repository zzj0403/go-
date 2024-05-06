package facade

import "fmt"

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a apiImpl) Test() string {
	aRet := a.a.testA()
	bRet := a.b.testB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type AModuleAPI interface {
	testA() string
}
type aModuleImpl struct {
}

func (i *aModuleImpl) testA() string {
	return "A module running"
}

type BModuleAPI interface {
	testB() string
}

// NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

type bModuleImpl struct {
}

func (i *bModuleImpl) testB() string {
	return "B module running"
}
