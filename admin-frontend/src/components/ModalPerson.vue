<script setup lang="ts">
import {
  MDBModal,
  MDBModalHeader,
  MDBModalTitle,
  MDBModalBody,
  MDBModalFooter,
  MDBBtn,
  MDBInput,
  MDBTextarea,
  MDBFile
} from 'mdb-vue-ui-kit';
import { ref } from 'vue';

const isModalOpen = ref(false);
</script>

<template>
  <MDBModal
      id="addPersonModal"
      tabindex="-1"
      labelledby="addPersonModalLabel"
      size="lg"
      v-model="isModalOpen"
  >
    <MDBModalHeader>
      <MDBModalTitle id="addPersonModalLabel">Добавить личность</MDBModalTitle>
    </MDBModalHeader>
    <MDBModalBody>
      <MDBInput
          type="text"
          label="Личность"
          id="formName"
          wrapperClass="mb-4"
      />

      <MDBInput
          type="email"
          label="Траектория"
          id="formEmail"
          wrapperClass="mb-4"
      />

      <MDBTextarea
          label="Описание"
          id="formTextarea"
          wrapperClass="mb-4"
      />

      <MDBFile label="Добавить фото" />
    </MDBModalBody>
    <MDBModalFooter>
      <MDBBtn color="secondary" @click="$emit('close')">Закрыть</MDBBtn>
      <MDBBtn color="success" @click="$emit('close')">Сохранить</MDBBtn>
    </MDBModalFooter>
  </MDBModal>
</template>

<style scoped>

</style>