<script lang="ts">
  import type { BalanceYear } from "../../models/balanceModels";
  import YearCard from "./YearCard.svelte";

  export let params: { year: string | undefined } = { year: undefined };
  $: {
    const url = `./api/balance/${params.year || new Date().getFullYear()}`;
    fetch(url)
      .then((r) => r.json())
      .then((r) => (model = r));
  }

  let model: BalanceYear = {
    Year: "",
    Balances: [],
    EnableYears: [],
  };
</script>

<div>
  {#each model.EnableYears as year}
    <YearCard {year} {model} />
  {/each}
</div>
