import { reactive } from 'vue';

export const eventBus = reactive({
  cartCount: JSON.parse(localStorage.getItem('cart'))?.length || 0,
  updateCartCount(count) {
    this.cartCount = count;
  },
}); 