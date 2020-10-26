<template>
  <div id="login-page" class="row ">
    <div class="col s12 l4 offset-l4 z-depth-6 card-panel">
      <form class="login-form" @submit.prevent="register">
        <div class="row"></div>
        <div class="row">
          <h4>Registro de Adoptante</h4>
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
          <div class="input-field col s12">
            <i class="material-icons prefix">lock_outline</i>
            <input
              id="confirm-password"
              type="password"
              v-model="confirmPassword"
            />
            <label for="password">Confirmar Password</label>
          </div>
        </div>

        <div v-if="errors.length" class="row">
          <div class="input-field col s12 m12 l12">
            <p
              v-for="(error, index) in errors"
              :key="index"
              class="margin left-align medium-small red-text text-darken-1"
            >
              {{ error }}
            </p>
          </div>
        </div>

        <div class="row">
          <div class="input-field col s12">
            <button class="btn waves-effect waves-light col s12">
              Registrarme
            </button>
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
    confirmPassword: null,
    errors: [],
    user: {
      username: null,
      password: null,
    },
  }),

  methods: {
    register() {
      if (this.validar()) {
        const axios = require("axios").default;
        this.user.username = this.email;
        this.user.password = this.password;

        axios
          .post(process.env.VUE_APP_RUTA_API + "adoptante/registrar", this.user)
          .then(
            (response) => (
              alert("se ha registrado el nuevo usuario"),
              console.log(response.data),
              (this.email = null),
              (this.password = null),
              (this.confirmPassword = null)
            )
          )
          .catch((e) => {
            console.log(e);
            this.errors.push("Correo ya se encuentra registrado");
          });
      } else {
        alert("formulario no valido");
      }
    },
    validar() {
      this.errors = [];
      if (
        this.email &&
        this.password &&
        this.confirmPassword &&
        this.password == this.confirmPassword
      ) {
        return true;
      }

      if (!this.email) {
        this.errors.push("El ingreso del correo es obligatorio");
      }

      if (!this.password) {
        this.errors.push("El ingreso de la contraseña es obligarotio");
      }

      if (!this.confirmPassword) {
        this.errors.push("Porfavor confirmar contraseña");
      }

      if (this.password != this.confirmPassword) {
        this.errors.push("Las dos contraseñas deben ser iguales");
      }

      return false;
    },
  },
};
</script>
