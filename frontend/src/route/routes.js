import Home from '../views/Home.vue'
import About from '../views/About.vue'
import NotFound from '../views/NotFound.vue'
import Login from '../views/login/index.vue'
import Page from '../views/page/index.vue'
import Cart from '../views/cart/index.vue'
import SignUp from '../views/signup/index.vue'

/** @type {import('vue-router').RouterOptions['routes']} */
export const routes = [
  { path: '/', component: Home, meta: { title: 'Home' } },
  {
    path: '/about',
    meta: { title: 'About' },
    component: About,
    // example of route level code-splitting
    // this generates a separate chunk (About.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    // component: () => import('./views/About.vue')
  },
  { path: '/:path(.*)', component: NotFound },
  { path: '/login', component: Login, meta: { title: 'Login' } },
  { path: '/products', component: Page, meta: { title: 'Page' } },
  { path: '/cart', component: Cart, meta: { title: 'Cart' } },
  { path: '/signup', component: SignUp, meta: {title: 'SignUp'}},
  { path: '/products/:user_id', component: Page, meta: { title: 'Page' } },
  { path: '/products/:user_id/:platform_id', component: Page, meta: { title: 'Page' } },
  { path: '/products/like/:user_id/:product_id', component: Page, meta: { title: 'Page' } },
  { path: '/products/dislike/:user_id/:product_id', component: Page, meta: { title: 'Page' } },
  { path: '/products/likes/:user_id', component: Page, meta: { title: 'Page' } },
  { path: '/products/details/:user_id/:product_id', component: Page, meta: { title: 'Page' } },
  { path: '/products/search/:user_id', component: Page, meta: { title: 'Page' } },
  { path: '/cart/:user_id/:product_id', component: Page, meta: { title: 'Page' } },
  { path: '/cart/:user_id', component: Cart, meta: { title: 'Cart' } },
  { path: '/cart/delete/:user_id/:product_id', component: Cart, meta: { title: 'Cart' } },
  { path: '/cart/compare/:user_id', component: Cart, meta: { title: 'Cart' } },
  
]
