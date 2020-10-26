<template>
  <ul
    id="slide-in"
    class="sidenav sidenav-fixed"
    style="transform: translateX(0%);"
  >
    <li>
      <div class="user-view">
        <div class="background">
          <img
            src="http://www.dauro.es/wp-content/uploads/2017/06/mobiliario-de-oficina-43.jpg"
          />
        </div>
        <a href="#user">
          <img
            class="circle"
            src="https://i1.sndcdn.com/avatars-000587714706-vjdrog-t200x200.jpg"
          />
        </a>
        <a href="#name">
          <span class="white-text name">John Doe</span>
        </a>
        <a href="#email">
          <span class="white-text email">jdandturk@gmail.com</span>
        </a>
      </div>
    </li>
    <li>
      <a href="#!"> <i class="material-icons">cloud</i>First Link With Icon</a>
    </li>
    <li class="no-padding">
      <ul class="collapsible collapsible-accordion">
        <li>
          <a class="collapsible-header"
            >Fundaciones<i class="material-icons">arrow_drop_down</i></a
          >
          <div class="collapsible-body">
            <ul>
              <li>
                <router-link
                  class="waves-effect waves-light"
                  to="/admin/fundacion/registrar"
                  >Agregar</router-link
                >
              </li>
              <li>
                <router-link
                  class="waves-effect waves-light"
                  to="/admin/fundacion/lista"
                  >Listar</router-link
                >
              </li>

             
            </ul>
          </div>
        </li>
      </ul>
    </li>

    <li class="no-padding">
      <ul class="collapsible collapsible-accordion">
        <li>
          <a class="collapsible-header"
            >Mascotas<i class="material-icons">arrow_drop_down</i></a
          >
          <div class="collapsible-body">
            <ul>
              <li>
                <router-link
                  class="waves-effect waves-light"
                  to="/fundacion/mascota/registrar"
                  >Agregar</router-link
                >
              </li>

              <li><a href="#!">Listar</a></li>
            </ul>
          </div>
        </li>
      </ul>
    </li>
    <li>
      <a href="#!">Second Link</a>
    </li>
    <li>
      <div class="divider"></div>
    </li>
    <li>
      <a class="subheader">Subheader</a>
    </li>
    <li>
      <a class="waves-effect" href="#!">Third Link With Waves</a>
    </li>

    <li>
      <a class="sidenav-close" href="#!">Clicking this will close Sidenav</a>
    </li>
  </ul>
</template>
<script>
import M from "materialize-css";
export default {
  data: () => ({
    isVisible: false,
    elemento: null,
    instances: null,
  }),
  mounted() {
    M.AutoInit();
  },
  methods: {
    cambiar() {
      this.isVisible = !this.isVisible;
      console.log("cerrando menu");
    },
  },
};
</script>
<style scoped>
.visible {
  transform: translateX(0%);
}
.invisible {
  transform: translateX(-105%);
}
</style>
