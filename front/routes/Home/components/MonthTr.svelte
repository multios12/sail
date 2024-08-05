<script lang="ts">
  import { onMount } from "svelte";
  import type { BalanceItem } from "../../../models/balanceModels";
  import EditTd from "../../../components/EditTd.svelte";

  export let editMonth: String;
  export let model: BalanceItem;

  let saving = model.Saving;
  let memo = model.Memo;

  $: month =
    model.Month.substring(0, 4) + "年" + model.Month.substring(4) + "月";
  $: isEdit = editMonth === model.Month;

  onMount(async () => {
    if (editMonth === "" || editMonth === model.Month) {
      document
        .querySelector(`#edit${model.Month}`)
        ?.classList.remove("is-hidden");
    } else {
      document.querySelector(`#edit${model.Month}`)?.classList.add("is-hidden");
    }
  });

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
    class="has-text-right"
    class:has-background-danger-light={model.IsNotCost}
  >
    {model.Cost.toLocaleString()}
  </td>
  <EditTd bind:value={saving} {isEdit} />
  <EditTd bind:value={memo} {isEdit} isText />
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
