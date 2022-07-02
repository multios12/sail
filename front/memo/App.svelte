<script lang="ts">
  import "bulma/css/bulma.css";
  import Router from "svelte-spa-router";
  import List from "./routes/List.svelte";
  import Edit from "./routes/Edit.svelte";

  const routes = {
    "/": List,
    "/:id": Edit,
  };
</script>

<Router {routes} />
