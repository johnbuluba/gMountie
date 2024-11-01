<script lang="ts">
    import logo from "$lib/assets/images/logo-full-cropped.png";
    import LoginForm from "./login-form.svelte";
    import { Avatar } from '@skeletonlabs/skeleton';

</script>

<div class="flex justify-center">
    <Avatar src={logo} width="w-32" rounded="rounded-3xl" />
</div>
<div class="flex justify-center py-12">
    <LoginForm />
</div>
