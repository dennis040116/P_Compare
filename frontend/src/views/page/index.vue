<template>


<div class="bg-gradient-to-b from-[#dee6f7] to-white">
       <!-- <img src="./assets/img/blob.svg" alt="blob" class="absolute top-0 right-0 pointer-events-none"> -->
        <!-- <img src="https://raw.githubusercontent.com/roma-marshall/shiny-octo-succotash/main/src/assets/img/blob.svg" alt="blob" class="absolute top-0 right-0 pointer-events-none"> -->
        <div class="px-6 mx-auto max-w-6xl">
          <div class="flex justify-between items-center py-10">


            
            <!-- <div @click="goHome" class="cursor-pointer">
            <Logo/>
            </div> -->



          </div>
        </div>

        <a href="#" class="absolute top-4 right-12 flex-shrink-0 group block"> <!-- Use absolute positioning -->
    <div class="flex items-center">
      <div>
        <img class="inline-block h-9 w-9 rounded-full" src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80" alt="" />
      </div>
      <div class="ml-3">
        <p class="text-sm font-medium text-gray-700 group-hover:text-gray-900">{{user_id}}</p>
        <p class="text-xs font-medium text-indigo-500 group-hover:text-indigo-700">View profile</p>
      </div>
    </div>
  </a>
    </div>

    <div v-if = "errorMessage" class="flex items-center p-4 mb-4 text-sm text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400" role="alert">
      <svg class="flex-shrink-0 inline w-4 h-4 me-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
        <path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z"/>
      </svg>
      <span class="sr-only">Info</span>
      <div>
        <span class="font-medium">{{ errorMessage }}</span>
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

    <section class="py-14">
    <div class="max-w-screen-xl mx-auto px-4 text-center text-gray-600 md:px-8">
      <div class="max-w-2xl mx-auto">
        <h3 class="text-gray-800 text-3xl font-semibold sm:text-4xl">
          The fastest way to launch a product
        </h3>
        <p class="mt-3">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec congue, nisl eget molestie varius, enim ex faucibus purus.
        </p>
      </div>
    </div>
  </section>





    <div class="flex justify-center items-center py-3">
    <div class="relative w-full max-w-lg flex items-center">
      <!-- 搜索输入框 -->
      <input
        v-model="searchQuery"
        @input="onInput"
        type="text"
        placeholder="Search..."
        class="w-full py-3 px-4 pl-10 text-gray-700 bg-white rounded-full shadow-xl border border-transparent focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500 focus:outline-none transition duration-300 ease-in-out"
      />

      <!-- 搜索图标 -->
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
        class="w-5 h-5 absolute top-1/2 left-3 transform -translate-y-1/2 text-gray-500"
      >
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
          d="M8 12a4 4 0 100-8 4 4 0 000 8zm6 8l-4-4m0 0l4-4m-4 4H6"></path>
      </svg>


      <!-- 搜索按钮 -->
      <button
        @click="SearchProducts(searchQuery)"
        class="absolute top-1/2 right-3 transform -translate-y-1/2 bg-indigo-600 text-white px-4 py-2 rounded-full hover:bg-indigo-700 transition duration-300 ease-in-out"
      >
        Search
      </button>
    </div>
  </div>    
 





 

    <!-- <section class="mt-12 mx-auto px-4 max-w-screen-xl lg:px-8">
    <div class="mt-12 grid gap-2 sm:grid-cols-2 lg:grid-cols-4">
      <article v-for="items in posts" :key="items.id" class="max-w-md mx-auto mt-4 shadow-lg border rounded-md duration-300 hover:shadow-sm">
        <div class="w-64 p-2 m-auto bg-white shadow-lg rounded-2xl">
    <img src="https://flowbite.s3.amazonaws.com/docs/gallery/masonry/image-8.jpg" alt="adidas" class="w-full h-auto object-cover"/>
    <div class="p-4 m-3 bg-white rounded-lg">
        <p class="text-xl font-bold text-black ">
            Adidas
        </p>
        <div class="flex items-center justify-between ">
            <p class="text-black">
                $98.00
            </p>

            
            <div class="group relative inline-block">
              <button type="button" class="w-10 h-10 text-base font-medium text-white bg-indigo-500 rounded-full hover:bg-indigo-700">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" class="mx-auto" fill="white" viewBox="0 0 1792 1792">
                    <path d="M1600 736v192q0 40-28 68t-68 28h-416v416q0 40-28 68t-68 28h-192q-40 0-68-28t-28-68v-416h-416q-40 0-68-28t-28-68v-192q0-40 28-68t68-28h416v-416q0-40 28-68t68-28h192q40 0 68 28t28 68v416h416q40 0 68 28t28 68z">
                    </path>
                </svg>
            </button>
              <div
                class="bg-primary absolute bottom-full left-1/2 z-20 mb-3 -translate-x-1/2 bg-black whitespace-nowrap rounded-[5px] py-1.5 px-3.5 text-sm text-white opacity-0 group-hover:opacity-100"
              >
                <span
                  class="bg-primary absolute bg-black bottom-[-3px] left-1/2 -z-10 h-2 w-2 -translate-x-1/2 rotate-45"
                ></span>
                Compare
              </div>
            </div>

        </div>
    </div>
