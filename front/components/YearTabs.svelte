<script lang="ts">
  import { link } from "svelte-spa-router";

  export let years: string[] = [];
  export let currentYear = "";
  export let basePath = "";
  export let ariaLabel = "年選択";

  const getUrl = (year: string) => {
    return basePath ? `${basePath}/${year}` : `/${year}`;
  };
</script>

<nav class="year-tabs" aria-label={ariaLabel}>
  <ul>
    {#each years as year}
      <li class:active-tab={year === currentYear}>
        <a href={getUrl(year)} use:link>{year}年</a>
      </li>
    {/each}
  </ul>
</nav>

<style>
  .year-tabs {
    margin-bottom: 0;
    overflow-x: auto;
  }

  .year-tabs ul {
    display: flex;
    flex-wrap: nowrap;
    margin: 0;
    border-bottom: 1px solid #374151;
    padding: 0;
    list-style: none;
  }

  .year-tabs li {
    flex: 0 0 auto;
  }

  .year-tabs a {
    display: block;
    position: relative;
    border: 1px solid #374151;
    border-radius: 8px 8px 0 0;
    padding: 0.45rem 0.75rem;
    color: #cbd5e1;
    text-decoration: none;
    white-space: nowrap;
  }

  .year-tabs li.active-tab a {
    z-index: 1;
    background: #111827;
    border-color: #4b5563;
    border-width: 2px 2px 1px;
    border-bottom-color: transparent;
    color: #f9fafb;
  }

  .year-tabs li.active-tab a::after {
    position: absolute;
    right: 0;
    bottom: -2px;
    left: 0;
    height: 3px;
    background: #111827;
    content: "";
  }

  .year-tabs li.active-tab a:hover {
    border-bottom-color: transparent;
  }

  .year-tabs a:hover {
    background: #1f2937;
    border-bottom-color: #4b5563;
    color: #fff;
  }
</style>
