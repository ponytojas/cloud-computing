<template>
    <SheetContent>
        <SheetTitle>
            <p v-if="!userStore?.user?.Username" class="text-center mb-1">Get inside the app</p>
            <p v-else class="text-center mb-1">Welcome {{ userStore?.user?.Username }}</p>
        </SheetTitle>
        <div v-if="!userStore.isLogging">
            <Tabs default-value="login">
                <div class="flex flex-col justify-center w-full items-center mb-2">
                    <TabsList>
                        <TabsTrigger value="login" @click="tab = 'login'">
                            Login
                        </TabsTrigger>
                        <TabsTrigger value="register" @click="tab = 'register'">
                            Register
                        </TabsTrigger>
                    </TabsList>
                </div>
                <TabsContent value="login">
                    <form>
                        <FormField name="username">
                            <FormItem class="mb-4">
                                <FormLabel>Username</FormLabel>
                                <FormControl>
                                    <Input type="text" placeholder="rickAstley" v-model="username" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        </FormField>
                        <FormField name="password">
                            <FormItem class="mb-4">
                                <FormLabel>Password</FormLabel>
                                <FormControl>
                                    <Input type="password" placeholder="" v-model="password" autocomplete="on" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        </FormField>
                    </form>
                </TabsContent>
                <TabsContent value="register">
                    <form>
                        <FormField name="username">
                            <FormItem class="mb-4">
                                <FormLabel>Username</FormLabel>
                                <FormControl>
                                    <Input type="text" placeholder="jonhdoe" v-model="username" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        </FormField>
                        <FormField name="email">
                            <FormItem class="mb-4">
                                <FormLabel>Email</FormLabel>
                                <FormControl>
                                    <Input type="email" placeholder="example@example.com" v-model="email" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        </FormField>
                        <FormField name="password">
                            <FormItem class="mb-4">
                                <FormLabel>Password</FormLabel>
                                <FormControl>
                                    <Input type="password" placeholder="" v-model="password" autocomplete="on" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        </FormField>
                    </form>
                </TabsContent>
                <div class="flex flex-row justify-center">
                    <Button class="mt-6" :class="loading ? 'bg-gray-300 text-white hover:bg-gray-300 hover:text-white cursor-wait'
                        : 'bg-green-600 text-white hover:bg-green-500 hover:text-black'" @click="onSubmit">
                        <span v-if="!loading">Submit</span>
                        <div v-else class="flex flex-row">
                            <Loader2 class="w-4 h-4 mr-2 animate-spin" />
                            Please wait
                        </div>
                    </Button>
                </div>
            </Tabs>
        </div>
        <div v-else class="flex flex-row align-middle justify-center mt-6">
            <Button variant="destructive">
                <p @click="logout" class="">Logout</p>
            </Button>
        </div>
    </SheetContent>
</template>

<script setup>
import { Loader2 } from 'lucide-vue-next';
import { toast } from 'vue-sonner';

import { useBagStore } from '../../store/bag';
import { useUserStore } from '../../store/user';

const config = useRuntimeConfig()

const userStore = useUserStore()
const bagStore = useBagStore()

const tab = ref('login')
const loading = ref(false)

const username = ref('')
const email = ref('')
const password = ref('')

const errorToastMissingFields = () => {
    toast.error('Missing fields', {
        description: `The following fields are required: ${tab.value === 'login' ? 'username and password' : 'username, email and password'}`,
    })
}

const saveToken = (data) => {
    userStore.login(data)
    loading.value = false
    toast.success('Logged in', {
        description: `You are now logged in`,
    })
}

const onSubmit = async () => {
    if (!username.value || !password.value) {
        errorToastMissingFields()
        return
    }
    if (tab.value === 'register' && !email.value) {
        errorToastMissingFields()
        return
    }
    let url = tab.value === 'login' ? `${config.public.apiBase}/v1/user/login` : `${config.public.apiBase}/v1/user/register`
    const body = tab.value === 'login' ? { username: username.value, password: password.value } : { username: username.value, email: email.value, password: password.value }
    loading.value = true
    const res = await $fetch(url, {
        method: 'POST',
        body: JSON.stringify(body),
        mode: 'no-cors'
    });
    if (!res) {
        toast.error('Something went wrong', {
            description: `Please review your credentials and try again`,
        })
        loading.value = false
        return
    }
    if (res.status === 'OK' && tab.value === 'register') {
        toast.success('Account created', {
            description: `Your account has been created`,
        })
        url = `${config.public.apiBase}/v1/user/login`
        const res2 = await $fetch(url, {
            method: 'POST',
            body: JSON.stringify({ username: username.value, password: password.value }),
            mode: 'no-cors'
        });
        saveToken(res2)
    } else {
        saveToken(res)
    }
}

const logout = () => {
    userStore.logout()
    bagStore.clearBag()
}

</script>