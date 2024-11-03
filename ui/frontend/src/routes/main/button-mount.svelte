<script lang="ts">
    import {Volume} from "bindings/gmountie/pkg/common";
    import {IsMounted, Mount, Unmount} from 'bindings/gmountie/ui3/app'

    export let volume: Volume;

    let icon = "fa-play";
    let disabled = true;
    let mounted;

    IsMounted(volume).then((result) => {
        mounted = result;
        if (mounted) {
            icon = "fa-stop";
            disabled = false;
        } else {
            icon = "fa-play";
            disabled = false;
        }
    });

    function mountUnmount() {
        disabled = true;
        if (mounted) {
            Unmount(volume).then(() => {
                icon = "fa-play";
                mounted = false;
            }).finally(() => {
                disabled = false;
            });
        } else {
            Mount(volume).then(() => {
                icon = "fa-stop";
                mounted = true;
            }).finally(() => {
                disabled = false;
            });
        }
    }
</script>

<button type="button" class="btn-icon variant-filled" aria-label="status" disabled={disabled} on:click={mountUnmount}>
    <i class="fa-solid {icon}"></i>
</button>

