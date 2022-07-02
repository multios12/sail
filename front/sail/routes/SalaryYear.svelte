<script lang="ts">
  import type { SalaryModel } from "../models";
  import { link } from "svelte-spa-router";

  export let params: { year: string } = { year: undefined };
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
                  <td
                    ><a
                      href="/salary/{Value.Month.substring(
                        0,
                        4
                      )}/{Value.Month.substring(4)}"
                      use:link>{Value.Title}</a
                    ></td
                  >
                  <td
                    >{Value.Totals
                      ? Value.Totals[0].Value.toLocaleString()
                      : 0}</td
                  >
                  <td
                    >{Value.Totals
                      ? Value.Totals[2].Value.toLocaleString()
                      : 0}</td
                  >
                  <td
                    >{Value.Month.length === 6
                      ? Value.Expense.toLocaleString()
                      : ""}</td
                  >
                  <td
                    >{Value.Month.length === 6
                      ? Value.Counts
                        ? Value.Counts[0].Value
                        : 0
                      : ""}</td
                  >
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
