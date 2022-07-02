<script lang="ts">
  import { onMount } from "svelte";
  import type { BalanceItem } from "src/models";

  export let editMonth: String;
  export let model: BalanceItem;

  let saving = model.Saving;
  let memo = model.Memo;

  onMount(async () => {
    if (editMonth === "" || editMonth === model.Month) {
      document
        .querySelector(`#edit${model.Month}`)
        ?.classList.remove("is-hidden");
    } else {
      document.querySelector(`#edit${model.Month}`)?.classList.add("is-hidden");
    }
  });

  const saveClick = (e: MouseEvent) => {
    model.Saving = saving;
    model.Memo = memo;
    const url = `./api/balance/${model.Month.toString().substring(
      0,
      4
    )}/${model.Month.toString().substring(4)}`;
    fetch(url, { method: "post", body: JSON.stringify(model) });
    editMonth = "";
  };
</script>

{#if editMonth === model.Month}
  <tr>
    <td class="MuNumber p-0 has-text-right"
      >{model.Month.substring(0, 4)}年{model.Month.substring(4)}月</td
    >
    <td class="MuNumber p-0 has-text-right">{model.Salary.toLocaleString()}</td>
    <td class="MuNumber p-0 has-text-right">{model.Paid.toLocaleString()}</td>
    <td class={model.IsNotCost ? "has-background-danger-light" : ""}
      >{model.Cost.toLocaleString()}</td
    >
    <td class="MuNumber p-0 has-text-right">
      <input
        type="number"
        class="input px-0 has-text-right"
        bind:value={saving}
      />
    </td>
    <td class="MuNumber p-0 has-text-right">
      <input type="text" class="input px-0 has-text-right" bind:value={memo} />
    </td>
    <td
      ><button
        class="button is-primary is-small material-icons"
        on:click={saveClick}>save</button
      ></td
    >
  </tr>
{:else}
  <tr>
    <td>{model.Month.substring(0, 4)}年{model.Month.substring(4)}月</td>
    <td class="MuNumber has-text-right">{model.Salary.toLocaleString()}</td>
    <td class="MuNumber has-text-right">{model.Paid.toLocaleString()}</td>
    <td
      class={(model.IsNotCost ? "has-background-danger-light" : "") +
        " has-text-right"}>{model.Cost.toLocaleString()}</td
    >
    <td class="MuNumber px-0 has-text-right pr-4">{saving.toLocaleString()}</td>
    <td>{memo}</td>
    <td
      ><button
        id={`edit${model.Month}`}
        class="button is-info is-small is-inverted material-icons"
        on:click={() => {
          editMonth = model.Month;
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
