<template>
  <div class="container">
    <h1>Recipe E-liquid</h1>
    </header>
    <hr>
    <div class=" form-group">
    <button class="btn btn btn-success btn-sm" role="button" @click="save_recipe">Save</button>
    <router-link to="/" class="btn btn-secondary btn-sm" role="button">Back</router-link>
    </div>
    <div class="row">
      <div class="col-7">
        <div class="input-group form-group">
          <span class="input-group-addon">Recipe Name</span>
          <input type="text" class="form-control" v-model="new_recipe.recipe_name">
        </div>
        <div class="input-group form-group">
          <span class="input-group-addon">Botte ID</span>
          <input type="text" class="form-control" v-model="new_recipe.bottle_id">
        </div>
        <div class="input-group form-group">
          <span class="input-group-addon">Bottle Size</span>
          <input type="number" min="1" max="200" step="1" class="form-control" v-model.number="new_recipe.bottle_size">
          <span class="input-group-addon">ml</span>
        </div>
        <div class="input-group form-group">
          <span class="input-group-addon">PG</span>
          <input type="number" min="1" max="100" step="1" class="form-control" v-model.number="new_recipe.pg">
          <span class="input-group-addon">%</span>
        </div>
        <div class="input-group form-group">
          <span class="input-group-addon">VG</span>
          <input type="number" min="1" max="100" step="1" class="form-control" v-model.number="new_recipe.vg">
          <span class="input-group-addon">%</span>
        </div>
        <div class="input-group form-group">
          <span class="input-group-addon">Water | Alcohol | Other</span>
          <input type="number" min="0" max="100" step="1" class="form-control" v-model.number="new_recipe.other">
          <span class="input-group-addon">%</span>
        </div>
        <div class="input-group form-group">
          <span class="input-group-addon">Nicotine</span>
          <input type="number" min="0" class="form-control" v-model.number="new_recipe.nicotine">
          <span class="input-group-addon">mg</span>
        </div>
        <div class="input-group form-group">
          <span class="input-group-addon">Date</span>
          <input type="date" class="form-control" v-model="new_recipe.date">
        </div>
        <div class="input-group form-group">
          <span class="input-group-addon">Steep</span>
          <input type="number" min="0" class="form-control" v-model.number="new_recipe.steep_time">
          <span class="input-group-addon">days</span>
        </div>
        <div class="input-group form-group">
          <span class="input-group-addon">Note</span>
          <textarea name="" cols="10" rows="5" class="form-control" v-model="new_recipe.note" id="note"></textarea>
        </div>
        <div class="input-group form-group">
          <button type="button" class="btn btn-primary btn-sm" @click="add_new_flavor">Add flavor</button>
        </div>
        <div class="input-group form-group" v-for="(flavor,index) in new_flavor" :key="index">
          <span class="input-group-addon">Flavor {{index+1}}</span>
          <input type="text" class="form-control" v-model="flavor.flavor_name">
          <span class="input-group-addon"></span>
          <input type="number" min="0" class="form-control" v-model.number="flavor.flavor_percent">
          <span class="input-group-addon">%</span>
          <span class="input-group-btn">
            <button type="button" class="btn btn-danger btn-sm" @click="delete_flavor(index)">Delete</button>
          </span>
        </div>
      </div>
      <div class="col-5">
        <p>{{new_flavor}}</p>
        <p>{{new_recipe}}</p>
        <p>{{errors}}</p>
        <p>{{message}}</p>
      </div>
    </div>
  </div>
</template>



<script>
  import axios from 'axios';

  export default {
    name: 'CreatedNewRecipe',
    data() {
      return {
        new_flavor: [],
        new_recipe: {},
        message: [],
        errors: [],
      };
    },
    methods: {
      add_new_flavor() {
        this.new_flavor.push({});
      },
      delete_flavor(index) {
        this.new_flavor.splice(index, 1);
      },
      save_recipe() {
        axios.post('http://localhost:8000/recipe/new', this.new_recipe)
        .then((recipe) => {
          this.message.push(recipe.data);
          Object.keys(this.new_flavor).forEach((element) => {
            this.$set(this.new_flavor[element], 'recipe_id', Object.values(recipe.data)[0]);
          });
          axios.post('http://localhost:8000/flavor/new', this.new_flavor)
          .then((flavor) => {
            this.message.push(flavor.data);
          });
        })
        .catch((e) => {
          this.errors.push(e.message);
        });
      },
    },
  };

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
