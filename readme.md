
### Build App & Database

![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white)
![Cloudflare](https://img.shields.io/badge/Cloudflare-F38020?style=for-the-badge&logo=Cloudflare&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)
![Midtrans](https://img.shields.io/badge/Midtrans-FF6C37?style=for-the-badge&logo=midtrans&logoColor=white)

# YSHOP APP

This is a golang rest api project group organized by Alterra Academy. This API is used to run YSHOP applications. This application has features as below.


# Features
## User:
- Register
- Login
- Show profile
- Edit profile
- Deactive account

<div>

<details>

| Feature User | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| POST | /register | - | NO | This is how users register their account. |
| POST | /login  | - | NO | This is how users log in.  |
| GET | /users | - | YES | Users obtain their account information in this form. |
| PUT | /users | - | YES | This is how users Update their profile. |
| DELETE | /users | - | YES | This is how users Delete their profile. |

</details>

<div>

## Product :
- Add product
- Show all product
- Edit product
- Show detail product
- Delete product

<div>

<details>

| Feature Product | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| POST | /products | - | YES | This is how users add product in their account. |
| GET | /products  | - | NO | This is how all products show in homepage.  |
| PUT | /products | ID PRODUCT | YES | Users edit their product information in this form. |
| GET | /products | ID PRODUCT | NO | This is how users show detail product. |
| DELETE | /products | ID PRODUCT | YES | This is how users Delete their product. |

</details>

</div>

## Cart :
- Add product in cart
- Show all product in cart
- Delete all product in cart
- Edit or Update quantity product in cart

<div>

<details>

| Feature Cart | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| POST | /carts | - | YES | This is how users add product in their cart. |
| GET | /carts  | - | YES | This is how show all product in cart.  |
| DELETE | /carts | ID CART | YES | This is how users Delete their all products in cart. |
| PUT | /carts | ID CART | YES | Users edit their product quantity in cart. |

</details>

</div>


## Oder :
- Add order to payment gateway
- Show order history
- Show sales history
- Recieve payment notification

<div>

<details>

| Feature Cart | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| POST | /orders | - | YES | This is how users add orders to transaction. |
| GET | /orders  | - | YES | This is how users show order history.  |
| DELETE | /sales | - | YES | This is how seller sales history. |
| POST | /paymentnotification | - | - | Handling payment notification from midtrans. |

</details>

</div>


# ERD
<img src="image/ERD.png">

# API Documentations

[Click here](https://app.swaggerhub.com/apis-docs/icxz1/E-commerceAPI/1.0.0#/) to see documentations.


## How to Install To Your Local

- Clone it

```
$ git clone https://github.com/ALTA-Project3-Group6/ecommerceAPI-project
```

- Go to directory

```
$ cd ecommerceAPI-project
```

# UNIT TEST COVERAGE BY FEATURE

<div>
- USER
</div>
<div>
<img src="features/user/services/usercoverage.png">
</div>

<div>
- PRODUCT
</div>
<div>
<img src="features/product/services/productcoverage.png">
</div>

<div>
- CART
</div>
<div>
<img src="features/cart/services/cartcoverage.png">
</div>

<div>
- ORDER
</div>
<div>
<img src="features/order/services/ordercoverage.png">
</div>

# UNIT TEST COVERAGE ALL
<img src="image/unittest.png">

## Authors 👑

-   Muh Fauzan Putra  [![GitHub](https://img.shields.io/badge/fauzan-putra-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/mfauzanptra)

-  Alfian Aditya [![GitHub](https://img.shields.io/badge/alfian-aditya-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/icxz1)

 <p align="right">(<a href="#top">back to top</a>)</p>
<h3>
<p align="center">:copyright: January 2023 </p>
</h3>
<!-- end -->
<!-- comment -->
