<template>
  <div class="container">
    <header>
      <h1>Recipe E-liquid</h1>
<hr>
    </header>
    <div class="form-group">
      <input type="text" class="form-control" disabled placeholder="Search">
    </div>
    <router-link to="/4" class="btn btn-secondary btn-sm" role="button">Add new flavor</router-link>
    <router-link to="/" class="btn btn-secondary btn-sm" role="button">Back</router-link>
    <div class="class-group">
    </div>
  </div>
</template>



<script>

export default {
  name: 'FlavorView',
  data() {
    return {
    };
  },
  created() {

  },
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>


</style>
