<template>
  <div class="w-full h-screen">
    <Toaster richColors position="top-left" />
    <div class="h-auto flex flex-row justify-end align-middle">
      <Sheet v-if="userStore.isLogging">
        <SheetTrigger>
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger>
                <div class="flex flex-row items-center">
                  <p v-if="bagStore.totalItems > 0" class="text-xl font-inter font-extralight mr-1 cursor-pointer">{{
        bagStore.totalItems }}</p>
                  <ShoppingBag @click="isLogging = false" :size="28" :class="bagStore.totalItems > 0 ? '' : 'mt-2'"
                    class="mr-4 cursor-pointer" />
                </div>
              </TooltipTrigger>
              <TooltipContent side="bottom" delay="100">
                <p>Open shooping bag</p>
              </TooltipContent>
            </Tooltip>
          </TooltipProvider>
        </SheetTrigger>
        <ContentSheetShoppingBag />
      </Sheet>
      <Sheet>
        <SheetTrigger>
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger>
                <Menu @click="isLogging = true" :size="32" class="mr-2 mt-2 cursor-pointer" />
              </TooltipTrigger>
              <TooltipContent side="bottom" delay="100">
                <p>Login</p>
              </TooltipContent>
            </Tooltip>
          </TooltipProvider>
        </SheetTrigger>
        <ContentSheetLoginRegister />
      </Sheet>
    </div>
    <div class="w-full flex flex-row flex-wrap min-h-full h-full overflow-y-auto">
      <p class="text-5xl font-inter font-extralight ml-4">Products</p>
      <div class="w-full flex flex-row flex-wrap justify-center">
        <p v-if="data">{{ data || 'PEEP' }}</p>
        <p v-if="error">{{ error }}</p>s
        <!-- <div v-for="product in productsWithStock" :key="product.id"
          class="w-1/4 h-1/2 flex flex-col justify-center items-center">
          <ProductCard :product="product" />
        </div> -->
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, computed, onMounted } from 'vue';
import { ContentSheetLoginRegister, ContentSheetShoppingBag, ProductCard } from '#components';
import { Toaster } from '@/components/ui/sonner';
import { Menu, ShoppingBag } from 'lucide-vue-next';
import { useBagStore } from '../store/bag';
import { useProductsStore } from '../store/products';
import { useUserStore } from '../store/user';

const config = useRuntimeConfig();
const userStore = useUserStore();
const bagStore = useBagStore();
const productStore = useProductsStore();

const { data, pending, error, status } = await useFetch('https://localhost/api/v1/product', {
  method: 'GET',
  mode: 'no-cors'
},
  { server: false, inmediate: false, default: [] }
)


const stocks = ref([]);
const loading = ref(true);

// const fetchProducts = async () => {
//   // const { data: products } = await useFetch(`${config.public.apiBase}/v1/product/`, { method: 'GET', mode: 'no-cors' });
//   // console.log('products:', products);

//   // try {
//   //   const productPromise = $fetch(`${config.public.apiBase}/v1/product/`, { method: 'GET', mode: 'no-cors' });
//   //   const stockPromise = $fetch(`${config.public.apiBase}/v1/stock/`, { method: 'GET', mode: 'no-cors' });

//   //   const [productRes, stockRes] = await Promise.all([productPromise, stockPromise]);
//   //   console.log('productRes:', productRes);
//   //   products.value = productRes.Data;
//   //   stocks.value = stockRes.Data;
//   // } catch (error) {
//   //   console.error('Failed to fetch data:', error);
//   // } finally {
//   //   loading.value = false;
//   // }
// };

// onMounted(fetchProducts);

const productsWithStock = computed(() => {
  if (products.value.length === 0) return [];
  const temp = products.value.map(product => {
    const stockItem = stocks.value.find(stockItem => stockItem.ProductID === product.ProductID);
    return {
      ...product,
      stock: stockItem?.Quantity || 0
    };
  });
  productStore.updateProducts(temp);
  return temp;
});
</script>
