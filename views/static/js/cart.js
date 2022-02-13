document.addEventListener("DOMContentLoaded", ready);

function ready() {
    var addToCartButtons = document.getElementsByClassName("add-cart");
    for (var i = 0; i < addToCartButtons.length; i++) {
        var button = addToCartButtons[i];
        button.addEventListener('click', addToCartClicked);
    }

    var quantityInputs = document.getElementsByClassName("input-number");
    for (var i = 0; i < quantityInputs.length; i++) {
        var input = quantityInputs[i];
        input.addEventListener('change', quantityChanged);
    }

    var removeCartItemButtons = document.getElementsByClassName("remove-icon");
    for (var i = 0; i < removeCartItemButtons.length; i++) {
        var input = removeCartItemButtons[i];
        input.addEventListener('click', removeCartItem);
    }
}

function onLoadNumbers() {
    var productNumbers = localStorage.getItem("cartNumbers");

    if (productNumbers) {
        document.querySelector(".single-icon span").textContent = productNumbers;
    }
}

function cartNumbers() {
    var productNumbers = localStorage.getItem("cartNumbers");
    productNumbers = parseInt(productNumbers);

    if (productNumbers) {
        localStorage.setItem("cartNumbers", productNumbers + 1);
        document.querySelector(".single-icon span").textContent = productNumbers + 1;
    } else {
        localStorage.setItem("cartNumbers", 1);
        document.querySelector(".single-icon span").textContent = 1;
    }
}

function addToCartClicked(event) {
    cartNumbers();
    var button = event.target;
    var shopItem = button.parentElement.parentElement.parentElement.parentElement;
    var title = shopItem.querySelector(".product-name").innerText;
    var price = shopItem.querySelector(".product-price").innerText;
    var image = shopItem.querySelector(".hover-img").src;
    var itemsList = {
        Title: title,
        Price: price,
        Image: image,
    }

    let cartItems = localStorage.getItem("productsInCart");
    cartItems = cartItems ? JSON.parse(cartItems) : {};
    cartItems[itemsList.Title] = itemsList;
    localStorage.setItem("productsInCart", JSON.stringify(cartItems));
    cartItems = localStorage.getItem("productsInCart");
    cartItems = JSON.parse(cartItems);
}

function quantityChanged(event) {
    var input = event.target;
    if (isNaN(input.value) || input.value <= 0) {
        input.value = 1;
    }
    updateCartTotal();
}

function removeCartItem(event) {
    var buttonClicked = event.target;
    buttonClicked.parentElement.parentElement.parentElement.remove();
    var productName = buttonClicked.parentElement.parentElement.parentElement.querySelector(".product-name").innerText;
    const cartItems = JSON.parse(localStorage.productsInCart || "{}")
    if (cartItems.hasOwnProperty(productName)) {
        delete cartItems[productName];
        localStorage.productsInCart = JSON.stringify(cartItems);
    }
    updateCartTotal();
}

function updateCartTotal() {
    var cartItemContainer = document.querySelector(".cart-items");
    var cartRows = cartItemContainer.querySelector(".cart-row");
    var cartRows = cartItemContainer.getElementsByClassName("cart-row");
    var subTotal = 0;
    for(var i=0; i < cartRows.length; i++) {
        var cartRow = cartRows[i]; 
        var priceElement = cartRow.querySelector(".price");
        var quantityElement = cartRow.querySelector(".input-number");
        var price = parseInt(priceElement.innerText.replace("$", ""));
        var quantity = quantityElement.value;
        subTotal = subTotal + (price * quantity);
        localStorage.setItem("subTotal", subTotal);
    }
    document.querySelector(".subTotal").innerText = "$" + subTotal + ".00";
    document.querySelector(".totalPrice").innerText = "$" + (subTotal + 20) + ".00";    
}

onLoadNumbers();