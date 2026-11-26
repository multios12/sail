<script lang="ts">
  import type { BalanceItem } from "../../../models/balanceModels";

  export let editMonth: String;
  export let model: BalanceItem;

  let saving = model.Saving;
  let memo = model.Memo;

  $: month =
    model.Month.substring(0, 4) + "年" + model.Month.substring(4) + "月";
  $: isEdit = editMonth === model.Month;
</script>

<tr><td colspan="2">{month}</td></tr>
<tr>
  <td>
    <table>
      <tbody>
        <tr>
          <td>給与総額</td>
          <td>{model.Salary.toLocaleString()}</td>
        </tr>
        <tr>
          <td>振込額</td>
          <td>{model.Paid.toLocaleString()}</td>
        </tr>
        <tr>
          <td>固定支出額</td>
          <td>{model.Cost.toLocaleString()}</td>
        </tr>
        <tr>
          <td>貯蓄</td>
          <td>{saving}</td>
        </tr>
        <tr>
          <td>メモ</td>
          <td>{memo}</td>
        </tr>
      </tbody>
    </table>
  </td>
</tr>
