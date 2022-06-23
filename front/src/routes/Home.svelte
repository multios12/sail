<script lang="ts">
  import type { BalanceYear } from "../models";
  import YearCard from "../lib/HomeYearCard.svelte";

  export let params: { year: string } = { year: undefined };
  $: {
    const url = `./api/${params.year || new Date().getFullYear()}`;
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

<div>
  {#each model.EnableYears as year}
    <YearCard {year} {model} />
  {/each}
</div>
