<template>
    <SheetContent>
        <div v-if="!userStore.isLogging">
            <Tabs default-value="login">
                <div class="flex flex-col justify-center w-full items-center mb-2">
                    <TabsList>
                        <TabsTrigger value="login">
                            Login
                        </TabsTrigger>
                        <TabsTrigger value="register">
                            Register
                        </TabsTrigger>
                    </TabsList>
                </div>
                <TabsContent value="login">
                    <FormField v-slot="{ componentField }" name="username">
                        <FormItem class="mb-4">
                            <FormLabel>Username</FormLabel>
                            <FormControl>
                                <Input type="text" placeholder="rickAstley" v-model="username" />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    </FormField>
                    <FormField v-slot="{ componentField }" name="password">
                        <FormItem class="mb-4">
                            <FormLabel>Password</FormLabel>
                            <FormControl>
                                <Input type="password" placeholder="" v-model="password" />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    </FormField>
                </TabsContent>
                <TabsContent value="register">
                    <FormField v-slot="{ componentField }" name="username">
                        <FormItem class="mb-4">
                            <FormLabel>Username</FormLabel>
                            <FormControl>
                                <Input type="text" placeholder="jonhdoe" v-model="username" />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    </FormField>
                    <FormField v-slot="{ componentField }" name="email">
                        <FormItem class="mb-4">
                            <FormLabel>Email</FormLabel>
                            <FormControl>
                                <Input type="email" placeholder="example@example.com" v-model="email" />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    </FormField>
                    <FormField v-slot="{ componentField }" name="password">
                        <FormItem class="mb-4">
                            <FormLabel>Password</FormLabel>
                            <FormControl>
                                <Input type="password" placeholder="" v-model="password" />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    </FormField>
                </TabsContent>
                <div class="flex flex-row justify-center">
                    <Button class="mt-6 bg-green-600 text-white hover:bg-green-500 hover:text-black" @click="onSubmit">
                        Submit
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
import { useForm } from 'vee-validate'

import { useUserStore } from '../../store/user';
import { useBagStore } from '../../store/bag';

const userStore = useUserStore()
const bagStore = useBagStore()
const username = ref('')
const email = ref('')
const password = ref('')

const onSubmit = () => {
    console.debug('onSubmit: ', username.value, email.value, password.value)
}

const logout = () => {
    userStore.logout()
    bagStore.clearBag()
}

</script>