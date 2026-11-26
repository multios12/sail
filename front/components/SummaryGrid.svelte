<script lang="ts">
  type SummaryItem = {
    Name: string;
    Value: number;
  };

  export let items: SummaryItem[] = [];

  const format = (value: number) => {
    return value.toLocaleString();
  };
</script>

<div class="summary-grid">
  {#each items as item}
    <article class="summary-card">
      <span>{item.Name}</span>
      <strong>{format(item.Value)}</strong>
    </article>
  {/each}
</div>

<style>
  .summary-grid {
    display: grid;
    gap: 0.75rem;
    grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
    padding: 1rem 1.25rem;
  }

  .summary-card {
    background: #172033;
    border: 1px solid #2f3b52;
    border-radius: 8px;
    padding: 0.875rem 1rem;
  }

  .summary-card span {
    color: #a8b3c7;
    display: block;
    font-size: 0.85rem;
  }

  .summary-card strong {
    color: #f8fafc;
    display: block;
    font-size: 1.35rem;
    font-variant-numeric: tabular-nums;
    line-height: 1.2;
    margin-top: 0.25rem;
    text-align: right;
  }

  @media screen and (max-width: 768px) {
    .summary-grid {
      grid-template-columns: 1fr;
      padding: 1rem;
    }
  }
</style>
