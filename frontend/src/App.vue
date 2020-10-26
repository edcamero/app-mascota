<template>
<div class="container-fluid ">
  <div class="navbar-fixed  "> 
    <nav class="teal lighten-2">
      <div class="nav-wrapper "> 
        <a href="#" data-target="mobile-demo" class="sidenav-trigger"><i class="material-icons">menu</i></a>
        <a href="#" data-target="slide-in" class="top-nav sidenav-trigger full hide-on-large-only"><i class="material-icons">menu</i></a>
        <ul id="nav-mobile" class=" right hide-on-med-and-down">
             <li> <router-link  class="waves-effect waves-light" to="/">Home</router-link></li>
             <li><router-link class="waves-effect waves-light" to="/admin/" v-if="rol=='admin'" >Admin</router-link></li>
             <li> <router-link class="waves-effect waves-light" to="/fundacion/" v-if="rol=='fundacion'">Fundacion</router-link> </li>
             <li><router-link class="waves-effect waves-light" to="/adoptante/" v-if="rol=='adoptante'">Adoptante</router-link></li>
             <li v-if="isLogin" class="waves-effect waves-light"><a   v-on:click="logOut" >Cerrar</a></li>
             <li v-else><router-link class="waves-effect waves-light" to="/login/" >Ingresar</router-link></li>
             <li> <router-link class="waves-effect waves-light" to="/register/" v-if="!isLogin">Registrarme</router-link></li>
        </ul>
     
  
      
     </div>
    </nav>

  </div>
  <router-view class="row"/>
   <ul class="sidenav " id="mobile-demo">
             <li> <router-link  class="waves-effect waves-light" to="/">Home</router-link></li>
             <li><router-link to="/admin/" v-if="rol=='admin'" >Admin</router-link></li>
             <li> <router-link to="/fundacion/" v-if="rol=='fundacion'">Fundacion</router-link> </li>
             <li><router-link to="/adoptante/" v-if="rol=='adoptante'">Adoptante</router-link></li>
             <li v-if="isLogin"><a   v-on:click="logOut" >Cerrar</a></li>
             <li v-else><router-link to="/login/" >Ingresar</router-link></li>
             <li> <router-link to="/register/" v-if="!isLogin">Registrarme</router-link></li>
      </ul>
  </div>
          
</template>

<script>
import M from 'materialize-css'
export default {

  data: () => ({
    isLogin:false,
    user:null,
    rol:null
  }),
  created() {
    if (localStorage.isLogin=='true') {
      this.isLogin=true,
      this.rol=localStorage.getItem("rol")
    }
  },
  mounted(){
       M.AutoInit()
       
  },
  watch: {
    isLogin(newLogin) {
      localStorage.isLogin = newLogin;
    }
  },
  methods: {
    logOut() {
        this.isLogin=false,
        this.user=null,
        this.rol=null,
        localStorage.setItem('isLogin', 'true'),
        //localStorage.setItem('token', this.token),
        localStorage.removeItem('token')
        localStorage.removeItem('rol')
        localStorage.removeItem('username')
        this.$router.push('/')
      
    }
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}




</style>
