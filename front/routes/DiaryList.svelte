<script lang="ts">
  import { push } from "svelte-spa-router";
  import type { listType } from "../models/diaryModels.js";
  export let selectMonth: string | null = `${new Date().getFullYear()}-${(
    "00" +
    (new Date().getMonth() + 1)
  ).slice(-2)}`;
  export const showList = () => {
    let url = selectMonth !== null ? selectMonth.replace("-", "/") : null;
    url = `./api/diary/${url}`;
    fetch(url)
      .then((r) => r.json())
      .then((r) => r as listType)
      .then((r) => {
        model = r;
        if (r.Lines.length == 0 && r.WritedMonths.length > 0) {
          selectMonth = r.WritedMonths[0];
        }
      });
  };

  let model: listType = { WritedMonths: [], Lines: [] };

  /** 追加ボタンクリックイベント */
  const addClick = () => push("/d/add");

  /** リストクリックイベント */
  const listClick = async (e: any, l: string) => push("/d/" + l);

  $: {
    let url = selectMonth !== null ? selectMonth.replace("-", "/") : null;
    url = `./api/diary/${url}`;
    fetch(url)
      .then((r) => r.json())
      .then((r) => r as listType)
      .then((r) => {
        model = r;
        if (r.Lines.length == 0 && r.WritedMonths.length > 0) {
          selectMonth = r.WritedMonths[0];
        }
      });
  }
</script>

<div class="card px-10">
  <div class="card-content">
    <div class="columns">
      <div class="column">
        <div class="select">
          <select class="select" bind:value={selectMonth}>
            {#each model.WritedMonths as v}
              <option value={v}>{v}</option>
            {/each}
          </select>
        </div>
      </div>
      <div class="column">
        <button class="button is-primary" on:click={addClick}
          ><i class="material-icons">add</i>add</button
        >
      </div>
    </div>
    <table class="table is-hoverable is-fullwidth">
      <tbody>
        {#each model.Lines as v}
          <tr on:click={(e) => listClick(e, v.Day)}>
            <td>
              <span>
                {v.Day}
                {v.Outline}{#if v.IsDetail}<i
                    class="material-icons has-text-grey-light"
                    style="vertical-align:middle">note</i
                  >{/if}</span
              >
            </td>
            <td>
              <div class="tags are-medium">
                {#each v.Tags as t}
                  <span class="tag">{t}</span>
                {/each}
                {#if v.HCount > 0}
                  <i class="material-icons">local_activity</i>
                {/if}
              </div>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>
