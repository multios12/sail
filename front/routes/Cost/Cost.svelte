<script lang="ts">
  import { fetchJson } from "../../lib/api";
  import type { BalanceYear } from "../../models/balanceModels";
  import MonthTr from "./MonthList.svelte";
  import { link } from "svelte-spa-router";

  export let params: { year: string | undefined } = { year: undefined };
  $: {
    let url = `/api/balance/${params.year || new Date().getFullYear()}`;
    fetchJson<BalanceYear>(url).then((r) => {
      if (r) {
        model = r;
      }
    });
  }
  let model: BalanceYear = {
    Year: new Date().getFullYear().toString(),
    EnableYears: [],
    Balances: [],
  };

  // 現在編集している行
  let editMonth = "";
</script>

<div class="cost-page">
  {#each model.EnableYears as y}
    <section class="cost-card">
      <header class="cost-card-header">
        <div class="cost-card-title">
          {#if y === model.Year}
            {model.Year}年支出
          {:else}
            <a href="/cost/{y}" use:link>{y}年</a>
          {/if}
        </div>
      </header>
      {#if y === model.Year}
        <div class="cost-card-content">
          <table class="cost-table">
            <thead>
              <tr>
                <th></th>
                <th>合計</th>
                <th>住宅</th>
                <th>水道</th>
                <th>電気</th>
                <th>ガス</th>
                <th>携帯</th>
                <th>通信</th>
                <th>納税</th>
              </tr>
            </thead>
            <tbody>
              {#each model.Balances as v}
                <MonthTr Value={v} {editMonth} />
              {/each}
            </tbody>
          </table>
        </div>
      {/if}
    </section>
  {/each}
</div>

<style>
  .cost-page {
    display: grid;
    gap: 1rem;
    color: #e5e7eb;
  }

  .cost-card {
    overflow: hidden;
    border: 1px solid #374151;
    border-radius: 8px;
    background: #111827;
    box-shadow: 0 18px 45px rgb(0 0 0 / 22%);
  }

  .cost-card-header {
    border-bottom: 1px solid #1f2937;
    background: #0f172a;
    padding: 0.85rem 1rem;
  }

  .cost-card-title {
    color: #f8fafc;
    font-size: 1.1rem;
    font-weight: 700;
  }

  .cost-card-title a {
    color: #93c5fd;
    text-decoration: none;
  }

  .cost-card-title a:hover {
    color: #bfdbfe;
  }

  .cost-card-content {
    overflow-x: auto;
  }

  .cost-table {
    width: 100%;
    border-collapse: collapse;
    background: #111827;
    color: #e5e7eb;
  }

  .cost-table th {
    border-bottom: 1px solid #263244;
    background: #0f172a;
    color: #cbd5e1;
    padding: 0.55rem 0.65rem;
    text-align: right;
    white-space: nowrap;
  }

  .cost-table th:first-child {
    text-align: left;
  }

  :global(.cost-table td) {
    border-bottom: 1px solid #263244;
    padding: 0.45rem 0.55rem;
    vertical-align: middle;
    white-space: nowrap;
  }

  :global(.cost-table tbody tr:nth-child(even)) {
    background: #151f31;
  }

  :global(.cost-table tbody tr:hover) {
    background: #1f2a3d;
  }
</style>
