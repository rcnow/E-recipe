<template>
  <div class="container">
    <header class="page-header">
      <h1>Recipe E-liquid</h1>
      <hr>
    </header>
    <div class="form-group">
      <router-link to="/" class="btn btn-secondary btn-sm" role="button">Back</router-link>
      <button class="btn btn-secondary btn-sm" role="button" @click="isEdit = !isEdit">Edit</button>
      <button class="btn btn-success btn-sm " v-if="isEdit" @click="update_recipe">Update</button>
      <button class="btn btn-danger btn-sm " v-if="isEdit" :id="recipe.id" @click="delete_recipe">Delete</button>
    </div>
    <div class="row">
      <div class="col-7">
        <div class="input-group form-group ">
          <span class="input-group-addon">Recipe ID</span>
          <input type="text" class="form-control" disabled :value="recipe.id">
        </div>
        <div class="input-group form-group ">
          <span class="input-group-addon">Recipe Name</span>
          <input type="text" class="form-control" disabled :value="recipe.recipe_name">
        </div>
        <div class="input-group form-group ">
          <span class="input-group-addon">Bottle ID </span>
          <input type="text" class="form-control" disabled :value="recipe.bottle_id">
        </div>
        <div class="input-group form-group ">
          <span class="input-group-addon">Bottle Size</span>
          <input type="number" class="form-control" disabled :value="recipe.bottle_size">
        </div>
        <div class="input-group form-group ">
          <span class="input-group-addon">PG</span>
          <input type="number" class="form-control" disabled :value="recipe.pg">
          <span class="input-group-addon">%</span>
        </div>
        <div class="input-group form-group ">
          <span class="input-group-addon">VG</span>
          <input type="number" class="form-control" disabled :value="recipe.vg">
          <span class="input-group-addon">%</span>
        </div>
        <div class="input-group form-group ">
          <span class="input-group-addon">Water | Alcohol | Other</span>
          <input type="number" class="form-control" disabled :value="recipe.other">
          <span class="input-group-addon">%</span>
        </div>
        <div class="input-group form-group ">
          <span class="input-group-addon">Nicotine</span>
          <input type="number" class="form-control" disabled :value="recipe.nicotine">
          <span class="input-group-addon">mg</span>
        </div>
        <div class="input-group form-group ">
          <span class="input-group-addon">Date</span>
          <input type="date" class="form-control" disabled :value="recipe.date">
        </div>
        <div class="input-group form-group ">
          <span class="input-group-addon">Steep </span>
          <input type="text" class="form-control" disabled :value="recipe.steep_time">
          <span class="input-group-addon">days</span>
        </div>
        <div class="input-group form-group">
          <span class="input-group-addon">Note</span>
          <textarea :disabled="!isEdit" cols="10" rows="5" class="form-control" v-model="recipe.note" id="note"></textarea>
        </div>
        <div class="input-group form-group" v-for="(flavor, index) in flavors" :key="flavor.id">
          <span class="input-group-addon">Flavor {{index+1}}</span>
          <input type="text" class="form-control" disabled v-model="flavor.flavor_name">
          <span class="input-group-addon"></span>
          <input type="number" min="0" class="form-control" disabled v-model.number="flavor.flavor_percent">
          <span class="input-group-addon">%</span>
        </div>
      </div>
      <div class="col-5">
        <p>{{recipe}}</p>
        <p>{{flavors}}</p>
        <p>{{errors}}</p>
        <p>{{message}}</p>
      </div>
    </div>
  </div>

</template>


<script>
import axios from 'axios';

export default {
  name: 'RecipeView',
  data() {
    return {
      errors: [],
      recipe: {},
      flavors: {},
      message: [],
      isEdit: false,
    };
  },
  created() {
    axios.get(`http://localhost:8000/recipeview/${this.$route.params.id}`)
    .then((response) => {
      this.recipe = response.data.recipe[0];
      this.flavors = response.data.flavors;
    })
    .catch((e) => {
      this.errors.push(e);
    });
  },
  methods: {
    delete_recipe() {
      axios.delete(`http://localhost:8000/recipe/delete/${this.$route.params.id}`)
      .then((response) => {
        this.message = response.data;
      })
      .catch((e) => {
        this.errors.push(e);
      });
    },
    update_recipe() {
      axios.put(`http://localhost:8000/recipe/update/${this.$route.params.id}`, this.recipe)
      .then((response) => {
        this.message = response.data;
      })
      .catch((e) => {
        this.errors.push(e);
      });
    },
  },

};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>


</style>
