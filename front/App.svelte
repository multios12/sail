<script lang="ts">
  import "bulma/css/bulma.css";
  import "mu-ui-lib/mu-theme.css";
  import { MuHeader } from "mu-ui-lib";
  import Router, { location, push } from "svelte-spa-router";
  import Home from "./routes/Home/Home.svelte";
  import SalaryList from "./routes/SalaryList.svelte";
  import SalaryDetail from "./routes/SalaryDetail.svelte";
  import Cost from "./routes/Cost/Cost.svelte";
  let activeTab = "";

  const routes = {
    "/": Home,
    "/:year": Home,
    "/salary/:year": SalaryList,
    "/salary/:year/:month": SalaryDetail,
    "/cost/:year": Cost,
  };

  const getCurrentYear = () => new Date().getFullYear();

  $: tabs = [
    { id: "b", label: "balance", active: activeTab === "b" },
    { id: "s", label: "salary", active: activeTab === "s" },
    { id: "c", label: "cost", active: activeTab === "c" },
  ];

  $: {
    if ($location.indexOf("/s") >= 0) {
      activeTab = "s";
    } else if ($location.indexOf("/c") >= 0) {
      activeTab = "c";
    } else {
      activeTab = "b";
    }
  }

  const onTabSelect = (e: CustomEvent<{ id?: string }>) => {
    activeTab = e.detail.id ?? "b";
    if (e.detail.id === "s") {
      push(`/salary/${getCurrentYear()}`);
    } else if (e.detail.id === "c") {
      push(`/cost/${getCurrentYear()}`);
    } else {
      push("/");
    }
  };
</script>

<MuHeader
  title="sail"
  headerButton={{ href: "/", name: "トップに戻る", icon: "sail-home-icon" }}
  {tabs}
  showLogout={false}
  ontabselect={onTabSelect}
/>
<main>
  <div class="box">
    <Router {routes} />
  </div>
</main>

<style>
  :global(a:has(> .sail-home-icon)) {
    --button-background: transparent;
    --button-border-color: transparent;
    --button-box-shadow: none;
    --button-hover-border-color: transparent;
    --button-hover-box-shadow: none;
    --button-hover-transform: none;
    --button-padding: 0;
    --button-transition: color 0.15s ease;
    color: var(--header-sub-title-color);
    font-size: 0;
    gap: 0;
    justify-content: center;
    margin-inline-end: 0.375rem;
    min-block-size: 1.75rem;
    min-inline-size: 1.75rem;
    vertical-align: middle;
  }

  :global(a:has(> .sail-home-icon):hover),
  :global(a:has(> .sail-home-icon):focus-visible) {
    color: var(--color-primary-contract);
  }

  :global(a:has(> .sail-home-icon):focus-visible) {
    outline: 2px solid currentColor;
    outline-offset: 2px;
  }

  :global(.sail-home-icon)::before {
    content: "home" / "";
    font: normal 1.125rem / 1 "Material Icons";
  }
</style>
