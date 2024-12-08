<script setup lang="ts">
  import {ref} from "vue";

  const props = defineProps(["name", "img", "road", "description"])
</script>

<template>
  <tr class="fw-normal">
    <th>
      <div class="person-img-name-container">
        <img
            :src="img"
            alt="avatar 1" class="person-img">
        <span>{{ name }}</span>
      </div>
    </th>
    <td class="align-middle text-center">
      <!-- TODO add color    -->
      <h5 class="mb-0">{{ road }}</h5>
    </td>
    <td class="align-middle">
      {{ description }}
    </td>
    <td class="align-middle">
      <a href="#!" data-mdb-tooltip-init title="Done"><i
          class="fas fa-edit fa-lg text-warning me-3"></i></a>
      <a href="#!" data-mdb-tooltip-init title="Remove"><i
          class="fas fa-trash-alt fa-lg text-danger"></i></a>
    </td>
  </tr>
</template>

<style scoped>

</style>