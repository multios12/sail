<script lang="ts">
  import { createEventDispatcher, onMount } from "svelte";
  import TagsInput from "./TagsInput.svelte";

  const dispatch = createEventDispatcher();
  export let isDayEdit: boolean;
  export let Day = "";
  export let Outline = "";
  export let Detail = "";
  export let Tags = [];

  /** 送信 クリックイベント */
  const sendClick = async () => {
    if (Day == null) {
      return;
    }

    let url = Day.replaceAll("-", "/");
    url = `./api/diary/${url}`;
    const init = {
      method: "post",
      body: JSON.stringify({ Tags, Outline, Detail }),
    };
    await fetch(url, init);
    Day = "";
    dispatch("update");
  };

  /** キャンセル クリックイベント */
  const cancelClick = () => (Day = "");

  /** 削除 クリックイベント */
  const deleteClick = async () => {
    let url = Day.replaceAll("-", "/");
    url = `./api/diary/${url}`;
    await fetch(url, { method: "delete" });
    Day = "";
    dispatch("update");
  };
</script>

<div class="modal" class:is-active={Day != ""}>
  <div class="modal-background" />
  <div class="modal-card">
    <header class="modal-card-head sp-panel-heading">
      <input type="date" class="input" bind:value={Day} disabled={!isDayEdit} />
      {#if !isDayEdit}
        <button
          class="button is-inverted is-small has-text-danger"
          on:click={deleteClick}
        >
          <i class="material-icons">delete</i>
        </button>
      {/if}
    </header>
    <section class="modal-card-body">
      <!-- Card Content -->
      <input
        type="text"
        placeholder="outline"
        class="input"
        bind:value={Outline}
      />
      <div class="field">
        <div class="control">
          <TagsInput bind:items={Tags} />
        </div>
        <textarea class="textarea" placeholder="detail" bind:value={Detail} />
      </div>
    </section>
    <footer class="modal-card-foot">
      <button class="button is-primary" on:click={sendClick}>send</button>
      <button class="button" on:click={cancelClick}>cancel</button>
    </footer>
  </div>
</div>
