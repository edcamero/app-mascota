<template>
  <div class="fundacion container">
    <div class="col s12 m3">
      <Menu></Menu>
    </div>

    <main>
      <router-view></router-view>
    </main>
  </div>
</template>
<script>
import Menu from "./fundacion/MenuFundacion.vue";
export default {
  components: {
    Menu,
  },
};
</script>
<style>
header,
main,
footer {
  padding-left: 300px;
}

@media only screen and (max-width: 992px) {
  header,
  main,
  footer {
    padding-left: 0;
  }
  main {
    width: 100%;
  }
}
</style>
