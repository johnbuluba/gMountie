<div class="py-8">
    <!-- Show the router -->
    <Router {routes} />

</div>

<script>
    // Import the router component
    // Normally, this would be: `import Router from 'svelte-spa-router'`
    import Router from 'svelte-spa-router'
    import { AppBar } from '@skeletonlabs/skeleton';
    // Import the list of routes
    import routes from './routes'

    import '@fortawesome/fontawesome-free/css/fontawesome.css';
    import '@fortawesome/fontawesome-free/css/brands.css';
    import '@fortawesome/fontawesome-free/css/solid.css';
</script>
