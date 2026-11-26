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

<section class:current-year-card={year === model.Year} class="year-card">
  <header class="year-card-header">
    <div class="year-card-title">
      {model.Year}年
    </div>
  </header>
  <div class="year-card-content desktop-content">
    <table class="balance-table">
      <thead>
        <tr>
          <th></th>
          <th>給与総額</th>
          <th><a href="/salary/{year}" use:link>振込額</a></th>
          <th><a href="/cost/{year}" use:link>固定支出額</a></th>
          <th>貯蓄額</th>
          <th>メモ</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {#each model.Balances as balance}
          <HomeMonthTr model={balance} {editMonth} />
        {/each}
      </tbody>
    </table>
  </div>
  <div class="year-card-content mobile-content">
    <table class="balance-table mobile-balance-table">
      <tbody>
        {#each model.Balances as balance}
          <HomeMonthTr2 model={balance} {editMonth} />
        {/each}
      </tbody>
    </table>
  </div>
</section>

<style>
  .year-card {
    overflow: hidden;
    border: 1px solid #374151;
    border-radius: 8px;
    background: #111827;
    color: #e5e7eb;
    box-shadow: 0 18px 45px rgb(0 0 0 / 22%);
  }

  .current-year-card {
    padding: 0.6rem;
  }

  .year-card-header {
    border-bottom: 1px solid #1f2937;
    background: #0f172a;
    padding: 0.85rem 1rem;
  }

  .year-card-title {
    color: #f8fafc;
    font-size: 1.1rem;
    font-weight: 700;
  }

  .year-card-content {
    overflow-x: auto;
  }

  .balance-table {
    width: 100%;
    border-collapse: collapse;
    background: #111827;
    color: #e5e7eb;
  }

  .balance-table th {
    border-bottom: 1px solid #263244;
    background: #0f172a;
    color: #cbd5e1;
    padding: 0.55rem 0.65rem;
    text-align: right;
    white-space: nowrap;
  }

  .balance-table th:first-child {
    text-align: left;
  }

  .balance-table a {
    color: #93c5fd;
    text-decoration: none;
  }

  .balance-table a:hover {
    color: #bfdbfe;
  }

  :global(.balance-table td) {
    border-bottom: 1px solid #263244;
    padding: 0.45rem 0.55rem;
    vertical-align: middle;
  }

  :global(.balance-table tbody tr:nth-child(even)) {
    background: #151f31;
  }

  :global(.balance-table tbody tr:hover) {
    background: #1f2a3d;
  }

  @media screen and (max-width: 768px) {
    .desktop-content {
      display: none;
    }
  }

  @media screen and (min-width: 769px) {
    .mobile-content {
      display: none;
    }
  }
</style>
