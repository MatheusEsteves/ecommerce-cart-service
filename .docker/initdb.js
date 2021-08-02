db.createUser(
    {
        user: "ecommerce_cart_user",
        pwd: "ecommerce_cart_pass",
        roles: [ { role: "readWrite", db: "ecommerce-cart-products"} ]
    }
);