package order

func (s *service) New(bookId, userId int64) (orderId int64, err error) {
	panic("implement me")
}

func (s *service) UpdateOrderState(orderId int64, state int) (err error) {
	panic("implement me")
}
