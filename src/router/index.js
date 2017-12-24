import Vue from 'vue';
import Router from 'vue-router';
import MainView from '@/components/MainView';
import CreatedNewRecipe from '@/components/CreatedNewRecipe';
import RecipeView from '@/components/RecipeView';
import FlavorView from '@/components/FlavorView';

Vue.use(Router);

export default new Router({
  routes: [{
    path: '/',
    name: 'Main',
    component: MainView,
  },
  {
    path: '/created',
    name: 'CreatedRecipe',
    component: CreatedNewRecipe,
  },
  {
    path: '/recipeview/:id',
    name: 'RecipeView',
    component: RecipeView,
  },
  {
    path: '/flavor',
    name: 'FlavorView',
    component: FlavorView,
  },
  ],
});
