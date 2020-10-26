import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
//Rutas de Admin
import Admin from '../views/Admin.vue'
import HomeAdmin from '../views/admin/HomeAdmin.vue'
import RegistrarFundacion from '../views/admin/fundaciones/RegistrarFundacion.vue'
import ListarFundacionesAdmin from '../views/admin/fundaciones/ListarFundaciones.vue'
//rutas de Fundaciones
import Fundacion from '../views/Fundacion.vue'
import HomeFundacion from '../views/fundacion/HomeFundacion.vue'
import RegistrarMascotaFundacion from '../views/fundacion/RegistrarMascotaFundacion.vue'
import Adoptante from '../views/Adoptante.vue'
import Login from '../views/Auth/Login.vue'
import Register from '../views/Auth/Register.vue'
const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },

  {
    path: '/admin',
    name: 'Admin',
    component: Admin,
    meta:{
      requiresAuth: true
    },
    children: [
      {
        path: '',
        component: HomeAdmin
        },
        {
          path: 'fundacion/registrar',
          component: RegistrarFundacion
        },
        {
          path: 'fundacion/lista',
          component: ListarFundacionesAdmin
        }
    ]
  },

  {
    path: '/fundacion',
    name: 'Fundacion',
    component: Fundacion,
    meta:{
      requiresAuth: true
    },
    children: [
      {
      path: '',
      component: HomeFundacion
      },
      {
        path: 'mascota/registrar',
        component: RegistrarMascotaFundacion 
        }

    ]
  },

  {
    path: '/adoptante',
    name: 'Adoptante',
    component: Adoptante,
    meta:{
      requiresAuth: true
    }
  },

  {
    path: '/login',
    name: 'Login',
    component: Login,
    
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to,from,next)=>{
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if(localStorage.getItem('isLogin')==='true'){
      var urlArray=to.path.split('/')
      console.log(urlArray[1])
      if(localStorage.getItem('rol')===urlArray[1]){
        next()
      }else{
        console.log(to.path )
        next('/'+localStorage.getItem('rol'))
      }
      

    }else{
      next('/login')
    }
    }else{
      next()
    }
  }
)

export default router
