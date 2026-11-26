<script lang="ts">
  import { fetchJson } from "../../lib/api";
  import YearTabs from "../../components/YearTabs.svelte";
  import type { BalanceYear } from "../../models/balanceModels";
  import YearCard from "./components/YearCard.svelte";

  export let params: { year: string | undefined } = { year: undefined };
  $: {
    const url = `/api/balance/${params.year || new Date().getFullYear()}`;
    fetchJson<BalanceYear>(url).then((r) => {
      if (r) {
        model = r;
      }
    });
  }

  let model: BalanceYear = {
    Year: "",
    Balances: [],
    EnableYears: [],
  };
</script>

<div class="home-page">
  <YearTabs
    years={model.EnableYears}
    currentYear={model.Year}
    ariaLabel="収支年選択"
  />

  {#if model.Year}
    <YearCard year={model.Year} {model} />
  {/if}
</div>

<style>
  .home-page {
    display: grid;
    gap: 1rem;
  }

</style>
