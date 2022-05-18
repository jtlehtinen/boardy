<script lang='ts'>
  import FaShoppingCart from 'svelte-icons/fa/FaShoppingCart.svelte'

  export let title = '';
  export let img = '';
  export let id = '';

  function addToCart() {
    console.log('cart');
  }
</script>

<div class='card'>
  <div class='zoom'>
    <a href='/game/{id}'>
      <img class='object-contain' src='{img}' alt='@TODO'>
    </a>
    <p>{title}</p>

    <p class='add-to-cart' on:click={addToCart}>
      <span class='icon'><FaShoppingCart/></span> ADD TO CART
    </p>
  </div>
</div>

<style>
  .add-to-cart {
    background-color: #eee;
    padding: 5px;
    cursor: pointer;
  }

  .add-to-cart:hover {
    background-color: black;
    color: white;
  }

  img {
    width: 100%;
    height: 220px;
    border: 1px solid blue; /* DEBUG */
  }

  .object-contain {
    object-fit: contain;
    object-position: center;
  }

  .card {
    text-align: center;
    flex-basis: 25%;
    padding-left: 0.5em;
    padding-right: 0.5em;
    margin-bottom: 2em;

    border: 1px solid red; /* DEBUG */
    /* @TODO: box shadow */
  }

  .zoom {
    transition-duration: 0.2s;
    transition-timing-function: ease-in-out;
  }

  .zoom:hover {
    transform: scale(1.05, 1.05);
  }
</style>