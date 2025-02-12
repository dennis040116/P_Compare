<!-- This example requires Tailwind CSS v2.0+ -->
<template>




      <div class="max-w-2xl mx-auto py-16 px-4 sm:py-24 sm:px-6 lg:px-0">
        <h1 class="text-3xl font-extrabold text-center tracking-tight text-gray-900 sm:text-4xl">Shopping Cart</h1>
  
        <div class="mt-12">
          <section aria-labelledby="cart-heading">
            <h2 id="cart-heading" class="sr-only">Items in your shopping cart</h2>
  
            <ul role="list" class="border-t border-b border-gray-200 divide-y divide-gray-200">
              <li v-for="product in cartProducts" :key="product.id" class="flex py-6">
                <div class="flex-shrink-0">
                  <img :src="product.ImageURL" class="w-24 h-24 rounded-md object-center object-cover sm:w-32 sm:h-32" />
                </div>
  
                <div class="ml-4 flex-1 flex flex-col sm:ml-6">
                  <div>
                    <div class="flex justify-between">
                      <h4 class="text-sm">
                        <a :href="product.url" class="font-medium text-gray-700 hover:text-gray-800">
                          {{ product.Description }}
                        </a>
                      </h4>
                      <p class="ml-4 text-sm font-medium text-gray-900">{{ product.Price }}</p>
                    </div>
                    <p class="mt-1 text-sm text-gray-500">
                      {{ platformName(product.PlatformID) }}
                    </p>
                    <p class="mt-1 text-sm text-gray-500">
                      {{ }}
                    </p>
                  </div>
  
                  <div class="mt-4 flex-1 flex items-end justify-between">
                    <p class="flex items-center text-sm text-gray-700 space-x-2">
                      <CheckIcon v-if="product.inStock" class="flex-shrink-0 h-5 w-5 text-green-500" aria-hidden="true" />
                      <ClockIcon v-else class="flex-shrink-0 h-5 w-5 text-gray-300" aria-hidden="true" />
                      <span>{{}}</span>
                    </p>
                    <div class="ml-4">
                      <button @click = removeProduct(product.product_id) type="button" class="text-sm font-medium text-indigo-600 hover:text-indigo-500">
                        <span>Remove</span>
                      </button>
                    </div>
                  </div>
                </div>
              </li>
            </ul>
          </section>
  
          <!-- Order summary -->
          <section aria-labelledby="summary-heading" class="mt-10">

            <!-- <h2 id="summary-heading" class="sr-only">Order summary</h2>
  
            <div>
              <dl class="space-y-4">
                <div class="flex items-center justify-between">
                  <dt class="text-base font-medium text-gray-900">Subtotal</dt>
                  <dd class="ml-4 text-base font-medium text-gray-900">$96.00</dd>
                </div>
              </dl>
              <p class="mt-1 text-sm text-gray-500">Shipping and taxes will be calculated at checkout.</p>
            </div> -->
  
            <div class="mt-10">
              <button @click="compare()" class="w-full bg-indigo-600 border border-transparent rounded-md shadow-sm py-3 px-4 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-50 focus:ring-indigo-500">Go Compare</button>
            </div>
            <div class="mt-6 text-sm text-center">
              <p>
                or <a href="/products" class="text-indigo-600 font-medium hover:text-indigo-500">Continue Shopping<span aria-hidden="true"> &rarr;</span></a>
              </p>
            </div>
          </section>
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
        <h3 class="text-lg font-medium text-gray-800 mb-4">Price Compare</h3>

        <!-- 图表容器 -->
        <div id="hs-curved-area-charts"></div>
      </div>
    </div>
   
 

  </template>
  
  <script setup>
  import { ref, nextTick, onMounted } from 'vue';
  import ApexCharts from 'apexcharts';
  import axios from '../../plugins/axios';
  
  // const products = ref([
  //   {
  //     id: 1,
  //     name: 'Fusion',
  //     category: 'UI Kit',
  //     href: '/login',
  //     price: '$49',
  //     imageSrc: 'https://flowbite.s3.amazonaws.com/docs/gallery/masonry/image-8.jpg',
  //     imageAlt:
  //       'Payment application dashboard screenshot with transaction table, financial highlights, and main clients on colorful purple background.',
  //   },
  //   {
  //     id: 2,
  //     name: 'Fusion',
  //     category: 'UI Kit',
  //     href: '#',
  //     price: '$49',
  //     imageSrc: 'https://tailwindui.com/img/ecommerce-images/product-page-05-related-product-01.jpg',
  //     imageAlt:
  //       'Payment application dashboard screenshot with transaction table, financial highlights, and main clients on colorful purple background.',
  //   },
  //   {
  //     id: 3,
  //     name: 'Fusion',
  //     category: 'UI Kit',
  //     href: '#',
  //     price: '$49',
  //     imageSrc: 'https://tailwindui.com/img/ecommerce-images/product-page-05-related-product-01.jpg',
  //     imageAlt:
  //       'Payment application dashboard screenshot with transaction table, financial highlights, and main clients on colorful purple background.',
  //   },
  //   {
  //     id: 4,
  //     name: 'Fusion',
  //     category: 'UI Kit',
  //     href: '#',
  //     price: '$49',
  //     imageSrc: 'https://tailwindui.com/img/ecommerce-images/product-page-05-related-product-01.jpg',
  //     imageAlt:
  //       'Payment application dashboard screenshot with transaction table, financial highlights, and main clients on colorful purple background.',
  //   },
  //   {
  //     id: 5,
  //     name: 'Fusion',
  //     category: 'UI Kit',
  //     href: '#',
  //     price: '$49',
  //     imageSrc: 'https://tailwindui.com/img/ecommerce-images/product-page-05-related-product-01.jpg',
  //     imageAlt:
  //       'Payment application dashboard screenshot with transaction table, financial highlights, and main clients on colorful purple background.',
  //   },
  //   {
  //     id: 6,
  //     name: 'Fusion',
  //     category: 'UI Kit',
  //     href: '#',
  //     price: '$49',
  //     imageSrc: 'https://tailwindui.com/img/ecommerce-images/product-page-05-related-product-01.jpg',
  //     imageAlt:
  //       'Payment application dashboard screenshot with transaction table, financial highlights, and main clients on colorful purple background.',
  //   },
  // ]);
  
  // 对话框是否可见的状态
  const isDialogVisible = ref(false);
  
  