</div>
      </article>
    </div>
  </section> -->



  <div class="max-w-2xl mx-auto py-16 px-4 sm:py-24 sm:px-6 lg:max-w-7xl lg:px-8">
    <!-- Tabs for different product categories -->
    <div class="flex justify-center space-x-4 mb-8">
      <button
        v-for="tab in tabs"
        :key="tab"
        @click="setActiveTab(tab)"
        :class="{'text-indigo-600 font-semibold border-b-2 border-indigo-600': activeTab === tab, 'text-gray-600': activeTab !== tab}"
        class="px-4 py-2 text-lg"
      >
        {{ tab }}
      </button>
    </div>

    <!-- Product List -->
    <div class="grid grid-cols-1 gap-y-12 sm:grid-cols-2 sm:gap-x-6 lg:grid-cols-4 xl:gap-x-8">
      <div v-for="product in filteredProducts" :key="product.id">
        <div class="relative group">
          <!-- Image container -->
          <div class="relative w-full h-72 rounded-lg overflow-hidden">
            <img :src="product.ImageURL" class="w-full h-full object-center object-cover" />

            <!-- View Details button, hidden initially, shown on hover -->
            <div class="absolute inset-0 bg-black bg-opacity-50 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity z-10">
              <button  @click = getDetail(product.product_id) class="bg-gray-100 text-gray-900 font-semibold py-2 px-4 rounded-md">
                View Details
              </button>
            </div>


        <!-- Favourite icon positioned at the top right -->
        <button @click="toggleFavorite(product)" class="absolute top-3 right-3 p-1 rounded-full focus:outline-none z-10">
          <svg 
            xmlns="http://www.w3.org/2000/svg" 
            class="h-5 w-5" 
            :class="isFavorite(product) ? 'text-red-500' : 'text-gray-400'" 
            fill="none" 
            viewBox="0 0 24 24" 
            stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3 3h18v18H3V3z" />
            <path d="M12 21l-3.09-3.14A6.978 6.978 0 015 12a6.978 6.978 0 014.91-5.86A4.992 4.992 0 0112 7a4.992 4.992 0 012.09-1.86A6.978 6.978 0 0119 12a6.978 6.978 0 01-3.91 5.86L12 21z" />
          </svg>
        </button>

          </div>

          <!-- Product information -->
          <div class="relative mt-4">
            <h3 class="text-sm font-medium text-gray-900">{{ product.Description }}</h3>
            <!-- <p class="mt-1 text-sm text-gray-500">{{ product. }}</p> -->
          </div>

          <!-- Price overlay -->
          <div class="absolute top-0 inset-x-0 h-72 rounded-lg p-4 flex items-end justify-end overflow-hidden">
            <div aria-hidden="true" class="absolute inset-x-0 bottom-0 h-36 bg-gradient-to-t from-black opacity-50"></div>
            <p class="relative text-lg font-semibold text-white">{{ product.Price }}</p>
          </div>
        </div>

        <!-- Add to bag button -->
        <div class="mt-6">
          <button @click = "addToCart(product.product_id)" class="w-full relative flex bg-gray-100 border border-transparent rounded-md py-2 px-8 items-center justify-center text-sm font-medium text-gray-900 hover:bg-gray-200">
            Add to cart
          </button>
        </div>
      </div>
    </div>
  </div>




