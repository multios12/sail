<script lang="ts">
  import type { BalanceItem } from "../models";
  import { onMount } from "svelte";
  export let Value: BalanceItem;

  export let editMonth: string;
  let Water = Value.CostWater;
  let Electric = Value.CostElectric;
  let Gas = Value.CostGas;
  let Mobile = Value.CostMobile;
  let Line = Value.CostLine;
  let Tax = Value.CostTax;

  onMount(async () => {
    if (editMonth === "" || editMonth === Value.Month) {
      document
        .querySelector(`#edit${Value.Month}`)
        ?.classList.remove("is-hidden");
    } else {
      document.querySelector(`#edit${Value.Month}`)?.classList.add("is-hidden");
    }
  });

  const saveClick = () => {
    const url = `./api/${Value.Month.toString().substring(
      0,
      4
    )}/${Value.Month.toString().substring(4)}`;
    const d = {
      Month: Value.Month,
      CostWater: Water,
      CostElectric: Electric,
      CostGas: Gas,
      CostMobile: Mobile,
      CostLine: Line,
      CostTax: Tax,
    };
    fetch(url, { method: "post", body: JSON.stringify(d) });
    editMonth = "";
  };
</script>

{#if editMonth === Value.Month}
  <tr>
    <td>{`${Value.Month.substring(0, 4)}年${Value.Month.substring(4)}月`}</td>
    <td class="MuNumber px-0 has-text-right"
      >{Water + Electric + Gas + Mobile + Line + Tax}</td
    >
    <td class="MuNumber p-0 has-text-right"
      ><input
        type="number"
        class="input px-0 has-text-right"
        bind:value={Water}
      /></td
    >
    <td class="MuNumber p-0 has-text-right"
      ><input
        type="number"
        class="input px-0 has-text-right"
        bind:value={Electric}
      /></td
    >
    <td class="MuNumber p-0 has-text-right"
      ><input
        type="number"
        class="input px-0 has-text-right"
        bind:value={Gas}
      /></td
    >
    <td class="MuNumber p-0 has-text-right"
      ><input
        type="number"
        class="input px-0 has-text-right"
        bind:value={Mobile}
      /></td
    >
    <td class="MuNumber p-0 has-text-right"
      ><input
        type="number"
        class="input px-0 has-text-right"
        bind:value={Line}
      /></td
    >
    <td class="MuNumber p-0 has-text-right"
      ><input
        type="number"
        class="input px-0 has-text-right"
        bind:value={Tax}
      /></td
    >
    <td
      ><button
        class="button is-primary is-small material-icons"
        on:click={saveClick}>save</button
      ></td
    >
  </tr>
{:else}
  <tr>
    <td>{`${Value.Month.substring(0, 4)}年${Value.Month.substring(4)}月`}</td>
    <td class="MuNumber px-0 has-text-right"
      >{Water + Electric + Gas + Mobile + Line + Tax}</td
    >
    <td class="MuNumber px-0 has-text-right pr-4">{Water}</td>
    <td class="MuNumber px-0 has-text-right pr-4">{Electric}</td>
    <td class="MuNumber px-0 has-text-right pr-4">{Gas}</td>
    <td class="MuNumber px-0 has-text-right pr-4">{Mobile}</td>
    <td class="MuNumber px-0 has-text-right pr-4">{Line}</td>
    <td class="MuNumber px-0 has-text-right pr-4">{Tax}</td>
    <td
      ><button
        id={"edit" + Value.Month}
        class="button is-info is-small is-inverted material-icons"
        on:click={() => {
          editMonth = Value.Month;
        }}>edit</button
      ></td
    >
  </tr>
{/if}

<style>
  input.input {
    width: 70px;
  }

  td.MuNumber {
    width: 90px;
  }
</style>
