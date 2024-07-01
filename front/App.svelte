<script lang="ts">
  import "bulma/css/bulma.css";
  import Router, { location, link } from "svelte-spa-router";
  import Home from "./routes/Home/Home.svelte";
  import SalaryList from "./routes/SalaryList.svelte";
  import SalaryDetail from "./routes/SalaryDetail.svelte";
  import Cost from "./routes/Cost/Cost.svelte";
  import UploadButton from "./components/UploadButton.svelte";
  let page = "";

  const routes = {
    "/": Home,
    "/:year": Home,
    "/salary/:year": SalaryList,
    "/salary/:year/:month": SalaryDetail,
    "/cost/:year": Cost,
  };

  $: {
    if ($location.indexOf("/s") >= 0) {
      page = "s";
    } else if ($location.indexOf("/c") >= 0) {
      page = "c";
    } else {
      page = "b";
    }
  }

  document.addEventListener("DOMContentLoaded", () => {
    // Get all "navbar-burger" elements
    const $navbarBurgers = Array.prototype.slice.call(
      document.querySelectorAll(".navbar-burger"),
      0,
    );

    // Add a click event on each of them
    $navbarBurgers.forEach((el) => {
      el.addEventListener("click", () => {
        // Get the target from the "data-target" attribute
        const target = el.dataset.target;
        const _target = <HTMLElement>document.getElementById(target);

        // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
        el.classList.toggle("is-active");
        _target.classList.toggle("is-active");
      });
    });
  });
</script>

<nav class="navbar is-dark" aria-label="main navigation">
  <div class="navbar-brand">
    <div class="navbar-item is-unselectable has-text-weight-bold">sail</div>
    <div class="navbar-burger js-burger" data-target="navbarMMemo">
      <span></span>
      <span></span>
      <span></span>
      <span></span>
    </div>
  </div>

  <div id="navbarMMemo" class="navbar-menu">
    <div class="navbar-start">
      <a
        class="navbar-item is-tab"
        class:is-active={page == "b"}
        href="/"
        use:link
      >
        <i class="material-icons">balance</i>balance
      </a>
      <a
        href="/salary/{new Date().getFullYear()}"
        class="navbar-item is-tab"
        class:is-active={page == "s"}
        use:link
      >
        <i class="material-icons">attach_money</i>salary
      </a>
      <a
        href="/cost/{new Date().getFullYear()}"
        class="navbar-item is-tab"
        class:is-active={page == "c"}
        use:link
      >
        <i class="material-icons"> payment </i> cost
      </a>
      <div class="navbar-item is-tab">
        <UploadButton />
      </div>
    </div>
  </div>
</nav>
<main>
  <div class="box">
    <Router {routes} />
  </div>
</main>

<style>
</style>