<div class="dock-container flex justify-center items-end w-full py-4 z-20">
    <div class="dock flex space-x-3 p-2 bg-gray-100 rounded-lg shadow-lg">
      <div
        v-for="(icon, index) in icons"
        :key="index"
        class="dock-item transform transition-transform duration-200 ease-in-out relative"
        @mouseover="onHover(index)"
        @mouseleave="resetScale"
        @click="navigateTo(icon.link)"  
        :style="{ transform: hoverIndex === index ? 'scale(1.25)' : 'scale(1)' }"
        style="cursor: pointer"
      >
        <!-- 图标上方的小红点 -->
        <div v-if="icon.notifications > 0" class="absolute left-0 top-0 -m-1 bg-red-500 rounded-full">
          <!-- 固定大小红点，不显示数字 -->
          <span class="w-3 h-3 block rounded-full bg-red-500"></span>
        </div>

        <img
          :src="icon.src"
          alt="icon"
          class="dock-icon w-12 h-12 object-contain rounded-lg shadow-sm"
        />
      </div>
    </div>
  </div>

      <!-- Dialog 模态框 -->
      <div v-if="isDialogVisible" class="fixed inset-0 flex items-center justify-center z-10">
        <!-- 背景遮罩 -->
        <div class="fixed inset-0 bg-black opacity-40" @click="closeDialog"></div>
  
        <!-- 对话框内容 -->
        <div class="relative bg-white w-full max-w-lg mx-auto p-6 rounded-lg shadow-lg z-20">
          <!-- 关闭按钮 -->
          <button @click="closeDialog" class="absolute top-4 right-4 text-gray-500 hover:text-gray-800">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
  
          <!-- 对话框标题 -->
          <h3 class="text-lg font-medium text-gray-800 mb-4">Price History</h3>
  
          <!-- 图表容器 -->
          <div id="hs-curved-area-charts"></div>
        </div>
      </div>

    <footer class="pt-10">
    <div class="max-w-screen-xl mx-auto px-4 md:px-8">
      <div class="justify-between items-center gap-12 md:flex">
        <div class="flex-1 max-w-lg">
          <h3 class="text-2xl font-bold">
            Get our beautiful newsletter straight to your inbox.
          </h3>
        </div>
        <div class="flex-1 mt-6 md:mt-0">
          <form @submit.prevent class="flex items-center gap-x-3 md:justify-end">
            <div class="relative">
              <svg class="w-6 h-6 text-gray-400 absolute left-3 inset-y-0 my-auto" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 01-2.25 2.25h-15a2.25 2.25 0 01-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0019.5 4.5h-15a2.25 2.25 0 00-2.25 2.25m19.5 0v.243a2.25 2.25 0 01-1.07 1.916l-7.5 4.615a2.25 2.25 0 01-2.36 0L3.32 8.91a2.25 2.25 0 01-1.07-1.916V6.75" />
              </svg>
              <input
                type="email"
                required
                placeholder="Enter your email"
                class="w-full pl-12 pr-3 py-2 text-gray-500 bg-white outline-none border focus:border-indigo-600 shadow-sm rounded-lg"
              />
            </div>
            <button class="block w-auto py-3 px-4 font-medium text-sm text-center text-white bg-indigo-600 hover:bg-indigo-500 active:bg-indigo-700 active:shadow-none rounded-lg shadow">
              Subscribe
            </button>
          </form>
        </div>
      </div>
      <div class="flex-1 mt-16 space-y-6 justify-between sm:flex md:space-y-0">
        <div v-for="(item, idx) in footerNavs" :key="idx" class="space-y-4 text-gray-600">
          <h4 class="text-gray-800 font-semibold sm:pb-2">
            {{ item.label }}
          </h4>
          <ul>
            <li v-for="(el, elIdx) in item.items" :key="elIdx">
              <a :href="el.href" class="hover:text-gray-800 duration-150">
                {{ el.name }}
              </a>
            </li>
          </ul>
        </div>
      </div>
      <div class="mt-10 py-10 border-t items-center justify-between sm:flex">
        <p class="text-gray-600">© 2022 Float UI Inc. All rights reserved.</p>
        <div class="flex items-center gap-x-6 text-gray-400 mt-6">
          <a href="javascript:void()">
              <svg class="w-6 h-6 hover:text-gray-500 duration-150" fill="none" viewBox="0 0 48 48"><g clip-path="url(#a)"><path fill="currentColor" d="M48 24C48 10.745 37.255 0 24 0S0 10.745 0 24c0 11.979 8.776 21.908 20.25 23.708v-16.77h-6.094V24h6.094v-5.288c0-6.014 3.583-9.337 9.065-9.337 2.625 0 5.372.469 5.372.469v5.906h-3.026c-2.981 0-3.911 1.85-3.911 3.75V24h6.656l-1.064 6.938H27.75v16.77C39.224 45.908 48 35.978 48 24z" /></g><defs><clipPath id="a"><path fill="#fff" d="M0 0h48v48H0z" /></clipPath></defs></svg>
          </a>
          <a href="javascript:void()">
              <svg class="w-6 h-6 hover:text-gray-500 duration-150" fill="none" viewBox="0 0 48 48"><g clip-path="url(#clip0_17_80)"><path fill="currentColor" d="M15.1 43.5c18.11 0 28.017-15.006 28.017-28.016 0-.422-.01-.853-.029-1.275A19.998 19.998 0 0048 9.11c-1.795.798-3.7 1.32-5.652 1.546a9.9 9.9 0 004.33-5.445 19.794 19.794 0 01-6.251 2.39 9.86 9.86 0 00-16.788 8.979A27.97 27.97 0 013.346 6.299 9.859 9.859 0 006.393 19.44a9.86 9.86 0 01-4.462-1.228v.122a9.844 9.844 0 007.901 9.656 9.788 9.788 0 01-4.442.169 9.867 9.867 0 009.195 6.843A19.75 19.75 0 010 39.078 27.937 27.937 0 0015.1 43.5z" /></g><defs><clipPath id="clip0_17_80"><path fill="#fff" d="M0 0h48v48H0z" /></clipPath></defs></svg>
          </a>
          <a href="javascript:void()">
              <svg class="w-6 h-6 hover:text-gray-500 duration-150" fill="none" viewBox="0 0 48 48"><g fill="currentColor" clip-path="url(#clip0_910_44)"><path fill-rule="evenodd" d="M24 1A24.086 24.086 0 008.454 6.693 23.834 23.834 0 00.319 21.044a23.754 23.754 0 003.153 16.172 23.98 23.98 0 0012.938 10.29c1.192.221 1.641-.518 1.641-1.146 0-.628-.024-2.45-.032-4.442-6.676 1.443-8.087-2.817-8.087-2.817-1.089-2.766-2.663-3.493-2.663-3.493-2.178-1.478.163-1.45.163-1.45 2.413.17 3.68 2.461 3.68 2.461 2.138 3.648 5.616 2.593 6.983 1.976.215-1.545.838-2.596 1.526-3.193-5.333-.6-10.937-2.647-10.937-11.791a9.213 9.213 0 012.472-6.406c-.246-.6-1.069-3.026.234-6.322 0 0 2.015-.64 6.602 2.446a22.904 22.904 0 0112.017 0c4.583-3.086 6.594-2.446 6.594-2.446 1.307 3.288.484 5.714.238 6.322a9.194 9.194 0 012.476 6.414c0 9.163-5.615 11.183-10.957 11.772.859.742 1.626 2.193 1.626 4.421 0 3.193-.028 5.762-.028 6.548 0 .636.433 1.38 1.65 1.146a23.98 23.98 0 0012.938-10.291 23.754 23.754 0 003.151-16.175A23.834 23.834 0 0039.56 6.69 24.086 24.086 0 0024.009 1H24z" clip-rule="evenodd" /><path d="M9.089 35.264c-.052.119-.243.154-.398.071-.155-.083-.27-.237-.214-.36.056-.122.242-.154.397-.07.155.082.274.24.215.359zM10.063 36.343a.4.4 0 01-.493-.11c-.155-.167-.187-.396-.068-.499.12-.102.334-.055.489.11.155.167.19.396.072.499zM11.008 37.714c-.147.103-.397 0-.536-.206a.395.395 0 010-.569c.147-.098.397 0 .537.202.139.202.143.47 0 .573zM12.292 39.042c-.131.146-.397.106-.616-.091-.219-.198-.27-.467-.139-.609.131-.142.397-.102.624.091.226.194.27.466.131.609zM14.092 39.816c-.06.186-.33.269-.6.19-.27-.08-.449-.3-.397-.49.051-.19.326-.277.6-.19.274.087.449.297.397.49zM16.056 39.95c0 .194-.223.36-.509.364-.286.004-.52-.154-.52-.348 0-.193.222-.36.508-.363.286-.004.52.15.52.347zM17.884 39.646c.036.194-.163.395-.45.443-.285.047-.536-.067-.572-.257-.035-.19.171-.395.45-.447.278-.05.536.068.572.261z" /></g><defs><clipPath id="clip0_910_44"><path fill="#fff" d="M0 0h48v48H0z" /></clipPath></defs></svg>
          </a>
          <a href="javascript:void()">
              <svg class="w-6 h-6 hover:text-gray-500 duration-150" fill="currentColor" viewBox="0 0 48 48"><g clip-path="url(#clip0_17_63)"><path d="M24 4.322c6.413 0 7.172.028 9.694.14 2.343.104 3.61.497 4.453.825 1.116.432 1.922.957 2.756 1.791.844.844 1.36 1.64 1.79 2.756.329.844.723 2.12.826 4.454.112 2.53.14 3.29.14 9.693 0 6.413-.028 7.172-.14 9.694-.103 2.344-.497 3.61-.825 4.453-.431 1.116-.957 1.922-1.79 2.756-.845.844-1.642 1.36-2.757 1.791-.844.328-2.119.722-4.453.825-2.532.112-3.29.14-9.694.14-6.413 0-7.172-.028-9.694-.14-2.343-.103-3.61-.497-4.453-.825-1.115-.431-1.922-.956-2.756-1.79-.844-.844-1.36-1.641-1.79-2.757-.329-.844-.723-2.119-.826-4.453-.112-2.531-.14-3.29-.14-9.694 0-6.412.028-7.172.14-9.694.103-2.343.497-3.609.825-4.453.431-1.115.957-1.921 1.79-2.756.845-.844 1.642-1.36 2.757-1.79.844-.329 2.119-.722 4.453-.825 2.522-.113 3.281-.141 9.694-.141zM24 0c-6.516 0-7.331.028-9.89.14-2.55.113-4.304.526-5.822 1.116-1.585.619-2.926 1.435-4.257 2.775-1.34 1.332-2.156 2.672-2.775 4.247C.666 9.806.253 11.55.141 14.1.028 16.669 0 17.484 0 24s.028 7.331.14 9.89c.113 2.55.526 4.304 1.116 5.822.619 1.585 1.435 2.925 2.775 4.257a11.732 11.732 0 004.247 2.765c1.528.591 3.272 1.003 5.822 1.116 2.56.112 3.375.14 9.89.14 6.516 0 7.332-.028 9.891-.14 2.55-.113 4.303-.525 5.822-1.116a11.732 11.732 0 004.247-2.765 11.732 11.732 0 002.766-4.247c.59-1.528 1.003-3.272 1.115-5.822.113-2.56.14-3.375.14-9.89 0-6.516-.027-7.332-.14-9.891-.112-2.55-.525-4.303-1.115-5.822-.591-1.594-1.407-2.935-2.747-4.266a11.732 11.732 0 00-4.247-2.765C38.194.675 36.45.262 33.9.15 31.331.028 30.516 0 24 0z" /><path d="M24 11.672c-6.806 0-12.328 5.522-12.328 12.328 0 6.806 5.522 12.328 12.328 12.328 6.806 0 12.328-5.522 12.328-12.328 0-6.806-5.522-12.328-12.328-12.328zm0 20.325a7.998 7.998 0 010-15.994 7.998 7.998 0 010 15.994zM39.694 11.184a2.879 2.879 0 11-2.878-2.878 2.885 2.885 0 012.878 2.878z" /></g><defs><clipPath id="clip0_17_63"><path d="M0 0h48v48H0z" /></clipPath></defs></svg>
          </a>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { ref, nextTick, computed, watch,onMounted} from 'vue';
