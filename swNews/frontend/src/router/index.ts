import { createRouter, createWebHistory } from 'vue-router'
import { createDiscreteApi, type LoadingBarApi } from 'naive-ui'
import { useTheme } from '@/stores/theme'
import { nextTick } from 'vue'

let loadingBar: LoadingBarApi
nextTick(() => {
  const { themeColor } = useTheme()

  const { loadingBar: bar } = createDiscreteApi(['loadingBar'], {
    loadingBarProviderProps: {
      themeOverrides: {
        colorLoading: themeColor
      }
    }
  })
  loadingBar = bar
})


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '',
      name: 'home',
      component: () => import('@/views/HomeView.vue')
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('@/views/AboutView.vue')
    },
    {
      path: '/athletes',
      name: 'athletes',
      component: () => import('@/views/AthletesView.vue')
    },
    {
      path: '/schedule',
      name: 'schedule',
      component: () => import('@/views/ScheduleView.vue')
    },
    {
      path: '/results/:id?',
      name: 'results',
      component: () => import('@/views/ResultsView.vue')
    },
    {
      path: '/medal',
      name: 'medal',
      component: () => import('@/views/MedalView.vue')
    },
    {
      path: '/comment',
      name: 'comment',
      component: () => import('@/views/CommentView.vue')
    }
  ]
})

router.beforeEach(() => {
  loadingBar?.start()
})

router.afterEach(() => {
  loadingBar?.finish()
})


export default router
