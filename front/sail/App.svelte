<script lang="ts">
  import "bulma/css/bulma.css";
  import AppBar from "./lib/AppBar.svelte";
  import Router from "svelte-spa-router";
  import Home from "./routes/Home.svelte";
  import SalaryYear from "./routes/SalaryYear.svelte";
  import SalaryMonth from "./routes/SalaryMonth.svelte";
  import CostYear from "./routes/Cost.svelte";

  const routes = {
    "/": Home,
    "/:year": Home,
    "/salary/:year": SalaryYear,
    "/salary/:year/:month": SalaryMonth,
    "/cost/:year": CostYear,
  };
</script>

<AppBar />
<main>
  <Router {routes} />
</main>
