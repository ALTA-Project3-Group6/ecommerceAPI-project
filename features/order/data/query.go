package data

import (
	"ecommerceapi/config"
	cart "ecommerceapi/features/cart/data"
	"ecommerceapi/features/order"
	"ecommerceapi/features/orderproduct"
	product "ecommerceapi/features/product/data"
	"fmt"
	"log"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type orderQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) order.OrderData {
	return &orderQuery{
		db: db,
	}
}

func (oq *orderQuery) Add(userId uint, totalPrice float64) (order.Core, string, error) {
	tx := oq.db.Begin()

	// membuat order
	orderinput := Order{
		BuyerId:     userId,
		OrderStatus: "Waiting For Payment",
		CreatedAt:   time.Now(),
		TotalPrice:  totalPrice,
	}
	if err := tx.Create(&orderinput).Error; err != nil {
		tx.Rollback()
		log.Println("error add order query: ", err.Error())
		return order.Core{}, "", err
	}

	// update transactionid di database
	// tx.First(&orderinput)
	orderinput.TransactionId = "Transaction-" + fmt.Sprint(orderinput.ID)
	tx.Save(&orderinput)

	// mengambil cart user
	userCart := []cart.Cart{}
	if err := tx.Where("user_id = ?", userId).Find(&userCart).Error; err != nil {
		tx.Rollback()
		log.Println("error retrieve user cart: ", err.Error())
		return order.Core{}, "", err
	}

	// mengisi seller_id di order
	product := product.Product{}
	tx.First(&product, userCart[0].ProductId)
	orderinput.SellerId = product.UserId
	tx.Save(&orderinput)

	// membuat orderproduct
	orderProducts := []orderproduct.Core{}
	for _, item := range userCart {
		orderProduct := orderproduct.Core{
			OrderId:   orderinput.ID,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		orderProducts = append(orderProducts, orderProduct)
	}
	if err := oq.db.Create(&orderProducts).Error; err != nil {
		tx.Rollback()
		log.Println("error create orderproduct: ", err.Error())
		return order.Core{}, "", err
	}

	// menghapus cart user
	if err := tx.Where("user_id = ?", userId).Delete(cart.Cart{}).Error; err != nil {
		tx.Rollback()
		log.Println("error delete user cart: ", err.Error())
		return order.Core{}, "", err
	}

	// membuat pembayaran midtrans
	s := config.MidtransSnapClient()
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderinput.TransactionId,
			GrossAmt: int64(totalPrice),
		},
	}
	snapResp, _ := s.CreateTransaction(req)

	// commit tx transaksi
	tx.Commit()

	return DataToCore(orderinput), snapResp.RedirectURL, nil
}
func (oq *orderQuery) GetOrderHistory(userId uint) ([]order.Core, error) {
	orders := []order.Core{}

	err := oq.db.Raw("SELECT o.id , o.buyer_id , u.name buyer_name , o.seller_id , u2.name seller_name , total_price , o.created_at , order_status FROM orders o JOIN users u ON o.buyer_id = u.id JOIN users u2 ON o.seller_id = u.id WHERE o.buyer_id = ?", userId).Scan(&orders).Error
	if err != nil {
		log.Println("error query select order hisoty: ", err.Error())
		return []order.Core{}, err
	}

	return orders, nil
}
func (oq *orderQuery) GetSellingHistory(userId uint) ([]order.Core, error) {
	orders := []order.Core{}

	err := oq.db.Raw("SELECT o.id , o.buyer_id , u.name buyer_name , o.seller_id , u2.name seller_name , total_price , o.created_at , order_status FROM orders o JOIN users u ON o.buyer_id = u.id JOIN users u2 ON o.seller_id = u.id WHERE o.seller_id = ?", userId).Scan(&orders).Error
	if err != nil {
		log.Println("error query select order hisoty: ", err.Error())
		return []order.Core{}, err
	}

	return orders, nil
}
func (oq *orderQuery) GetTransactionStatus(orderId uint) (string, error) {
	return "", nil
}
