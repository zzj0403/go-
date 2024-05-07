package abstractfactory

import "testing"

func getMainAndDetail(factor DAOFactory) {
	factor.CreateOrderMainDAO().SaveOrderMain()
	factor.CreateOrderDetailDAO().SaveOrderDetail()
}

func Test_RdbFactory(t *testing.T) {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
	// Output:
	// rdb main save
	// rdb detail save
}

func Test_XmlFactory(t *testing.T) {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
	// Output:
	// xml main save
	// xml detail save
}
