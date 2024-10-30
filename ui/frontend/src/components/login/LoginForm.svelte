<script lang="ts">
    import {push} from 'svelte-spa-router'
    import {Login} from '../../../wailsjs/go/main/App'


    let endpoint: string = "127.0.0.1:9449"
    let username: string = "john"
    let password: string = "123456"

    function login(): void {
        Login(endpoint, username, password).then((result) => {
            console.log("Logged in")
            push('/volumes')
        })
    }
</script>


<div class="w-full max-w-sm">
    <div class="card p-4 space-y-2">
        <label class="label">
            <span>Endpoint</span>
            <input class="input" type="text" placeholder="Endpoint" bind:value={endpoint} />
        </label>
        <label class="label">
            <span>Username</span>
            <input class="input" type="text" placeholder="Username" bind:value={username}/>
        </label>
        <label class="label">
            <span>Password</span>
            <input class="input" type="password" placeholder="Password" bind:value={password} />
        </label>
        <div class="flex justify-center">
            <button class="btn variant-filled" on:click={login}>Log in</button>
        </div>
    </div>
</div>

