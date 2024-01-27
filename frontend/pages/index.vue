<template>
    <div class="w-full h-screen">
        <div class="h-auto flex flex-row justify-end align-middle">
            <Sheet v-if="userStore.isLogging">
                <SheetTrigger>
                    <TooltipProvider>
                        <Tooltip>
                            <TooltipTrigger>
                                <div class="flex flex-row items-center">
                                    <p v-if="bagStore.totalItems > 0"
                                        class="text-xl font-inter font-extralight mr-1 cursor-pointer">3</p>
                                    <ShoppingBag @click="isLogging = false" :size="28"
                                        :class="bagStore.totalItems > 0 ? '' : 'mt-2'" class="mr-4 cursor-pointer" />
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
                <div v-for="product in productsWithStock" :key="product.id"
                    class="w-1/4 h-1/2 flex flex-col justify-center items-center">
                    <ProductCard :product="product" />
                </div>
            </div>
        </div>
    </div>
</template>
<script setup>
import { ContentSheetLoginRegister, ContentSheetShoppingBag, ProductCard } from '#components';
import { Menu, ShoppingBag } from 'lucide-vue-next';
import { useBagStore } from '../store/bag';
import { useProductsStore } from '../store/products';
import { useUserStore } from '../store/user';

const config = useRuntimeConfig()
const userStore = useUserStore()
const bagStore = useBagStore()
const productStore = useProductsStore()

// const products = ref([]);
// const stocks = ref([]);
// const fetchProducts = async () => {
//     const res = await $fetch(`${config.public.apiBase}/v1/product`, {
//         method: 'GET',
//         mode: 'no-cors'
//     });
//     products.value = res.Data;

//     const res2 = await $fetch(`${config.public.apiBase}/v1/stock`, {
//         method: 'GET',
//         mode: 'no-cors'
//     });
//     stocks.value = res2.Data;
// };

// onMounted(() => {
//     fetchProducts();
// });

const products = ref([
    {
        "ProductID": 1,
        "Name": "Test",
        "Pricing": 5,
        "Description": "test"
    },
    {
        "ProductID": 2,
        "Name": "Test2",
        "Pricing": 10,
        "Description": "test"
    }
])

const stocks = ref([
    {
        "ProductStockID": 1,
        "ProductID": 1,
        "Quantity": 0
    },
    {
        "ProductStockID": 2,
        "ProductID": 2,
        "Quantity": 5
    }
])

const productsWithStock = computed(() => {
    const temp = products.value.map(product => {
        const stockItem = stocks.value.find(stockItem => stockItem.ProductID === product.ProductID)
        return {
            ...product,
            stock: stockItem?.Quantity || 0
        }
    })
    productStore.updateProducts(temp);
    return temp;
})
</script>