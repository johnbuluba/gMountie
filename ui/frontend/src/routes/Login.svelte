<script lang="ts">
    import {Login} from '../../wailsjs/go/main/App.js'
    import {push} from 'svelte-spa-router'


    let resultText: string = "Please enter your name below 👇"
    let host: string = "test"
    let port: number = 9449
    let name: string = "test"
    let password: string = "test"

    function login(): void {
        Login(host, name, password).then(result => push('/volumes'))
    }
</script>

<div  class="w-full max-w-sm">
    <form class="w-full max-w-sm" onsubmit={login}>
        <div class="md:flex md:items-center mb-6">
            <div class="md:w-1/3">
                <label class="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4"
                       for="inline-full-name">
                    Full Name
                </label>
            </div>
            <div class="md:w-2/3">
                <input class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
                       id="inline-full-name" type="text" value="Jane Doe">
            </div>
        </div>
        <div class="md:flex md:items-center mb-6">
            <div class="md:w-1/3">
                <label class="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4"
                       for="inline-password">
                    Password
                </label>
            </div>
            <div class="md:w-2/3">
                <input class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
                       id="inline-password" type="password"
                       placeholder="******************">
            </div>
        </div>

        <div class="md:flex md:items-center">
            <div class="md:w-1/3"></div>
            <div class="md:w-2/3">
                <button class="shadow bg-purple-500 hover:bg-purple-400 focus:shadow-outline focus:outline-none text-white font-bold py-2 px-4 rounded"
                        type="button" onclick="{login}">
                    Log in
                </button>
            </div>
        </div>
    </form>
</div>