import ApexCharts from 'apexcharts';
import { useRouter } from 'vue-router'  // 导入 useRouter
import axios from '../../plugins/axios';

const router = useRouter()  // 初始化 router

const tabs = ref(['淘宝', '京东', 'Favorites']);
const activeTab = ref('淘宝');

const setActiveTab = (tab) => {
  activeTab.value = tab;
};

const hasFetchedPlatformData = ref({
  '淘宝': false,
  '京东': false,
  'Favorites': false
});

const products = ref([
  // {
  //   id: 1,
  //   name: 'Fusion',
  //   category: 'UI Kit',
  //   href: '/login',
  //   price: '$49',
  //   imageSrc: 'https://flowbite.s3.amazonaws.com/docs/gallery/masonry/image-8.jpg',
  //   imageAlt:
  //     'Payment application dashboard screenshot with transaction table, financial highlights, and main clients on colorful purple background.',
  // },
  // {
  //   product_id: 1,
  //   platform_id: 1,
  //   description: "Adidas",
  //   image_url: "https://flowbite.s3.amazonaws.com/docs/gallery/masonry/image-8.jpg",
  //   price: 98.00,
  //   is_favorite: false,
  // }
]);

// 对话框是否可见的状态
const isDialogVisible = ref(false);

const navigateTo = (link) => {
  router.push(link)  // 使用 router.push 跳转到指定路径
}

