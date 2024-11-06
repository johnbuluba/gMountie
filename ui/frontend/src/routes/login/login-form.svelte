<script lang="ts">
    import { goto } from '$app/navigation'
    import {
        Login,
        IsLoggedIn
    } from "bindings/gmountie/pkg/ui/controller/logincontrollerimpl";
    import {LogInInfo} from "bindings/gmountie/pkg/ui/controller";

    let address: string = "127.0.0.1"
    let port: number = 9449
    let username: string = "admin"
    let password: string = "admin"

    function login(): void {
        const loginInfo = {
            Address: address,
            Port: port,
            Username: username,
            Password: password
        } as LogInInfo;
        console.log(loginInfo)
        Login(loginInfo).then((result) => {
            if (result) {
                console.log("Logged in")
                goto('/main')
            }
            console.log("Failed to log in")
        })
    }

    IsLoggedIn().then((result) => {
        if (result) {
            goto('/main')
        }
    })

</script>


<div class="w-full max-w-sm">
    <div class="card p-4 space-y-2">
        <label class="label">
            <span>Endpoint</span>
            <input class="input" type="text" placeholder="Address" bind:value={address} />
        </label>
        <label class="label">
            <span>Endpoint</span>
            <input class="input" type="number" placeholder="Port" bind:value={port} />
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

