<template>
  <div class="container">
    <header>
      <h1>Recipe E-liquid</h1>
    </header>
    <hr>
    <div class="form-group">
      <input type="text" class="form-control " disabled placeholder="Search">
    </div>
    <div class="form-group ">
      <router-link to="/created" class="btn btn-secondary btn-sm" role="button">Created new</router-link>
      <router-link to="/flavor" class="btn btn-secondary btn-sm" role="button">Flavor list</router-link>
    </div>
    <div class="class-group">
      <div v-for="recipe in recipes" v-bind:key="recipe.id" :id="recipe.id">
        <p>
          <router-link class="list-group-item list-group-item-action" :to="`/recipeview/${recipe.id}`">{{recipe.recipe_name}} </router-link>
        </p>
      </div>
    </div>
    {{recipes}}
  </div>
</template>



<script>
import axios from 'axios';

export default {
  name: 'MainView',
  data() {
    return {
      errors: [],
      recipes: [],
    };
  },
  created() {
    axios.get('http://localhost:8000/recipe')
    .then((response) => {
      this.recipes = response.data.recipe;
    })
    .catch((e) => {
      this.errors.push(e);
    });
  },
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>


</style>
