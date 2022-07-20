<script lang="ts">
  import "bulma/css/bulma.css";
  import { push } from "svelte-spa-router";
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
