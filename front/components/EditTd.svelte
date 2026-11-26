<script lang="ts">
    /** テキストとして編集する場合true、数値として編集する場合、false */
    export let isText: boolean = false;
    /** 表示する値 */
    export let value: Number | string;
    /** 編集モードの場合、true */
    export let isEdit: boolean = false;
</script>

<td class="edit-cell">
    {#if isEdit && isText}
        <input type="text" class="edit-input" bind:value />
    {:else if isEdit}
        <input type="number" class="edit-input" bind:value />
    {:else}
        {value}
    {/if}
</td>

<style>
    td {
        width: 90px;
    }

    .edit-cell {
        padding-left: 0;
        padding-right: 0;
        text-align: right;
    }

    .edit-input {
        box-sizing: border-box;
        width: 100%;
        min-height: 2rem;
        border: 1px solid #d1d5db;
        border-radius: 4px;
        padding: 0;
        background: #fff;
        color: #111827;
        font: inherit;
        text-align: right;
    }

    .edit-input:focus {
        border-color: #2563eb;
        outline: 2px solid rgb(37 99 235 / 20%);
        outline-offset: 0;
    }
</style>
