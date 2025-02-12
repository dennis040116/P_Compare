<template>
  <div>
    <!-- 按钮触发对话框显示 -->
    <button @click="openDialog" class="w-32 mx-auto py-2 ml-2 shadow-sm rounded-md bg-indigo-600 text-white mt-4 flex items-center justify-center">
      Click me
    </button>

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
        <h3 class="text-lg font-medium text-gray-800 mb-4">Terms and agreements</h3>

        <!-- 图表容器 -->
        <div id="hs-curved-area-charts"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, nextTick } from 'vue';
import ApexCharts from 'apexcharts';

// 对话框是否可见的状态
const isDialogVisible = ref(false);

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

// 渲染 ApexCharts 图表
const renderChart = () => {
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
      { name: 'Outcome', data: [26000, 39000, 60600, 77660, 88000] }
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
          if (title) {
            let t = title.split(' ');
            return `${t[0]} ${t[1].slice(0, 3)}`;
          }
          return '';
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
</script>

<style scoped>

</style>