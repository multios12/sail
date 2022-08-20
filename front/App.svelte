<script lang="ts">
  import Router, { location, link } from "svelte-spa-router";
  import Home from "./routes/BalanceHome.svelte";
  import SalaryYear from "./routes/BalanceSalaryYear.svelte";
  import SalaryMonth from "./routes/BalanceSalaryMonth.svelte";
  import CostYear from "./routes/BalanceCost.svelte";
  import DiaryList from "./routes/DiaryList.svelte";
  import DiaryEdit from "./routes/DiaryDetail.svelte";
  import HMemoList from "./routes/HMemoList.svelte";
  import HMemoEdit from "./routes/HMemoDetail.svelte";
  let page = "";

  const routes = {
    "/": Home,
    "/balance/:year": Home,
    "/balance/salary/:year": SalaryYear,
    "/balance/salary/:year/:month": SalaryMonth,
    "/balance/cost/:year": CostYear,
    "/d/": DiaryList,
    "/d/:id": DiaryEdit,
    "/d/add": DiaryEdit,
    "/h/": HMemoList,
    "/h/:id": HMemoEdit,
    "/h/add": HMemoEdit,
  };

  $: {
    if ($location.indexOf("/d") >= 0) {
      page = "d";
    } else if ($location.indexOf("/h") >= 0) {
      page = "h";
    } else {
      page = "b";
    }
  }
</script>

<nav class="navbar is-dark" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/" use:link>sail</a>
  </div>
  <div class="tabs is-boxed">
    <ul>
      <li class:is-active={page == "b"}>
        <a class="navbar-item" href="/" use:link>balance</a>
      </li>
      <li class:is-active={page == "d"}>
        <a class="navbar-item" href="/d/" use:link>diary</a>
      </li>
      <li class:is-active={page == "h"}>
        <a class="navbar-item" href="/h/" use:link>memo</a>
      </li>
    </ul>
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
