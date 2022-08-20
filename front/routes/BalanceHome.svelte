<script lang="ts">
  import type { BalanceYear } from "../../diary/models/balanceModels";
  import YearCard from "../components/BalanceHomeYearCard.svelte";
  import BalanceTabs from "../components/BalanceTabs.svelte";

  export let params: { year: string } = { year: undefined };
  $: {
    const url = `./api/balance/${params.year || new Date().getFullYear()}`;
    fetch(url)
      .then((r) => r.json())
      .then((r) => (model = r));
  }

  let model: BalanceYear = {
    Year: undefined,
    Balances: [],
    EnableYears: [],
  };
</script>

<BalanceTabs />
<div>
  {#each model.EnableYears as year}
    <YearCard {year} {model} />
  {/each}
</div>
