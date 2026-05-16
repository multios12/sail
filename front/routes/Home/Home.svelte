<script lang="ts">
  import { fetchJson } from "../../lib/api";
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

<div>
  {#each model.EnableYears as year}
    <YearCard {year} {model} />
  {/each}
</div>
