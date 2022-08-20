<script lang="ts">
  import { link, location } from "svelte-spa-router";
  import UploadButton from "./BalanceUploadButton.svelte";
  let page = "";

  $: {
    if ($location.indexOf("/balance/salary") >= 0) {
      page = "s";
    } else if ($location.indexOf("/balance/cost") >= 0) {
      page = "c";
    } else {
      page = "";
    }
  }
</script>

<div class="tabs is-centered">
  <ul>
    <li class:is-active={page == "s"}>
      <a href="/balance/salary/{new Date().getFullYear()}" use:link
        ><i class="material-icons">attach_money</i>salary</a
      >
    </li>
    <li class:is-active={page == "c"}>
      <a href="/balance/cost/{new Date().getFullYear()}" use:link
        ><i class="material-icons">payment</i>cost</a
      >
    </li>
    <li class="is-active">
      <UploadButton />
    </li>
  </ul>
</div>