const user_id = ref(localStorage.getItem('user_id'))

const searchQuery = ref('');

const footerNavs = ref([
        {
            label: "Resources",
            items: [
                {
                    href: 'javascript:void()',
                    name: 'contact'
                },
                {
                    href: 'javascript:void()',
                    name: 'Support'
                },
                {
                    href: 'javascript:void()',
                    name: 'Documentation'
                },
                {
                    href: 'javascript:void()',
                    name: 'Pricing'
                },
            ],
        },
        {
            label: "About",
            items: [
                {
                    href: 'javascript:void()',
                    name: 'Terms'
                },
                {
                    href: 'javascript:void()',
                    name: 'License'
                },
                {
                    href: 'javascript:void()',
                    name: 'Privacy'
                },
                {
                    href: 'javascript:void()',
                    name: 'About US'
                },
            ]
        },
        {
            label: "Explore",
            items: [
                {
                    href: 'javascript:void()',
                    name: 'Showcase'
                },
                {
                    href: 'javascript:void()',
                    name: 'Roadmap'
                },
                {
                    href: 'javascript:void()',
                    name: 'Languages'
                },
                {
                    href: 'javascript:void()',
                    name: 'Blog'
                },
            ]
        },
        {
            label: "Company",
            items: [
                {
                    href: 'javascript:void()',
                    name: 'Partners'
                },
                {
                    href: 'javascript:void()',
                    name: 'Team'
                },
                {
                    href: 'javascript:void()',
                    name: 'Careers'
                },
            ],
        }
      ])


      const icons = ref([
  {
    src: 'https://flowbite.s3.amazonaws.com/docs/gallery/masonry/image-8.jpg',
    link: '/cart',   // 路由路径
    notifications: 0,
  },
  {
    src: 'https://example.com/icon2.png',
    link: '/',
    notifications: 0,
  },
  {
    src: 'https://example.com/icon3.png',
    link: '/page3',
    notifications: 11,
  },
  {
    src: 'https://example.com/icon3.png',
    link: '/page3',
    notifications: 5,
  },
  {
    src: 'https://example.com/icon3.png',
    link: '/page3',
    notifications: 5,
  },
  {
    src: 'https://example.com/icon3.png',
    link: '/page3',
    notifications: 0,
  },
])

const favorites = ref(new Set());

const toggleFavorite = (product) => {
  console.log(product.IsFavorited)
  if (product.IsFavorited){
    dislikeProduct(product);
   
  }else{
    likeProduct(product);
    
  }
};

const isFavorite = (product) => {
  return product.IsFavorited
};


watch(activeTab, (newTab) => {
  if (newTab === '淘宝' && !hasFetchedPlatformData.value['淘宝']) {
    fetchProductsByPlatform('675bc9d0960209b29601aea0');
    hasFetchedPlatformData.value['淘宝'] = true;
    hasFetchedPlatformData.value['Favorites'] = false;
    hasFetchedPlatformData.value['京东'] = false;
  }

  if (newTab === '京东' && !hasFetchedPlatformData.value['京东']) {
    fetchProductsByPlatform('675bc9d0960209b29601aea1');
    hasFetchedPlatformData.value['京东'] = true;
    hasFetchedPlatformData.value['Favorites'] = false;
    hasFetchedPlatformData.value['淘宝'] = false;
  }

  if (newTab === 'Favorites' && !hasFetchedPlatformData.value['Favorites']) {
    fetchLikeProducts();
    hasFetchedPlatformData.value['Favorites'] = true;
    hasFetchedPlatformData.value['淘宝'] = false;
    hasFetchedPlatformData.value['京东'] = false;
  }
});

const filteredProducts = computed(() => {
  if (activeTab.value === '淘宝') {
    return products.value.filter((product) => product.PlatformID === '675bc9d0960209b29601aea0');
  }

  if(activeTab.value === '京东'){
    return products.value.filter((product) => product.PlatformID === '675bc9d0960209b29601aea1');
  }

  if (activeTab.value === 'Favorites') {
    return products.value;
  }

  return [];
});

const hoverIndex = ref(-1); // 当前悬停的索引

const openDialog = () => {
  isDialogVisible.value = true;

  // 使用 nextTick 确保 DOM 更新完成后再渲染图表
  nextTick(() => {
    renderChart();
  });
};

