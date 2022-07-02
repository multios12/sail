<script lang="ts">
  import "bulma/css/bulma.css";
  import UploadButton from "./UploadButton.svelte";
  import { link } from "svelte-spa-router";

  const burgerClick = () => {
    document.querySelector("#menu")?.classList.toggle("is-active");
    document.querySelector(".navbar-burger")?.classList.toggle("is-active");
  };
</script>

<nav class="navbar is-dark" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/" use:link>sail</a>
    <a class="navbar-item" href="/salary/{new Date().getFullYear()}" use:link
      ><i class="material-icons">attach_money</i></a
    >
    <a class="navbar-item" href="/cost/{new Date().getFullYear()}" use:link
      ><i class="material-icons">payment</i></a
    >

    <button
      class="navbar-burger"
      on:click={burgerClick}
      aria-label="menu"
      aria-expanded="false"
    >
      <span aria-hidden="true" />
      <span aria-hidden="true" />
      <span aria-hidden="true" />
    </button>
  </div>

  <div id="menu" class="navbar-menu">
    <div class="navbar-start" />

    <div class="navbar-end">
      <UploadButton />
    </div>
  </div>
</nav>
