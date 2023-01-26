package data

import (
	"ecommerceapi/config"
	cart "ecommerceapi/features/cart/data"
	"ecommerceapi/features/order"
	product "ecommerceapi/features/product/data"
	"errors"
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

	// mengambil cart user
	userCart := []cart.Cart{}
	if err := tx.Where("user_id = ?", userId).Find(&userCart).Error; err != nil {
		tx.Rollback()
		log.Println("error retrieve user cart: ", err.Error())
		return order.Core{}, "", err
	}

	// membuat order
	orderinput := Order{
		BuyerId:     userId,
		OrderStatus: "Waiting For Payment",
		CreatedAt:   time.Now(),
		TotalPrice:  totalPrice,
	}
	// mengisi seller_id di order
	product := product.Product{}
	tx.First(&product, userCart[0].ProductID)
	orderinput.SellerId = product.UserId
	// tx.Save(&orderinput)
	//input order ke tabel
	if err := tx.Create(&orderinput).Error; err != nil {
		tx.Rollback()
		log.Println("error add order query: ", err.Error())
		return order.Core{}, "", err
	}

	// update transactionid di database
	// tx.First(&orderinput)
	orderinput.TransactionId = "Transaction-" + fmt.Sprint(orderinput.ID)
	tx.Save(&orderinput)

	// membuat orderproduct
	orderProducts := []OrderProduct{}
	for _, item := range userCart {
		orderProduct := OrderProduct{
			OrderId:   orderinput.ID,
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		orderProducts = append(orderProducts, orderProduct)
	}
	if err := tx.Create(&orderProducts).Error; err != nil {
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
	snapResp, err := s.CreateTransaction(req)
	if err != nil {
		tx.Rollback()
		log.Println("error making midtrans transaction: ", err.Error())
		return order.Core{}, "", err
	}

	// commit tx transaksi
	tx.Commit()

	return DataToCore(orderinput), snapResp.RedirectURL, nil
}
func (oq *orderQuery) GetOrderHistory(userId uint) ([]order.Core, error) {
	orders := []order.Core{}

	err := oq.db.Raw("SELECT o.id , o.buyer_id , u.name buyer_name , o.seller_id , u2.name seller_name , total_price , o.created_at , order_status FROM orders o JOIN users u ON o.buyer_id = u.id JOIN users u2 ON o.seller_id = u2.id WHERE o.buyer_id = ?", userId).Scan(&orders).Error
	if err != nil {
		log.Println("error query select order hisoty: ", err.Error())
		return []order.Core{}, err
	}

	return orders, nil
}
func (oq *orderQuery) GetSellingHistory(userId uint) ([]order.Core, error) {
	orders := []order.Core{}

	err := oq.db.Raw("SELECT o.id , o.buyer_id , u.name buyer_name , o.seller_id , u2.name seller_name , total_price , o.created_at , order_status FROM orders o JOIN users u ON o.buyer_id = u.id JOIN users u2 ON o.seller_id = u2.id WHERE o.seller_id = ?", userId).Scan(&orders).Error
	if err != nil {
		log.Println("error query select order hisoty: ", err.Error())
		return []order.Core{}, err
	}

	return orders, nil
}
func (oq *orderQuery) NotificationTransactionStatus(transactionId, transStatus string) error {
	order := Order{}

	oq.db.First(&order, "transaction_id = ?", transactionId)

	// 5. Do set transaction status based on response from check transaction status
	if transStatus == "capture" {
		if transStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			order.OrderStatus = "challenge"
		} else if transStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			order.OrderStatus = "success"
		}
	} else if transStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		order.OrderStatus = "success"
	} else if transStatus == "cancel" || transStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		order.OrderStatus = "failure"
	} else if transStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		order.OrderStatus = "waiting payment"
	} else {
		order.OrderStatus = transStatus
	}

	aff := oq.db.Save(&order)
	if aff.RowsAffected <= 0 {
		log.Println("error update order status, no rows affected")
		return errors.New("error update order status")
	}

	//update product stock
	if order.OrderStatus == "success" {
		orderProducts := []OrderProduct{}
		oq.db.Find(&orderProducts, "order_id = ?", order.ID)
		for _, item := range orderProducts {
			prod := product.Product{}
			oq.db.First(&prod, item.ProductId)
			prod.Stock -= item.Quantity
			oq.db.Save(&prod)
		}
	}

	return nil
}
