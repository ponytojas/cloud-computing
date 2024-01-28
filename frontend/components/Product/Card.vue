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
                <Button v-if="userStore.isLogging && product?.stock > 0" @click="addToCart">
                    <ShoppingCart class="mr-4" />
                    Add to cart
                </Button>
                <Button v-else-if="userStore.isLogging && product?.stock <= 0" variant="outline" class="cursor-not-allowed">
                    <HeartCrack class="mr-4" color="red" />
                    Not available
                </Button>
            </div>
        </CardFooter>
    </Card>
</template>

<script setup>
import { DollarSign, HeartCrack, Package, PackageOpen, ShoppingCart } from 'lucide-vue-next';
import { toast } from 'vue-sonner';

import { useBagStore } from '../../store/bag';
import { useUserStore } from '../../store/user';
const headers = useRequestHeaders(['Authorization'])
const userStore = useUserStore()
const bagStore = useBagStore()
const config = useRuntimeConfig()

const { product } = defineProps({
    product: {
        type: Object,
        required: true
    }
})

const addToCart = async () => {
    const token = userStore.token;
    const userId = userStore.user.UserId;
    const productId = product.ProductID;
    const quantity = bagStore.howManyInBag(product.ProductID) === 0 ? 1 : bagStore.howManyInBag(product.ProductID) + 1;
    const body = { userId, productId, quantity };
    const res = await $fetch(`${config.public.cartBase}/v1/add-to-cart`, {
        method: 'POST',
        body: JSON.stringify(body),
        headers: {
            Authorization: `${token}`
        },
    });

    if (!res) {
        toast.error('Something went wrong')
        return;
    }
    if (res.status === 'OK') {
        toast.success('Added to the bag')
        bagStore.add(product)
    }

}
</script>