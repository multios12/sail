<script lang="ts">
  import type { BalanceYear } from "../../models/balanceModels";
  import MonthTr from "./MonthList.svelte";
  import { link } from "svelte-spa-router";

  export let params: { year: string | undefined } = { year: undefined };
  $: {
    let url = `./api/balance/${params.year || new Date().getFullYear()}`;
    fetch(url)
      .then((r) => r.json())
      .then((r) => (model = r));
  }
  let model: BalanceYear = {
    Year: new Date().getFullYear().toString(),
    EnableYears: [],
    Balances: [],
  };

  // 現在編集している行
  let editMonth = "";
</script>

<div>
  {#each model.EnableYears as y}
    <div class="card">
      <div class="card-header">
        <div class="card-header-title">
          {#if y === model.Year}
            {model.Year}年支出
          {:else}
            <a href="/cost/{y}" use:link>{y}年</a>
          {/if}
        </div>
      </div>
      {#if y === model.Year}
        <div class="card-content">
          <table class="table is-striped is-hoverable">
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
    </div>
  {/each}
</div>