const platformName = (platform_id) =>
{
  if(platform_id === '675bc9d0960209b29601aea0')
  {
    return 'TaoBao'
  }

  if(platform_id === '675bc9d0960209b29601aea1')
  {
    return 'JingDong'
  }

  return ''
}
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
  
const errorMessage = ref(''); // 错误消息

const cartProducts = ref([
  // {
  //   product_id: 1,
  //   platform_id: 1,
  //   description: 'Fusion',
  //   image_url: 'https://flowbite.s3.amazonaws.com/docs/gallery/masonry/image-8.jpg',
  //   price: 49,
  // }
]); // 购物车商品

  const getCart = async() =>{
    isLoading.value = true;
  errorMessage.value = '';
    console.log('getCart')
  try {
    const userId = localStorage.getItem('userId');
    console.log(userId)
    const response = await axios.get(`/cart/${userId}`);
    console.log(response.data)
    cartProducts.value = response.data;
    console.log(response.data);
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
  }

  const isLoading = ref('')

  const removeProduct = async(productId) => {
    isLoading.value = true;
    errorMessage.value = '';

    try {
      const userId = localStorage.getItem('userId');
      await axios.delete(`/cart/delete/${userId}/${productId}`);
      getCart();
    } catch (error) {
      if (error.response) {
        const { status, data } = error.response;
        errorMessage.value = data.message || '删除商品失败';
      } else {
        errorMessage.value = '网络错误，请检查网络连接';
      }
      setTimeout(() => {
        errorMessage.value = '';
      }, 5000);
    } finally {
      isLoading.value = false;
    }
  }

  const compare_price = ref([])
  
  const compare = async() => {
    isLoading.value = true;
    errorMessage.value = '';
  
    try {
      const userId = localStorage.getItem('userId');
      const response = await axios.get(`/cart/compare/${userId}`);
      compare_price.value = response.data;
      console.log(response.data);
      openDialog();
    } catch (error) {
      if (error.response) {
        const { status, data } = error.response;
        errorMessage.value = data.message || '比价失败';
      } else {
        errorMessage.value = '网络错误，请检查网络连接';
      }
      setTimeout(() => {
        errorMessage.value = '';
      }, 5000);
    } finally {
      isLoading.value = false;
    }
  }


  
  // 渲染 ApexCharts 图表
  const renderChart = () => {
    const allDates = [...new Set(compare_price.value.flatMap(product => Object.keys(product.PriceHistory)))];

// 2. 创建多个 series 数据
const seriesData = compare_price.value.map(product => ({
  name: product.ProductID, // 产品名称
  data: allDates.map(date => product.PriceHistory[date] || null)  // 如果某个日期没有数据，填充 null
}));
console.log(allDates)
console.log(seriesData)
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
      series: seriesData,
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
        categories: allDates,
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

  onMounted(() => {
    getCart();
  })
  </script>