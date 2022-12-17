<script setup>
  import ProductThumbsVue from './ProductThumbs.vue';
  import { ref, onMounted } from 'vue'
  import Webservice from '@/webservice';

  let products = ref(null);

  onMounted(async() => {
    Webservice.getData('products').then((res) => {
      products.value = res.data;
    }).catch((err) => {
      console.log(err)
    })
  })
</script>

<template>
  <h2>Popular products</h2>
  <div class="popular-products">
    <ProductThumbsVue v-for="product in products" :key="product.id" :product="product"></ProductThumbsVue>
  </div>
</template>


<style scoped>
  .popular-products {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    margin-top: 20px;
  }

</style>