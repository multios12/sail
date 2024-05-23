<script lang="ts">
  import "bulma/css/bulma.css";
  import Router, { location, link } from "svelte-spa-router";
  import Home from "./routes/Home/Home.svelte";
  import SalaryYear from "./routes/BalanceSalaryYear.svelte";
  import SalaryMonth from "./routes/BalanceSalaryMonth.svelte";
  import CostYear from "./routes/Cost/Cost.svelte";
  import BalanceUploadButton from "./components/UploadButton.svelte";
  let page = "";

  const routes = {
    "/": Home,
    "/:year": Home,
    "/salary/:year": SalaryYear,
    "/salary/:year/:month": SalaryMonth,
    "/cost/:year": CostYear,
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
</script>

<nav class="navbar is-dark" aria-label="main navigation">
  <div class="navbar-brand">
    <div class="navbar-item is-unselectable has-text-weight-bold">sail</div>
  </div>
  <div class="navbar-item is-tab" class:is-active={page == "b"}>
    <a href="/" use:link> balance </a>
  </div>
  <div class="navbar-item is-tab" class:is-active={page == "s"}>
    <a href="/salary/{new Date().getFullYear()}" use:link>
      <i class="material-icons">attach_money</i>salary
    </a>
  </div>
  <div class="navbar-item is-tab" class:is-active={page == "c"}>
    <a href="/cost/{new Date().getFullYear()}" use:link>
      <i class="material-icons">payment</i>cost
    </a>
  </div>
  <div class="navbar-item is-tab">
    <BalanceUploadButton />
  </div>
</nav>
<main>
  <div class="box">
    <Router {routes} />
  </div>
</main>

<style>
</style>
