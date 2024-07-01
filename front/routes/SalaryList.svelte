<script lang="ts">
  import type { SalaryModel } from "../models/balanceModels.js";
  import { link } from "svelte-spa-router";

  export let params: { year: string | undefined } = { year: undefined };
  $: {
    params.year = params.year || String(new Date().getFullYear());

    const url = `./api/salary/${params.year}`;
    fetch(url)
      .then((r) => r.json())
      .then((r) => (model = r));
  }
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

<div>
  {#each model.EnableYears as year}
    {#if year === model.Year}
      <div class="card px-10">
        <div class="card-header">
          <div class="card-header-title">{model.Year}年収入</div>
        </div>
        <div class="card-content">
          <nav class="level">
            {#each model.Totals as Value}
              <article class="tile is-child box">
                <p class="is-size-6">{Value.Name}</p>
                <p class="is-size-4">{Value.Value.toLocaleString()}</p>
              </article>
            {/each}
          </nav>
          <table class="table is-striped is-hoverable">
            <thead>
              <tr>
                <th />
                <th>支給額</th>
                <th>差引支給額</th>
                <th>経費支給額</th>
                <th>出勤日数</th>
              </tr>
            </thead>
            <tbody>
              {#each model.Details as Value}
                <tr class={Value.IsError ? "has-background-danger-light" : ""}>
                  <td>
                    <a href={getSalaryUrl(Value.Month)} use:link>
                      {Value.Title}
                    </a>
                  </td>
                  <td>{format(Value.Totals[0].Value)}</td>
                  <td>{format(Value.Totals[2].Value)}</td>
                  {#if Value.Month.length === 6}
                    <td>{Value.Expense.toLocaleString()}</td>
                    <td>{Value.Counts[0].Value}</td>
                  {:else}
                    <td></td>
                    <td></td>
                  {/if}
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
    {:else}
      <div class="card">
        <div class="card-header">
          <div class="card-header-title">
            <a href="/salary/{year}" use:link>{year}年</a>
          </div>
        </div>
      </div>
    {/if}
  {/each}
</div>
