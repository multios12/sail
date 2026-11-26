<script lang="ts">
  import { onMount } from "svelte";
  import { link } from "svelte-spa-router";
  import { fetchJson } from "../lib/api";
  import type { SalaryMonthModel } from "../models/balanceModels.js";

  export let params: { year: string | undefined; month: string | undefined } = {
    year: undefined,
    month: undefined,
  };

  let loading = false;
  let model: SalaryMonthModel = {
    Month: "202010",
    Counts: [],
    Times: [],
    Salarys: [],
    Costs: [],
    Totals: [],
    Expense: 0,
    Expenses: [],
    Images: [],
  };

  $: title = model.Title || `${params.year ?? ""}年${params.month ?? ""}月`;
  $: backUrl = `/salary/${params.year ?? new Date().getFullYear()}`;

  const fetchSalary = async () => {
    loading = true;
    const url = `./api/salary/${params.year}/${params.month}`;
    try {
      const r = await fetchJson<SalaryMonthModel>(url);
      if (r) {
        model = r;
      }
    } finally {
      loading = false;
    }
  };

  const format = (value: number) => {
    return value.toLocaleString();
  };

  onMount(() => {
    fetchSalary();
  });
</script>

<div class="salary-detail-page">
  <div class="detail-toolbar">
    <a class="back-button" href={backUrl} use:link>
      <span class="material-icons">arrow_back</span>
      <span>一覧へ</span>
    </a>
  </div>

  <section class="detail-panel">
    <header class="detail-header">
      <div>
        <p class="detail-caption">Salary detail</p>
        <h2>{title}</h2>
      </div>
    </header>

    <div class:content-loading={loading} class="detail-content">
      <section class="detail-section salary-section">
        <div class="section-heading">
          <div>
            <span class="section-label">支給</span>
            <strong>{format(model.Totals[0]?.Value ?? 0)}</strong>
          </div>
        </div>
        <div class="item-grid">
          {#each model.Salarys as item}
            <article class="detail-item">
              <span>{item.Name}</span>
              <strong>{format(item.Value)}</strong>
            </article>
          {/each}
        </div>
      </section>

      <section class="detail-section cost-section">
        <div class="section-heading">
          <div>
            <span class="section-label">控除</span>
            <strong>{format(model.Totals[1]?.Value ?? 0)}</strong>
          </div>
        </div>
        <div class="item-grid">
          {#each model.Costs as item}
            <article class="detail-item">
              <span>{item.Name}</span>
              <strong>{format(item.Value)}</strong>
            </article>
          {/each}
        </div>
      </section>

      <section class="detail-section paid-section">
        <div class="section-heading">
          <div>
            <span class="section-label">差引</span>
            <strong>{format(model.Totals[2]?.Value ?? 0)}</strong>
          </div>
        </div>
        <div class="paid-spacer"></div>
      </section>

      {#each model.Images as i}
        <img
          src={`./api/salary/${params.year}/${params.month}/images/${i}`}
          alt="salary"
        />
      {/each}
    </div>
  </section>
</div>

<style>
  .salary-detail-page {
    color: #e5e7eb;
    display: grid;
    gap: 0.6rem;
  }

  .detail-toolbar {
    align-items: center;
    display: flex;
    gap: 0.5rem;
    justify-content: space-between;
  }

  .back-button {
    align-items: center;
    border: 1px solid #374151;
    border-radius: 4px;
    background: #111827;
    color: #e5e7eb;
    display: inline-flex;
    gap: 0.35rem;
    min-height: 2.5rem;
    padding: 0.5rem 0.8rem;
    text-decoration: none;
  }

  .back-button:hover {
    background: #1f2937;
    color: #fff;
  }

  .detail-panel {
    background: #111827;
    border: 1px solid #374151;
    border-radius: 8px;
    box-shadow: 0 18px 45px rgb(0 0 0 / 22%);
    overflow: hidden;
  }

  .detail-header {
    align-items: center;
    background: #0f172a;
    border-bottom: 1px solid #1f2937;
    display: flex;
    gap: 0.75rem;
    justify-content: space-between;
    padding: 0.7rem 0.9rem;
  }

  .detail-caption {
    color: #94a3b8;
    font-size: 0.78rem;
    font-weight: 700;
    letter-spacing: 0;
    margin-bottom: 0.05rem;
    text-transform: uppercase;
  }

  .detail-header h2 {
    color: #f8fafc;
    font-size: 1.1rem;
    font-weight: 700;
    line-height: 1.3;
    margin: 0;
  }

  .detail-content {
    display: grid;
    gap: 0.65rem;
    padding: 0.75rem;
  }

  .content-loading {
    opacity: 0.45;
    pointer-events: none;
  }

  .detail-section {
    background: #172033;
    border: 1px solid #2f3b52;
    border-radius: 8px;
    overflow: hidden;
  }

  .section-heading {
    align-items: center;
    border-left: 4px solid #64748b;
    display: flex;
    justify-content: space-between;
    padding: 0.55rem 0.75rem;
  }

  .salary-section .section-heading {
    border-left-color: #22c55e;
  }

  .cost-section .section-heading {
    border-left-color: #f43f5e;
  }

  .paid-section .section-heading {
    border-left-color: #38bdf8;
  }

  .section-heading div {
    display: grid;
    gap: 0.2rem;
    width: 100%;
  }

  .section-label {
    color: #a8b3c7;
    font-size: 0.78rem;
    font-weight: 700;
  }

  .section-heading strong {
    color: #f8fafc;
    font-size: 1.2rem;
    font-variant-numeric: tabular-nums;
    line-height: 1.2;
    text-align: right;
  }

  .item-grid {
    display: grid;
    gap: 0.5rem;
    grid-template-columns: repeat(auto-fit, minmax(140px, 1fr));
    padding: 0.65rem;
  }

  .detail-item {
    background: #101827;
    border: 1px solid #263244;
    border-radius: 6px;
    padding: 0.5rem 0.6rem;
  }

  .detail-item span {
    color: #a8b3c7;
    display: block;
    font-size: 0.78rem;
  }

  .detail-item strong {
    color: #f8fafc;
    display: block;
    font-variant-numeric: tabular-nums;
    margin-top: 0.1rem;
    text-align: right;
  }

  .detail-content > img {
    display: block;
    justify-self: center;
    max-width: 100%;
    width: 60vw;
  }

  @media screen and (min-width: 769px) {
    .salary-section,
    .cost-section,
    .paid-section {
      display: grid;
      grid-template-columns: 150px minmax(0, 1fr);
    }

    .salary-section .section-heading,
    .cost-section .section-heading,
    .paid-section .section-heading {
      align-items: stretch;
    }

    .salary-section .item-grid,
    .cost-section .item-grid {
      align-items: stretch;
      display: flex;
      overflow-x: auto;
    }

    .salary-section .detail-item,
    .cost-section .detail-item {
      flex: 1 0 120px;
      min-width: 0;
    }
  }

  @media screen and (max-width: 768px) {
    .detail-toolbar {
      align-items: stretch;
      flex-direction: column;
    }

    .detail-toolbar .back-button {
      justify-content: center;
      width: 100%;
    }

    .detail-header {
      align-items: flex-start;
      flex-direction: column;
      padding: 0.7rem 0.75rem;
    }

    .detail-content {
      padding: 0.65rem;
    }

    .item-grid {
      grid-template-columns: 1fr;
    }

    .detail-content > img {
      width: 100%;
    }
  }
</style>
