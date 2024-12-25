<template>
  <div class="books-by-category">
    <div v-for="category in categories" :key="category.category_id" class="category-section">
      <h2 v-if="getProductsByCategory(category.category_id).length > 0">{{ category.category_name }}</h2>
      <div class="books-list" v-if="getProductsByCategory(category.category_id).length > 0">
        <div v-for="(product) in getProductsByCategory(category.category_id).slice(0, 5)" :key="product.category_id" class="book-card" @click="viewProductDetails(product.category_id)">
          <img :src="product.image_url || 'https://cdn0.fahasa.com/media/catalog/product/8/9/8935246937136.jpg'" alt="Product cover" class="book-image" />
          <h3>{{ product.product_name }}</h3>
          <p>{{ product.description }}</p>
          <p class="price">{{ formatPrice(product.price) }} đ</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getProductsByCategory } from '../apis/productsApi';
import { getAllCategories } from '../apis/categoriesApi';

export default {
  name: 'ProductByCategory',
  data() {
    return {
      categories: [],
      booksByCategory: {},
    };
  },
  async created() {
    this.categories = await getAllCategories();
    await this.fetchAllProducts();
  },
  methods: {
    async fetchAllProducts() {
      for (const category of this.categories) {
        console.log({
          category_id: category.category_id,
          category_name: category.category_name,
          description: category.description,
          created_at: new Date().toISOString() // Assuming created_at is the current date for demonstration
        });
        this.booksByCategory[category.category_id] = await getProductsByCategory(category.category_id);
      }
    },
    getProductsByCategory(categoryId) {
      return this.booksByCategory[categoryId] || [];
    },
    formatPrice(value) {
      return new Intl.NumberFormat('vi-VN').format(value);
    },
    viewProductDetails(productId) {
      this.$router.push({ name: 'ProductDetails', params: { id: productId } });
    },
  },
};
</script>

<style scoped>
.books-by-category {
  font-family: Arial, sans-serif;
}

.category-section {
  margin-bottom: 40px;
}

.books-by-category h2 {
  font-size: 2em;
  margin-bottom: 15px;
  color: #2c3e50;
  text-align: center;
  position: relative;
  padding-bottom: 10px;
}

.books-by-category h2::after {
  content: '';
  display: block;
  width: 50px;
  height: 4px;
  background-color: #3498db;
  margin: 10px auto 0;
  border-radius: 2px;
}

.books-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 20px;
}

.book-card {
  background-color: white;
  border-radius: 8px;
  padding: 15px;
  text-align: center;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s;
  cursor: pointer;
}

.book-card:hover {
  transform: translateY(-5px);
}

.book-image {
  width: 100%;
  height: auto;
  border-radius: 5px;
  margin-bottom: 10px;
}

.price {
  color: #e74c3c;
  font-weight: bold;
  margin-top: 5px;
}
</style> 