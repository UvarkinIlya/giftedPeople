<script setup>
import {
  MDBBtn,
  MDBFile,
  MDBInput,
  MDBModal,
  MDBModalBody,
  MDBModalFooter,
  MDBModalHeader,
  MDBModalTitle,
  MDBTextarea
} from 'mdb-vue-ui-kit';
import {ref} from 'vue';
import axios from 'axios';

const emits = defineEmits(['close', 'update'])

const isModalOpen = ref(false);

const personName = ref('');
const personRoad = ref('');
const personDescription = ref('');
const personPhoto = ref();

const savePerson = async () => {
  const person = {
    name: personName.value,
    road: personRoad.value,
    description: personDescription.value,
    file: personPhoto.value,
  }

  console.log(person);
  try {
    const response = await axios.post("http://localhost:8080/person", person, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
  } catch (error) {
    console.log(error);
  }
  emits('close')
  emits('update')
};
</script>

<template>
  <MDBModal
      id="addPersonModal"
      tabindex="-1"
      labelledby="addPersonModalLabel"
      size="lg"
      v-on:submit.prevent="savePerson"
      v-model="isModalOpen"
  >
    <form>
      <MDBModalHeader>
        <MDBModalTitle id="addPersonModalLabel">Добавить личность</MDBModalTitle>
      </MDBModalHeader>
      <MDBModalBody>
        <MDBInput
            type="text"
            label="Личность"
            id="formName"
            wrapperClass="mb-4"
            v-model="personName"
        />

        <MDBInput
            type="text"
            label="Траектория"
            id="formEmail"
            wrapperClass="mb-4"
            v-model="personRoad"
        />

        <MDBTextarea
            label="Описание"
            id="formTextarea"
            wrapperClass="mb-4"
            v-model="personDescription"
        />

        <MDBFile label="Добавить фото" v-model="personPhoto"/>
      </MDBModalBody>
      <MDBModalFooter>
        <MDBBtn color="secondary" @click="$emit('close')">Закрыть</MDBBtn>
        <MDBBtn color="success" type="submit">Сохранить</MDBBtn>
      </MDBModalFooter>
    </form>
  </MDBModal>
</template>

<style scoped>

</style>