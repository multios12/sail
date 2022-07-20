<script lang="ts">
  import Router from "svelte-spa-router";
  import List from "./routes/List.svelte";
  import Edit from "./routes/Edit.svelte";
  import { link } from "svelte-spa-router";

  const routes = {
    "/": List,
    "/:id": Edit,
  };
</script>

<nav class="navbar is-dark" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/" use:link>memo</a>
  </div>

  <div id="menu" class="navbar-menu">
    <div class="navbar-start" />
    <div class="navbar-end">
      <a class="navbar-item" href="/add" use:link
        ><i class="material-icons">add</i></a
      >
    </div>
  </div>
</nav>
<main>
  <div class="box">
    <Router {routes} />
  </div>
</main>

<style>
  @import url("bulma/css/bulma.css");
</style>
