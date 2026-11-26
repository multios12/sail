<script lang="ts">
  import { fetchJson } from "../lib/api";
  import type { SalaryModel } from "../models/balanceModels.js";
  import SummaryGrid from "../components/SummaryGrid.svelte";
  import UploadButton from "../components/UploadButton.svelte";
  import YearTabs from "../components/YearTabs.svelte";
  import { link } from "svelte-spa-router";

  export let params: { year: string | undefined } = { year: undefined };
  const currentYear = String(new Date().getFullYear());
  let requestId = 0;

  async function loadSalary(year: string) {
    const id = ++requestId;
    const result = await fetchJson<SalaryModel>(`./api/salary/${year}`);
    if (result && id === requestId) model = result;
  }

  $: loadSalary(params.year || currentYear);
  export let model: SalaryModel = {
    Year: new Date().getFullYear().toString(),
    Details: [],
    Totals: [],
    EnableYears: [],
  };
  const getSalaryUrl = (month: string) => {
    return `/salary/${month.substring(0, 4)}/${month.substring(4)}`;
  };
  const format = (value: number) => {
    return value.toLocaleString();
  };
</script>

<div class="salary-page">
  <div class="salary-toolbar">
    <div class="salary-year-tabs">
      <YearTabs
        years={model.EnableYears}
        currentYear={model.Year}
        basePath="/salary"
        ariaLabel="給与年選択"
      />
    </div>

    <div class="salary-upload">
      <UploadButton onImported={() => loadSalary(params.year || currentYear)} />
    </div>
  </div>

  <section class="salary-panel">
    <header class="salary-header">
      <div>
        <h2>{model.Year}年収入</h2>
      </div>
      <span class="salary-count">{model.Details.length}件</span>
    </header>

    <SummaryGrid items={model.Totals} />

    <div class="table-wrap desktop-table">
      <table class="salary-table">
        <thead>
          <tr>
            <th>対象月</th>
            <th class="amount">総支給額</th>
            <th class="amount">差引支給額</th>
            <th class="amount">経費支給額</th>
            <th class="amount">出勤日数</th>
          </tr>
        </thead>
        <tbody>
          {#each model.Details as Value}
            <tr class:salary-row-error={Value.IsError}>
              <td>
                <a class="month-link" href={getSalaryUrl(Value.Month)} use:link>
                  {Value.Title}
                </a>
                {#if Value.IsError}
                  <span class="error-badge inline-error-badge">取得エラー</span>
                {/if}
              </td>
              <td class="amount">{format(Value.Totals[0].Value)}</td>
              <td class="amount">{format(Value.Totals[2].Value)}</td>
              {#if Value.Month.length === 6}
                <td class="amount">{format(Value.Expense)}</td>
                <td class="amount">{format(Value.Counts?.[0]?.Value ?? 0)}</td>
              {:else}
                <td></td>
                <td></td>
              {/if}
            </tr>
          {/each}
        </tbody>
      </table>
    </div>

    <div class="mobile-list">
      {#each model.Details as Value}
        <a
          class:salary-row-error={Value.IsError}
          class="mobile-salary-card"
          href={getSalaryUrl(Value.Month)}
          use:link
        >
          <div class="mobile-card-header">
            <span>{Value.Title}</span>
            {#if Value.IsError}
              <span class="error-badge">取得エラー</span>
            {/if}
          </div>
          <dl>
            <div>
              <dt>支給</dt>
              <dd>{format(Value.Totals[0].Value)}</dd>
            </div>
            <div>
              <dt>差引</dt>
              <dd>{format(Value.Totals[2].Value)}</dd>
            </div>
            {#if Value.Expense != 0}
              <div>
                <dt>経費</dt>
                <dd>{format(Value.Expense)}</dd>
              </div>
            {/if}
            {#if Value.Month.length === 6}
              <div>
                <dt>出勤</dt>
                <dd>{format(Value.Counts?.[0]?.Value ?? 0)}日</dd>
              </div>
            {/if}
          </dl>
        </a>
      {/each}
    </div>
  </section>
</div>

<style>
  .salary-page {
    color: #e5e7eb;
    display: grid;
    gap: 1rem;
  }

  .salary-toolbar {
    align-items: flex-start;
    display: flex;
    gap: 1rem;
    justify-content: space-between;
  }

  .salary-year-tabs {
    flex: 1 1 auto;
    min-width: 0;
  }

  .salary-upload {
    flex: 0 0 auto;
  }

  .salary-panel {
    background: #111827;
    border: 1px solid #374151;
    border-radius: 8px;
    box-shadow: 0 18px 45px rgb(0 0 0 / 22%);
    overflow: hidden;
  }

  .salary-header {
    align-items: center;
    background: #0f172a;
    border-bottom: 1px solid #1f2937;
    display: flex;
    justify-content: space-between;
    gap: 1rem;
    padding: 1rem 1.25rem;
  }

  .salary-header h2 {
    color: #f8fafc;
    font-size: 1.25rem;
    font-weight: 700;
    line-height: 1.3;
    margin: 0;
  }

  .salary-count {
    background: #1f2937;
    border-radius: 999px;
    color: #dbeafe;
    font-size: 0.85rem;
    font-weight: 700;
    padding: 0.25rem 0.65rem;
    white-space: nowrap;
  }

  .table-wrap {
    border-top: 1px solid #1f2937;
    overflow-x: auto;
  }

  .salary-table {
    background: #111827;
    border-collapse: collapse;
    color: #e5e7eb;
    margin-bottom: 0;
    width: 100%;
  }

  .salary-table thead th {
    background: #0f172a;
    border-color: #263244;
    color: #cbd5e1;
  }

  .salary-table th,
  .salary-table td {
    border-color: #263244;
    padding: 0.5rem 0.875rem;
    vertical-align: middle;
  }

  .salary-table th.amount {
    text-align: right;
  }

  .salary-table tbody tr {
    background: #111827;
  }

  .salary-table tbody tr:nth-child(even) {
    background: #151f31;
  }

  .salary-table tbody tr:hover {
    background: #1f2a3d;
  }

  .amount {
    font-variant-numeric: tabular-nums;
    text-align: right;
    white-space: nowrap;
  }

  .month-link {
    color: #93c5fd;
    font-weight: 700;
  }

  .month-link:hover {
    color: #bfdbfe;
  }

  .salary-row-error {
    box-shadow: inset 4px 0 0 #f14668;
  }

  .mobile-list {
    display: grid;
    gap: 0.75rem;
    padding: 0 1rem 1rem;
  }

  .mobile-salary-card {
    background: #172033;
    border: 1px solid #2f3b52;
    border-radius: 8px;
    color: #e5e7eb;
    display: block;
    padding: 1rem;
  }

  .mobile-salary-card:hover {
    border-color: #4b6386;
    color: #f8fafc;
  }

  .mobile-card-header {
    align-items: center;
    display: flex;
    font-weight: 700;
    gap: 0.5rem;
    justify-content: space-between;
    margin-bottom: 0.75rem;
  }

  .mobile-salary-card dl {
    display: grid;
    gap: 0.5rem;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    margin: 0;
  }

  .mobile-salary-card dl div {
    background: #101827;
    border-radius: 6px;
    padding: 0.625rem 0.75rem;
  }

  .mobile-salary-card dt {
    color: #94a3b8;
    font-size: 0.78rem;
  }

  .mobile-salary-card dd {
    color: #f8fafc;
    font-variant-numeric: tabular-nums;
    font-weight: 700;
    margin: 0.15rem 0 0;
    text-align: right;
  }

  .error-badge {
    background: #3f1d2d;
    border-radius: 999px;
    color: #fecdd3;
    display: inline-flex;
    font-size: 0.75rem;
    font-weight: 700;
    line-height: 1;
    padding: 0.28rem 0.5rem;
    white-space: nowrap;
  }

  .inline-error-badge {
    margin-left: 0.5rem;
  }

  @media screen and (width < 720px) {
    .desktop-table {
      display: none;
    }

    .salary-toolbar {
      align-items: stretch;
      flex-direction: column;
      gap: 0.75rem;
    }

    .salary-upload {
      display: flex;
      justify-content: flex-end;
    }

    .salary-header {
      padding: 0.875rem 1rem;
    }
  }

  @media screen and (width >= 720px) {
    .mobile-list {
      display: none;
    }
  }
</style>