const closeDialog = () => {
  isDialogVisible.value = false;
};

const onHover = (index) => {
  hoverIndex.value = index;
};

const resetScale = () => {
  hoverIndex.value = -1;
};

// const productsPriceHistory = [
//   {
//     productName: 'Product A',
//     priceHistory: {
//       "2024-12-13 05:44:48": 155,
//       "2024-12-27 06:05:34": 156
//     }
//   },
//   {
//     productName: 'Product B',
//     priceHistory: {
//       "2024-12-13 05:44:48": 120,
//       "2024-12-27 06:05:34": 125
//     }
//   }
// ];

// // 提取所有日期
// const allDates = [...new Set(productsPriceHistory.flatMap(product => Object.keys(product.priceHistory)))];

// // 创建多个 series 数据
// const seriesData = productsPriceHistory.map(product => ({
//   name: product.productName,
//   data: allDates.map(date => product.priceHistory[date] || null)  // 如果某个日期没有数据，填充 null
// }));

const price_history = ref({
  ID :'',
  PriceHistory:{},
  ProductID:''
})


// 渲染 ApexCharts 图表
const renderChart = () => {
  const dates = Object.keys(price_history.value.PriceHistory).map(dateTime => dateTime.split(' ')[0]);  // 获取所有日期
  const prices = Object.values(price_history.value.PriceHistory); 
  console.log(price_history.value.PriceHistory)
  console.log(prices)
  console.log(dates)
  const options = {
    chart: {
      height: 300,
      type: 'area',
      toolbar: { show: false },
      zoom: { enabled: false },
    },
    // series: [
    //   { name: 'Income', data: [18000, 51000, 60000, 38000, 88000] },
    //   { name: 'Outcome', data: [27000, 38000, 60000, 77000, 40000] },
    //   { name: 'autcome', data: [26000, 39000, 60600, 77660, 33000] }
    // ],
    series: [
      { name: price_history.value.ProductID, data: prices }
    ],
    stroke: { curve: 'smooth', width: 2 },
    fill: {
      type: 'gradient',
      gradient: {
        type: 'vertical',
        shadeIntensity: 1,
        opacityFrom: 0.1,
        opacityTo: 0.8
      }
    },
    xaxis: {
      // categories: [
      //   '25 January 2023', '26 January 2023', '27 January 2023',
      //   '28 January 2023', '29 January 2023', '30 January 2023',
      //   '31 January 2023', '1 February 2023', '2 February 2023',
      //   '3 February 2023', '4 February 2023', '5 February 2023'
      // ],
      categories: dates,
      labels: {
        style: {
          colors: '#9ca3af',
          fontSize: '13px',
          fontFamily: 'Inter, ui-sans-serif',
          fontWeight: 400,
        },
        formatter: (title) => {
          if (title && typeof title === 'string') {
    let t = title.split(' ');
    if (t.length > 1) {
      return `${t[0]} ${t[1].slice(0, 3)}`;
    }
    return t[0]; // 如果没有第二部分，只返回第一个部分
  }
  return ''; // 如果 title 无效，返回空字符串
        }
      }
    },
    yaxis: {
      labels: {
        align: 'left',
        style: {
          colors: '#9ca3af',
          fontSize: '13px',
          fontFamily: 'Inter, ui-sans-serif',
          fontWeight: 400,
        },
        formatter: (value) => (value >= 1000 ? `${value / 1000}k` : value)
      }
    },
    tooltip: {
      x: { format: 'MMMM yyyy' },
      y: { formatter: (value) => `$${value >= 1000 ? `${value / 1000}k` : value}` },
    }
  };

  const chart = new ApexCharts(document.querySelector("#hs-curved-area-charts"), options);
  chart.render();
};

const isLoading = ref(false);

const errorMessage = ref('');

const fetchProducts = async ()=>
{

  isLoading.value = true;
  errorMessage.value = '';

  try {
    const userId = localStorage.getItem('userId');
   const response = await axios.get(`/products/${userId}`)

    products.value = response.data; // 假设返回的产品列表在 `response.data`
  } catch (error) {
    if (error.response) {
      const { status, data } = error.response;
      errorMessage.value = data.message || '获取数据失败';
    } else {
      errorMessage.value = '网络错误，请检查网络连接';
    }
    setTimeout(() => {
      errorMessage.value = '';
    }, 5000);
  } finally {
    isLoading.value = false;
  }
};

const successMessage = ref('');

const addToCart = async (productId) => {
  errorMessage.value = '';
  successMessage.value = '';
  try{
    const userId = localStorage.getItem('userId');
    const response = await axios.post(`/cart/${userId}/${productId}`);
    if(response.status === 201){
      successMessage.value = '添加到购物车成功';
    }
    setTimeout(() => {
      successMessage.value = '';
    }, 5000);
  }catch(error){
    if (error.response) {
      const { status, data } = error.response;
      errorMessage.value = data.message || '添加到购物车失败';
    } else {
      errorMessage.value = '网络错误，请检查网络连接';
    }
    setTimeout(() => {
      errorMessage.value = '';
    }, 5000);
  }
}

