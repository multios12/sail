<script lang="ts">
    /** テキストとして編集する場合true、数値として編集する場合、false */
    export let isText: boolean = false;
    /** 表示する値 */
    export let value: Number | string;
    /** 編集モードの場合、true */
    export let isEdit: boolean = false;
</script>

<td class="px-0 has-text-right">
    {#if isEdit && isText}
        <input type="text" class="input p-0 has-text-right" bind:value />
    {:else if isEdit}
        <input type="number" class="input p-0 has-text-right" bind:value />
    {:else}
        {value}
    {/if}
</td>

<style>
    td {
        width: 90px;
    }
</style>
