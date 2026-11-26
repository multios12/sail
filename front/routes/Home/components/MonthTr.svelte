<script lang="ts">
  import type { BalanceItem } from "../../../models/balanceModels";
  import EditTd from "../../../components/EditTd.svelte";

  export let editMonth: String;
  export let model: BalanceItem;

  let saving = model.Saving;
  let memo = model.Memo;

  $: month =
    model.Month.substring(0, 4) + "年" + model.Month.substring(4) + "月";
  $: isEdit = editMonth === model.Month;

  const onClick = (e: MouseEvent) => {
    model.Saving = saving;
    model.Memo = memo;
    const url = `./api/balance/${model.Month.toString().substring(
      0,
      4,
    )}/${model.Month.toString().substring(4)}`;
    fetch(url, { method: "post", body: JSON.stringify(model) });
    editMonth = "";
  };
</script>

<tr>
  <EditTd value={month} />
  <EditTd value={model.Salary.toLocaleString()} />
  <EditTd value={model.Paid.toLocaleString()} />
  <td
    class:cost-warning={model.IsNotCost}
    class="amount-cell"
  >
    {model.Cost.toLocaleString()}
  </td>
  <EditTd bind:value={saving} {isEdit} />
  <EditTd bind:value={memo} {isEdit} isText />
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
  .amount-cell {
    text-align: right;
  }

  .cost-warning {
    background: #fee2e2;
    color: #991b1b;
  }

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
