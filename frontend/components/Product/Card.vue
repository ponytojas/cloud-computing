<template>
    <Card class="w-72">
        <CardHeader>
            <CardTitle>{{ product.Name }}</CardTitle>
            <CardDescription>{{ product.Description }}</CardDescription>
        </CardHeader>
        <CardContent>
            <div class="flex flex-row">
                <div class="mr-3">
                    <Package v-if="product?.stock > 0" />
                    <PackageOpen v-if="product?.stock <= 0" />
                </div>
                <p class="text-xl">
                    Stock: {{ product?.stock }}
                </p>
            </div>
        </CardContent>
        <CardFooter>
            <div class="flex flex-row justify-between align-middle items-center">
                <div class="flex flex-row align-middle items-center">
                    <DollarSign :size="20" />
                    <p class="text-xl">{{ product.Pricing }}</p>
                </div>
                <Button v-if="userStore.isLogging">
                    <ShoppingCart class="mr-4" />
                    Add to cart
                </Button>
            </div>
        </CardFooter>
    </Card>
</template>

<script setup>
import { DollarSign, ShoppingCart, Package, PackageOpen } from 'lucide-vue-next';
import { useUserStore } from '../../store/user';
const userStore = useUserStore()

const { product } = defineProps({
    product: {
        type: Object,
        required: true
    }
})
</script>