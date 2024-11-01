<script lang="ts">
    import MountButton from "./button-mount.svelte";
    import {common} from "$lib/wailsjs/go/models";

    export let volumes: common.Volume[];

</script>


<section class="w-full text-token p-4 space-y-4">
    <div class="table-container">
        <table class="table table-compact">
            <thead>
            <tr>
                <th>Name</th>
                <th class="table-cell-fit">Mount</th>
            </tr>
            </thead>
            <tbody>
            {#each volumes as row, i}
                <tr>
                    <td><h2>{row.Name}</h2></td>
                    <td>
                        <div class="flex justify-center">
                            <MountButton bind:volume={row}/>
                        </div>
                    </td>
                </tr>
            {/each}
            </tbody>
        </table>
    </div>
</section>
