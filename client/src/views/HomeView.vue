<template>
  <div class="container">
    <ImageCarousel />
    <div class="promo-container">
      <PromoCard
        v-for="(promo, index) in promotions"
        :key="index"
        :title="promo.title"
        :image="promo.image"
      />
    </div>
    <ProductByCategory category="Fiction" />
  </div>
</template>

<script>
import PromoCard from '@/components/PromoCard.vue';
import ImageCarousel from '@/components/ImageCarousel.vue';
import ProductByCategory from '../components/ProductByCategory.vue';

export default {
  name: 'HomeView',
  components: {
    PromoCard,
    ImageCarousel,
    ProductByCategory,
  },
  data() {
    return {
      promotions: [
        { title: '#', image: 'https://file.hstatic.net/200000504579/file/kidsland-early-development-toys-banner-01_cdc69f7dd98844d9bf4fec055facf024.jpg' },
        { title: '#', image: 'https://file.hstatic.net/200000504579/file/kidsland_castorland-homepage-banner-01_3ca596243dbe49a08b079d3be16e9400.jpg' },
        { title: '#', image: 'https://file.hstatic.net/200000504579/file/kidsland-gundam-homepage-banner-01_239b582a494c4abca8f634f62e67e488.jpg' },
        { title: '#', image: 'https://file.hstatic.net/200000504579/file/new_arrivals_homepage_banner_3d193f40fbd641aba4828782db3f8a07.jpg' },

      ],
    };
  },
};
</script>

<style scoped>
.container {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.promo-container {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 16px;
}
</style>
