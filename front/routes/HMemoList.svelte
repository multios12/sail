<script lang="ts">
  import { link, push } from "svelte-spa-router";
  import { onMount } from "svelte";
  import type { memoType } from "../models/memoModels.js";
  let memos: memoType[] = [];
  const showEdit = (id: string | undefined) => {
    push(`/h/${id}`);
  };
  onMount(async () => {
    const r = await fetch("./api/memos");
    memos = await r.json();
  });
</script>

<div class="card px-10">
  <div class="card-content">
    <div class="columns">
      <div class="column">
        <a class="button is-primary" href="/h/add" use:link
          ><i class="material-icons">add</i>add</a
        >
      </div>
    </div>

    <table class="table is-striped is-hoverable is-fullwidth">
      <tbody>
        {#each memos as m}
          <tr>
            <td
              on:click={() => showEdit(m.Id)}
              on:keypress={() => showEdit(m.Id)}>{m.Date}&nbsp;{m.Name}</td
            >
            <td
              on:click={() => showEdit(m.Id)}
              on:keypress={() => showEdit(m.Id)}>{m.Shop}</td
            >
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>
