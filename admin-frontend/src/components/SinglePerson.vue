<script setup>
  import {ref} from "vue";

  const props = defineProps(["name", "img", "road", "description", "isLast"])

  const getImageUrl = (imgID) => {
    return `http://localhost:8080/person-image/${imgID}`;
  };
</script>

<template>
  <tr class="fw-normal">
    <th :class="isLast ? 'border-0' : ''">
      <div class="person-img-name-container">
        <img
            :src="getImageUrl(img)"
            alt="avatar 1" class="person-img">
        <span>{{ name }}</span>
      </div>
    </th>
    <td :class="isLast ? 'border-0' : ''" class="align-middle text-center">
      <!-- TODO add color    -->
      <h5 class="mb-0">{{ road }}</h5>
    </td>
    <td :class="isLast ? 'border-0' : ''" class="align-middle">
      {{ description }}
    </td>
    <td :class="isLast ? 'border-0' : ''" class="align-middle">
      <a href="#!" data-mdb-tooltip-init title="Done"><i
          class="fas fa-edit fa-lg text-warning me-3"></i></a>
      <a href="#!" data-mdb-tooltip-init title="Remove"><i
          class="fas fa-trash-alt fa-lg text-danger"></i></a>
    </td>
  </tr>
</template>

<style scoped>

</style>