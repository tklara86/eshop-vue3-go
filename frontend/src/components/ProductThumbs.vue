<script setup>
const props = defineProps(['product'])

</script>

<template>
    <div>
      <p>{{product.name}}</p>
      <img :src="product.image" alt="">
      <p>{{product.total_price.display}}</p>
      <p>{{product.description}}</p>
     
    </div>
    
</template>


<style scoped>
  img {
    width: 70px;
  }

</style>