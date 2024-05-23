<script lang="ts">
  import type { BalanceItem } from "../../models/balanceModels";
  import { onMount } from "svelte";
  import EditTd from "../../components/EditTd.svelte";
  export let Value: BalanceItem;

  export let editMonth: string;
  let Housing = Value.CostHousing;
  let Water = Value.CostWater;
  let Electric = Value.CostElectric;
  let Gas = Value.CostGas;
  let Mobile = Value.CostMobile;
  let Line = Value.CostLine;
  let Tax = Value.CostTax;
  $: Total = Housing + Water + Electric + Gas + Mobile + Line + Tax;
  $: isEdit = editMonth === Value.Month;
  onMount(async () => {
    if (editMonth === "" || isEdit) {
      document
        .querySelector(`#edit${Value.Month}`)
        ?.classList.remove("is-hidden");
    } else {
      document.querySelector(`#edit${Value.Month}`)?.classList.add("is-hidden");
    }
  });

  const onClick = () => {
    if (!isEdit) {
      editMonth = Value.Month;
      return;
    }
    let url = `./api/balance/${Value.Month.toString().substring(0, 4)}`;
    url = `${url}/${Value.Month.toString().substring(4)}`;
    const d = {
      Month: Value.Month,
      CostHousing: Housing,
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

<tr>
  <td>{`${Value.Month.substring(0, 4)}年${Value.Month.substring(4)}月`}</td>
  <EditTd value={Total} />
  <EditTd bind:value={Housing} {isEdit} />
  <EditTd bind:value={Water} {isEdit} />
  <EditTd bind:value={Electric} {isEdit} />
  <EditTd bind:value={Gas} {isEdit} />
  <EditTd bind:value={Mobile} {isEdit} />
  <EditTd bind:value={Line} {isEdit} />
  <EditTd bind:value={Tax} {isEdit} />
  <td>
    <button
      class="button is-small material-icons"
      class:is-primary={isEdit}
      class:is-info={!isEdit}
      class:is-inverted={!isEdit}
      on:click={onClick}
    >
      {isEdit ? "save" : "edit"}
    </button>
  </td>
</tr>
