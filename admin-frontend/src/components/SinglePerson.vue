<script setup>
  import {ref} from "vue";
  import axios from "axios";
  import ModalPerson from "./ModalPerson.vue";

  const props = defineProps(["id", "name", "img", "road", "description", "isLast"])

  const put = "PUT"
  const isModelOpen = ref(false);
  const getImageUrl = (imgID) => {
    return `http://localhost:8080/person-image/${imgID}`;
  };

  const emits = defineEmits(['updatePersons'])

  const removeItem = async (id) => {
    try {
      const response = await axios.delete("http://localhost:8080/person/"+id);
    } catch (error) {
      console.log(error);
    }
    emits('updatePersons')
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
      <a href="javascript:void(0)" @click="isModelOpen=true" data-mdb-tooltip-init title="Done"><i
          class="fas fa-edit fa-lg text-warning me-3"></i></a>
      <a href="javascript:void(0)" @click="removeItem(id)" data-mdb-tooltip-init title="Remove"><i
          class="fas fa-trash-alt fa-lg text-danger"></i></a>
    </td>
  </tr>

  <ModalPerson :person-id="id" :personName="name" :road="road" :description="description" :action="put" v-model="isModelOpen" @close="isModelOpen=false" @update="emits('updatePersons')"/>
</template>

<style scoped>

</style>