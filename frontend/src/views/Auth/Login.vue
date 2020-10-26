<template>
  <div id="login-page" class="row">
    <div class="col s12 l4 offset-l4 z-depth-6 card-panel">
      <form class="login-form" @submit.prevent="login">
        <div class="row"></div>
        <div class="row">
          <h4>Login</h4>
          <div class="input-field col s12">
            <i class="material-icons prefix">mail_outline</i>
            <input class="validate" id="email" type="email" v-model="email" />
            <label for="email" data-error="wrong" data-success="right"
              >Email</label
            >
          </div>
        </div>
        <div class="row">
          <div class="input-field col s12">
            <i class="material-icons prefix">lock_outline</i>
            <input id="password" type="password" v-model="password" />
            <label for="password">Password</label>
          </div>
        </div>
        <div class="row">
          <div class="input-field col s12 m12 l12 login-text">
            <label>
              <input type="checkbox" />
              <span>Recordarme</span>
            </label>
          </div>
        </div>
        <div class="row">
          <div class="input-field col s12">
            <button class="btn waves-effect waves-light col s12">
              Ingresar
            </button>
          </div>
        </div>
        <div class="row">
          <div class="input-field col s6 m6 l6">
            <p class="margin medium-small"><a href="#">Registrarme!</a></p>
          </div>
          <div class="input-field col s6 m6 l6">
            <p class="margin right-align medium-small">
              <a href="#">Olvide mi contraseña</a>
            </p>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data: () => ({
    email: null,
    password: null,
    token: "hola",
    user: null,
  }),

  methods: {
    login() {
      const axios = require("axios").default;
      let formData = new FormData();
      if (this.email && this.password) {
        //console.log(process.env.VUE_APP_RUTA_API+'login')
        console.log(this.password);
        formData.append("username", this.email);
        formData.append("password", this.password);
        axios
          .post(process.env.VUE_APP_RUTA_API + "login", formData)
          .then(
            (response) => (
              (this.token = response.data.tokenUser),
              console.log(response.data),
              localStorage.setItem("isLogin", "true"),
              localStorage.setItem("token", this.token),
              (this.$root.isLogin = true),
              (this.user = response.data.user),
              (this.$root.user = this.user),
              (this.$root.rol = this.user.rol.nombre),
              localStorage.setItem("username", this.user.username),
              localStorage.setItem("rol", this.user.rol.nombre),
              this.redireccionar(this.user.rol.nombre)
            )
          )
          .catch(function(error) {
            console.log(error);
          });
      }
    },
    redireccionar(rol) {
      console.log("redireccionando a:" + rol);
      switch (rol) {
        case "adoptante":
          this.$router.push("/adoptante");
          break;
        case "admin":
          this.$router.push("/admin");
          break;
        case "fundacion":
          this.$router.push("/fundacion");
          break;
      }
    },
  },
};
</script>
