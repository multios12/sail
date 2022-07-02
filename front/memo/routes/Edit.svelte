<script lang="ts">
  import "bulma/css/bulma.css";
  import { push, link } from "svelte-spa-router";
  import { onMount } from "svelte";

  type memoType = {
    Id: string | undefined;
    Name: string;
    Shop: string;
    Page: string;
    Date: string;
    Play: string;
    Talk: string;
  };

  const m: memoType = {
    Id: undefined,
    Name: "",
    Shop: "",
    Page: "",
    Date: new Date().toISOString().substring(0, 10),
    Play: "",
    Talk: "",
  };

  export let params: { id: string } = { id: undefined };
  let memo = m;
  let isErr = false;
  let errMessage = "";
  let isLoading = false;

  const regist = () => {
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
      .then(() => push("/"))
      .catch((res) => {
        errMessage = res.response.data.error;
        isErr = true;
      })
      .finally(() => {
        isLoading = false;
      });
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
      {#if errMessage != ""}
        <div class="notification is-danger">{errMessage}</div>
      {/if}
      <div class="field">
        <label class="label" for="name">name</label>
        <div class="control">
          <input class="input" type="text" name="name" bind:value={memo.Name} />
        </div>
      </div>
      <div class="field">
        <label class="label" for="shop">shop</label>
        <div class="control">
          <input class="input" type="text" name="shop" bind:value={memo.Shop} />
        </div>
        <label class="label" for="page">home page</label>
        <div class="control">
          <input class="input" type="text" name="page" bind:value={memo.Page} />
        </div>
      </div>
      <div class="field">
        <label class="label" for="date">date</label>
        <div class="control">
          <input class="input" type="text" name="date" bind:value={memo.Date} />
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
        <button
          class="card-footer-item"
          class:is-loading={isLoading}
          on:click={() => regist()}>ok</button
        >
        <button class="card-footer-item" on:click={() => push("/")}
          >cancel</button
        >
      </footer>
    </div>
  </div>
</main>
