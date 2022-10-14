package ddd

//
//import (
//	"github.com/google/uuid"
//	"time"
//	Cart "waffletime/internal/cart/domain"
//	Coupon "waffletime/internal/coupon/domain"
//	Point "waffletime/internal/point/domain"
//	Product "waffletime/internal/product/domain"
//	Review "waffletime/internal/review/domain"
//	User "waffletime/internal/user/domain"
//)
//
//type Domain interface {
//	User.User | Product.Product | Review.Review | Coupon.Coupon | Point.Point | Cart.Cart
//}
//
//type DomainMessage struct {
//	Message string
//}
//
//type Event[D Domain, M DomainMessage] struct {
//	ID         string
//	Domain     D
//	Message    M
//	OccurredAt time.Time
//}
//
//func NewDomainEvent(domain any, message DomainMessage) (Event, error) {
//	uuid, _ := uuid.NewUUID()
//	return &Event[Domain, DomainMessage]{ID: uuid, Domain: domain, Message: message, OccurredAt: time.Now()}, nil
//}
//
//func PublishEvent() {}
//
//func SubscribeEvent() {}
