import Login from "./routes/Login.svelte";
import Volumes from "./routes/Volumes.svelte";

// Export the route definition object
export default {
    // Exact path
    '/': Login,
    '/volumes': Volumes,
}
