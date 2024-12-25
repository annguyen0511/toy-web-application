<template>
  <div class="carousel">
    <div class="carousel-row">
      <div class="carousel-item single-image">
        <img :src="banners[0].image" alt="Banner Image" class="banner-image rounded-corners" />
      </div>
      <div class="carousel-item double-image">
        <img v-for="(banner, index) in banners.slice(1, 3)" :key="index" :src="banner.image" alt="Banner Image" class="banner-image equal-height rounded-corners" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ImageCarousel',
  data() {
    return {
      banners: [
        { image: 'https://file.hstatic.net/1000202622/file/2_footer-website-1_8f8c151515d24bd7a99d78c3d3bd20c8.png' },
        { image: 'https://file.hstatic.net/1000202622/file/_2__1920_x_602_px_pc__7.png' },
        { image: 'https://file.hstatic.net/1000202622/file/1920x602_ts_2x_51898236936d4e59aab94ea3c661b36d.png' },
      ],
    };
  },
};
</script>

<style scoped>
.carousel {
  display: flex;
  flex-wrap: wrap;
  overflow-x: hidden;
  margin-bottom: 20px;
}

.carousel-row {
  display: flex;
  width: 100%;
}

.carousel-item.single-image {
  flex: 0 0 70%;
  scroll-snap-align: start;
  margin-right: 10px; /* Added margin to the right of the first column */
}

.carousel-item.double-image {
  flex: 0 0 30%;
  display: flex;
  flex-direction: column;
  scroll-snap-align: start;
}

.banner-image {
  width: 100%;
  height: auto;
  margin-bottom: 10px;
  border-radius: 8px; /* Added border-radius to round the corners of the images */
}

.banner-image:last-child {
  margin-bottom: 0;
}

.equal-height {
  height: 50%; /* Ensures both images in the double-image column have equal height */
  object-fit: cover; /* Ensures the image covers the allocated space */
}

.rounded-corners {
  border-radius: 8px; /* Ensures all images have rounded corners */
}
</style> 