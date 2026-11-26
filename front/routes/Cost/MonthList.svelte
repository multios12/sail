<script lang="ts">
  import type { BalanceItem } from "../../models/balanceModels";
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

  const onClick = () => {
    if (!isEdit) {
      editMonth = Value.Month;
      return;
    }
    let url = `/api/balance/${Value.Month.toString().substring(0, 4)}`;
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
      class:edit-action-active={isEdit}
      class="edit-action material-icons"
      onclick={onClick}
    >
      {isEdit ? "save" : "edit"}
    </button>
  </td>
</tr>

<style>
  .edit-action {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 2rem;
    height: 2rem;
    border: 1px solid #0ea5e9;
    border-radius: 4px;
    background: transparent;
    color: #7dd3fc;
    cursor: pointer;
    font-size: 1.15rem;
    line-height: 1;
  }

  .edit-action:hover {
    background: #0c4a6e;
    color: #e0f2fe;
  }

  .edit-action-active {
    border-color: #22c55e;
    background: #16a34a;
    color: #fff;
  }

  .edit-action-active:hover {
    background: #15803d;
  }
</style>
