<template>
    <SheetContent>
        <SheetHeader>
            <SheetTitle>Shopping bag</SheetTitle>
            <div v-for="item in bagStore.items" class="flex flex-row justify-between">
                <p>{{ item.Name }}</p>
                <div class="flex flex-row align-middle items-center">
                    <p>{{ item.Added }}</p>
                    <span>x</span>
                    <p> {{ item.Pricing }}</p>
                    <DollarSign :size="16" />
                </div>
            </div>
            <div class="flex flex-row mt-6 mb-2 align-middle items-center">
                <p class="text-xl">Total: {{ bagStore.total }}</p>
                <DollarSign :size="18" />
            </div>
            <div v-if="bagStore.items.length > 0" class="flex flex-row-reverse justify-around mt-2">
                <Button @click="bagStore.clearBag" variant="outlined">Clear</Button>
                <Dialog>
                    <DialogTrigger as-child>
                        <Button>
                            <CreditCard class="mr-2" color="#F79F77" />
                            <p>Payout</p>
                        </Button>

                    </DialogTrigger>
                    <DialogContent class="sm:max-w-md">
                        <DialogHeader>
                            <DialogTitle>Payment</DialogTitle>
                            <DialogDescription>
                                Add your credit card
                            </DialogDescription>
                        </DialogHeader>
                        <div class="flex items-center space-x-2">
                            <div class="flex flex-1 gap-2 flex-row items-center">
                                <WalletCards :size="36" color="#90D8CC" />
                                <Input id="link" default-value="47** **** **** ****" read-only />
                            </div>
                        </div>
                        <DialogFooter class="sm:justify-start">
                            <div class="flex flex-row justify-between w-full">
                                <DialogClose as-child>
                                    <Button type="button" variant="outlined">
                                        Close
                                    </Button>
                                </DialogClose>
                                <DialogClose as-child>
                                    <Button @click="pay" type="button" variant="secondary">
                                        Continue
                                    </Button>
                                </DialogClose>
                            </div>
                        </DialogFooter>
                    </DialogContent>
                </Dialog>
            </div>
        </SheetHeader>
    </SheetContent>
</template>

<script setup>
import { CreditCard, DollarSign, WalletCards } from 'lucide-vue-next';
import { toast } from 'vue-sonner';

import { useBagStore } from '../../store/bag';
import { useUserStore } from '../../store/user';

const config = useRuntimeConfig()

const bagStore = useBagStore()
const userStore = useUserStore()

const clearBag = async () => {
    const token = userStore.user.Token;
    const userId = userStore.user.UserId;
    const res = await $fetch(`${config.public.cartBase}/v1/${userId}`, {
        method: 'DELETE',
        headers: {
            Authorization: `${token}`
        },
    });
    if (!res) {
        toast.error('Something went wrong')
        return;
    }
    bagStore.clearBag()
}

const pay = async () => {
    const token = userStore.token;
    const userId = userStore.user.UserId;
    const res = await $fetch(`${config.public.paymentBase}/v1/pay/${userId}`, {
        method: 'POST',
        headers: {
            Authorization: `${token}`
        },
    });
    if (!res) {
        toast.error('Something went wrong')
        return;
    }
    toast.success('Payment successful')
    bagStore.clearBag()
}
</script>