const likeProduct = async (product) => {
  errorMessage.value = '';
  successMessage.value = '';
  try{
    const userId = localStorage.getItem('userId');
    const response = await axios.put(`/products/like/${userId}/${product.product_id}`);
    fetchProductsByPlatform(product.PlatformID)
    if(response.status === 200){
      successMessage.value = '添加到收藏夹成功';
    }
    setTimeout(() => {
      successMessage.value = '';
    }, 5000);
  }catch(error){
    if (error.response) {
      const { status, data } = error.response;
      errorMessage.value = data.message || '添加到收藏夹失败';
    } else {
      errorMessage.value = '网络错误，请检查网络连接';
    }
    setTimeout(() => {
      errorMessage.value = '';
    }, 5000);
  }
}

const dislikeProduct = async (product) => {
  errorMessage.value = '';
  successMessage.value = '';
  try{
    const userId = localStorage.getItem('userId');
    const response = await axios.put(`/products/dislike/${userId}/${product.product_id}`);
    if(activeTab.value === 'Favorites'){
      fetchLikeProducts();
    }
    else{
    fetchProductsByPlatform(product.PlatformID)
    }
    if(response.status === 200){
      successMessage.value = '从收藏夹移除成功';
    }
    setTimeout(() => {
      successMessage.value = '';
    }, 5000);
  }catch(error){
    if (error.response) {
      const { status, data } = error.response;
      errorMessage.value = data.message || '从收藏夹移除失败';
    } else {
      errorMessage.value = '网络错误，请检查网络连接';
    }
    setTimeout(() => {
      errorMessage.value = '';
    }, 5000);
  }
}

const fetchLikeProducts = async () => {
  isLoading.value = true;
  errorMessage.value = '';

  try {
    const userId = localStorage.getItem('userId');
    const response = await axios.get(`/products/likes/${userId}`);
    products.value = response.data;
  } catch (error) {
    if (error.response) {
      const { status, data } = error.response;
      errorMessage.value = data.message || '获取数据失败';
    } else {
      errorMessage.value = '网络错误，请检查网络连接';
    }
    setTimeout(() => {
      errorMessage.value = '';
    }, 5000);
  } finally {
    isLoading.value = false;
  }
};

const fetchProductsByPlatform = async (platformId) => {
  isLoading.value = true;
  errorMessage.value = '';

  try {
    const userId = localStorage.getItem('userId');
    const response = await axios.get(`/products/${userId}/${platformId}`);
    products.value = response.data;
    console.log(products.value);
  } catch (error) {
    if (error.response) {
      const { status, data } = error.response;
      errorMessage.value = data.message || '获取数据失败';
    } else {
      errorMessage.value = '网络错误，请检查网络连接';
    }
    setTimeout(() => {
      errorMessage.value = '';
    }, 5000);
  } finally {
    isLoading.value = false;
  }
};

const SearchProducts = async (query) => {
  isLoading.value = true;
  errorMessage.value = '';

  if (!query) {
    fetchProducts();
    return;
  }
  try {
    const userId = localStorage.getItem('userId');
    const response = await axios.get(`/products/search/${userId}`, {
      params: {
        query
      }
    });
    products.value = response.data;
  } catch (error) {
    if (error.response) {
      const { status, data } = error.response;
      errorMessage.value = data.message || '获取数据失败';
    } else {
      errorMessage.value = '网络错误，请检查网络连接';
    }
    setTimeout(() => {
      errorMessage.value = '';
    }, 5000);
  } finally {
    isLoading.value = false;
  }
};

const getDetail = async (productId) => {
  isLoading.value = true;
  errorMessage.value = '';

  try {
    const userId = localStorage.getItem('userId');
    const response = await axios.get(`/products/details/${userId}/${productId}`);
    price_history.value = response.data.price_history;
    console.log(price_history.value);
    openDialog();
  } catch (error) {
    if (error.response) {
      const { status, data } = error.response;
      errorMessage.value = data.message || '获取数据失败';
    } else {
      errorMessage.value = '网络错误，请检查网络连接';
    }
    setTimeout(() => {
      errorMessage.value = '';
    }, 5000);
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  fetchProductsByPlatform('675bc9d0960209b29601aea0');
});

</script>

<!-- <script>


import ApexCharts from 'apexcharts';
import Statistic from './compare_pic.vue';
import { Static } from 'vue';

