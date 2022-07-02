<script lang="ts">
  import "bulma/css/bulma.css";
  import { push, link } from "svelte-spa-router";
  import { onMount } from "svelte";
  let memos = [];
  const showEdit = (id: string) => {
    push(`/${id}`);
  };
  onMount(async () => {
    const r = await fetch("./api/memos");
    memos = await r.json();
  });
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
  <div class="card px-10">
    <div class="card-content">
      <table class="table is-striped is-hoverable is-fullwidth">
        <thead>
          <tr>
            <th>date</th>
            <th class="is-half">name</th>
            <th>shop</th>
            <th />
          </tr>
        </thead>
        <tbody>
          {#each memos as m}
            <tr>
              <td on:click={() => showEdit(m.Id)}>{m.Date}</td>
              <td on:click={() => showEdit(m.Id)}>{m.Name}</td>
              <td on:click={() => showEdit(m.Id)}>{m.Shop}</td>
              <td><button /></td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
</main>
