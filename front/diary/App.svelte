<script lang="ts">
  import Router from "svelte-spa-router";
  import List from "./routes/List.svelte";
  const routes = {
    "/": List,
  };
</script>

<nav class="navbar is-dark" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="./">la-diary</a>
  </div>
</nav>
<main>
  <Router {routes} />
</main>

<style>
  @import url("bulma/css/bulma.css");
</style>
