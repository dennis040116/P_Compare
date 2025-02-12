<template>
    <main
    class="w-full h-screen flex flex-col items-center justify-center px-4">
    <div class="max-w-sm w-full text-gray-600 space-y-8">

      <div v-if = "errorMessage" class="flex items-center p-4 mb-4 text-sm text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400" role="alert">
  <svg class="flex-shrink-0 inline w-4 h-4 me-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
    <path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z"/>
  </svg>
  <span class="sr-only">Info</span>
  <div>
    <span class="font-medium">{{ errrorMessage }}</span>
  </div>
</div>

<div v-if = "successMessage" class="flex items-center p-4 mb-4 text-sm text-green-800 border border-green-300 rounded-lg bg-green-50 dark:bg-gray-800 dark:text-green-400 dark:border-green-800" role="alert">
  <svg class="flex-shrink-0 inline w-4 h-4 me-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
    <path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z"/>
  </svg>
  <span class="sr-only">Info</span>
  <div>
    <span class="font-medium">{{ successMessage }}</span> 
  </div>
</div>

      <div class="text-center">
        <!-- <img src="https://floatui.com/logo.svg" width="150" class="mx-auto" /> -->
        <div @click="goHome" class="cursor-pointer">
            <Logo width="150" class="mx-auto"/>
        </div>
        <div class="mt-5 space-y-2">
          <h3 class="text-gray-800 text-2xl font-bold sm:text-3xl">
            Sign up for an account
          </h3>
          <p class="">
            Already have an account?
            <a
              href="/login"
              class="font-medium text-indigo-600 hover:text-indigo-500"
              >Login</a
            >
          </p>
        </div>
      </div>
      <form @submit.prevent>
        <div>
          <label class="font-medium"> UserName </label>
          <input v-model="name"
            type="text"
            required
            class="w-full mt-2 px-3 py-2 text-gray-500 bg-transparent outline-none border focus:border-indigo-600 shadow-sm rounded-lg"
          />
          <label class="font-medium"> Email </label>
          <input v-model="email"
            type="email"
            required
            class="w-full mt-2 px-3 py-2 text-gray-500 bg-transparent outline-none border focus:border-indigo-600 shadow-sm rounded-lg"
          />
          <label class="font-medium"> Password </label>
          <input v-model="password"
            type="password"
            required
            class="w-full mt-2 px-3 py-2 text-gray-500 bg-transparent outline-none border focus:border-indigo-600 shadow-sm rounded-lg"
          />
        </div>
        <button
          @click="signup_({email:email,password:password,name:name})"
          class="w-full mt-4 px-4 py-2 text-white font-medium bg-indigo-600 hover:bg-indigo-500 active:bg-indigo-600 rounded-lg duration-150"
        >
          Create account
        </button>
      </form>
      <div class="relative">
        <span class="block w-full h-px bg-gray-300"></span>
        <p
          class="inline-block w-fit text-sm bg-white px-2 absolute -top-2 inset-x-0 mx-auto"
        >
          Or continue with
        </p>
      </div>
      <div class="space-y-4 text-sm font-medium">
        <!-- Google Button -->
        <button
          class="w-full flex items-center justify-center gap-x-3 py-2.5 border rounded-lg hover:bg-gray-50 duration-150 active:bg-gray-100"
        >
          <!-- SVG for Google -->
          <img
            src="https://raw.githubusercontent.com/sidiDev/remote-assets/7cd06bf1d8859c578c2efbfda2c68bd6bedc66d8/google-icon.svg"
            alt="Google"
            class="w-5 h-5"
          />
          <!-- Comment: Google Icon SVG here -->
          Continue with Google
        </button>
        <!-- Twitter Button -->
        <button
          class="w-full flex items-center justify-center gap-x-3 py-2.5 border rounded-lg hover:bg-gray-50 duration-150 active:bg-gray-100"
        >
          <!-- SVG for Twitter -->
          <img
            src="https://raw.githubusercontent.com/sidiDev/remote-assets/f7119b9bdd8c58864383802fb92c7fc3a25c0646/twitter-icon.svg"
            alt="Twitter"
            class="w-5 h-5"
          />
          <!-- Comment: Twitter Icon SVG here -->
          Continue with Twitter
        </button>
        <!-- Github Button -->
        <button
          class="w-full flex items-center justify-center gap-x-3 py-2.5 border rounded-lg hover:bg-gray-50 duration-150 active:bg-gray-100"
        >
          <!-- SVG for Github -->
          <img
            src="https://raw.githubusercontent.com/sidiDev/remote-assets/0d3b55a09c6bb8155ca19f43283dc6d88ff88bf5/github-icon.svg"
            alt="Github"
            class="w-5 h-5"
          />
          <!-- Comment: Github Icon SVG here -->
          Continue with Github
        </button>
      </div>
    </div>
    </main>
  </template>

<script>
import Logo from '../../assets/img/logo.svg'
import {signup} from '../../api/auth/auth'

export default {
  data(){
    return {
      email:'',
      password:'',
      name:'',
      errorMessage: "", // 错误消息
      successMessage: "", // 成功消息
    }
  },
  components: {
    Logo
  },
    methods: {
      async singup_(data){
        this.errorMessage = "";
          this.successMessage = "";

          try {
            const response = await signup(data);

        // 注册成功
        const { accessToken, refreshToken } = response.data;
        this.successMessage = "注册成功！请使用你的账号登录。";
        console.log("Access Token:", accessToken);
        console.log("Refresh Token:", refreshToken);

        setTimeout(() => {
          this.successMessage = "";
          this.$router.push('/login')
        }, 3000); // 显示 3 秒
        // 可根据需求保存 token（如存入 localStorage 或 vuex）
        localStorage.setItem("accessToken", accessToken);
      } catch (error) {
        if (error.response) {
          const { status, data } = error.response;

          // 处理各种状态码
          if (status === 400) {
            this.errorMessage = "请求参数错误：" + data.message;
          } else if (status === 409) {
            this.errorMessage = "用户已存在：" + data.message;
          } else if (status === 500) {
            this.errorMessage = "服务器内部错误：" + data.message;
          } else {
            this.errorMessage = "未知错误，请稍后再试";
          }
        } else {
          this.errorMessage = "网络错误，请检查你的网络连接";
        }
        setTimeout(() => {
          this.errorMessage = "";
        }, 3000); // 显示 3 秒
      }
        },
        goHome() {
          this.$router.push('/')
        },
        goMain(){
          this.$router.push('/page')
        }
    }
};

</script>