<script lang="ts">
  import { onMount } from "svelte";
  import type { detailType, lineType, listType } from "../components/models";
  import Detail from "../components/Detail.svelte";
  let edit: detailType = { Day: "", Outline: "", Tags: [], Detail: "" };
  let model: listType = { WritedMonths: [], Lines: [] };
  let isDayEdit: boolean;
  let selectMonth = "2022-02";

  const addClick = () => {
    isDayEdit = true;
    const dt = new Date();
    edit.Day = `${dt.getFullYear()}-`;
    edit.Day += ("00" + (dt.getMonth() + 1)).slice(-2);
    edit.Day += `-${("00" + dt.getDate()).slice(-2)}`;
    edit.Outline = "";
    edit.Tags = [];
    edit.Detail = "";
  };

  /** リストクリックイベント */
  const listClick = async (e: any, l: lineType) => {
    isDayEdit = false;

    let url = l.Day.replaceAll("-", "/");
    url = `./api/diary/${url}`;
    const r = await fetch(url, { method: "get" }).then((r) => r.json());
    const s = r as detailType;
    edit = s;
  };

  const showLines = async () => {
    let url = selectMonth.replace("-", "/");
    url = `./api/diary/${url}`;
    const r = await fetch(url).then((r) => r.json());
    model = r as listType;
    if (model.Lines != null && model.WritedMonths.length > 0) {
      selectMonth = model.WritedMonths[0];
      let url = selectMonth.substring(0, 4) + "/" + selectMonth.substring(5, 7);
      url = `./api/diary/${url}`;
      const r = await fetch(url).then((r) => r.json());
      model = r as listType;
    } else {
      model.Lines = [];
    }
  };

  onMount(() => {
    const dt = new Date();
    selectMonth = ("00" + (dt.getMonth() + 1)).slice(-2);
    selectMonth = `${dt.getFullYear()}-${selectMonth}`;
    showLines();
  });
</script>

<Detail
  {isDayEdit}
  bind:Day={edit.Day}
  Outline={edit.Outline}
  Tags={edit.Tags}
  Detail={edit.Detail}
  on:update={showLines}
/>
<div class="box">
  <div class="card px-10">
    <div class="card-content">
      <div class="columns">
        <div class="column">
          <div class="select">
            <select
              class="select"
              bind:value={selectMonth}
              on:change={showLines}
            >
              {#each model.WritedMonths as v}
                <option value={v}>{v}</option>
              {/each}
            </select>
          </div>
        </div>
        <div class="column">
          <button class="button is-primary" on:click={addClick}>add</button>
        </div>
      </div>
      <table class="table is-hoverable is-fullwidth">
        <tbody>
          {#each model.Lines as v}
            <tr on:click={(e) => listClick(e, v)}>
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
</div>
