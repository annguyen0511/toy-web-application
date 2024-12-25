<template>
  <div class="book-details-container">
    <div class="book-card">
      <div class="book-image-section">
        <img :src="product.image_url" alt="Product cover" class="book-image" />
      </div>
      <div class="book-info-section">
        <h2 class="book-title">{{ product.product_name }}</h2>
        <p class="book-author"><strong>Tồn kho:</strong> {{ product.stock }}</p>
       
        <p class="book-price"><strong>Giá:</strong> {{ formatPrice(product.price) }} đ</p>
        <button class="add-to-cart" @click="addToCart">Thêm vào giỏ hàng</button>
        <button class="add-to-wishlist" @click="addToWishlist">Thêm vào danh sách yêu thích</button>
      </div>
    </div>
    <div class="book-description-section">
      <h3>Mô tả</h3>
      <p>{{ product.description }}</p>
    </div>
    <div class="related-products-section">
      <h3>Sản phẩm liên quan</h3>
      <div class="related-products-list">
        <div
          v-for="relatedProduct in relatedProducts"
          :key="relatedProduct.id"
          class="related-book-card"
          @click="viewProductDetails(relatedProduct.id)"
        >
          <img :src="relatedProduct.image_url" alt="Product cover" class="related-book-image" />
          <h4>{{ relatedProduct.product_name }}</h4>
          <p class="related-book-price">{{ formatPrice(relatedProduct.price) }} đ</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getProductById, getAllProducts } from '../apis/productsApi';
import { addToWishlist } from '../apis/wishlistsApi';
import { eventBus } from '../eventBus';

export default {
  name: 'ProductDetails',
  data() {
    return {
      product: {},
      relatedProducts: [],
    };
  },
  async created() {
    const productId = this.$route.params.id;
    this.product = await getProductById(productId);
    const allProducts = await getAllProducts();
    this.relatedProducts = this.getRandomProducts(allProducts, 5);
  },
  methods: {
    formatPrice(value) {
      return new Intl.NumberFormat('vi-VN').format(value);
    },
    getRandomProducts(products, count) {
      const shuffled = products.sort(() => 0.5 - Math.random());
      return shuffled.slice(0, count);
    },
    viewProductDetails(productId) {
      this.$router.push({ name: 'ProductDetails', params: { id: productId } });
    },
    addToCart() {
      let cart = JSON.parse(localStorage.getItem('cart')) || [];
      const existingItem = cart.find(item => item.id === this.product.id);
      if (existingItem) {
        existingItem.quantity += 1;
      } else {
        cart.push({ ...this.product, quantity: 1 });
      }
      localStorage.setItem('cart', JSON.stringify(cart));
      eventBus.updateCartCount(cart.length);
    },
    async addToWishlist() {
      try {
        const user = JSON.parse(localStorage.getItem('user'));
        if (!user || !user.id) {
          alert('Vui lòng đăng nhập để thêm vào danh sách yêu thích.');
          return;
        }
        const response = await addToWishlist({ productId: this.product.id, userId: user.id });
        if (response.success) {
          alert('Đã thêm vào danh sách yêu thích!');
        } else if (response.message === 'Item already in wishlist') {
          alert('Sản phẩm đã có trong danh sách yêu thích.');
        } else {
          alert('Có lỗi xảy ra khi thêm vào danh sách yêu thích.');
        }
      } catch (error) {
        console.error('Error adding to wishlist:', error);
        alert('Có lỗi xảy ra khi thêm vào danh sách yêu thích.');
      }
    },
  },
};
</script>

<style scoped>
.book-details-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px;
  background-color: #f0f4f8;
}

.book-card {
  display: flex;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
  overflow: hidden;
  max-width: 900px;
  width: 100%;
  margin-bottom: 20px;
}

.book-image-section {
  flex: 1;
  padding: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #e9ecef;
}

.book-image {
  max-width: 90%;
  max-height: 400px;
  border-radius: 8px;
  object-fit: cover;
}

.book-info-section {
  flex: 2;
  padding: 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  background-color: #fff;
  border-left: 1px solid #ddd;
}

.book-title {
  font-size: 1.8em;
  margin-bottom: 10px;
  color: #2c3e50;
  font-weight: bold;
}

.book-info-section p {
  margin: 8px 0;
  color: #555;
  line-height: 1.6;
}

.book-price {
  font-size: 1.5em;
  color: #e74c3c;
  font-weight: bold;
  margin-top: 10px;
}

.add-to-cart, .add-to-wishlist {
  margin-top: 20px;
  padding: 12px 25px;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 1em;
  transition: background-color 0.3s, transform 0.2s;
}

.add-to-cart:hover, .add-to-wishlist:hover {
  background-color: #2980b9;
  transform: translateY(-2px);
}

.add-to-cart:active, .add-to-wishlist:active {
  transform: translateY(0);
}

.book-description-section {
  max-width: 900px;
  width: 100%;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
  padding: 20px;
  margin-top: 20px;
}

.book-description-section h3 {
  font-size: 1.5em;
  margin-bottom: 10px;
  color: #2c3e50;
}

.book-description-section p {
  color: #555;
  line-height: 1.6;
}

.related-products-section {
  max-width: 900px;
  width: 100%;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
  padding: 20px;
  margin-top: 20px;
}

.related-products-section h3 {
  font-size: 1.5em;
  margin-bottom: 10px;
  color: #2c3e50;
}

.related-products-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 20px;
}

.related-book-card {
  text-align: center;
  background-color: #f9f9f9;
  border-radius: 8px;
  padding: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  transition: transform 0.2s;
}

.related-book-card:hover {
  transform: translateY(-5px);
}

.related-book-image {
  width: 100%;
  height: auto;
  border-radius: 5px;
  margin-bottom: 10px;
}

.related-book-price {
  color: #e74c3c;
  font-weight: bold;
}
</style>