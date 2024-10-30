<script lang="ts">
    import {push} from 'svelte-spa-router'
    import {GetVolumes, Logout} from '../../wailsjs/go/main/App'
    import VolumesTable from "../components/volumes/VolumesTable.svelte";
    import logo from "../assets/images/logo-full-cropped.png";
    import {Avatar} from '@skeletonlabs/skeleton';

    let results = []


    function getVolumes(): void {
        GetVolumes().then((result) => {
            console.log(result)
            results = result
        })
    }

    function logout(): void {
        Logout().then(() => {
            console.log('Logged out')
            push('/')
        })
    }
    getVolumes()
</script>


<div class="grid grid-cols-6 mx-4">
    <div class="col-start-1 col-end-6">
        <Avatar src={logo}></Avatar>
    </div>
    <div class="col-span-1">
        <div class="flex justify-end">
            <button class="btn variant-filled-primary" on:click={logout}>Logout</button>
        </div>
    </div>
</div>

<div class="p-4">
    <VolumesTable bind:volumes={results}></VolumesTable>
</div>
