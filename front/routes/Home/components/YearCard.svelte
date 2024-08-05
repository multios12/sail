<script lang="ts">
  import type { BalanceYear } from "../../../models/balanceModels";
  import HomeMonthTr from "./MonthTr.svelte";
  import HomeMonthTr2 from "./MonthTr2.svelte";
  import { link } from "svelte-spa-router";

  export let year: string;
  export let model: BalanceYear;

  // 現在編集している行
  let editMonth = "";
</script>

<div class="card" class:px-10={year === model.Year}>
  <div class="card-header">
    <div class="card-header-title">
      {#if year === model.Year}
        {model.Year}年
      {:else}
        <a href="/{year}" use:link>{year}年</a>
      {/if}
    </div>
  </div>
  {#if year === model.Year}
    <div class="card-content is-hidden-mobile">
      <table class="table is-striped is-hoverable">
        <thead>
          <tr>
            <th />
            <th>給与総額</th>
            <th><a href="/salary/{year}" use:link>振込額</a></th>
            <th><a href="/cost/{year}" use:link>固定支出額</a></th>
            <th>貯蓄額</th>
            <th>メモ</th>
            <th />
          </tr>
        </thead>
        <tbody>
          {#each model.Balances as balance}
            <HomeMonthTr model={balance} {editMonth} />
          {/each}
        </tbody>
      </table>
    </div>
    <div class="card-content is-hidden-tablet">
      <table class="table is-striped is-hoverable">
        <tbody>
          {#each model.Balances as balance}
            <HomeMonthTr2 model={balance} {editMonth} />
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</div>

<style>
  .table {
    width: 100%;
  }
</style>