export default {
  name: 'App',
  mounted() {
    // 等待 DOM 渲染完毕后执行
    this.renderChart();
  },
  Components: {
    Statistic
  },
  data() {
    return {
        posts: [
        {
          title: "What is SaaS? Software as a Service Explained11111111111111111111111111111111111111111111111111",
          desc: "Going into this journey, I had a standard therapy regimen, based on looking at the research literature. After I saw the movie, I started to ask other people what they did for their anxiety, and some1111111111111111111111",
          img: "https://flowbite.s3.amazonaws.com/docs/gallery/masonry/image-11.jpg",
          authorLogo: "https://api.uifaces.co/our-content/donated/xZ4wg2Xj.jpg",
          authorName: "Sidi dev",
          date: "Jan 4 2022",
          href: "javascript:void(0)"
        },
        {
          title: "A Quick Guide to WordPress Hosting",
          desc: "According to him, â€œI'm still surprised that this has happened. But we are surprised because we are so surprised.â€More revelations about Whittington will be featured in the film",
          img: "https://images.unsplash.com/photo-1620287341056-49a2f1ab2fdc?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
          authorLogo: "https://api.uifaces.co/our-content/donated/FJkauyEa.jpg",
          authorName: "Micheal",
          date: "Jan 4 2022",
          href: "javascript:void(0)"
        },
        {
          title: "7 Promising VS Code Extensions Introduced in 2022",
          desc: "I hope I remembered all the stuff that they needed to know. They're like, 'okay,' and write it in their little reading notebooks. I realized today that I have all this stuff that",
          img: "https://images.unsplash.com/photo-1517694712202-14dd9538aa97?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
          authorLogo: "https://randomuser.me/api/portraits/men/46.jpg",
          authorName: "Luis",
          date: "Jan 4 2022",
          href: "javascript:void(0)"
        },
        {
          title: "How to Use Root C++ Interpreter Shell to Write C++ Programs",
          desc: "The powerful gravity waves resulting from the impact of the planets' moons â€” four in total â€” were finally resolved in 2015 when gravitational microlensing was used to observe the",
          img: "https://images.unsplash.com/photo-1617529497471-9218633199c0?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
          authorLogo: "https://api.uifaces.co/our-content/donated/KtCFjlD4.jpg",
          authorName: "Lourin",
          date: "Jan 4 2022",
          href: "javascript:void(0)"
        },
        {
          title: "How to Use Root C++ Interpreter Shell to Write C++ Programs",
          desc: "The powerful gravity waves resulting from the impact of the planets' moons â€” four in total â€” were finally resolved in 2015 when gravitational microlensing was used to observe the",
          img: "https://images.unsplash.com/photo-1617529497471-9218633199c0?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
          authorLogo: "https://api.uifaces.co/our-content/donated/KtCFjlD4.jpg",
          authorName: "Lourin",
          date: "Jan 4 2022",
          href: "javascript:void(0)"
        },
        {
          title: "How to Use Root C++ Interpreter Shell to Write C++ Programs",
          desc: "The powerful gravity waves resulting from the impact of the planets' moons â€” four in total â€” were finally resolved in 2015 when gravitational microlensing was used to observe the",
          img: "https://images.unsplash.com/photo-1617529497471-9218633199c0?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80",
          authorLogo: "https://api.uifaces.co/our-content/donated/KtCFjlD4.jpg",
          authorName: "Lourin",
          date: "Jan 4 2022",
          href: "javascript:void(0)"
        }
      ],
        icons: [
        '', // Example Icon
        'https://cdn-icons-png.flaticon.com/512/179/179328.png', // Example Icon
        '', // Example Icon
        '', // Example Icon
        ''  // Example Icon
      ],
      hoverIndex: -1, // 当前悬停的索引
    }
  },
 
  methods: {
    onHover(index) {
      this.hoverIndex = index;
    },
    resetScale() {
      this.hoverIndex = -1;
    },
    renderChart() {
  const options = {
    chart: {
      height: 300,
      type: 'area',
      toolbar: { show: false },
      zoom: { enabled: false },
    },
    series: [
      { name: 'Income', data: [18000, 51000, 60000, 38000, 88000] },
      { name: 'Outcome', data: [27000, 38000, 60000, 77000, 40000] },
      { name: 'Oucome', data: [26000, 39000, 60600, 77660, 33000] }
    ],
    stroke: { curve: 'smooth', width: 2 },
    fill: {
      type: 'gradient',
      gradient: {
        type: 'vertical',
        shadeIntensity: 1,
        opacityFrom: 0.1,
        opacityTo: 0.8
      }
    },
    xaxis: {
      categories: [
        '25 January 2023', '26 January 2023', '27 January 2023',
        '28 January 2023', '29 January 2023', '30 January 2023',
        '31 January 2023', '1 February 2023', '2 February 2023',
        '3 February 2023', '4 February 2023', '5 February 2023'
      ],
      labels: {
        style: {
          colors: '#9ca3af',
          fontSize: '13px',
          fontFamily: 'Inter, ui-sans-serif',
          fontWeight: 400,
        },
        formatter: (title) => {
          // 确保 title 存在
          if (title) {
            let t = title.split(' ');
            return `${t[0]} ${t[1].slice(0, 3)}`;
          }
          return ''; // 如果 title 不存在，返回空字符串
        }
      }
    },
    yaxis: {
      labels: {
        align: 'left',
        style: {
          colors: '#9ca3af',
          fontSize: '13px',
          fontFamily: 'Inter, ui-sans-serif',
          fontWeight: 400,
        },
        formatter: (value) => (value >= 1000 ? `${value / 1000}k` : value)
      }
    },
    tooltip: {
      x: { format: 'MMMM yyyy' },
      y: { formatter: (value) => `$${value >= 1000 ? `${value / 1000}k` : value}` },
    }
  };

  const chart = new ApexCharts(document.querySelector("#hs-curved-area-charts"), options);
  chart.render();
}
  }
}
</script> -->

<style scoped>
.dock-container {
  position: fixed;
  bottom: 20px; /* 让 dock 离底部远一点 */
  width: 100%;
}

.dock {
  background-color: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(8px);
  border-radius: 10px;
  padding: 10px;
}

.dock-item {
  cursor: pointer;
  transition: transform 0.2s ease-in-out;
}

.dock-item:hover {
  filter: drop-shadow(0 4px 6px rgba(0, 0, 0, 0.1));
}

.dock-icon {
  transition: box-shadow 0.3s ease;
  box-shadow: 0px 3px 8px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
}
</style>
