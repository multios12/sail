<script lang="ts">
  import List from "./routes/List.svelte";
</script>

<nav class="navbar is-dark" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="./">la-diary</a>
  </div>
</nav>
<main>
  <List />
</main>

<style>
  @import url("bulma/css/bulma.css");
</style>
