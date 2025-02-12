<template>
  <div>
    <main>
      <router-view />
    </main>
  </div>
</template>


<script>
/* SCRIPT setup CAN NOT CONTAIN export default */
import { useStore } from './stores/store';
import { storeToRefs } from 'pinia'
import { useRoute } from 'vue-router';


export default {
  setup() {
    const route = useRoute()
    const isLoginPage = route.path === '/login'
    const menu = ['home', 'about us', 'blogs', 'contact']
    const store = useStore()
    // const {count} = store // NOT REACTIVE 
    const { count } = storeToRefs(store);
    console.log(count.value);
    return {count,menu,isLoginPage}
  },
  methods: {
    goHome() {
      this.$router.push('/')
    }
  }
};

</script>
