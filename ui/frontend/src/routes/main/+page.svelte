<script lang="ts">
    import VolumesTable from "./table-volumes.svelte";
    import logo from "$lib/assets/images/logo-full-cropped.png";
    import {Avatar} from '@skeletonlabs/skeleton';
    import { goto } from '$app/navigation'
    import {
        GetVolumes
    } from "bindings/gmountie/pkg/ui/controller/volumecontrollerimpl";
    import {
        Logout
    } from "bindings/gmountie/pkg/ui/controller/logincontrollerimpl";


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
            goto('/login')
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
