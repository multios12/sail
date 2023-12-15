<script lang="ts">
  import { pop, push } from "svelte-spa-router";
  import type { memoType } from "../models/memoModels.js";
  import { onMount } from "svelte";

  const m: memoType = {
    Id: undefined,
    Name: "",
    Shop: "",
    Page: "",
    Date: new Date().toISOString().substring(0, 10),
    Play: "",
    Talk: "",
  };

  export let params: { id: string | undefined } = { id: undefined };
  let memo = m;
  let isErr = false;
  let errMessage = "";
  let isLoading = false;

  const sendClick = () => {
    isLoading = true;
    let url = "./api/memos";
    url += params.id === "add" ? "" : `/${params.id}`;
    let o = {};
    if (params.id === "add") {
      console.log(memo);
      o = { method: "put", body: JSON.stringify(memo) };
    } else {
      o = { method: "post", body: JSON.stringify(memo) };
    }
    fetch(url, o)
      .then(() => push("/h/"))
      .catch((res) => {
        errMessage = res.response.data.error;
        isErr = true;
      })
      .finally(() => {
        isLoading = false;
      });
  };

  const deleteClick = async () => {
    let url = `./api/memos/${params.id}`;
    await fetch(url, { method: "delete" });
    pop();
  };

  onMount(async () => {
    if (params.id === "add") {
      return;
    }
    isLoading = true;
    const r = await fetch(`./api/memos/${params.id}`);
    const v = await r.json();
    memo = v[0];
    isLoading = false;
  });
</script>

<div class="card px-10">
  <header class="card-header">
    <p class="card-header-title" />
    <button
      class="button is-inverted is-small has-text-danger sp-right"
      on:click={deleteClick}
    >
      <i class="material-icons">delete</i>
    </button>
  </header>
  <div class="card-content">
    {#if errMessage != ""}
      <div class="notification is-danger">{errMessage}</div>
    {/if}
    <div class="columns m-0">
      <div class="field column m-0">
        <label class="label" for="name">name</label>
        <div class="control">
          <input class="input" type="text" name="name" bind:value={memo.Name} />
        </div>
      </div>
      <div class="field column">
        <label class="label" for="date">date</label>
        <div class="control">
          <input class="input" type="text" name="date" bind:value={memo.Date} />
        </div>
      </div>
    </div>
    <div class="columns m-0">
      <div class="field column">
        <label class="label" for="shop">shop</label>
        <div class="control">
          <input class="input" type="text" name="shop" bind:value={memo.Shop} />
        </div>
      </div>
      <div class="field column">
        <label class="label" for="page">home page</label>
        <div class="control">
          <input class="input" type="text" name="page" bind:value={memo.Page} />
        </div>
      </div>
    </div>
    <div class="field">
      <label class="label" for="play">play</label>
      <div class="control">
        <textarea class="textarea" name="play" bind:value={memo.Play} />
      </div>
    </div>
    <div class="field">
      <label class="label" for="talk">talk</label>
      <div class="control">
        <textarea class="textarea" name="talk" bind:value={memo.Talk} />
      </div>
    </div>
    <footer class="card-footer">
      <div class="card-footer-item">
        <button
          class="button is-primary mx-5"
          class:is-loading={isLoading}
          on:click={() => sendClick()}>ok</button
        >
        <button class="button is-light mx-5" on:click={() => push("/h/")}
          >cancel</button
        >
      </div>
    </footer>
  </div>
</div>
