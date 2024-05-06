package abstractfactory

func getMainAndDetail(factor DAOFactory) {
	factor.CreateOrderMainDAO().SaveOrderMain()
	factor.CreateOrderDetailDAO().SaveOrderDetail()
}

func ExampleRdbFactory() {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
	// Output:
	// rdb main save
	// rdb detail save
}

func ExampleXmlFactory() {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
	// Output:
	// xml main save
	// xml detail save
}
